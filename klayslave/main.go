package main

//go:generate abigen --sol cpuHeavyTC/CPUHeavy.sol --pkg cpuHeavyTC --out cpuHeavyTC/CPUHeavy.go
//go:generate abigen --sol userStorageTC/UserStorage.sol --pkg userStorageTC --out userStorageTC/UserStorage.go

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/config"
	"github.com/kaiachain/kaia-load-tester/klayslave/vaultsetup"
	"github.com/kaiachain/kaia-load-tester/testcase"
	"github.com/kaiachain/kaia-load-tester/testcase/perpHalfFillTxTC"
	"github.com/kaiachain/kaia-load-tester/testcase/perpNoTradeTxTC"
	"github.com/kaiachain/kaia-load-tester/testcase/perpVaultMMTxTC"
	"github.com/myzhan/boomer"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Name = filepath.Base(os.Args[0])
	app.Usage = "This is for kaia load testing."
	app.Version = config.GetVersionWithCommit() // To see the version, run 'klayslave -v'
	app.HideVersion = false
	app.Copyright = "Copyright 2024 Kaia-load-tester authors"
	app.Flags = append(config.Flags, config.BoomerFlags...)

	// This app doesn't provide any subcommand
	//		app.Commands = []*cli.Command{}
	//		sort.Sort(cli.CommandsByName(app.Commands))
	//		app.CommandNotFound = nodecmd.CommandNotExist
	// app.OnUsageError = nodecmd.OnUsageError
	app.Before = func(cli *cli.Context) error {
		// runtime.GOMAXPROCS(runtime.NumCPU())
		if runtime.GOOS == "darwin" {
			return nil
		}
		return config.SetRLimit()
	}
	app.Action = RunAction
	app.After = func(cli *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func RunAction(ctx *cli.Context) {
	cfg := config.NewConfig(ctx)
	accGrp := account.NewAccGroup(cfg.GetChainID(), cfg.GetGasPrice(), cfg.GetBaseFee(), cfg.GetBatchSize(), cfg.InTheTcList("transferUnsignedTx"))
	var nUserForGaslessRevertTx, nUserForGaslessApproveTx int = 0, 0
	if cfg.InTheTcList("gaslessRevertTransactionTC") {
		nUserForGaslessRevertTx = cfg.GetNUserForSigned() // same as nUserForSignedTx
	}
	if cfg.InTheTcList("gaslessOnlyApproveTC") {
		nUserForGaslessApproveTx = cfg.GetNUserForSigned() // same as nUserForSignedTx
	}
	accGrp.CreateAccountsPerAccGrp(cfg.GetNUserForSigned(), cfg.GetNUserForUnsigned(), cfg.GetNUserForNewAccounts(), nUserForGaslessRevertTx, nUserForGaslessApproveTx, cfg.GetTcStrList(), cfg.GetGEndpoint())

	createTestAccGroupsAndPrepareContracts(cfg, accGrp)
	tasks := cfg.GetExtendedTasks()
	initializeTasks(cfg, accGrp, tasks)
	boomer.Run(toBoomerTasks(tasks)...)
}

// TODO-kaia-load-tester: remove global variables in the tc packages
func setSmartContractAddressPerPackage(a *account.AccGroup) {
}

// onlyPerpTCs reports whether every enabled test case is a perp one. Perp orders
// only need the USDT margin token, so setup can skip charging the other tokens.
func onlyPerpTCs(cfg *config.Config) bool {
	tcs := cfg.GetTcStrList()
	if len(tcs) == 0 {
		return false
	}
	for _, name := range tcs {
		if name != perpNoTradeTxTC.Name && name != perpHalfFillTxTC.Name && name != perpVaultMMTxTC.Name {
			return false
		}
	}
	return true
}

// createTestAccGroupsAndPrepareContracts do every init steps before task.Init
// those steps are about deploying test contracts and
func createTestAccGroupsAndPrepareContracts(cfg *config.Config, accGrp *account.AccGroup) *account.Account {
	if len(cfg.GetChargeValue().Bits()) == 0 {
		return nil
	}

	// 1. Import global reservoir Account and create local reservoir account
	globalReservoirAccount := account.GetAccountFromKey(0, cfg.GetRichWalletPrivateKey())
	localReservoirAccount := account.NewAccount(0)

	// 2. charge local reservoir
	_ = globalReservoirAccount.GetNonce(cfg.GetGCli())
	revertGroupChargeValue := new(big.Int).Mul(cfg.GetChargeValue(), big.NewInt(int64(len(accGrp.GetAccListByName(account.AccListForGaslessRevertTx)))))
	approveGroupChargeValue := new(big.Int).Mul(cfg.GetChargeValue(), big.NewInt(int64(len(accGrp.GetAccListByName(account.AccListForGaslessApproveTx)))))
	tx := globalReservoirAccount.TransferSignedTxWithGuaranteeRetry(cfg.GetGCli(), localReservoirAccount, new(big.Int).Add(cfg.GetTotalChargeValue(), new(big.Int).Add(revertGroupChargeValue, approveGroupChargeValue)))
	receipt, err := bind.WaitMined(context.Background(), cfg.GetGCli(), tx)
	if err != nil {
		log.Fatalf("receipt failed, err:%v", err.Error())
	}
	if receipt.Status != 1 {
		log.Fatalf("transfer for reservoir failed, localReservoir")
	}

	targetTokens := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
	if cfg.InTheTcList("ethLegacyTxTC") {
		targetTokens = []string{}
	} else if onlyPerpTCs(cfg) {
		// Perp test cases only need the USDT margin token ("2"); charging tokens
		// 3-15 to every account is pure setup overhead, so skip them.
		targetTokens = []string{"2"}
	} else if cfg.InTheTcList("tokenTransferTxTC") {
		targetTokens = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
	}

	// 3. charge KAIA
	if cfg.InTheTcList("transferTxTC") || cfg.InTheTcList("ethLegacyTxTC") {
		log.Printf("Start charging KLAY to test accounts because transferTxTC and/or ethLegacyTxTC is enabled")
		accs := accGrp.GetValidAccGrp()
		accs = append(accs, accGrp.GetAccListByName(account.AccListForGaslessRevertTx)...)  // for avoid validation
		accs = append(accs, accGrp.GetAccListByName(account.AccListForGaslessApproveTx)...) // for avoid validation
		gasFee := big.NewInt(25e9 * 21000)
		account.HierarchicalDistribute(accs, localReservoirAccount, cfg.GetChargeValue(), gasFee, func(from, to *account.Account, value *big.Int) {
			from.TransferSignedTxWithGuaranteeRetry(cfg.GetGCli(), to, value)
		})
		log.Printf("Finished charging KLAY to %d test account(s)\n", len(accs))
	} else {
		log.Printf("Start charging Tokens [%s] to test accounts in parallel", strings.Join(targetTokens, ","))
		accs := accGrp.GetValidAccGrp()
		accs = append(accs, accGrp.GetAccListByName(account.AccListForGaslessRevertTx)...)  // for avoid validation
		accs = append(accs, accGrp.GetAccListByName(account.AccListForGaslessApproveTx)...) // for avoid validation

		var wg sync.WaitGroup
		for _, token := range targetTokens {
			token := token
			wg.Add(1)
			go func() {
				defer wg.Done()
				// top up this token to local reservoir
				tx := globalReservoirAccount.TransferTokenSignedTxWithGuaranteeRetry(cfg.GetGCli(), localReservoirAccount, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18)), token)
				receipt, err := bind.WaitMined(context.Background(), cfg.GetGCli(), tx)
				if err != nil {
					log.Fatalf("receipt failed, err:%v", err.Error())
				}
				if receipt.Status != 1 {
					log.Fatalf("transfer for reservoir failed, localReservoir")
				}
				// distribute to all accounts
				value := new(big.Int).Mul(big.NewInt(1e10), big.NewInt(1e18))
				gasFee := big.NewInt(25e9 * 21000)
				account.HierarchicalDistribute(accs, localReservoirAccount, value, gasFee, func(from, to *account.Account, value *big.Int) {
					from.TransferTokenSignedTxWithGuaranteeRetry(cfg.GetGCli(), to, value, token)
				})
				log.Printf("Finished charging Token \"%s\" to %d test account(s)\n", token, len(accs))
			}()
		}
		wg.Wait()
	}

	// Wait, charge KAIA happen in 100% of all created test accounts
	// But, from here including prepareTestContracts like MintERC721, only 20% of account happens
	accGrp.SetAccGrpByActivePercent(cfg.GetActiveUserPercent())

	// Perp test cases (perpNoTradeTxTC, perpHalfFillTxTC, perpVaultMMTxTC) share the same
	// market and margin setup: inject the market id / reference price / tick size, then
	// move USDT (token "2", the only genesis-whitelisted perp deposit token) from each
	// active account's spot wallet into its perp wallet so orders have order margin.
	// Mirrors the spot token charging pre-work above; the accounts already hold USDT
	// from it.
	if cfg.InTheTcList(perpNoTradeTxTC.Name) || cfg.InTheTcList(perpHalfFillTxTC.Name) || cfg.InTheTcList(perpVaultMMTxTC.Name) {
		mktId, ref, tick := cfg.GetPerpMarketId(), cfg.GetPerpRefPrice(), cfg.GetPerpTickSize()
		if cfg.InTheTcList(perpNoTradeTxTC.Name) {
			perpNoTradeTxTC.SetMarketId(mktId)
			perpNoTradeTxTC.SetRefPrice(ref)
			perpNoTradeTxTC.SetTickSize(tick)
		}
		if cfg.InTheTcList(perpHalfFillTxTC.Name) {
			perpHalfFillTxTC.SetMarketId(mktId)
			perpHalfFillTxTC.SetRefPrice(ref)
			perpHalfFillTxTC.SetTickSize(tick)
		}

		// Generous margin so orders never exhaust available balance during a run.
		// Must not exceed each account's charged USDT spot balance.
		perpDepositAmount := new(big.Int).Mul(big.NewInt(1e9), big.NewInt(1e18))
		accs := accGrp.GetValidAccGrp()
		log.Printf("Start perp-depositing USDT to %d test account(s)", len(accs))
		var wg sync.WaitGroup
		for _, acc := range accs {
			acc := acc
			wg.Add(1)
			go func() {
				defer wg.Done()
				acc.PerpDepositWithGuaranteeRetry(cfg.GetGCli(), "2", perpDepositAmount)
			}()
		}
		wg.Wait()
		log.Printf("Finished perp-depositing USDT to %d test account(s)", len(accs))

		// perpVaultMMTxTC needs an on-chain PerpVault plus a set of registered executor
		// session keys before it can run: deploy the vault, deposit owner margin into
		// it, and register the executors, then hand the vault + keys to the TC.
		if cfg.InTheTcList(perpVaultMMTxTC.Name) {
			perpVaultMMTxTC.SetMarketId(mktId)
			perpVaultMMTxTC.SetRefPrice(ref)
			perpVaultMMTxTC.SetTickSize(tick)

			owner := accs[0] // holds spot USDT from the token charging above

			// The L2 runs with a zero gas price, so the owner needs no native KAIA:
			// vaultsetup's regular EVM txs (deploy impl/proxy, N x addExecutor) cost
			// nothing, and the margin deposit is a gas-free VaultDeposit dex command.
			log.Printf("Deploying PerpVault + %d executors (owner=%s)", cfg.GetVaultExecutorCount(), owner.GetAddress().Hex())
			vault, execKeys, err := vaultsetup.Deploy(
				cfg.GetGCli(), owner,
				cfg.GetVaultTokenId(), "LoadTestVault", "LTV",
				cfg.GetVaultLockupPeriod(), cfg.GetVaultDepositAmount(),
				cfg.GetVaultExecutorCount(),
				mktId, // marketId: restricts registered executors to this market
			)
			if err != nil {
				log.Fatalf("PerpVault setup failed: %v", err)
			}
			perpVaultMMTxTC.SetVault(vault, execKeys)
			log.Printf("PerpVault ready: vault=%s executors=%d", vault.Hex(), len(execKeys))
		}
	}

	// Set SmartContractAddress value in each packages if needed
	setSmartContractAddressPerPackage(accGrp)
	return localReservoirAccount
}

func initializeTasks(cfg *config.Config, accGrp *account.AccGroup, tasks []*testcase.ExtendedTask) {
	println("Initializing tasks")

	// Tc package initializes the task
	for _, extendedTask := range tasks {
		accs := accGrp.GetAccListByName(account.AccListForSignedTx)
		if extendedTask.Name == "transferUnsignedTx" {
			accs = accGrp.GetAccListByName(account.AccListForUnsignedTx)
		} else if extendedTask.Name == "gaslessRevertTransactionTC" {
			accs = accGrp.GetAccListByName(account.AccListForGaslessRevertTx)
		} else if extendedTask.Name == "gaslessOnlyApproveTC" {
			accs = accGrp.GetAccListByName(account.AccListForGaslessApproveTx)
		}
		extendedTask.Init(accs, cfg.GetGEndpoint(), cfg.GetGasPrice())
		println("=> " + extendedTask.Name + " extendedTask is initialized.")
	}
}

func toBoomerTasks(tasks []*testcase.ExtendedTask) []*boomer.Task {
	var boomerTasks []*boomer.Task
	for _, task := range tasks {
		boomerTasks = append(boomerTasks, &boomer.Task{Weight: task.Weight, Fn: task.Fn, Name: task.Name})
	}
	return boomerTasks
}

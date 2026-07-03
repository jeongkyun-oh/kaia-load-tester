package vaultsetup

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/contracts/vaultgen"
)

const permPerpOnly = uint8(1) // types.SessionPermPerpOnly

// waitMined polls until the tx receipt is available and status==success.
func waitMined(cli *ethclient.Client, tx *types.Transaction, what string) error {
	for i := 0; i < 60; i++ {
		r, err := cli.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			if r.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("%s tx %s failed status=%d", what, tx.Hash().Hex(), r.Status)
			}
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("%s tx %s not mined", what, tx.Hash().Hex())
}

// Deploy performs the full per-slave vault setup: deploy impl + proxy(initialize),
// deposit margin, and register `execCount` fresh executor keys restricted to
// `marketId` (PerpOnly permission).
func Deploy(cli *ethclient.Client, owner *account.Account, tokenId, name, symbol string, lockup, deposit *big.Int, execCount int, marketId uint64) (common.Address, []*ecdsa.PrivateKey, error) {
	tid, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return common.Address{}, nil, fmt.Errorf("invalid tokenId %q", tokenId)
	}

	auth, err := owner.TransactOpts()
	if err != nil {
		return common.Address{}, nil, err
	}

	// 1) deploy implementation
	implAddr, implTx, _, err := vaultgen.DeployPerpVault(auth, cli)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("deploy impl: %w", err)
	}
	if err := waitMined(cli, implTx, "deploy impl"); err != nil {
		return common.Address{}, nil, err
	}

	// 2) encode initialize(name, symbol, tokenId, lockup) as proxy constructor data
	parsed, err := vaultgen.PerpVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, err
	}
	initData, err := parsed.Pack("initialize", name, symbol, tid.Uint64(), lockup)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("pack initialize: %w", err)
	}

	// 3) deploy proxy (runs initialize in constructor → owner = this sender)
	auth2, err := owner.TransactOpts()
	if err != nil {
		return common.Address{}, nil, err
	}
	proxyAddr, proxyTx, err := vaultgen.DeployPerpVaultProxy(auth2, cli, implAddr, initData)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("deploy proxy: %w", err)
	}
	if err := waitMined(cli, proxyTx, "deploy proxy"); err != nil {
		return common.Address{}, nil, err
	}

	vault, err := vaultgen.NewPerpVault(proxyAddr, cli)
	if err != nil {
		return common.Address{}, nil, err
	}

	// 4) deposit margin into the vault. PerpVault.deposit → ArbDex.vaultDeposit
	// only accepts deposits authorized by the dex-command vault handlers (they
	// stash a tx-scoped VaultDepositAuth before delegating; DEX-2045/ASA-60), so
	// a direct EVM call to deposit() always reverts "vault deposit: unauthorized".
	// Submit a VaultDeposit dex command instead (Path A: outer tx signed by the
	// owner, gas-free like every dex-command tx).
	depTx, err := owner.GenVaultDepositTx(proxyAddr, deposit)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("vault deposit: %w", err)
	}
	if _, err := owner.SendTx(cli, depTx); err != nil {
		return common.Address{}, nil, fmt.Errorf("vault deposit: %w", err)
	}
	if err := waitMined(cli, depTx, "vault deposit"); err != nil {
		return common.Address{}, nil, err
	}

	// 5) register executors, restricted to the single target perp market
	expiresAt := uint64(time.Now().Unix()) + 180*24*60*60 // 180d cap
	allowedMarkets := []uint64{marketId}
	execKeys := make([]*ecdsa.PrivateKey, 0, execCount)
	for i := 0; i < execCount; i++ {
		k, err := crypto.GenerateKey()
		if err != nil {
			return common.Address{}, nil, err
		}
		execKeys = append(execKeys, k)
		authN, err := owner.TransactOpts()
		if err != nil {
			return common.Address{}, nil, err
		}
		addTx, err := vault.AddExecutor(authN, crypto.PubkeyToAddress(k.PublicKey), expiresAt, permPerpOnly, []uint8{}, allowedMarkets)
		if err != nil {
			return common.Address{}, nil, fmt.Errorf("addExecutor %d: %w", i, err)
		}
		if err := waitMined(cli, addTx, fmt.Sprintf("addExecutor %d", i)); err != nil {
			return common.Address{}, nil, err
		}
	}

	return proxyAddr, execKeys, nil
}

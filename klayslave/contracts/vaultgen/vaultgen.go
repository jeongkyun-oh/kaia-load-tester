// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultgen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PerpVaultMetaData contains all meta data concerning the PerpVault contract.
var PerpVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllDexPermissionNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositCapExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"cmd\",\"type\":\"uint8\"}],\"name\":\"InvalidPerpCommand\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LockupExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LockupNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketRulesLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxMarginUtilizationTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"MinDepositNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"MinWithdrawNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelfAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShareTransferDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SyncWithdrawDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyPendingRequests\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferOfEscrowedShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"tokenId\",\"type\":\"uint64\"}],\"name\":\"UnregisteredPerpDepositToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WhitelistLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroShares\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BackstopDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BackstopEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Donated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"LockupPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"marketId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"MarketRuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"MaxMarginUtilizationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinWithdrawUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executableAt\",\"type\":\"uint256\"}],\"name\":\"RedeemRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"RequestExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"RequestExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhitelistEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"WhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executableAt\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_LOCKUP_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCKUP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PENDING_PER_OWNER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"permType\",\"type\":\"uint8\"},{\"internalType\":\"uint8[]\",\"name\":\"allowedCmds\",\"type\":\"uint8[]\"},{\"internalType\":\"uint64[]\",\"name\":\"allowedMarkets\",\"type\":\"uint64[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backstopEnabledAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBackstop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBackstop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"escrowedSharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"executeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"marketId\",\"type\":\"uint64\"}],\"name\":\"getMarketCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMarginUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"requestedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executableAt\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"tokenId_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lockupPeriod_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBackstopEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"removeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"requestRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCap\",\"type\":\"uint256\"}],\"name\":\"setDepositCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockupPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"marketIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256[]\",\"name\":\"caps\",\"type\":\"uint256[]\"}],\"name\":\"setMarketRules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"setMaxMarginUtilization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMin\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMin\",\"type\":\"uint256\"}],\"name\":\"setMinWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"allowed\",\"type\":\"bool[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"permType\",\"type\":\"uint8\"},{\"internalType\":\"uint8[]\",\"name\":\"allowedCmds\",\"type\":\"uint8[]\"},{\"internalType\":\"uint64[]\",\"name\":\"allowedMarkets\",\"type\":\"uint64[]\"}],\"name\":\"updateExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000e8565b600054610100900460ff1615620000935760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051615c63620001206000396000818161175a015281816117f001528181611b9101528181611c270152611d220152615c636000f3fe6080604052600436106104ec5760003560e01c80636e417c6911610294578063c58343ef1161015e578063d96c1284116100d6578063ee947a7c1161008a578063f14faf6f1161006f578063f14faf6f14610f21578063f2fde38b14610f41578063f3ae241514610f6157600080fd5b8063ee947a7c14610eed578063ef8b30f714610d1057600080fd5b8063e699734a116100bb578063e699734a14610e7e578063e81cc3cc14610e9e578063ea77501714610ed557600080fd5b8063d96c128414610e04578063dd62ed3e14610e3857600080fd5b8063ccc143b81161012d578063d157485511610112578063d157485514610d90578063d685da8d14610db0578063d905777e14610de457600080fd5b8063ccc143b814610d50578063ce96cb7714610d7057600080fd5b8063c58343ef14610c80578063c63d75b614610cf0578063c6e6f59214610d10578063c771c39014610d3057600080fd5b806394bf804d1161020c578063b187bd26116101c0578063b460af94116101a5578063b460af9414610c40578063b4b78d7714610c60578063ba08765214610c4057600080fd5b8063b187bd2614610c0b578063b3d7f6b914610c2057600080fd5b8063a457c2d7116101f1578063a457c2d714610bab578063a9059cbb14610bcb578063ac18de4314610beb57600080fd5b806394bf804d14610b7657806395d89b4114610b9657600080fd5b80638456cb59116102635780638da5cb5b116102485780638da5cb5b14610ae25780638fcc9cfb14610b0057806393ad3e0314610b2057600080fd5b80638456cb5914610aad5780638665120314610ac257600080fd5b80636e417c6914610a225780636e553f6514610a4257806370a0823114610a62578063715018a614610a9857600080fd5b80633507e3cb116103d5578063402d267d1161034d57806352d1902d116103015780635c975abb116102e65780635c975abb146109d55780635ef12042146109ed5780636a7d051d14610a0257600080fd5b806352d1902d146109ab578063595f427e146109c057600080fd5b80634d93b6aa116103325780634d93b6aa1461092a5780634f1ef2861461096157806351fb012d1461097457600080fd5b8063402d267d1461090a5780634cdad506146105e257600080fd5b806338d52e0f116103a45780633af32abf116103895780633af32abf1461087d5780633b99adf7146108d55780633f4ba83a146108f557600080fd5b806338d52e0f1461082b578063395093511461085d57600080fd5b80633507e3cb146107b657806335aa134a146107d65780633659cfe6146107f6578063375126b61461081657600080fd5b80630eaad3f11161046857806323b872dd1161043757806327f69b931161041c57806327f69b93146107655780632d06177a1461077a578063313ce5671461079a57600080fd5b806323b872dd14610725578063247884291461074557600080fd5b80630eaad3f114610686578063107703ab146106ba57806317d70f7c146106da57806318160ddd1461071057600080fd5b806306fdde03116104bf578063095ea7b3116104a4578063095ea7b3146106025780630a28a477146106325780630d2431bf1461065257600080fd5b806306fdde03146105c057806307a2d13a146105e257600080fd5b806301a598da146104f157806301e1d11414610534578063052d9e7e1461054957806306161f4c1461056b575b600080fd5b3480156104fd57600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad02545b6040519081526020015b60405180910390f35b34801561054057600080fd5b50610521610fb9565b34801561055557600080fd5b506105696105643660046153ef565b61103a565b005b34801561057757600080fd5b50610521610586366004615428565b6001600160a01b031660009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad05602052604090205490565b3480156105cc57600080fd5b506105d56110a8565b60405161052b9190615467565b3480156105ee57600080fd5b506105216105fd36600461549a565b61113a565b34801561060e57600080fd5b5061062261061d3660046154b3565b61114d565b604051901515815260200161052b565b34801561063e57600080fd5b5061052161064d36600461549a565b611168565b34801561065e57600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0e54610521565b34801561069257600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0a54610521565b3480156106c657600080fd5b506105216106d53660046154dd565b611175565b3480156106e657600080fd5b50600080516020615be78339815191525460405167ffffffffffffffff909116815260200161052b565b34801561071c57600080fd5b50603554610521565b34801561073157600080fd5b50610622610740366004615509565b6113c3565b34801561075157600080fd5b50610569610760366004615428565b6113e0565b34801561077157600080fd5b50610521600081565b34801561078657600080fd5b50610569610795366004615428565b6114f2565b3480156107a657600080fd5b506040516012815260200161052b565b3480156107c257600080fd5b506105696107d13660046155ba565b611565565b3480156107e257600080fd5b506105696107f136600461549a565b6116d7565b34801561080257600080fd5b50610569610811366004615428565b611750565b34801561082257600080fd5b50610521608081565b34801561083757600080fd5b506065546001600160a01b03165b6040516001600160a01b03909116815260200161052b565b34801561086957600080fd5b506106226108783660046154b3565b6118ed565b34801561088957600080fd5b50610622610898366004615428565b6001600160a01b031660009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0f602052604090205460ff1690565b3480156108e157600080fd5b506105696108f036600461565c565b611936565b34801561090157600080fd5b50610569611ab2565b34801561091657600080fd5b50610521610925366004615428565b611b24565b34801561093657600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0c5460ff16610622565b61056961096f366004615754565b611b87565b34801561098057600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad105460ff16610622565b3480156109b757600080fd5b50610521611d15565b3480156109cc57600080fd5b50610569611dda565b3480156109e157600080fd5b5060c95460ff16610622565b3480156109f957600080fd5b50610569611f0c565b348015610a0e57600080fd5b50610569610a1d36600461565c565b612065565b348015610a2e57600080fd5b50610521610a3d36600461549a565b612215565b348015610a4e57600080fd5b50610521610a5d3660046154dd565b6125f6565b348015610a6e57600080fd5b50610521610a7d366004615428565b6001600160a01b031660009081526033602052604090205490565b348015610aa457600080fd5b506105696128cf565b348015610ab957600080fd5b506105696128e1565b348015610ace57600080fd5b50610569610add36600461549a565b612951565b348015610aee57600080fd5b506097546001600160a01b0316610845565b348015610b0c57600080fd5b50610569610b1b36600461549a565b6129dd565b348015610b2c57600080fd5b50610521610b3b3660046157b6565b67ffffffffffffffff1660009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad07602052604090205490565b348015610b8257600080fd5b50610521610b913660046154dd565b612a56565b348015610ba257600080fd5b506105d5612d1e565b348015610bb757600080fd5b50610622610bc63660046154b3565b612d2d565b348015610bd757600080fd5b50610622610be63660046154b3565b612de2565b348015610bf757600080fd5b50610569610c06366004615428565b612df6565b348015610c1757600080fd5b50610622612e66565b348015610c2c57600080fd5b50610521610c3b36600461549a565b612e74565b348015610c4c57600080fd5b50610521610c5b3660046157d1565b612e81565b348015610c6c57600080fd5b50610569610c7b36600461549a565b612eb5565b348015610c8c57600080fd5b50610ca0610c9b36600461549a565b612fce565b604080516001600160a01b03988916815297909616602088015294860193909352606085019190915267ffffffffffffffff90811660808501521660a083015260ff1660c082015260e00161052b565b348015610cfc57600080fd5b50610521610d0b366004615428565b61307e565b348015610d1c57600080fd5b50610521610d2b36600461549a565b6130a4565b348015610d3c57600080fd5b50610569610d4b36600461549a565b6130b1565b348015610d5c57600080fd5b50610521610d6b3660046154dd565b6131c8565b348015610d7c57600080fd5b50610521610d8b366004615428565b61341d565b348015610d9c57600080fd5b50610569610dab36600461582d565b61348d565b348015610dbc57600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0b54610521565b348015610df057600080fd5b50610521610dff366004615428565b6137f9565b348015610e1057600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0d54610521565b348015610e4457600080fd5b50610521610e533660046158a9565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b348015610e8a57600080fd5b50610569610e993660046155ba565b613854565b348015610eaa57600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad085460ff16610622565b348015610ee157600080fd5b506105216301e1338081565b348015610ef957600080fd5b507f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0154610521565b348015610f2d57600080fd5b50610569610f3c36600461549a565b6139c6565b348015610f4d57600080fd5b50610569610f5c366004615428565b613b23565b348015610f6d57600080fd5b50610622610f7c366004615428565b6001600160a01b031660009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff1690565b6040517f792e6c4d00000000000000000000000000000000000000000000000000000000815230600482015260009060cb9063792e6c4d90602401602060405180830381865afa158015611011573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103591906158d3565b905090565b611042613bb0565b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad10805482151560ff19909116811790915560408051918252517f49d3057180a80162d2a0381be6848c15e0d117e900366482dd3b5443ca8db9749181900360200190a150565b6060603680546110b7906158ec565b80601f01602080910402602001604051908101604052809291908181526020018280546110e3906158ec565b80156111305780601f1061110557610100808354040283529160200191611130565b820191906000526020600020905b81548152906001019060200180831161111357829003601f168201915b5050505050905090565b6000611147826000613c0a565b92915050565b6000611157613c42565b6111618383613c95565b9392505050565b6000611147826001613ca3565b600061117f613c42565b600260fb54036111d65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b600260fb557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad1054600080516020615be78339815191529060ff16801561122e5750336000908152600f8201602052604090205460ff16155b1561124e57604051636f8bf18b60e11b81523360048201526024016111cd565b6001600160a01b038316331461127757604051636c16790960e01b815260040160405180910390fd5b8360000361129857604051639811e0c760e01b815260040160405180910390fd5b60006112a38561113a565b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0b54909150600080516020615be78339815191529080158015906112e657508083105b15611327576040517f621206b300000000000000000000000000000000000000000000000000000000815260048101829052602481018490526044016111cd565b611332878488613cd2565b6000818152600484810160209081526040928390209091015482518b815291820187905268010000000000000000900467ffffffffffffffff16918101919091529095506001600160a01b03871690339087907f3b8064ce836b1010de5808e04b451d4ebfe1157c492ad0e1b6d8ccffb329eb39906060015b60405180910390a45050600160fb5550909392505050565b60006113cd613c42565b6113d8848484613f92565b949350505050565b6097546001600160a01b0316331480159061142a57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15611448576040516305d1403760e01b815260040160405180910390fd5b6040517f913004740000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260cb90639130047490602401600060405180830381600087803b1580156114a357600080fd5b505af11580156114b7573d6000803e3d6000fd5b50506040516001600160a01b03841692507f4a2cf608bfb427f53279ec7f0eadf48913b9346ccefc3af138dbdec14ea0907d9150600090a250565b6114fa613bb0565b6001600160a01b03811660008181527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad096020526040808220805460ff19166001179055517f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a9190a250565b6097546001600160a01b031633148015906115af57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b156115cd576040516305d1403760e01b815260040160405180910390fd5b8460ff1660000361160a576040517faeb47d8100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460ff1660020361161f5761161f8484613fab565b6040517fced9f79800000000000000000000000000000000000000000000000000000000815260cb9063ced9f79890611668908a908a908a908a908a908a908a90600401615926565b600060405180830381600087803b15801561168257600080fd5b505af1158015611696573d6000803e3d6000fd5b50506040516001600160a01b038a1692507fae5b7c3b000f575c241001dc9bcb3d8778376889353b07121115574eceff78c59150600090a250505050505050565b6116df613bb0565b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0b5460408051918252602082018390528051600080516020615be7833981519152927f3c4f4d8cd2a65b4b1f4eeaf43669b14ab54e43d4842aa0ac8f0e4f9fe0bf5bf992908290030190a1600b0155565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036117ee5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016111cd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166118497f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146118c55760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f7879000000000000000000000000000000000000000060648201526084016111cd565b6118ce8161405f565b604080516000808252602082019092526118ea91839190614067565b50565b3360008181526034602090815260408083206001600160a01b038716845290915281205490919061192c90829086906119279087906159f3565b614207565b5060019392505050565b61193e613bb0565b828114611977576040517fe05561a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080516020615be783398151915260005b84811015611aaa578383828181106119a3576119a3615a06565b90506020020160208101906119b891906153ef565b82600f0160008888858181106119d0576119d0615a06565b90506020020160208101906119e59190615428565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055858582818110611a1f57611a1f615a06565b9050602002016020810190611a349190615428565b6001600160a01b03167ff93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d858584818110611a7057611a70615a06565b9050602002016020810190611a8591906153ef565b604051901515815260200160405180910390a280611aa281615a1c565b915050611989565b505050505050565b6097546001600160a01b03163314801590611afc57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15611b1a576040516305d1403760e01b815260040160405180910390fd5b611b2261435f565b565b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0254600090808203611b5a575060001992915050565b6000611b64610fb9565b905081811015611b7d57611b788183615a36565b6113d8565b6000949350505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003611c255760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016111cd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611c807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614611cfc5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f7879000000000000000000000000000000000000000060648201526084016111cd565b611d058261405f565b611d1182826001614067565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611db55760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016111cd565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6097546001600160a01b03163314801590611e2457503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15611e42576040516305d1403760e01b815260040160405180910390fd5b600080516020615be78339815191526000611e5b610fb9565b600c8301805460ff19166001179055600d8301819055604080517f99aa23e7000000000000000000000000000000000000000000000000000000008152905191925060cb916399aa23e79160048082019260009290919082900301818387803b158015611ec757600080fd5b505af1158015611edb573d6000803e3d6000fd5b50506040517f8c5c1631fda9e60c3d1c47a8effa57d678e664e85a7f797124006c9d93d444c3925060009150a15050565b6097546001600160a01b03163314801590611f5657503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15611f74576040516305d1403760e01b815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0c805460ff1916905560007f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0d819055604080517f821778810000000000000000000000000000000000000000000000000000000081529051600080516020615be78339815191529260cb926382177881926004808301939282900301818387803b15801561202157600080fd5b505af1158015612035573d6000803e3d6000fd5b50506040517f6e407d4f9fe5e6e3d2a2ad7fd3a065a65774c4b0b22c8bbec1b99476258bb7f0925060009150a150565b6097546001600160a01b031633148015906120af57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b156120cd576040516305d1403760e01b815260040160405180910390fd5b828114612106576040517fba5173a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080516020615be783398151915260005b84811015611aaa5783838281811061213257612132615a06565b9050602002013582600701600088888581811061215157612151615a06565b905060200201602081019061216691906157b6565b67ffffffffffffffff16815260208101919091526040016000205585858281811061219357612193615a06565b90506020020160208101906121a891906157b6565b67ffffffffffffffff167fb5c53be95670e712b0ed22b769b3b8ae589ed033c496933be419cdf1c9ee955a8585848181106121e5576121e5615a06565b905060200201356040516121fb91815260200190565b60405180910390a28061220d81615a1c565b915050612118565b600061221f613c42565b600260fb54036122715760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016111cd565b600260fb81905560008381527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad04602052604090206004810154600080516020615be783398151915292600191700100000000000000000000000000000000900460ff16908111156122e4576122e4615a49565b1461231b576040517f2efb6afd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600481015468010000000000000000900467ffffffffffffffff1642101561236f576040517f3db5d6a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600381015482546040517f9328d7b900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482018190526024820183905291945060cb90639328d7b990604401600060405180830381600087803b1580156123df57600080fd5b505af19250505080156123f0575060015b61244957600182015482546040518681526001600160a01b03928316929091169087907f9aaaffd6862f8fa67d255c221adc58895928765ead274e59afe9ff06cd0d27d69060200160405180910390a4600093506125e9565b6004820180547fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff16700200000000000000000000000000000000179055600282015482546001600160a01b03166000908152600585016020526040812080549091906124b6908490615a36565b909155505081546001600160a01b0316600090815260068401602052604081208054916124e283615a5f565b909155505081546002830154612501916001600160a01b0316906143b1565b60018201546040517fc37581ef0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015267ffffffffffffffff821660248201526044810185905260cb9063c37581ef90606401600060405180830381600087803b15801561257957600080fd5b505af115801561258d573d6000803e3d6000fd5b5050505060018201548254600284015460408051918252602082018890526001600160a01b03938416939092169188917fa7a5da0a86cebe498b65e58861c57870380d8ea9e66683ec37d6ad37181cab84910160405180910390a45b5050600160fb5550919050565b6000612600613c42565b600260fb54036126525760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016111cd565b600260fb557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad1054600080516020615be78339815191529060ff1680156126aa5750336000908152600f8201602052604090205460ff16155b156126ca57604051636f8bf18b60e11b81523360048201526024016111cd565b6001600160a01b03831633146126f357604051636c16790960e01b815260040160405180910390fd5b6126fc83611b24565b841115612735576040517f935d630c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0a54801580159061276557508085105b156127a6576040517f5af7f70700000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016111cd565b6127af856130a4565b9250826000036127d257604051639811e0c760e01b815260040160405180910390fd5b6127dc8484614542565b60cb63ff3e9aa933600080516020615be78339815191525460405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b03909216600483015267ffffffffffffffff16602482015260448101889052606401600060405180830381600087803b15801561286357600080fd5b505af1158015612877573d6000803e3d6000fd5b505060408051888152602081018790526001600160a01b03881693503392507fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d791015b60405180910390a35050600160fb5592915050565b6128d7613bb0565b611b22600061462d565b6097546001600160a01b0316331480159061292b57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15612949576040516305d1403760e01b815260040160405180910390fd5b611b2261468c565b6097546001600160a01b0316331480159061299b57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b156129b9576040516305d1403760e01b815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0255565b6129e5613bb0565b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0a5460408051918252602082018390528051600080516020615be7833981519152927fb566d3df2587c9e70b06b6419bdeeeeec8ca8cd60e4c48c6baad0d94c46809c792908290030190a1600a0155565b6000612a60613c42565b600260fb5403612ab25760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016111cd565b600260fb557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad1054600080516020615be78339815191529060ff168015612b0a5750336000908152600f8201602052604090205460ff16155b15612b2a57604051636f8bf18b60e11b81523360048201526024016111cd565b6001600160a01b0383163314612b5357604051636c16790960e01b815260040160405180910390fd5b612b5c8361307e565b841115612b95576040517f935d630c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b9e84612e74565b915081600003612bc157604051630cb65c7760e21b815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0a548015801590612bf157508083105b15612c32576040517f5af7f70700000000000000000000000000000000000000000000000000000000815260048101829052602481018490526044016111cd565b612c3c8486614542565b60cb63ff3e9aa933600080516020615be78339815191525460405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b03909216600483015267ffffffffffffffff16602482015260448101869052606401600060405180830381600087803b158015612cc357600080fd5b505af1158015612cd7573d6000803e3d6000fd5b505060408051868152602081018990526001600160a01b03881693503392507fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d791016128ba565b6060603780546110b7906158ec565b3360008181526034602090815260408083206001600160a01b038716845290915281205490919083811015612dca5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016111cd565b612dd78286868403614207565b506001949350505050565b6000612dec613c42565b61116183836146c9565b612dfe613bb0565b6001600160a01b03811660008181527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad096020526040808220805460ff19169055517fef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd319190a250565b600061103560c95460ff1690565b6000611147826001613c0a565b60006040517f7a5d846a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6097546001600160a01b03163314801590612eff57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15612f1d576040516305d1403760e01b815260040160405180910390fd5b612710811115612f59576040517f817ded3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0e80549082905560408051828152602081018490528151600080516020615be783398151915293927f04ae0db7c0aa77a9404f22d44aa47d647c6427822272a8a6d5fd43dc5c204e07928290030190a1505050565b600080808080808080600080516020615be783398151915260008a8152600491820160205260409020805460018201546002808401546003850154958501549496506001600160a01b0393841695929093169367ffffffffffffffff808216926801000000000000000083049091169160ff700100000000000000000000000000000000909104169081111561306657613066615a49565b959f949e50929c50909a509850965090945092505050565b60008061308a83611b24565b9050600019810361309f575060001992915050565b611161815b6000611147826000613ca3565b6097546001600160a01b031633148015906130fb57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b15613119576040516305d1403760e01b815260040160405180910390fd5b6301e13380811115613157576040517f73397ddd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad015460408051918252602082018390528051600080516020615be7833981519152927fa4711e25376fde3c9a55e3dba876110179242e2dcb40fd6644c6d7884642a1f592908290030190a160010155565b60006131d2613c42565b600260fb54036132245760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016111cd565b600260fb557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad1054600080516020615be78339815191529060ff16801561327c5750336000908152600f8201602052604090205460ff16155b1561329c57604051636f8bf18b60e11b81523360048201526024016111cd565b6001600160a01b03831633146132c557604051636c16790960e01b815260040160405180910390fd5b836000036132e657604051630cb65c7760e21b815260040160405180910390fd5b7f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0b54600080516020615be783398151915290801580159061332657508086105b15613367576040517f621206b300000000000000000000000000000000000000000000000000000000815260048101829052602481018790526044016111cd565b600061337287611168565b90508060000361339557604051639811e0c760e01b815260040160405180910390fd5b6133a0818888613cd2565b6000818152600485810160209081526040928390209091015482518581529182018b905268010000000000000000900467ffffffffffffffff16918101919091529095506001600160a01b03871690339087907f74a8d0df732a141e45f44b230aadff310598bbeabb761b313e96ae2e27337a79906060016113ab565b6001600160a01b0381166000908152603360209081526040808320547f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad059092528220548281831161346f576000613479565b6134798284615a36565b90506134848161113a565b95945050505050565b600054610100900460ff16158080156134ad5750600054600160ff909116105b806134c75750303b1580156134c7575060005460ff166001145b6135395760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016111cd565b6000805460ff19166001179055801561355c576000805461ff0019166101001790555b61356660006146d7565b613570858561474b565b6135786147c0565b613580614833565b6135886148a6565b613590614919565b6000821561359e57826135a1565b60005b90506301e133808111156135e1576040517f73397ddd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805167ffffffffffffffff861660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f877fa889000000000000000000000000000000000000000000000000000000001790529051600091829160ce9161366591615a76565b600060405180830381855afa9150503d80600081146136a0576040519150601f19603f3d011682016040523d82523d6000602084013e6136a5565b606091505b50915091508180156136b957506020815110155b1561371b576000818060200190518101906136d49190615a92565b905080613719576040517f1018d0a900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff881660048201526024016111cd565b505b5050600080516020615be783398151915280547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86161790557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad01557f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad08805460ff1916600117905580156137f2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6001600160a01b0381166000908152603360209081526040808320547f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0590925282205480821161384a5760006113d8565b6113d88183615a36565b6097546001600160a01b0316331480159061389e57503360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad09602052604090205460ff16155b156138bc576040516305d1403760e01b815260040160405180910390fd5b8460ff166000036138f9576040517faeb47d8100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460ff1660020361390e5761390e8484613fab565b6040517f7b91b75a00000000000000000000000000000000000000000000000000000000815260cb90637b91b75a90613957908a908a908a908a908a908a908a90600401615926565b600060405180830381600087803b15801561397157600080fd5b505af1158015613985573d6000803e3d6000fd5b50506040516001600160a01b038a1692507fd9436ef9ce00ffeabc5da2489701502d3bd1a5ed7b254a1981fc5ffef9828e119150600090a250505050505050565b6139ce613c42565b600260fb5403613a205760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016111cd565b600260fb556000819003613a4757604051630cb65c7760e21b815260040160405180910390fd5b60cb63ff3e9aa933600080516020615be78339815191525460405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b03909216600483015267ffffffffffffffff16602482015260448101849052606401600060405180830381600087803b158015613ace57600080fd5b505af1158015613ae2573d6000803e3d6000fd5b50506040518381523392507f2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543915060200160405180910390a250600160fb55565b613b2b613bb0565b6001600160a01b038116613ba75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016111cd565b6118ea8161462d565b6097546001600160a01b03163314611b225760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016111cd565b600080613c1660355490565b90506113d8613c23610fb9565b613c2e9060016159f3565b613c398360016159f3565b86919086614984565b60c95460ff1615611b225760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016111cd565b60003361192c818585614207565b600080613caf60355490565b90506113d8613cbf8260016159f3565b613cc7610fb9565b613c399060016159f3565b3360009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad066020526040812054600080516020615be783398151915290608011613d4a576040517f8bf0a57900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005820160209081526040808320546033909252822054613d719190615a36565b905080861115613dad576040517f5090213f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600090815260058301602052604081208054889290613dce9084906159f3565b90915550503360009081526006830160205260408120805491613df083615a1c565b9091555050600382018054906000613e0783615a1c565b9190505592506000826001015442613e1f91906159f3565b6040805160e0810182523381526001600160a01b03881660208201529081018990526060810188905267ffffffffffffffff4281166080830152821660a082015290915060c0810160019052600085815260048086016020908152604092839020845181546001600160a01b0391821673ffffffffffffffffffffffffffffffffffffffff19918216178355928601516001830180549190921693169290921790915591830151600280840191909155606084015160038401556080840151918301805460a086015167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009092169416939093179290921780835560c085015192917fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff90911690700100000000000000000000000000000000908490811115613f8057613f80615a49565b02179055509050505050509392505050565b600033613fa08582856149df565b612dd7858585614a71565b60005b8181101561405a57613fe5838383818110613fcb57613fcb615a06565b9050602002016020810190613fe09190615aaf565b614c93565b61404857828282818110613ffb57613ffb615a06565b90506020020160208101906140109190615aaf565b6040517f840aa02e00000000000000000000000000000000000000000000000000000000815260ff90911660048201526024016111cd565b8061405281615a1c565b915050613fae565b505050565b6118ea613bb0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561409a5761405a83614d00565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156140f4575060408051601f3d908101601f191682019092526140f1918101906158d3565b60015b6141665760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201527f6f6e206973206e6f74205555505300000000000000000000000000000000000060648201526084016111cd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146141fb5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f7860448201527f6961626c6555554944000000000000000000000000000000000000000000000060648201526084016111cd565b5061405a838383614dcb565b6001600160a01b0383166142825760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016111cd565b6001600160a01b0382166142fe5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016111cd565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b614367614df0565b60c9805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03821661442d5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016111cd565b61443982600083614e42565b6001600160a01b038216600090815260336020526040902054818110156144c85760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016111cd565b6001600160a01b03831660009081526033602052604081208383039055603580548492906144f7908490615a36565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b6001600160a01b0382166145985760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016111cd565b6145a460008383614e42565b80603560008282546145b691906159f3565b90915550506001600160a01b038216600090815260336020526040812080548392906145e39084906159f3565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b609780546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b614694613c42565b60c9805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586143943390565b60003361192c818585614a71565b600054610100900460ff166147425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b6118ea81614f41565b600054610100900460ff166147b65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b611d118282614fdb565b600054610100900460ff1661482b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b611b2261505f565b600054610100900460ff1661489e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b611b226150d3565b600054610100900460ff166149115760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b611b2261514a565b600054610100900460ff16611b225760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b6000806149928686866151bc565b905060018360028111156149a8576149a8615a49565b1480156149c55750600084806149c0576149c0615aca565b868809115b15613484576149d56001826159f3565b9695505050505050565b6001600160a01b038381166000908152603460209081526040808320938616835292905220546000198114614a6b5781811015614a5e5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016111cd565b614a6b8484848403614207565b50505050565b6001600160a01b038316614aed5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016111cd565b6001600160a01b038216614b695760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016111cd565b614b74838383614e42565b6001600160a01b03831660009081526033602052604090205481811015614c035760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016111cd565b6001600160a01b03808516600090815260336020526040808220858503905591851681529081208054849290614c3a9084906159f3565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051614c8691815260200190565b60405180910390a3614a6b565b60008160ff1660411480614caa57508160ff166042145b80614cb857508160ff166043145b80614cc657508160ff166045145b80614cd457508160ff166048145b80614ce257508160ff166049145b80614cf057508160ff16604a145b8061114757505060ff16604b1490565b6001600160a01b0381163b614d7d5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e74726163740000000000000000000000000000000000000060648201526084016111cd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b614dd48361526b565b600082511180614de15750805b1561405a57614a6b83836152ab565b60c95460ff16611b225760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016111cd565b6001600160a01b03831615801590614e6257506001600160a01b03821615155b15614e99576040517fc4d9ec0a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0383161561405a576001600160a01b03831660009081527f391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad0560209081526040808320546033909252822054600080516020615be78339815191529291614f0591615a36565b9050808311156137f2576040517f5090213f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff16614fac5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b6065805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b600054610100900460ff166150465760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b60366150528382615b26565b50603761405a8282615b26565b600054610100900460ff166150ca5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b611b223361462d565b600054610100900460ff1661513e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b60c9805460ff19169055565b600054610100900460ff166151b55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016111cd565b600160fb55565b60008080600019858709858702925082811083820303915050806000036151f6578382816151ec576151ec615aca565b0492505050611161565b80841161520257600080fd5b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b61527481614d00565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b61532a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016111cd565b600080846001600160a01b0316846040516153459190615a76565b600060405180830381855af49150503d8060008114615380576040519150601f19603f3d011682016040523d82523d6000602084013e615385565b606091505b50915091506134848282604051806060016040528060278152602001615c0760279139606083156153b7575081611161565b8251156153c75782518084602001fd5b8160405162461bcd60e51b81526004016111cd9190615467565b80151581146118ea57600080fd5b60006020828403121561540157600080fd5b8135611161816153e1565b80356001600160a01b038116811461542357600080fd5b919050565b60006020828403121561543a57600080fd5b6111618261540c565b60005b8381101561545e578181015183820152602001615446565b50506000910152565b6020815260008251806020840152615486816040850160208701615443565b601f01601f19169190910160400192915050565b6000602082840312156154ac57600080fd5b5035919050565b600080604083850312156154c657600080fd5b6154cf8361540c565b946020939093013593505050565b600080604083850312156154f057600080fd5b823591506155006020840161540c565b90509250929050565b60008060006060848603121561551e57600080fd5b6155278461540c565b92506155356020850161540c565b9150604084013590509250925092565b803567ffffffffffffffff8116811461542357600080fd5b803560ff8116811461542357600080fd5b60008083601f84011261558057600080fd5b50813567ffffffffffffffff81111561559857600080fd5b6020830191508360208260051b85010111156155b357600080fd5b9250929050565b600080600080600080600060a0888a0312156155d557600080fd5b6155de8861540c565b96506155ec60208901615545565b95506155fa6040890161555d565b9450606088013567ffffffffffffffff8082111561561757600080fd5b6156238b838c0161556e565b909650945060808a013591508082111561563c57600080fd5b506156498a828b0161556e565b989b979a50959850939692959293505050565b6000806000806040858703121561567257600080fd5b843567ffffffffffffffff8082111561568a57600080fd5b6156968883890161556e565b909650945060208701359150808211156156af57600080fd5b506156bc8782880161556e565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff808411156156f9576156f96156c8565b604051601f8501601f19908116603f01168101908282118183101715615721576157216156c8565b8160405280935085815286868601111561573a57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561576757600080fd5b6157708361540c565b9150602083013567ffffffffffffffff81111561578c57600080fd5b8301601f8101851361579d57600080fd5b6157ac858235602084016156de565b9150509250929050565b6000602082840312156157c857600080fd5b61116182615545565b6000806000606084860312156157e657600080fd5b833592506157f66020850161540c565b91506158046040850161540c565b90509250925092565b600082601f83011261581e57600080fd5b611161838335602085016156de565b6000806000806080858703121561584357600080fd5b843567ffffffffffffffff8082111561585b57600080fd5b6158678883890161580d565b9550602087013591508082111561587d57600080fd5b5061588a8782880161580d565b93505061589960408601615545565b9396929550929360600135925050565b600080604083850312156158bc57600080fd5b6158c58361540c565b91506155006020840161540c565b6000602082840312156158e557600080fd5b5051919050565b600181811c9082168061590057607f821691505b60208210810361592057634e487b7160e01b600052602260045260246000fd5b50919050565b600060a082016001600160a01b038a168352602067ffffffffffffffff808b168286015260ff808b16604087015260a060608701528389855260c0870190508a945060005b8a811015615990578261597d8761555d565b168252948401949084019060010161596b565b5086810360808801528781528301935087905060005b878110156159cb57826159b883615545565b16855293830193908301906001016159a6565b50929c9b505050505050505050505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611147576111476159dd565b634e487b7160e01b600052603260045260246000fd5b60006000198203615a2f57615a2f6159dd565b5060010190565b81810381811115611147576111476159dd565b634e487b7160e01b600052602160045260246000fd5b600081615a6e57615a6e6159dd565b506000190190565b60008251615a88818460208701615443565b9190910192915050565b600060208284031215615aa457600080fd5b8151611161816153e1565b600060208284031215615ac157600080fd5b6111618261555d565b634e487b7160e01b600052601260045260246000fd5b601f82111561405a57600081815260208120601f850160051c81016020861015615b075750805b601f850160051c820191505b81811015611aaa57828155600101615b13565b815167ffffffffffffffff811115615b4057615b406156c8565b615b5481615b4e84546158ec565b84615ae0565b602080601f831160018114615b895760008415615b715750858301515b600019600386901b1c1916600185901b178555611aaa565b600085815260208120601f198616915b82811015615bb857888601518255948401946001909101908401615b99565b5085821015615bd65787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe391c7ba7c740680e30817156af3441a6d715b70335ccfe9b8f16f839401dad00416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122051ae3b0bee86cc424e2f33f98c360f0774808b09b53c50153654a82b7c0a362b64736f6c63430008110033",
}

// PerpVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use PerpVaultMetaData.ABI instead.
var PerpVaultABI = PerpVaultMetaData.ABI

// PerpVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PerpVaultMetaData.Bin instead.
var PerpVaultBin = PerpVaultMetaData.Bin

// DeployPerpVault deploys a new Ethereum contract, binding an instance of PerpVault to it.
func DeployPerpVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PerpVault, error) {
	parsed, err := PerpVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PerpVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PerpVault{PerpVaultCaller: PerpVaultCaller{contract: contract}, PerpVaultTransactor: PerpVaultTransactor{contract: contract}, PerpVaultFilterer: PerpVaultFilterer{contract: contract}}, nil
}

// PerpVault is an auto generated Go binding around an Ethereum contract.
type PerpVault struct {
	PerpVaultCaller     // Read-only binding to the contract
	PerpVaultTransactor // Write-only binding to the contract
	PerpVaultFilterer   // Log filterer for contract events
}

// PerpVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerpVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerpVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerpVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerpVaultSession struct {
	Contract     *PerpVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerpVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerpVaultCallerSession struct {
	Contract *PerpVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PerpVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerpVaultTransactorSession struct {
	Contract     *PerpVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PerpVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerpVaultRaw struct {
	Contract *PerpVault // Generic contract binding to access the raw methods on
}

// PerpVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerpVaultCallerRaw struct {
	Contract *PerpVaultCaller // Generic read-only contract binding to access the raw methods on
}

// PerpVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerpVaultTransactorRaw struct {
	Contract *PerpVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerpVault creates a new instance of PerpVault, bound to a specific deployed contract.
func NewPerpVault(address common.Address, backend bind.ContractBackend) (*PerpVault, error) {
	contract, err := bindPerpVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerpVault{PerpVaultCaller: PerpVaultCaller{contract: contract}, PerpVaultTransactor: PerpVaultTransactor{contract: contract}, PerpVaultFilterer: PerpVaultFilterer{contract: contract}}, nil
}

// NewPerpVaultCaller creates a new read-only instance of PerpVault, bound to a specific deployed contract.
func NewPerpVaultCaller(address common.Address, caller bind.ContractCaller) (*PerpVaultCaller, error) {
	contract, err := bindPerpVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerpVaultCaller{contract: contract}, nil
}

// NewPerpVaultTransactor creates a new write-only instance of PerpVault, bound to a specific deployed contract.
func NewPerpVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*PerpVaultTransactor, error) {
	contract, err := bindPerpVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerpVaultTransactor{contract: contract}, nil
}

// NewPerpVaultFilterer creates a new log filterer instance of PerpVault, bound to a specific deployed contract.
func NewPerpVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*PerpVaultFilterer, error) {
	contract, err := bindPerpVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerpVaultFilterer{contract: contract}, nil
}

// bindPerpVault binds a generic wrapper to an already deployed contract.
func bindPerpVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PerpVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpVault *PerpVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpVault.Contract.PerpVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpVault *PerpVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.Contract.PerpVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpVault *PerpVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpVault.Contract.PerpVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpVault *PerpVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpVault *PerpVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpVault *PerpVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTLOCKUPPERIOD is a free data retrieval call binding the contract method 0x27f69b93.
//
// Solidity: function DEFAULT_LOCKUP_PERIOD() view returns(uint256)
func (_PerpVault *PerpVaultCaller) DEFAULTLOCKUPPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "DEFAULT_LOCKUP_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTLOCKUPPERIOD is a free data retrieval call binding the contract method 0x27f69b93.
//
// Solidity: function DEFAULT_LOCKUP_PERIOD() view returns(uint256)
func (_PerpVault *PerpVaultSession) DEFAULTLOCKUPPERIOD() (*big.Int, error) {
	return _PerpVault.Contract.DEFAULTLOCKUPPERIOD(&_PerpVault.CallOpts)
}

// DEFAULTLOCKUPPERIOD is a free data retrieval call binding the contract method 0x27f69b93.
//
// Solidity: function DEFAULT_LOCKUP_PERIOD() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) DEFAULTLOCKUPPERIOD() (*big.Int, error) {
	return _PerpVault.Contract.DEFAULTLOCKUPPERIOD(&_PerpVault.CallOpts)
}

// MAXLOCKUP is a free data retrieval call binding the contract method 0xea775017.
//
// Solidity: function MAX_LOCKUP() view returns(uint256)
func (_PerpVault *PerpVaultCaller) MAXLOCKUP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "MAX_LOCKUP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKUP is a free data retrieval call binding the contract method 0xea775017.
//
// Solidity: function MAX_LOCKUP() view returns(uint256)
func (_PerpVault *PerpVaultSession) MAXLOCKUP() (*big.Int, error) {
	return _PerpVault.Contract.MAXLOCKUP(&_PerpVault.CallOpts)
}

// MAXLOCKUP is a free data retrieval call binding the contract method 0xea775017.
//
// Solidity: function MAX_LOCKUP() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MAXLOCKUP() (*big.Int, error) {
	return _PerpVault.Contract.MAXLOCKUP(&_PerpVault.CallOpts)
}

// MAXPENDINGPEROWNER is a free data retrieval call binding the contract method 0x375126b6.
//
// Solidity: function MAX_PENDING_PER_OWNER() view returns(uint256)
func (_PerpVault *PerpVaultCaller) MAXPENDINGPEROWNER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "MAX_PENDING_PER_OWNER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENDINGPEROWNER is a free data retrieval call binding the contract method 0x375126b6.
//
// Solidity: function MAX_PENDING_PER_OWNER() view returns(uint256)
func (_PerpVault *PerpVaultSession) MAXPENDINGPEROWNER() (*big.Int, error) {
	return _PerpVault.Contract.MAXPENDINGPEROWNER(&_PerpVault.CallOpts)
}

// MAXPENDINGPEROWNER is a free data retrieval call binding the contract method 0x375126b6.
//
// Solidity: function MAX_PENDING_PER_OWNER() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MAXPENDINGPEROWNER() (*big.Int, error) {
	return _PerpVault.Contract.MAXPENDINGPEROWNER(&_PerpVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerpVault *PerpVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerpVault *PerpVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Allowance(&_PerpVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Allowance(&_PerpVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PerpVault *PerpVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PerpVault *PerpVaultSession) Asset() (common.Address, error) {
	return _PerpVault.Contract.Asset(&_PerpVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PerpVault *PerpVaultCallerSession) Asset() (common.Address, error) {
	return _PerpVault.Contract.Asset(&_PerpVault.CallOpts)
}

// BackstopEnabledAt is a free data retrieval call binding the contract method 0xd96c1284.
//
// Solidity: function backstopEnabledAt() view returns(uint256)
func (_PerpVault *PerpVaultCaller) BackstopEnabledAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "backstopEnabledAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BackstopEnabledAt is a free data retrieval call binding the contract method 0xd96c1284.
//
// Solidity: function backstopEnabledAt() view returns(uint256)
func (_PerpVault *PerpVaultSession) BackstopEnabledAt() (*big.Int, error) {
	return _PerpVault.Contract.BackstopEnabledAt(&_PerpVault.CallOpts)
}

// BackstopEnabledAt is a free data retrieval call binding the contract method 0xd96c1284.
//
// Solidity: function backstopEnabledAt() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) BackstopEnabledAt() (*big.Int, error) {
	return _PerpVault.Contract.BackstopEnabledAt(&_PerpVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerpVault *PerpVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerpVault *PerpVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PerpVault.Contract.BalanceOf(&_PerpVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PerpVault.Contract.BalanceOf(&_PerpVault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_PerpVault *PerpVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_PerpVault *PerpVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.ConvertToAssets(&_PerpVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_PerpVault *PerpVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.ConvertToAssets(&_PerpVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_PerpVault *PerpVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_PerpVault *PerpVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.ConvertToShares(&_PerpVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_PerpVault *PerpVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.ConvertToShares(&_PerpVault.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PerpVault *PerpVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PerpVault *PerpVaultSession) Decimals() (uint8, error) {
	return _PerpVault.Contract.Decimals(&_PerpVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PerpVault *PerpVaultCallerSession) Decimals() (uint8, error) {
	return _PerpVault.Contract.Decimals(&_PerpVault.CallOpts)
}

// EscrowedSharesOf is a free data retrieval call binding the contract method 0x06161f4c.
//
// Solidity: function escrowedSharesOf(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCaller) EscrowedSharesOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "escrowedSharesOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowedSharesOf is a free data retrieval call binding the contract method 0x06161f4c.
//
// Solidity: function escrowedSharesOf(address owner) view returns(uint256)
func (_PerpVault *PerpVaultSession) EscrowedSharesOf(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.EscrowedSharesOf(&_PerpVault.CallOpts, owner)
}

// EscrowedSharesOf is a free data retrieval call binding the contract method 0x06161f4c.
//
// Solidity: function escrowedSharesOf(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) EscrowedSharesOf(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.EscrowedSharesOf(&_PerpVault.CallOpts, owner)
}

// GetDepositCap is a free data retrieval call binding the contract method 0x01a598da.
//
// Solidity: function getDepositCap() view returns(uint256)
func (_PerpVault *PerpVaultCaller) GetDepositCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getDepositCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositCap is a free data retrieval call binding the contract method 0x01a598da.
//
// Solidity: function getDepositCap() view returns(uint256)
func (_PerpVault *PerpVaultSession) GetDepositCap() (*big.Int, error) {
	return _PerpVault.Contract.GetDepositCap(&_PerpVault.CallOpts)
}

// GetDepositCap is a free data retrieval call binding the contract method 0x01a598da.
//
// Solidity: function getDepositCap() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) GetDepositCap() (*big.Int, error) {
	return _PerpVault.Contract.GetDepositCap(&_PerpVault.CallOpts)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x93ad3e03.
//
// Solidity: function getMarketCap(uint64 marketId) view returns(uint256)
func (_PerpVault *PerpVaultCaller) GetMarketCap(opts *bind.CallOpts, marketId uint64) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getMarketCap", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCap is a free data retrieval call binding the contract method 0x93ad3e03.
//
// Solidity: function getMarketCap(uint64 marketId) view returns(uint256)
func (_PerpVault *PerpVaultSession) GetMarketCap(marketId uint64) (*big.Int, error) {
	return _PerpVault.Contract.GetMarketCap(&_PerpVault.CallOpts, marketId)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x93ad3e03.
//
// Solidity: function getMarketCap(uint64 marketId) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) GetMarketCap(marketId uint64) (*big.Int, error) {
	return _PerpVault.Contract.GetMarketCap(&_PerpVault.CallOpts, marketId)
}

// GetMaxMarginUtilization is a free data retrieval call binding the contract method 0x0d2431bf.
//
// Solidity: function getMaxMarginUtilization() view returns(uint256)
func (_PerpVault *PerpVaultCaller) GetMaxMarginUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getMaxMarginUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMarginUtilization is a free data retrieval call binding the contract method 0x0d2431bf.
//
// Solidity: function getMaxMarginUtilization() view returns(uint256)
func (_PerpVault *PerpVaultSession) GetMaxMarginUtilization() (*big.Int, error) {
	return _PerpVault.Contract.GetMaxMarginUtilization(&_PerpVault.CallOpts)
}

// GetMaxMarginUtilization is a free data retrieval call binding the contract method 0x0d2431bf.
//
// Solidity: function getMaxMarginUtilization() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) GetMaxMarginUtilization() (*big.Int, error) {
	return _PerpVault.Contract.GetMaxMarginUtilization(&_PerpVault.CallOpts)
}

// GetMinDeposit is a free data retrieval call binding the contract method 0x0eaad3f1.
//
// Solidity: function getMinDeposit() view returns(uint256)
func (_PerpVault *PerpVaultCaller) GetMinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getMinDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDeposit is a free data retrieval call binding the contract method 0x0eaad3f1.
//
// Solidity: function getMinDeposit() view returns(uint256)
func (_PerpVault *PerpVaultSession) GetMinDeposit() (*big.Int, error) {
	return _PerpVault.Contract.GetMinDeposit(&_PerpVault.CallOpts)
}

// GetMinDeposit is a free data retrieval call binding the contract method 0x0eaad3f1.
//
// Solidity: function getMinDeposit() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) GetMinDeposit() (*big.Int, error) {
	return _PerpVault.Contract.GetMinDeposit(&_PerpVault.CallOpts)
}

// GetMinWithdraw is a free data retrieval call binding the contract method 0xd685da8d.
//
// Solidity: function getMinWithdraw() view returns(uint256)
func (_PerpVault *PerpVaultCaller) GetMinWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getMinWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinWithdraw is a free data retrieval call binding the contract method 0xd685da8d.
//
// Solidity: function getMinWithdraw() view returns(uint256)
func (_PerpVault *PerpVaultSession) GetMinWithdraw() (*big.Int, error) {
	return _PerpVault.Contract.GetMinWithdraw(&_PerpVault.CallOpts)
}

// GetMinWithdraw is a free data retrieval call binding the contract method 0xd685da8d.
//
// Solidity: function getMinWithdraw() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) GetMinWithdraw() (*big.Int, error) {
	return _PerpVault.Contract.GetMinWithdraw(&_PerpVault.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns(address owner, address receiver, uint256 shares, uint256 assets, uint64 requestedAt, uint64 executableAt, uint8 status)
func (_PerpVault *PerpVaultCaller) GetRequest(opts *bind.CallOpts, requestId *big.Int) (struct {
	Owner        common.Address
	Receiver     common.Address
	Shares       *big.Int
	Assets       *big.Int
	RequestedAt  uint64
	ExecutableAt uint64
	Status       uint8
}, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "getRequest", requestId)

	outstruct := new(struct {
		Owner        common.Address
		Receiver     common.Address
		Shares       *big.Int
		Assets       *big.Int
		RequestedAt  uint64
		ExecutableAt uint64
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Shares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Assets = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RequestedAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.ExecutableAt = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns(address owner, address receiver, uint256 shares, uint256 assets, uint64 requestedAt, uint64 executableAt, uint8 status)
func (_PerpVault *PerpVaultSession) GetRequest(requestId *big.Int) (struct {
	Owner        common.Address
	Receiver     common.Address
	Shares       *big.Int
	Assets       *big.Int
	RequestedAt  uint64
	ExecutableAt uint64
	Status       uint8
}, error) {
	return _PerpVault.Contract.GetRequest(&_PerpVault.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns(address owner, address receiver, uint256 shares, uint256 assets, uint64 requestedAt, uint64 executableAt, uint8 status)
func (_PerpVault *PerpVaultCallerSession) GetRequest(requestId *big.Int) (struct {
	Owner        common.Address
	Receiver     common.Address
	Shares       *big.Int
	Assets       *big.Int
	RequestedAt  uint64
	ExecutableAt uint64
	Status       uint8
}, error) {
	return _PerpVault.Contract.GetRequest(&_PerpVault.CallOpts, requestId)
}

// IsBackstopEnabled is a free data retrieval call binding the contract method 0x4d93b6aa.
//
// Solidity: function isBackstopEnabled() view returns(bool)
func (_PerpVault *PerpVaultCaller) IsBackstopEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "isBackstopEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBackstopEnabled is a free data retrieval call binding the contract method 0x4d93b6aa.
//
// Solidity: function isBackstopEnabled() view returns(bool)
func (_PerpVault *PerpVaultSession) IsBackstopEnabled() (bool, error) {
	return _PerpVault.Contract.IsBackstopEnabled(&_PerpVault.CallOpts)
}

// IsBackstopEnabled is a free data retrieval call binding the contract method 0x4d93b6aa.
//
// Solidity: function isBackstopEnabled() view returns(bool)
func (_PerpVault *PerpVaultCallerSession) IsBackstopEnabled() (bool, error) {
	return _PerpVault.Contract.IsBackstopEnabled(&_PerpVault.CallOpts)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_PerpVault *PerpVaultCaller) IsManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "isManager", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_PerpVault *PerpVaultSession) IsManager(account common.Address) (bool, error) {
	return _PerpVault.Contract.IsManager(&_PerpVault.CallOpts, account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_PerpVault *PerpVaultCallerSession) IsManager(account common.Address) (bool, error) {
	return _PerpVault.Contract.IsManager(&_PerpVault.CallOpts, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PerpVault *PerpVaultCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PerpVault *PerpVaultSession) IsPaused() (bool, error) {
	return _PerpVault.Contract.IsPaused(&_PerpVault.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PerpVault *PerpVaultCallerSession) IsPaused() (bool, error) {
	return _PerpVault.Contract.IsPaused(&_PerpVault.CallOpts)
}

// IsVault is a free data retrieval call binding the contract method 0xe81cc3cc.
//
// Solidity: function isVault() view returns(bool)
func (_PerpVault *PerpVaultCaller) IsVault(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "isVault")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVault is a free data retrieval call binding the contract method 0xe81cc3cc.
//
// Solidity: function isVault() view returns(bool)
func (_PerpVault *PerpVaultSession) IsVault() (bool, error) {
	return _PerpVault.Contract.IsVault(&_PerpVault.CallOpts)
}

// IsVault is a free data retrieval call binding the contract method 0xe81cc3cc.
//
// Solidity: function isVault() view returns(bool)
func (_PerpVault *PerpVaultCallerSession) IsVault() (bool, error) {
	return _PerpVault.Contract.IsVault(&_PerpVault.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_PerpVault *PerpVaultCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_PerpVault *PerpVaultSession) IsWhitelisted(account common.Address) (bool, error) {
	return _PerpVault.Contract.IsWhitelisted(&_PerpVault.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_PerpVault *PerpVaultCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _PerpVault.Contract.IsWhitelisted(&_PerpVault.CallOpts, account)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_PerpVault *PerpVaultCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "lockupPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_PerpVault *PerpVaultSession) LockupPeriod() (*big.Int, error) {
	return _PerpVault.Contract.LockupPeriod(&_PerpVault.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) LockupPeriod() (*big.Int, error) {
	return _PerpVault.Contract.LockupPeriod(&_PerpVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PerpVault *PerpVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PerpVault *PerpVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxDeposit(&_PerpVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxDeposit(&_PerpVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_PerpVault *PerpVaultCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_PerpVault *PerpVaultSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxMint(&_PerpVault.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxMint(&_PerpVault.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_PerpVault *PerpVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxRedeem(&_PerpVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxRedeem(&_PerpVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_PerpVault *PerpVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxWithdraw(&_PerpVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _PerpVault.Contract.MaxWithdraw(&_PerpVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerpVault *PerpVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerpVault *PerpVaultSession) Name() (string, error) {
	return _PerpVault.Contract.Name(&_PerpVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerpVault *PerpVaultCallerSession) Name() (string, error) {
	return _PerpVault.Contract.Name(&_PerpVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpVault *PerpVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpVault *PerpVaultSession) Owner() (common.Address, error) {
	return _PerpVault.Contract.Owner(&_PerpVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpVault *PerpVaultCallerSession) Owner() (common.Address, error) {
	return _PerpVault.Contract.Owner(&_PerpVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PerpVault *PerpVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PerpVault *PerpVaultSession) Paused() (bool, error) {
	return _PerpVault.Contract.Paused(&_PerpVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PerpVault *PerpVaultCallerSession) Paused() (bool, error) {
	return _PerpVault.Contract.Paused(&_PerpVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewDeposit(&_PerpVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewDeposit(&_PerpVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewMint(&_PerpVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewMint(&_PerpVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewRedeem(&_PerpVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewRedeem(&_PerpVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewWithdraw(&_PerpVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _PerpVault.Contract.PreviewWithdraw(&_PerpVault.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PerpVault *PerpVaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PerpVault *PerpVaultSession) ProxiableUUID() ([32]byte, error) {
	return _PerpVault.Contract.ProxiableUUID(&_PerpVault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PerpVault *PerpVaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PerpVault.Contract.ProxiableUUID(&_PerpVault.CallOpts)
}

// Redeem is a free data retrieval call binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultCaller) Redeem(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "redeem", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeem is a free data retrieval call binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultSession) Redeem(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Redeem(&_PerpVault.CallOpts, arg0, arg1, arg2)
}

// Redeem is a free data retrieval call binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultCallerSession) Redeem(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Redeem(&_PerpVault.CallOpts, arg0, arg1, arg2)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerpVault *PerpVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerpVault *PerpVaultSession) Symbol() (string, error) {
	return _PerpVault.Contract.Symbol(&_PerpVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerpVault *PerpVaultCallerSession) Symbol() (string, error) {
	return _PerpVault.Contract.Symbol(&_PerpVault.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint64)
func (_PerpVault *PerpVaultCaller) TokenId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint64)
func (_PerpVault *PerpVaultSession) TokenId() (uint64, error) {
	return _PerpVault.Contract.TokenId(&_PerpVault.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint64)
func (_PerpVault *PerpVaultCallerSession) TokenId() (uint64, error) {
	return _PerpVault.Contract.TokenId(&_PerpVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PerpVault *PerpVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PerpVault *PerpVaultSession) TotalAssets() (*big.Int, error) {
	return _PerpVault.Contract.TotalAssets(&_PerpVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _PerpVault.Contract.TotalAssets(&_PerpVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerpVault *PerpVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerpVault *PerpVaultSession) TotalSupply() (*big.Int, error) {
	return _PerpVault.Contract.TotalSupply(&_PerpVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerpVault *PerpVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _PerpVault.Contract.TotalSupply(&_PerpVault.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_PerpVault *PerpVaultCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_PerpVault *PerpVaultSession) WhitelistEnabled() (bool, error) {
	return _PerpVault.Contract.WhitelistEnabled(&_PerpVault.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_PerpVault *PerpVaultCallerSession) WhitelistEnabled() (bool, error) {
	return _PerpVault.Contract.WhitelistEnabled(&_PerpVault.CallOpts)
}

// Withdraw is a free data retrieval call binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultCaller) Withdraw(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerpVault.contract.Call(opts, &out, "withdraw", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdraw is a free data retrieval call binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultSession) Withdraw(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Withdraw(&_PerpVault.CallOpts, arg0, arg1, arg2)
}

// Withdraw is a free data retrieval call binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 , address , address ) pure returns(uint256)
func (_PerpVault *PerpVaultCallerSession) Withdraw(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _PerpVault.Contract.Withdraw(&_PerpVault.CallOpts, arg0, arg1, arg2)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x3507e3cb.
//
// Solidity: function addExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultTransactor) AddExecutor(opts *bind.TransactOpts, executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "addExecutor", executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x3507e3cb.
//
// Solidity: function addExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultSession) AddExecutor(executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.Contract.AddExecutor(&_PerpVault.TransactOpts, executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x3507e3cb.
//
// Solidity: function addExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultTransactorSession) AddExecutor(executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.Contract.AddExecutor(&_PerpVault.TransactOpts, executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_PerpVault *PerpVaultTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_PerpVault *PerpVaultSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.AddManager(&_PerpVault.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_PerpVault *PerpVaultTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.AddManager(&_PerpVault.TransactOpts, manager)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Approve(&_PerpVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Approve(&_PerpVault.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerpVault *PerpVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerpVault *PerpVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.DecreaseAllowance(&_PerpVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerpVault *PerpVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.DecreaseAllowance(&_PerpVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_PerpVault *PerpVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_PerpVault *PerpVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.Deposit(&_PerpVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_PerpVault *PerpVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.Deposit(&_PerpVault.TransactOpts, assets, receiver)
}

// DisableBackstop is a paid mutator transaction binding the contract method 0x5ef12042.
//
// Solidity: function disableBackstop() returns()
func (_PerpVault *PerpVaultTransactor) DisableBackstop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "disableBackstop")
}

// DisableBackstop is a paid mutator transaction binding the contract method 0x5ef12042.
//
// Solidity: function disableBackstop() returns()
func (_PerpVault *PerpVaultSession) DisableBackstop() (*types.Transaction, error) {
	return _PerpVault.Contract.DisableBackstop(&_PerpVault.TransactOpts)
}

// DisableBackstop is a paid mutator transaction binding the contract method 0x5ef12042.
//
// Solidity: function disableBackstop() returns()
func (_PerpVault *PerpVaultTransactorSession) DisableBackstop() (*types.Transaction, error) {
	return _PerpVault.Contract.DisableBackstop(&_PerpVault.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 assets) returns()
func (_PerpVault *PerpVaultTransactor) Donate(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "donate", assets)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 assets) returns()
func (_PerpVault *PerpVaultSession) Donate(assets *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Donate(&_PerpVault.TransactOpts, assets)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 assets) returns()
func (_PerpVault *PerpVaultTransactorSession) Donate(assets *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Donate(&_PerpVault.TransactOpts, assets)
}

// EnableBackstop is a paid mutator transaction binding the contract method 0x595f427e.
//
// Solidity: function enableBackstop() returns()
func (_PerpVault *PerpVaultTransactor) EnableBackstop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "enableBackstop")
}

// EnableBackstop is a paid mutator transaction binding the contract method 0x595f427e.
//
// Solidity: function enableBackstop() returns()
func (_PerpVault *PerpVaultSession) EnableBackstop() (*types.Transaction, error) {
	return _PerpVault.Contract.EnableBackstop(&_PerpVault.TransactOpts)
}

// EnableBackstop is a paid mutator transaction binding the contract method 0x595f427e.
//
// Solidity: function enableBackstop() returns()
func (_PerpVault *PerpVaultTransactorSession) EnableBackstop() (*types.Transaction, error) {
	return _PerpVault.Contract.EnableBackstop(&_PerpVault.TransactOpts)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0x6e417c69.
//
// Solidity: function executeRequest(uint256 requestId) returns(uint256 assets)
func (_PerpVault *PerpVaultTransactor) ExecuteRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "executeRequest", requestId)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0x6e417c69.
//
// Solidity: function executeRequest(uint256 requestId) returns(uint256 assets)
func (_PerpVault *PerpVaultSession) ExecuteRequest(requestId *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.ExecuteRequest(&_PerpVault.TransactOpts, requestId)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0x6e417c69.
//
// Solidity: function executeRequest(uint256 requestId) returns(uint256 assets)
func (_PerpVault *PerpVaultTransactorSession) ExecuteRequest(requestId *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.ExecuteRequest(&_PerpVault.TransactOpts, requestId)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerpVault *PerpVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerpVault *PerpVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.IncreaseAllowance(&_PerpVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerpVault *PerpVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.IncreaseAllowance(&_PerpVault.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1574855.
//
// Solidity: function initialize(string name_, string symbol_, uint64 tokenId_, uint256 lockupPeriod_) returns()
func (_PerpVault *PerpVaultTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, tokenId_ uint64, lockupPeriod_ *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "initialize", name_, symbol_, tokenId_, lockupPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1574855.
//
// Solidity: function initialize(string name_, string symbol_, uint64 tokenId_, uint256 lockupPeriod_) returns()
func (_PerpVault *PerpVaultSession) Initialize(name_ string, symbol_ string, tokenId_ uint64, lockupPeriod_ *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Initialize(&_PerpVault.TransactOpts, name_, symbol_, tokenId_, lockupPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1574855.
//
// Solidity: function initialize(string name_, string symbol_, uint64 tokenId_, uint256 lockupPeriod_) returns()
func (_PerpVault *PerpVaultTransactorSession) Initialize(name_ string, symbol_ string, tokenId_ uint64, lockupPeriod_ *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Initialize(&_PerpVault.TransactOpts, name_, symbol_, tokenId_, lockupPeriod_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_PerpVault *PerpVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_PerpVault *PerpVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.Mint(&_PerpVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_PerpVault *PerpVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.Mint(&_PerpVault.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PerpVault *PerpVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PerpVault *PerpVaultSession) Pause() (*types.Transaction, error) {
	return _PerpVault.Contract.Pause(&_PerpVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PerpVault *PerpVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _PerpVault.Contract.Pause(&_PerpVault.TransactOpts)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address executor) returns()
func (_PerpVault *PerpVaultTransactor) RemoveExecutor(opts *bind.TransactOpts, executor common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "removeExecutor", executor)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address executor) returns()
func (_PerpVault *PerpVaultSession) RemoveExecutor(executor common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RemoveExecutor(&_PerpVault.TransactOpts, executor)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address executor) returns()
func (_PerpVault *PerpVaultTransactorSession) RemoveExecutor(executor common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RemoveExecutor(&_PerpVault.TransactOpts, executor)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_PerpVault *PerpVaultTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_PerpVault *PerpVaultSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RemoveManager(&_PerpVault.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_PerpVault *PerpVaultTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RemoveManager(&_PerpVault.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpVault *PerpVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpVault *PerpVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _PerpVault.Contract.RenounceOwnership(&_PerpVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpVault *PerpVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PerpVault.Contract.RenounceOwnership(&_PerpVault.TransactOpts)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultTransactor) RequestRedeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "requestRedeem", shares, receiver)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultSession) RequestRedeem(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RequestRedeem(&_PerpVault.TransactOpts, shares, receiver)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultTransactorSession) RequestRedeem(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RequestRedeem(&_PerpVault.TransactOpts, shares, receiver)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultTransactor) RequestWithdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "requestWithdraw", assets, receiver)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultSession) RequestWithdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RequestWithdraw(&_PerpVault.TransactOpts, assets, receiver)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets, address receiver) returns(uint256 requestId)
func (_PerpVault *PerpVaultTransactorSession) RequestWithdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.RequestWithdraw(&_PerpVault.TransactOpts, assets, receiver)
}

// SetDepositCap is a paid mutator transaction binding the contract method 0x86651203.
//
// Solidity: function setDepositCap(uint256 newCap) returns()
func (_PerpVault *PerpVaultTransactor) SetDepositCap(opts *bind.TransactOpts, newCap *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setDepositCap", newCap)
}

// SetDepositCap is a paid mutator transaction binding the contract method 0x86651203.
//
// Solidity: function setDepositCap(uint256 newCap) returns()
func (_PerpVault *PerpVaultSession) SetDepositCap(newCap *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetDepositCap(&_PerpVault.TransactOpts, newCap)
}

// SetDepositCap is a paid mutator transaction binding the contract method 0x86651203.
//
// Solidity: function setDepositCap(uint256 newCap) returns()
func (_PerpVault *PerpVaultTransactorSession) SetDepositCap(newCap *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetDepositCap(&_PerpVault.TransactOpts, newCap)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newPeriod) returns()
func (_PerpVault *PerpVaultTransactor) SetLockupPeriod(opts *bind.TransactOpts, newPeriod *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setLockupPeriod", newPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newPeriod) returns()
func (_PerpVault *PerpVaultSession) SetLockupPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetLockupPeriod(&_PerpVault.TransactOpts, newPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newPeriod) returns()
func (_PerpVault *PerpVaultTransactorSession) SetLockupPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetLockupPeriod(&_PerpVault.TransactOpts, newPeriod)
}

// SetMarketRules is a paid mutator transaction binding the contract method 0x6a7d051d.
//
// Solidity: function setMarketRules(uint64[] marketIds, uint256[] caps) returns()
func (_PerpVault *PerpVaultTransactor) SetMarketRules(opts *bind.TransactOpts, marketIds []uint64, caps []*big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setMarketRules", marketIds, caps)
}

// SetMarketRules is a paid mutator transaction binding the contract method 0x6a7d051d.
//
// Solidity: function setMarketRules(uint64[] marketIds, uint256[] caps) returns()
func (_PerpVault *PerpVaultSession) SetMarketRules(marketIds []uint64, caps []*big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMarketRules(&_PerpVault.TransactOpts, marketIds, caps)
}

// SetMarketRules is a paid mutator transaction binding the contract method 0x6a7d051d.
//
// Solidity: function setMarketRules(uint64[] marketIds, uint256[] caps) returns()
func (_PerpVault *PerpVaultTransactorSession) SetMarketRules(marketIds []uint64, caps []*big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMarketRules(&_PerpVault.TransactOpts, marketIds, caps)
}

// SetMaxMarginUtilization is a paid mutator transaction binding the contract method 0xb4b78d77.
//
// Solidity: function setMaxMarginUtilization(uint256 newBps) returns()
func (_PerpVault *PerpVaultTransactor) SetMaxMarginUtilization(opts *bind.TransactOpts, newBps *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setMaxMarginUtilization", newBps)
}

// SetMaxMarginUtilization is a paid mutator transaction binding the contract method 0xb4b78d77.
//
// Solidity: function setMaxMarginUtilization(uint256 newBps) returns()
func (_PerpVault *PerpVaultSession) SetMaxMarginUtilization(newBps *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMaxMarginUtilization(&_PerpVault.TransactOpts, newBps)
}

// SetMaxMarginUtilization is a paid mutator transaction binding the contract method 0xb4b78d77.
//
// Solidity: function setMaxMarginUtilization(uint256 newBps) returns()
func (_PerpVault *PerpVaultTransactorSession) SetMaxMarginUtilization(newBps *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMaxMarginUtilization(&_PerpVault.TransactOpts, newBps)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newMin) returns()
func (_PerpVault *PerpVaultTransactor) SetMinDeposit(opts *bind.TransactOpts, newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setMinDeposit", newMin)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newMin) returns()
func (_PerpVault *PerpVaultSession) SetMinDeposit(newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMinDeposit(&_PerpVault.TransactOpts, newMin)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newMin) returns()
func (_PerpVault *PerpVaultTransactorSession) SetMinDeposit(newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMinDeposit(&_PerpVault.TransactOpts, newMin)
}

// SetMinWithdraw is a paid mutator transaction binding the contract method 0x35aa134a.
//
// Solidity: function setMinWithdraw(uint256 newMin) returns()
func (_PerpVault *PerpVaultTransactor) SetMinWithdraw(opts *bind.TransactOpts, newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setMinWithdraw", newMin)
}

// SetMinWithdraw is a paid mutator transaction binding the contract method 0x35aa134a.
//
// Solidity: function setMinWithdraw(uint256 newMin) returns()
func (_PerpVault *PerpVaultSession) SetMinWithdraw(newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMinWithdraw(&_PerpVault.TransactOpts, newMin)
}

// SetMinWithdraw is a paid mutator transaction binding the contract method 0x35aa134a.
//
// Solidity: function setMinWithdraw(uint256 newMin) returns()
func (_PerpVault *PerpVaultTransactorSession) SetMinWithdraw(newMin *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.SetMinWithdraw(&_PerpVault.TransactOpts, newMin)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] accounts, bool[] allowed) returns()
func (_PerpVault *PerpVaultTransactor) SetWhitelist(opts *bind.TransactOpts, accounts []common.Address, allowed []bool) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setWhitelist", accounts, allowed)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] accounts, bool[] allowed) returns()
func (_PerpVault *PerpVaultSession) SetWhitelist(accounts []common.Address, allowed []bool) (*types.Transaction, error) {
	return _PerpVault.Contract.SetWhitelist(&_PerpVault.TransactOpts, accounts, allowed)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] accounts, bool[] allowed) returns()
func (_PerpVault *PerpVaultTransactorSession) SetWhitelist(accounts []common.Address, allowed []bool) (*types.Transaction, error) {
	return _PerpVault.Contract.SetWhitelist(&_PerpVault.TransactOpts, accounts, allowed)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool enabled) returns()
func (_PerpVault *PerpVaultTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "setWhitelistEnabled", enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool enabled) returns()
func (_PerpVault *PerpVaultSession) SetWhitelistEnabled(enabled bool) (*types.Transaction, error) {
	return _PerpVault.Contract.SetWhitelistEnabled(&_PerpVault.TransactOpts, enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool enabled) returns()
func (_PerpVault *PerpVaultTransactorSession) SetWhitelistEnabled(enabled bool) (*types.Transaction, error) {
	return _PerpVault.Contract.SetWhitelistEnabled(&_PerpVault.TransactOpts, enabled)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Transfer(&_PerpVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.Transfer(&_PerpVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.TransferFrom(&_PerpVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PerpVault *PerpVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerpVault.Contract.TransferFrom(&_PerpVault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PerpVault *PerpVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PerpVault *PerpVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.TransferOwnership(&_PerpVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PerpVault *PerpVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.TransferOwnership(&_PerpVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PerpVault *PerpVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PerpVault *PerpVaultSession) Unpause() (*types.Transaction, error) {
	return _PerpVault.Contract.Unpause(&_PerpVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PerpVault *PerpVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _PerpVault.Contract.Unpause(&_PerpVault.TransactOpts)
}

// UpdateExecutor is a paid mutator transaction binding the contract method 0xe699734a.
//
// Solidity: function updateExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultTransactor) UpdateExecutor(opts *bind.TransactOpts, executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "updateExecutor", executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// UpdateExecutor is a paid mutator transaction binding the contract method 0xe699734a.
//
// Solidity: function updateExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultSession) UpdateExecutor(executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.Contract.UpdateExecutor(&_PerpVault.TransactOpts, executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// UpdateExecutor is a paid mutator transaction binding the contract method 0xe699734a.
//
// Solidity: function updateExecutor(address executor, uint64 expiresAt, uint8 permType, uint8[] allowedCmds, uint64[] allowedMarkets) returns()
func (_PerpVault *PerpVaultTransactorSession) UpdateExecutor(executor common.Address, expiresAt uint64, permType uint8, allowedCmds []uint8, allowedMarkets []uint64) (*types.Transaction, error) {
	return _PerpVault.Contract.UpdateExecutor(&_PerpVault.TransactOpts, executor, expiresAt, permType, allowedCmds, allowedMarkets)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpVault *PerpVaultTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpVault *PerpVaultSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.UpgradeTo(&_PerpVault.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PerpVault *PerpVaultTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PerpVault.Contract.UpgradeTo(&_PerpVault.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PerpVault *PerpVaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PerpVault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PerpVault *PerpVaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PerpVault.Contract.UpgradeToAndCall(&_PerpVault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PerpVault *PerpVaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PerpVault.Contract.UpgradeToAndCall(&_PerpVault.TransactOpts, newImplementation, data)
}

// PerpVaultAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PerpVault contract.
type PerpVaultAdminChangedIterator struct {
	Event *PerpVaultAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultAdminChanged represents a AdminChanged event raised by the PerpVault contract.
type PerpVaultAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PerpVault *PerpVaultFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PerpVaultAdminChangedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PerpVaultAdminChangedIterator{contract: _PerpVault.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PerpVault *PerpVaultFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PerpVaultAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultAdminChanged)
				if err := _PerpVault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PerpVault *PerpVaultFilterer) ParseAdminChanged(log types.Log) (*PerpVaultAdminChanged, error) {
	event := new(PerpVaultAdminChanged)
	if err := _PerpVault.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PerpVault contract.
type PerpVaultApprovalIterator struct {
	Event *PerpVaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultApproval represents a Approval event raised by the PerpVault contract.
type PerpVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PerpVault *PerpVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PerpVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultApprovalIterator{contract: _PerpVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PerpVault *PerpVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PerpVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultApproval)
				if err := _PerpVault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PerpVault *PerpVaultFilterer) ParseApproval(log types.Log) (*PerpVaultApproval, error) {
	event := new(PerpVaultApproval)
	if err := _PerpVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultBackstopDisabledIterator is returned from FilterBackstopDisabled and is used to iterate over the raw logs and unpacked data for BackstopDisabled events raised by the PerpVault contract.
type PerpVaultBackstopDisabledIterator struct {
	Event *PerpVaultBackstopDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultBackstopDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultBackstopDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultBackstopDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultBackstopDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultBackstopDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultBackstopDisabled represents a BackstopDisabled event raised by the PerpVault contract.
type PerpVaultBackstopDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBackstopDisabled is a free log retrieval operation binding the contract event 0x6e407d4f9fe5e6e3d2a2ad7fd3a065a65774c4b0b22c8bbec1b99476258bb7f0.
//
// Solidity: event BackstopDisabled()
func (_PerpVault *PerpVaultFilterer) FilterBackstopDisabled(opts *bind.FilterOpts) (*PerpVaultBackstopDisabledIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "BackstopDisabled")
	if err != nil {
		return nil, err
	}
	return &PerpVaultBackstopDisabledIterator{contract: _PerpVault.contract, event: "BackstopDisabled", logs: logs, sub: sub}, nil
}

// WatchBackstopDisabled is a free log subscription operation binding the contract event 0x6e407d4f9fe5e6e3d2a2ad7fd3a065a65774c4b0b22c8bbec1b99476258bb7f0.
//
// Solidity: event BackstopDisabled()
func (_PerpVault *PerpVaultFilterer) WatchBackstopDisabled(opts *bind.WatchOpts, sink chan<- *PerpVaultBackstopDisabled) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "BackstopDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultBackstopDisabled)
				if err := _PerpVault.contract.UnpackLog(event, "BackstopDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackstopDisabled is a log parse operation binding the contract event 0x6e407d4f9fe5e6e3d2a2ad7fd3a065a65774c4b0b22c8bbec1b99476258bb7f0.
//
// Solidity: event BackstopDisabled()
func (_PerpVault *PerpVaultFilterer) ParseBackstopDisabled(log types.Log) (*PerpVaultBackstopDisabled, error) {
	event := new(PerpVaultBackstopDisabled)
	if err := _PerpVault.contract.UnpackLog(event, "BackstopDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultBackstopEnabledIterator is returned from FilterBackstopEnabled and is used to iterate over the raw logs and unpacked data for BackstopEnabled events raised by the PerpVault contract.
type PerpVaultBackstopEnabledIterator struct {
	Event *PerpVaultBackstopEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultBackstopEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultBackstopEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultBackstopEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultBackstopEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultBackstopEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultBackstopEnabled represents a BackstopEnabled event raised by the PerpVault contract.
type PerpVaultBackstopEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBackstopEnabled is a free log retrieval operation binding the contract event 0x8c5c1631fda9e60c3d1c47a8effa57d678e664e85a7f797124006c9d93d444c3.
//
// Solidity: event BackstopEnabled()
func (_PerpVault *PerpVaultFilterer) FilterBackstopEnabled(opts *bind.FilterOpts) (*PerpVaultBackstopEnabledIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "BackstopEnabled")
	if err != nil {
		return nil, err
	}
	return &PerpVaultBackstopEnabledIterator{contract: _PerpVault.contract, event: "BackstopEnabled", logs: logs, sub: sub}, nil
}

// WatchBackstopEnabled is a free log subscription operation binding the contract event 0x8c5c1631fda9e60c3d1c47a8effa57d678e664e85a7f797124006c9d93d444c3.
//
// Solidity: event BackstopEnabled()
func (_PerpVault *PerpVaultFilterer) WatchBackstopEnabled(opts *bind.WatchOpts, sink chan<- *PerpVaultBackstopEnabled) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "BackstopEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultBackstopEnabled)
				if err := _PerpVault.contract.UnpackLog(event, "BackstopEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackstopEnabled is a log parse operation binding the contract event 0x8c5c1631fda9e60c3d1c47a8effa57d678e664e85a7f797124006c9d93d444c3.
//
// Solidity: event BackstopEnabled()
func (_PerpVault *PerpVaultFilterer) ParseBackstopEnabled(log types.Log) (*PerpVaultBackstopEnabled, error) {
	event := new(PerpVaultBackstopEnabled)
	if err := _PerpVault.contract.UnpackLog(event, "BackstopEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PerpVault contract.
type PerpVaultBeaconUpgradedIterator struct {
	Event *PerpVaultBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultBeaconUpgraded represents a BeaconUpgraded event raised by the PerpVault contract.
type PerpVaultBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PerpVault *PerpVaultFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PerpVaultBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultBeaconUpgradedIterator{contract: _PerpVault.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PerpVault *PerpVaultFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PerpVaultBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultBeaconUpgraded)
				if err := _PerpVault.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PerpVault *PerpVaultFilterer) ParseBeaconUpgraded(log types.Log) (*PerpVaultBeaconUpgraded, error) {
	event := new(PerpVaultBeaconUpgraded)
	if err := _PerpVault.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PerpVault contract.
type PerpVaultDepositIterator struct {
	Event *PerpVaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultDeposit represents a Deposit event raised by the PerpVault contract.
type PerpVaultDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*PerpVaultDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultDepositIterator{contract: _PerpVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PerpVaultDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultDeposit)
				if err := _PerpVault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) ParseDeposit(log types.Log) (*PerpVaultDeposit, error) {
	event := new(PerpVaultDeposit)
	if err := _PerpVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultDonatedIterator is returned from FilterDonated and is used to iterate over the raw logs and unpacked data for Donated events raised by the PerpVault contract.
type PerpVaultDonatedIterator struct {
	Event *PerpVaultDonated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultDonated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultDonated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultDonated represents a Donated event raised by the PerpVault contract.
type PerpVaultDonated struct {
	Donor  common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonated is a free log retrieval operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 assets)
func (_PerpVault *PerpVaultFilterer) FilterDonated(opts *bind.FilterOpts, donor []common.Address) (*PerpVaultDonatedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultDonatedIterator{contract: _PerpVault.contract, event: "Donated", logs: logs, sub: sub}, nil
}

// WatchDonated is a free log subscription operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 assets)
func (_PerpVault *PerpVaultFilterer) WatchDonated(opts *bind.WatchOpts, sink chan<- *PerpVaultDonated, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultDonated)
				if err := _PerpVault.contract.UnpackLog(event, "Donated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDonated is a log parse operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 assets)
func (_PerpVault *PerpVaultFilterer) ParseDonated(log types.Log) (*PerpVaultDonated, error) {
	event := new(PerpVaultDonated)
	if err := _PerpVault.contract.UnpackLog(event, "Donated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultExecutorAddedIterator is returned from FilterExecutorAdded and is used to iterate over the raw logs and unpacked data for ExecutorAdded events raised by the PerpVault contract.
type PerpVaultExecutorAddedIterator struct {
	Event *PerpVaultExecutorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultExecutorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultExecutorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultExecutorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultExecutorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultExecutorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultExecutorAdded represents a ExecutorAdded event raised by the PerpVault contract.
type PerpVaultExecutorAdded struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecutorAdded is a free log retrieval operation binding the contract event 0xae5b7c3b000f575c241001dc9bcb3d8778376889353b07121115574eceff78c5.
//
// Solidity: event ExecutorAdded(address indexed executor)
func (_PerpVault *PerpVaultFilterer) FilterExecutorAdded(opts *bind.FilterOpts, executor []common.Address) (*PerpVaultExecutorAddedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "ExecutorAdded", executorRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultExecutorAddedIterator{contract: _PerpVault.contract, event: "ExecutorAdded", logs: logs, sub: sub}, nil
}

// WatchExecutorAdded is a free log subscription operation binding the contract event 0xae5b7c3b000f575c241001dc9bcb3d8778376889353b07121115574eceff78c5.
//
// Solidity: event ExecutorAdded(address indexed executor)
func (_PerpVault *PerpVaultFilterer) WatchExecutorAdded(opts *bind.WatchOpts, sink chan<- *PerpVaultExecutorAdded, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "ExecutorAdded", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultExecutorAdded)
				if err := _PerpVault.contract.UnpackLog(event, "ExecutorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutorAdded is a log parse operation binding the contract event 0xae5b7c3b000f575c241001dc9bcb3d8778376889353b07121115574eceff78c5.
//
// Solidity: event ExecutorAdded(address indexed executor)
func (_PerpVault *PerpVaultFilterer) ParseExecutorAdded(log types.Log) (*PerpVaultExecutorAdded, error) {
	event := new(PerpVaultExecutorAdded)
	if err := _PerpVault.contract.UnpackLog(event, "ExecutorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultExecutorRemovedIterator is returned from FilterExecutorRemoved and is used to iterate over the raw logs and unpacked data for ExecutorRemoved events raised by the PerpVault contract.
type PerpVaultExecutorRemovedIterator struct {
	Event *PerpVaultExecutorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultExecutorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultExecutorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultExecutorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultExecutorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultExecutorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultExecutorRemoved represents a ExecutorRemoved event raised by the PerpVault contract.
type PerpVaultExecutorRemoved struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecutorRemoved is a free log retrieval operation binding the contract event 0x4a2cf608bfb427f53279ec7f0eadf48913b9346ccefc3af138dbdec14ea0907d.
//
// Solidity: event ExecutorRemoved(address indexed executor)
func (_PerpVault *PerpVaultFilterer) FilterExecutorRemoved(opts *bind.FilterOpts, executor []common.Address) (*PerpVaultExecutorRemovedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "ExecutorRemoved", executorRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultExecutorRemovedIterator{contract: _PerpVault.contract, event: "ExecutorRemoved", logs: logs, sub: sub}, nil
}

// WatchExecutorRemoved is a free log subscription operation binding the contract event 0x4a2cf608bfb427f53279ec7f0eadf48913b9346ccefc3af138dbdec14ea0907d.
//
// Solidity: event ExecutorRemoved(address indexed executor)
func (_PerpVault *PerpVaultFilterer) WatchExecutorRemoved(opts *bind.WatchOpts, sink chan<- *PerpVaultExecutorRemoved, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "ExecutorRemoved", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultExecutorRemoved)
				if err := _PerpVault.contract.UnpackLog(event, "ExecutorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutorRemoved is a log parse operation binding the contract event 0x4a2cf608bfb427f53279ec7f0eadf48913b9346ccefc3af138dbdec14ea0907d.
//
// Solidity: event ExecutorRemoved(address indexed executor)
func (_PerpVault *PerpVaultFilterer) ParseExecutorRemoved(log types.Log) (*PerpVaultExecutorRemoved, error) {
	event := new(PerpVaultExecutorRemoved)
	if err := _PerpVault.contract.UnpackLog(event, "ExecutorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultExecutorUpdatedIterator is returned from FilterExecutorUpdated and is used to iterate over the raw logs and unpacked data for ExecutorUpdated events raised by the PerpVault contract.
type PerpVaultExecutorUpdatedIterator struct {
	Event *PerpVaultExecutorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultExecutorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultExecutorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultExecutorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultExecutorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultExecutorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultExecutorUpdated represents a ExecutorUpdated event raised by the PerpVault contract.
type PerpVaultExecutorUpdated struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecutorUpdated is a free log retrieval operation binding the contract event 0xd9436ef9ce00ffeabc5da2489701502d3bd1a5ed7b254a1981fc5ffef9828e11.
//
// Solidity: event ExecutorUpdated(address indexed executor)
func (_PerpVault *PerpVaultFilterer) FilterExecutorUpdated(opts *bind.FilterOpts, executor []common.Address) (*PerpVaultExecutorUpdatedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "ExecutorUpdated", executorRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultExecutorUpdatedIterator{contract: _PerpVault.contract, event: "ExecutorUpdated", logs: logs, sub: sub}, nil
}

// WatchExecutorUpdated is a free log subscription operation binding the contract event 0xd9436ef9ce00ffeabc5da2489701502d3bd1a5ed7b254a1981fc5ffef9828e11.
//
// Solidity: event ExecutorUpdated(address indexed executor)
func (_PerpVault *PerpVaultFilterer) WatchExecutorUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultExecutorUpdated, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "ExecutorUpdated", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultExecutorUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "ExecutorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutorUpdated is a log parse operation binding the contract event 0xd9436ef9ce00ffeabc5da2489701502d3bd1a5ed7b254a1981fc5ffef9828e11.
//
// Solidity: event ExecutorUpdated(address indexed executor)
func (_PerpVault *PerpVaultFilterer) ParseExecutorUpdated(log types.Log) (*PerpVaultExecutorUpdated, error) {
	event := new(PerpVaultExecutorUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "ExecutorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PerpVault contract.
type PerpVaultInitializedIterator struct {
	Event *PerpVaultInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultInitialized represents a Initialized event raised by the PerpVault contract.
type PerpVaultInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PerpVault *PerpVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*PerpVaultInitializedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PerpVaultInitializedIterator{contract: _PerpVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PerpVault *PerpVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PerpVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultInitialized)
				if err := _PerpVault.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PerpVault *PerpVaultFilterer) ParseInitialized(log types.Log) (*PerpVaultInitialized, error) {
	event := new(PerpVaultInitialized)
	if err := _PerpVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultLockupPeriodUpdatedIterator is returned from FilterLockupPeriodUpdated and is used to iterate over the raw logs and unpacked data for LockupPeriodUpdated events raised by the PerpVault contract.
type PerpVaultLockupPeriodUpdatedIterator struct {
	Event *PerpVaultLockupPeriodUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultLockupPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultLockupPeriodUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultLockupPeriodUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultLockupPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultLockupPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultLockupPeriodUpdated represents a LockupPeriodUpdated event raised by the PerpVault contract.
type PerpVaultLockupPeriodUpdated struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockupPeriodUpdated is a free log retrieval operation binding the contract event 0xa4711e25376fde3c9a55e3dba876110179242e2dcb40fd6644c6d7884642a1f5.
//
// Solidity: event LockupPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_PerpVault *PerpVaultFilterer) FilterLockupPeriodUpdated(opts *bind.FilterOpts) (*PerpVaultLockupPeriodUpdatedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "LockupPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpVaultLockupPeriodUpdatedIterator{contract: _PerpVault.contract, event: "LockupPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLockupPeriodUpdated is a free log subscription operation binding the contract event 0xa4711e25376fde3c9a55e3dba876110179242e2dcb40fd6644c6d7884642a1f5.
//
// Solidity: event LockupPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_PerpVault *PerpVaultFilterer) WatchLockupPeriodUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultLockupPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "LockupPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultLockupPeriodUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "LockupPeriodUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockupPeriodUpdated is a log parse operation binding the contract event 0xa4711e25376fde3c9a55e3dba876110179242e2dcb40fd6644c6d7884642a1f5.
//
// Solidity: event LockupPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_PerpVault *PerpVaultFilterer) ParseLockupPeriodUpdated(log types.Log) (*PerpVaultLockupPeriodUpdated, error) {
	event := new(PerpVaultLockupPeriodUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "LockupPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the PerpVault contract.
type PerpVaultManagerAddedIterator struct {
	Event *PerpVaultManagerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultManagerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultManagerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultManagerAdded represents a ManagerAdded event raised by the PerpVault contract.
type PerpVaultManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_PerpVault *PerpVaultFilterer) FilterManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*PerpVaultManagerAddedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultManagerAddedIterator{contract: _PerpVault.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_PerpVault *PerpVaultFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *PerpVaultManagerAdded, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultManagerAdded)
				if err := _PerpVault.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_PerpVault *PerpVaultFilterer) ParseManagerAdded(log types.Log) (*PerpVaultManagerAdded, error) {
	event := new(PerpVaultManagerAdded)
	if err := _PerpVault.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the PerpVault contract.
type PerpVaultManagerRemovedIterator struct {
	Event *PerpVaultManagerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultManagerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultManagerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultManagerRemoved represents a ManagerRemoved event raised by the PerpVault contract.
type PerpVaultManagerRemoved struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_PerpVault *PerpVaultFilterer) FilterManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*PerpVaultManagerRemovedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultManagerRemovedIterator{contract: _PerpVault.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_PerpVault *PerpVaultFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *PerpVaultManagerRemoved, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultManagerRemoved)
				if err := _PerpVault.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_PerpVault *PerpVaultFilterer) ParseManagerRemoved(log types.Log) (*PerpVaultManagerRemoved, error) {
	event := new(PerpVaultManagerRemoved)
	if err := _PerpVault.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultMarketRuleSetIterator is returned from FilterMarketRuleSet and is used to iterate over the raw logs and unpacked data for MarketRuleSet events raised by the PerpVault contract.
type PerpVaultMarketRuleSetIterator struct {
	Event *PerpVaultMarketRuleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultMarketRuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultMarketRuleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultMarketRuleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultMarketRuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultMarketRuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultMarketRuleSet represents a MarketRuleSet event raised by the PerpVault contract.
type PerpVaultMarketRuleSet struct {
	MarketId uint64
	Cap      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRuleSet is a free log retrieval operation binding the contract event 0xb5c53be95670e712b0ed22b769b3b8ae589ed033c496933be419cdf1c9ee955a.
//
// Solidity: event MarketRuleSet(uint64 indexed marketId, uint256 cap)
func (_PerpVault *PerpVaultFilterer) FilterMarketRuleSet(opts *bind.FilterOpts, marketId []uint64) (*PerpVaultMarketRuleSetIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "MarketRuleSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultMarketRuleSetIterator{contract: _PerpVault.contract, event: "MarketRuleSet", logs: logs, sub: sub}, nil
}

// WatchMarketRuleSet is a free log subscription operation binding the contract event 0xb5c53be95670e712b0ed22b769b3b8ae589ed033c496933be419cdf1c9ee955a.
//
// Solidity: event MarketRuleSet(uint64 indexed marketId, uint256 cap)
func (_PerpVault *PerpVaultFilterer) WatchMarketRuleSet(opts *bind.WatchOpts, sink chan<- *PerpVaultMarketRuleSet, marketId []uint64) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "MarketRuleSet", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultMarketRuleSet)
				if err := _PerpVault.contract.UnpackLog(event, "MarketRuleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketRuleSet is a log parse operation binding the contract event 0xb5c53be95670e712b0ed22b769b3b8ae589ed033c496933be419cdf1c9ee955a.
//
// Solidity: event MarketRuleSet(uint64 indexed marketId, uint256 cap)
func (_PerpVault *PerpVaultFilterer) ParseMarketRuleSet(log types.Log) (*PerpVaultMarketRuleSet, error) {
	event := new(PerpVaultMarketRuleSet)
	if err := _PerpVault.contract.UnpackLog(event, "MarketRuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultMaxMarginUtilizationUpdatedIterator is returned from FilterMaxMarginUtilizationUpdated and is used to iterate over the raw logs and unpacked data for MaxMarginUtilizationUpdated events raised by the PerpVault contract.
type PerpVaultMaxMarginUtilizationUpdatedIterator struct {
	Event *PerpVaultMaxMarginUtilizationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultMaxMarginUtilizationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultMaxMarginUtilizationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultMaxMarginUtilizationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultMaxMarginUtilizationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultMaxMarginUtilizationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultMaxMarginUtilizationUpdated represents a MaxMarginUtilizationUpdated event raised by the PerpVault contract.
type PerpVaultMaxMarginUtilizationUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxMarginUtilizationUpdated is a free log retrieval operation binding the contract event 0x04ae0db7c0aa77a9404f22d44aa47d647c6427822272a8a6d5fd43dc5c204e07.
//
// Solidity: event MaxMarginUtilizationUpdated(uint256 oldBps, uint256 newBps)
func (_PerpVault *PerpVaultFilterer) FilterMaxMarginUtilizationUpdated(opts *bind.FilterOpts) (*PerpVaultMaxMarginUtilizationUpdatedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "MaxMarginUtilizationUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpVaultMaxMarginUtilizationUpdatedIterator{contract: _PerpVault.contract, event: "MaxMarginUtilizationUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMarginUtilizationUpdated is a free log subscription operation binding the contract event 0x04ae0db7c0aa77a9404f22d44aa47d647c6427822272a8a6d5fd43dc5c204e07.
//
// Solidity: event MaxMarginUtilizationUpdated(uint256 oldBps, uint256 newBps)
func (_PerpVault *PerpVaultFilterer) WatchMaxMarginUtilizationUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultMaxMarginUtilizationUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "MaxMarginUtilizationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultMaxMarginUtilizationUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "MaxMarginUtilizationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxMarginUtilizationUpdated is a log parse operation binding the contract event 0x04ae0db7c0aa77a9404f22d44aa47d647c6427822272a8a6d5fd43dc5c204e07.
//
// Solidity: event MaxMarginUtilizationUpdated(uint256 oldBps, uint256 newBps)
func (_PerpVault *PerpVaultFilterer) ParseMaxMarginUtilizationUpdated(log types.Log) (*PerpVaultMaxMarginUtilizationUpdated, error) {
	event := new(PerpVaultMaxMarginUtilizationUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "MaxMarginUtilizationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultMinDepositUpdatedIterator is returned from FilterMinDepositUpdated and is used to iterate over the raw logs and unpacked data for MinDepositUpdated events raised by the PerpVault contract.
type PerpVaultMinDepositUpdatedIterator struct {
	Event *PerpVaultMinDepositUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultMinDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultMinDepositUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultMinDepositUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultMinDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultMinDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultMinDepositUpdated represents a MinDepositUpdated event raised by the PerpVault contract.
type PerpVaultMinDepositUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinDepositUpdated is a free log retrieval operation binding the contract event 0xb566d3df2587c9e70b06b6419bdeeeeec8ca8cd60e4c48c6baad0d94c46809c7.
//
// Solidity: event MinDepositUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) FilterMinDepositUpdated(opts *bind.FilterOpts) (*PerpVaultMinDepositUpdatedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpVaultMinDepositUpdatedIterator{contract: _PerpVault.contract, event: "MinDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDepositUpdated is a free log subscription operation binding the contract event 0xb566d3df2587c9e70b06b6419bdeeeeec8ca8cd60e4c48c6baad0d94c46809c7.
//
// Solidity: event MinDepositUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) WatchMinDepositUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultMinDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultMinDepositUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDepositUpdated is a log parse operation binding the contract event 0xb566d3df2587c9e70b06b6419bdeeeeec8ca8cd60e4c48c6baad0d94c46809c7.
//
// Solidity: event MinDepositUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) ParseMinDepositUpdated(log types.Log) (*PerpVaultMinDepositUpdated, error) {
	event := new(PerpVaultMinDepositUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultMinWithdrawUpdatedIterator is returned from FilterMinWithdrawUpdated and is used to iterate over the raw logs and unpacked data for MinWithdrawUpdated events raised by the PerpVault contract.
type PerpVaultMinWithdrawUpdatedIterator struct {
	Event *PerpVaultMinWithdrawUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultMinWithdrawUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultMinWithdrawUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultMinWithdrawUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultMinWithdrawUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultMinWithdrawUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultMinWithdrawUpdated represents a MinWithdrawUpdated event raised by the PerpVault contract.
type PerpVaultMinWithdrawUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawUpdated is a free log retrieval operation binding the contract event 0x3c4f4d8cd2a65b4b1f4eeaf43669b14ab54e43d4842aa0ac8f0e4f9fe0bf5bf9.
//
// Solidity: event MinWithdrawUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) FilterMinWithdrawUpdated(opts *bind.FilterOpts) (*PerpVaultMinWithdrawUpdatedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "MinWithdrawUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpVaultMinWithdrawUpdatedIterator{contract: _PerpVault.contract, event: "MinWithdrawUpdated", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawUpdated is a free log subscription operation binding the contract event 0x3c4f4d8cd2a65b4b1f4eeaf43669b14ab54e43d4842aa0ac8f0e4f9fe0bf5bf9.
//
// Solidity: event MinWithdrawUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) WatchMinWithdrawUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultMinWithdrawUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "MinWithdrawUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultMinWithdrawUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "MinWithdrawUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinWithdrawUpdated is a log parse operation binding the contract event 0x3c4f4d8cd2a65b4b1f4eeaf43669b14ab54e43d4842aa0ac8f0e4f9fe0bf5bf9.
//
// Solidity: event MinWithdrawUpdated(uint256 oldValue, uint256 newValue)
func (_PerpVault *PerpVaultFilterer) ParseMinWithdrawUpdated(log types.Log) (*PerpVaultMinWithdrawUpdated, error) {
	event := new(PerpVaultMinWithdrawUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "MinWithdrawUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PerpVault contract.
type PerpVaultOwnershipTransferredIterator struct {
	Event *PerpVaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultOwnershipTransferred represents a OwnershipTransferred event raised by the PerpVault contract.
type PerpVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpVault *PerpVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PerpVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultOwnershipTransferredIterator{contract: _PerpVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpVault *PerpVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PerpVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultOwnershipTransferred)
				if err := _PerpVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpVault *PerpVaultFilterer) ParseOwnershipTransferred(log types.Log) (*PerpVaultOwnershipTransferred, error) {
	event := new(PerpVaultOwnershipTransferred)
	if err := _PerpVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PerpVault contract.
type PerpVaultPausedIterator struct {
	Event *PerpVaultPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultPaused represents a Paused event raised by the PerpVault contract.
type PerpVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PerpVault *PerpVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*PerpVaultPausedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PerpVaultPausedIterator{contract: _PerpVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PerpVault *PerpVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PerpVaultPaused) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultPaused)
				if err := _PerpVault.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PerpVault *PerpVaultFilterer) ParsePaused(log types.Log) (*PerpVaultPaused, error) {
	event := new(PerpVaultPaused)
	if err := _PerpVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultRedeemRequestedIterator is returned from FilterRedeemRequested and is used to iterate over the raw logs and unpacked data for RedeemRequested events raised by the PerpVault contract.
type PerpVaultRedeemRequestedIterator struct {
	Event *PerpVaultRedeemRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultRedeemRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultRedeemRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultRedeemRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultRedeemRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultRedeemRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultRedeemRequested represents a RedeemRequested event raised by the PerpVault contract.
type PerpVaultRedeemRequested struct {
	RequestId    *big.Int
	Owner        common.Address
	Receiver     common.Address
	Shares       *big.Int
	Assets       *big.Int
	ExecutableAt *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemRequested is a free log retrieval operation binding the contract event 0x3b8064ce836b1010de5808e04b451d4ebfe1157c492ad0e1b6d8ccffb329eb39.
//
// Solidity: event RedeemRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) FilterRedeemRequested(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*PerpVaultRedeemRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "RedeemRequested", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultRedeemRequestedIterator{contract: _PerpVault.contract, event: "RedeemRequested", logs: logs, sub: sub}, nil
}

// WatchRedeemRequested is a free log subscription operation binding the contract event 0x3b8064ce836b1010de5808e04b451d4ebfe1157c492ad0e1b6d8ccffb329eb39.
//
// Solidity: event RedeemRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) WatchRedeemRequested(opts *bind.WatchOpts, sink chan<- *PerpVaultRedeemRequested, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "RedeemRequested", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultRedeemRequested)
				if err := _PerpVault.contract.UnpackLog(event, "RedeemRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemRequested is a log parse operation binding the contract event 0x3b8064ce836b1010de5808e04b451d4ebfe1157c492ad0e1b6d8ccffb329eb39.
//
// Solidity: event RedeemRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) ParseRedeemRequested(log types.Log) (*PerpVaultRedeemRequested, error) {
	event := new(PerpVaultRedeemRequested)
	if err := _PerpVault.contract.UnpackLog(event, "RedeemRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultRequestExecutedIterator is returned from FilterRequestExecuted and is used to iterate over the raw logs and unpacked data for RequestExecuted events raised by the PerpVault contract.
type PerpVaultRequestExecutedIterator struct {
	Event *PerpVaultRequestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultRequestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultRequestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultRequestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultRequestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultRequestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultRequestExecuted represents a RequestExecuted event raised by the PerpVault contract.
type PerpVaultRequestExecuted struct {
	RequestId *big.Int
	Owner     common.Address
	Receiver  common.Address
	Shares    *big.Int
	Assets    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestExecuted is a free log retrieval operation binding the contract event 0xa7a5da0a86cebe498b65e58861c57870380d8ea9e66683ec37d6ad37181cab84.
//
// Solidity: event RequestExecuted(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets)
func (_PerpVault *PerpVaultFilterer) FilterRequestExecuted(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*PerpVaultRequestExecutedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "RequestExecuted", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultRequestExecutedIterator{contract: _PerpVault.contract, event: "RequestExecuted", logs: logs, sub: sub}, nil
}

// WatchRequestExecuted is a free log subscription operation binding the contract event 0xa7a5da0a86cebe498b65e58861c57870380d8ea9e66683ec37d6ad37181cab84.
//
// Solidity: event RequestExecuted(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets)
func (_PerpVault *PerpVaultFilterer) WatchRequestExecuted(opts *bind.WatchOpts, sink chan<- *PerpVaultRequestExecuted, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "RequestExecuted", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultRequestExecuted)
				if err := _PerpVault.contract.UnpackLog(event, "RequestExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestExecuted is a log parse operation binding the contract event 0xa7a5da0a86cebe498b65e58861c57870380d8ea9e66683ec37d6ad37181cab84.
//
// Solidity: event RequestExecuted(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets)
func (_PerpVault *PerpVaultFilterer) ParseRequestExecuted(log types.Log) (*PerpVaultRequestExecuted, error) {
	event := new(PerpVaultRequestExecuted)
	if err := _PerpVault.contract.UnpackLog(event, "RequestExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultRequestExecutionFailedIterator is returned from FilterRequestExecutionFailed and is used to iterate over the raw logs and unpacked data for RequestExecutionFailed events raised by the PerpVault contract.
type PerpVaultRequestExecutionFailedIterator struct {
	Event *PerpVaultRequestExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultRequestExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultRequestExecutionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultRequestExecutionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultRequestExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultRequestExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultRequestExecutionFailed represents a RequestExecutionFailed event raised by the PerpVault contract.
type PerpVaultRequestExecutionFailed struct {
	RequestId *big.Int
	Owner     common.Address
	Receiver  common.Address
	Assets    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestExecutionFailed is a free log retrieval operation binding the contract event 0x9aaaffd6862f8fa67d255c221adc58895928765ead274e59afe9ff06cd0d27d6.
//
// Solidity: event RequestExecutionFailed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 assets)
func (_PerpVault *PerpVaultFilterer) FilterRequestExecutionFailed(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*PerpVaultRequestExecutionFailedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "RequestExecutionFailed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultRequestExecutionFailedIterator{contract: _PerpVault.contract, event: "RequestExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchRequestExecutionFailed is a free log subscription operation binding the contract event 0x9aaaffd6862f8fa67d255c221adc58895928765ead274e59afe9ff06cd0d27d6.
//
// Solidity: event RequestExecutionFailed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 assets)
func (_PerpVault *PerpVaultFilterer) WatchRequestExecutionFailed(opts *bind.WatchOpts, sink chan<- *PerpVaultRequestExecutionFailed, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "RequestExecutionFailed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultRequestExecutionFailed)
				if err := _PerpVault.contract.UnpackLog(event, "RequestExecutionFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestExecutionFailed is a log parse operation binding the contract event 0x9aaaffd6862f8fa67d255c221adc58895928765ead274e59afe9ff06cd0d27d6.
//
// Solidity: event RequestExecutionFailed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 assets)
func (_PerpVault *PerpVaultFilterer) ParseRequestExecutionFailed(log types.Log) (*PerpVaultRequestExecutionFailed, error) {
	event := new(PerpVaultRequestExecutionFailed)
	if err := _PerpVault.contract.UnpackLog(event, "RequestExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PerpVault contract.
type PerpVaultTransferIterator struct {
	Event *PerpVaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultTransfer represents a Transfer event raised by the PerpVault contract.
type PerpVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PerpVault *PerpVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PerpVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultTransferIterator{contract: _PerpVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PerpVault *PerpVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PerpVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultTransfer)
				if err := _PerpVault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PerpVault *PerpVaultFilterer) ParseTransfer(log types.Log) (*PerpVaultTransfer, error) {
	event := new(PerpVaultTransfer)
	if err := _PerpVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PerpVault contract.
type PerpVaultUnpausedIterator struct {
	Event *PerpVaultUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultUnpaused represents a Unpaused event raised by the PerpVault contract.
type PerpVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PerpVault *PerpVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PerpVaultUnpausedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PerpVaultUnpausedIterator{contract: _PerpVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PerpVault *PerpVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PerpVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultUnpaused)
				if err := _PerpVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PerpVault *PerpVaultFilterer) ParseUnpaused(log types.Log) (*PerpVaultUnpaused, error) {
	event := new(PerpVaultUnpaused)
	if err := _PerpVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PerpVault contract.
type PerpVaultUpgradedIterator struct {
	Event *PerpVaultUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultUpgraded represents a Upgraded event raised by the PerpVault contract.
type PerpVaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PerpVault *PerpVaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PerpVaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultUpgradedIterator{contract: _PerpVault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PerpVault *PerpVaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PerpVaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultUpgraded)
				if err := _PerpVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PerpVault *PerpVaultFilterer) ParseUpgraded(log types.Log) (*PerpVaultUpgraded, error) {
	event := new(PerpVaultUpgraded)
	if err := _PerpVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultWhitelistEnabledUpdatedIterator is returned from FilterWhitelistEnabledUpdated and is used to iterate over the raw logs and unpacked data for WhitelistEnabledUpdated events raised by the PerpVault contract.
type PerpVaultWhitelistEnabledUpdatedIterator struct {
	Event *PerpVaultWhitelistEnabledUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultWhitelistEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultWhitelistEnabledUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultWhitelistEnabledUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultWhitelistEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultWhitelistEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultWhitelistEnabledUpdated represents a WhitelistEnabledUpdated event raised by the PerpVault contract.
type PerpVaultWhitelistEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEnabledUpdated is a free log retrieval operation binding the contract event 0x49d3057180a80162d2a0381be6848c15e0d117e900366482dd3b5443ca8db974.
//
// Solidity: event WhitelistEnabledUpdated(bool enabled)
func (_PerpVault *PerpVaultFilterer) FilterWhitelistEnabledUpdated(opts *bind.FilterOpts) (*PerpVaultWhitelistEnabledUpdatedIterator, error) {

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "WhitelistEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &PerpVaultWhitelistEnabledUpdatedIterator{contract: _PerpVault.contract, event: "WhitelistEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistEnabledUpdated is a free log subscription operation binding the contract event 0x49d3057180a80162d2a0381be6848c15e0d117e900366482dd3b5443ca8db974.
//
// Solidity: event WhitelistEnabledUpdated(bool enabled)
func (_PerpVault *PerpVaultFilterer) WatchWhitelistEnabledUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultWhitelistEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "WhitelistEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultWhitelistEnabledUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "WhitelistEnabledUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistEnabledUpdated is a log parse operation binding the contract event 0x49d3057180a80162d2a0381be6848c15e0d117e900366482dd3b5443ca8db974.
//
// Solidity: event WhitelistEnabledUpdated(bool enabled)
func (_PerpVault *PerpVaultFilterer) ParseWhitelistEnabledUpdated(log types.Log) (*PerpVaultWhitelistEnabledUpdated, error) {
	event := new(PerpVaultWhitelistEnabledUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "WhitelistEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultWhitelistUpdatedIterator is returned from FilterWhitelistUpdated and is used to iterate over the raw logs and unpacked data for WhitelistUpdated events raised by the PerpVault contract.
type PerpVaultWhitelistUpdatedIterator struct {
	Event *PerpVaultWhitelistUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultWhitelistUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultWhitelistUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultWhitelistUpdated represents a WhitelistUpdated event raised by the PerpVault contract.
type PerpVaultWhitelistUpdated struct {
	Account common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistUpdated is a free log retrieval operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool allowed)
func (_PerpVault *PerpVaultFilterer) FilterWhitelistUpdated(opts *bind.FilterOpts, account []common.Address) (*PerpVaultWhitelistUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "WhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultWhitelistUpdatedIterator{contract: _PerpVault.contract, event: "WhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistUpdated is a free log subscription operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool allowed)
func (_PerpVault *PerpVaultFilterer) WatchWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *PerpVaultWhitelistUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "WhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultWhitelistUpdated)
				if err := _PerpVault.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistUpdated is a log parse operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool allowed)
func (_PerpVault *PerpVaultFilterer) ParseWhitelistUpdated(log types.Log) (*PerpVaultWhitelistUpdated, error) {
	event := new(PerpVaultWhitelistUpdated)
	if err := _PerpVault.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PerpVault contract.
type PerpVaultWithdrawIterator struct {
	Event *PerpVaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultWithdraw represents a Withdraw event raised by the PerpVault contract.
type PerpVaultWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*PerpVaultWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultWithdrawIterator{contract: _PerpVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PerpVaultWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultWithdraw)
				if err := _PerpVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PerpVault *PerpVaultFilterer) ParseWithdraw(log types.Log) (*PerpVaultWithdraw, error) {
	event := new(PerpVaultWithdraw)
	if err := _PerpVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpVaultWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the PerpVault contract.
type PerpVaultWithdrawRequestedIterator struct {
	Event *PerpVaultWithdrawRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PerpVaultWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpVaultWithdrawRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PerpVaultWithdrawRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PerpVaultWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpVaultWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpVaultWithdrawRequested represents a WithdrawRequested event raised by the PerpVault contract.
type PerpVaultWithdrawRequested struct {
	RequestId    *big.Int
	Owner        common.Address
	Receiver     common.Address
	Shares       *big.Int
	Assets       *big.Int
	ExecutableAt *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0x74a8d0df732a141e45f44b230aadff310598bbeabb761b313e96ae2e27337a79.
//
// Solidity: event WithdrawRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*PerpVaultWithdrawRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.FilterLogs(opts, "WithdrawRequested", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PerpVaultWithdrawRequestedIterator{contract: _PerpVault.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0x74a8d0df732a141e45f44b230aadff310598bbeabb761b313e96ae2e27337a79.
//
// Solidity: event WithdrawRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *PerpVaultWithdrawRequested, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PerpVault.contract.WatchLogs(opts, "WithdrawRequested", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpVaultWithdrawRequested)
				if err := _PerpVault.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawRequested is a log parse operation binding the contract event 0x74a8d0df732a141e45f44b230aadff310598bbeabb761b313e96ae2e27337a79.
//
// Solidity: event WithdrawRequested(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 shares, uint256 assets, uint256 executableAt)
func (_PerpVault *PerpVaultFilterer) ParseWithdrawRequested(log types.Log) (*PerpVaultWithdrawRequested, error) {
	event := new(PerpVaultWithdrawRequested)
	if err := _PerpVault.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

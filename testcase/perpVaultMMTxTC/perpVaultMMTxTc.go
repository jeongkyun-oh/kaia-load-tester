package perpVaultMMTxTC

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"math/rand"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/myzhan/boomer"
)

const Name = "perpVaultMMTxTC"

const (
	tifPost = uint8(2) // POST-only: rejected if it would cross → guarantees no trade
)
const (
	sideBuy  = uint8(0)
	sideSell = uint8(1)
)

type executor struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

var (
	endPoint string
	signer   *account.Account // any funded account; used only to build/sign session txs
	cliPool  clipool.ClientPool

	vaultAddr common.Address
	execs     []executor
	nExec     int
	cursor    uint32

	marketId uint64 = 1
	refPrice        = scaleUp(100)
	tickSize        = scaleUp(1)
)

func SetMarketId(id uint64) { marketId = id }
func SetRefPrice(p *big.Int) {
	if p != nil && p.Sign() > 0 {
		refPrice = p
	}
}
func SetTickSize(p *big.Int) {
	if p != nil && p.Sign() > 0 {
		tickSize = p
	}
}

// SetVault injects the deployed vault address and the executor session keys
// registered on it (from the setup phase).
func SetVault(vault common.Address, execKeys []*ecdsa.PrivateKey) {
	vaultAddr = vault
	execs = execs[:0]
	for _, k := range execKeys {
		execs = append(execs, executor{key: k, addr: crypto.PubkeyToAddress(k.PublicKey)})
	}
	nExec = len(execs)
}

func Init(accs []*account.Account, endpoint string, _ *big.Int) {
	endPoint = endpoint
	cliCreate := func() any {
		c, err := ethclient.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}
	cliPool.Init(1000, 3000, cliCreate)
	if len(accs) > 0 {
		signer = accs[0]
	}
}

// makerPrice returns a tick-aligned, non-crossing price: buys one tick below the
// aligned reference, sells one tick above. Same guarantee as perpNoTradeTxTC.
func makerPrice(side uint8) *big.Int {
	refAligned := new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
	if side == sideBuy {
		return new(big.Int).Sub(refAligned, tickSize)
	}
	return new(big.Int).Add(refAligned, tickSize)
}

func Run() {
	if nExec == 0 {
		return
	}
	cli := cliPool.Alloc().(*ethclient.Client)
	defer cliPool.Free(cli)

	idx := atomic.AddUint32(&cursor, 1) % uint32(nExec)
	exec := execs[idx]
	side := uint8(rand.Intn(2))
	price := makerPrice(side)
	quantity := scaleUp(1)

	start := boomer.Now()
	tx, err := signer.GenPerpOrderTxBySession(exec.key, vaultAddr, marketId, side, price, quantity, tifPost)
	if err != nil {
		log.Printf("Failed to gen vault MM perp order: err=%v, exec=%s, side=%d, price=%s", err, exec.addr.Hex(), side, price)
		return
	}
	_, err = signer.SendTx(cli, tx)
	elapsed := boomer.Now() - start
	if err != nil {
		log.Printf("Failed to send vault MM perp order: err=%v, exec=%s, side=%d, price=%s", err, exec.addr.Hex(), side, price)
		boomer.RecordFailure("http", "SendVaultMMPerpOrderTx to "+endPoint, elapsed, err.Error())
		return
	}
	boomer.RecordSuccess("http", "SendVaultMMPerpOrderTx to "+endPoint, elapsed, int64(10))
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}

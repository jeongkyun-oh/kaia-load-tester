package perpNoTradeTxTC

import (
	"log"
	"math/big"
	"math/rand"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/myzhan/boomer"
)

const Name = "perpNoTradeTxTC"

// Perp order TimeInForce values (see core/types dex_perp.go).
const (
	tifGTC  = uint8(0)
	tifPost = uint8(2) // POST-only: never takes liquidity; rejected if it would cross.
)

// Perp order side (see core/types dex_perp.go).
const (
	sideBuy  = uint8(0)
	sideSell = uint8(1)
)

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	cursor uint32

	// marketId is the perp market to trade against. It is assumed to be already
	// registered on the target L2 (an ArbPerpMarketRegistry admin op); set it via
	// SetMarketId from the -perpMarketId flag.
	marketId uint64 = 1

	// refPrice approximates the market's mark price (18-decimal fixed-point). Perp
	// orders are rejected by the price-band check if placed too far from the mark,
	// so refPrice must be tuned to the market via -perpRefPrice. Buy orders rest one
	// tick BELOW the (tick-aligned) reference and sells one tick ABOVE it, so opposite
	// sides never cross (canMatchLimit: BUY crosses when price>=ask, SELL when
	// price<=bid). That non-crossing spread, plus POST-only, guarantees no trade.
	refPrice = scaleUp(100)

	// tickSize is the market's price tick (18-decimal). Order prices MUST be exact
	// multiples of it or the node rejects them ("price ... is not a multiple of tick
	// size"). Override via -perpTickSize to match the target market.
	tickSize = scaleUp(1)
)

// SetMarketId overrides the target perp market id (from the -perpMarketId flag).
func SetMarketId(id uint64) {
	marketId = id
}

// SetRefPrice overrides the reference (mark) price used to derive non-crossing
// order prices (from the -perpRefPrice flag). Ignored when p is nil or non-positive.
func SetRefPrice(p *big.Int) {
	if p != nil && p.Sign() > 0 {
		refPrice = p
	}
}

// SetTickSize overrides the market's price tick used to align order prices (from
// the -perpTickSize flag). Ignored when p is nil or non-positive.
func SetTickSize(p *big.Int) {
	if p != nil && p.Sign() > 0 {
		tickSize = p
	}
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

	accGrp = append(accGrp, accs...)

	nAcc = len(accGrp)
}

func Run() {
	cli := cliPool.Alloc().(*ethclient.Client)
	defer cliPool.Free(cli)

	var (
		from     = accGrp[atomic.AddUint32(&cursor, 1)%uint32(nAcc)]
		quantity = scaleUp(1)
		side     = uint8(rand.Intn(2))
		price    *big.Int
	)

	// Align the reference to the tick grid, then rest one tick below (buy) or above
	// (sell). Tick-aligned so the node accepts it; buy < sell so opposite sides never
	// cross. POST-only is the belt-and-suspenders no-trade guarantee.
	refAligned := new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
	if side == sideBuy {
		price = new(big.Int).Sub(refAligned, tickSize)
	} else {
		price = new(big.Int).Add(refAligned, tickSize)
	}

	start := boomer.Now()
	tx, err := from.GenNewPerpOrderTx(marketId, side, price, quantity, tifPost)
	if err != nil {
		log.Printf("Failed to generate new perp order tx: error=%v, marketId=%d, side=%d, price=%s, quantity=%s",
			err, marketId, side, price.String(), quantity.String())
		return
	}
	_, err = from.SendTx(cli, tx)
	elapsed := boomer.Now() - start
	if err != nil {
		log.Printf("Failed to send new perp order tx: error=%v, marketId=%d, side=%d, price=%s, quantity=%s\n",
			err, marketId, side, price.String(), quantity.String())
		boomer.RecordFailure("http", "SendNewPerpOrderTx"+" to "+endPoint, elapsed, err.Error())
		return
	}

	boomer.RecordSuccess("http", "SendNewPerpOrderTx"+" to "+endPoint, elapsed, int64(10))
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}

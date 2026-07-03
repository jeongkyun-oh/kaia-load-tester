package perpTakerTxTC

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

const Name = "perpTakerTxTC"

// Perp order TimeInForce (see core/types dex_perp.go). Taker types only: GTC crosses
// and rests the remainder, IOC crosses and cancels the remainder, MARKET takes at best
// price. POST (2) is maker-only and intentionally excluded.
const (
	tifGTC    = uint8(0)
	tifIOC    = uint8(1)
	tifMarket = uint8(3)
)

// takerTifs is the set of order types this TC randomly chooses from each Run.
var takerTifs = []uint8{tifGTC, tifIOC, tifMarket}

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

	// marketId / refPrice / tickSize mirror perpNoTradeTxTC; set from the shared
	// -perpMarketId / -perpRefPrice / -perpTickSize flags in main.go.
	marketId uint64 = 1
	refPrice        = scaleUp(100)
	tickSize        = scaleUp(1)
)

// SetMarketId overrides the target perp market id (from the -perpMarketId flag).
func SetMarketId(id uint64) { marketId = id }

// SetRefPrice overrides the reference (mark) price (from the -perpRefPrice flag).
// Ignored when p is nil or non-positive.
func SetRefPrice(p *big.Int) {
	if p != nil && p.Sign() > 0 {
		refPrice = p
	}
}

// SetTickSize overrides the market's price tick (from the -perpTickSize flag).
// Ignored when p is nil or non-positive.
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

// takerPrice returns a tick-aligned crossing price: a taker BUY posts one tick ABOVE the
// aligned reference (lifting a resting ask at ref+tick), a taker SELL one tick BELOW
// (hitting a resting bid at ref-tick). This is the opposite of the maker (perpNoTradeTxTC)
// convention, so a taker order crosses the maker liquidity resting between the two levels.
// When refPrice <= tickSize the sell level would be zero or negative, which the node
// rejects; it is clamped to the smallest valid tick level instead.
func takerPrice(side uint8) *big.Int {
	refAligned := new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
	if side == sideBuy {
		return new(big.Int).Add(refAligned, tickSize)
	}
	sell := new(big.Int).Sub(refAligned, tickSize)
	if sell.Sign() <= 0 {
		return new(big.Int).Set(tickSize)
	}
	return sell
}

// accountSide pins an account (by its slot index) to a single order side: even
// indices trade BUY-only, odd indices SELL-only (same invariant as perpHalfFillTxTC).
// A GTC remainder can rest on the book, and a match needs opposite sides, so a
// pinned account can never cross its own resting order: self-trades are impossible.
func accountSide(idx uint32) uint8 {
	if idx%2 == 0 {
		return sideBuy
	}
	return sideSell
}

func Run() {
	if nAcc == 0 {
		return
	}
	cli := cliPool.Alloc().(*ethclient.Client)
	defer cliPool.Free(cli)

	idx := atomic.AddUint32(&cursor, 1) % uint32(nAcc)
	from := accGrp[idx]
	side := accountSide(idx)
	tif := takerTifs[rand.Intn(len(takerTifs))]
	price := takerPrice(side)
	quantity := scaleUp(1)

	start := boomer.Now()
	tx, err := from.GenNewPerpOrderTx(marketId, side, price, quantity, tif)
	if err != nil {
		log.Printf("Failed to generate perp taker tx: error=%v, marketId=%d, side=%d, tif=%d, price=%s, quantity=%s",
			err, marketId, side, tif, price.String(), quantity.String())
		return
	}
	_, err = from.SendTx(cli, tx)
	elapsed := boomer.Now() - start
	if err != nil {
		log.Printf("Failed to send perp taker tx: error=%v, marketId=%d, side=%d, tif=%d, price=%s, quantity=%s\n",
			err, marketId, side, tif, price.String(), quantity.String())
		boomer.RecordFailure("http", "SendPerpTakerTxTC to "+endPoint, elapsed, err.Error())
		return
	}
	boomer.RecordSuccess("http", "SendPerpTakerTxTC to "+endPoint, elapsed, int64(10))
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}

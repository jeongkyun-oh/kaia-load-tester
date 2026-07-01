package perpHalfFillTxTC

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

const Name = "perpHalfFillTxTC"

// Perp order TimeInForce (see core/types dex_perp.go). GTC lets taker orders cross
// and execute; POST-only would reject them, so it is not used here.
const tifGTC = uint8(0)

// Perp order side (see core/types dex_perp.go).
const (
	sideBuy  = uint8(0)
	sideSell = uint8(1)
)

// Order role: half of the orders rest as makers, half cross as takers.
const (
	roleMaker = 0
	roleTaker = 1
)

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	cursor uint32

	// See perpNoTradeTxTC for the meaning of these; they are set from the shared
	// -perpMarketId / -perpRefPrice / -perpTickSize flags in main.go.
	marketId uint64 = 1
	refPrice        = scaleUp(100)
	tickSize        = scaleUp(1)
)

// SetMarketId overrides the target perp market id (from the -perpMarketId flag).
func SetMarketId(id uint64) {
	marketId = id
}

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

// Run maintains a two-level book around the (tick-aligned) reference price:
//
//	LOW  = ref - tick   HIGH = ref + tick
//
// A maker posts on the passive side (buy@LOW, sell@HIGH) and rests; a taker posts on
// the aggressive side (buy@HIGH, sell@LOW) and crosses whatever resting maker sits on
// the opposite edge (canMatchLimit: BUY crosses when price>=ask, SELL when price<=bid).
// Role is a 50/50 coin flip, so on average half the orders trade and half rest as makers.
func Run() {
	cli := cliPool.Alloc().(*ethclient.Client)
	defer cliPool.Free(cli)

	var (
		from     = accGrp[atomic.AddUint32(&cursor, 1)%uint32(nAcc)]
		quantity = scaleUp(1)
		side     = uint8(rand.Intn(2))
		role     = rand.Intn(2)

		refAligned = new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
		low        = new(big.Int).Sub(refAligned, tickSize)
		high       = new(big.Int).Add(refAligned, tickSize)
		price      *big.Int
	)

	switch {
	case role == roleMaker && side == sideBuy: // passive bid, rests below the ask
		price = low
	case role == roleMaker && side == sideSell: // passive ask, rests above the bid
		price = high
	case role == roleTaker && side == sideBuy: // lifts the resting ask at HIGH
		price = high
	default: // roleTaker && sideSell: hits the resting bid at LOW
		price = low
	}

	start := boomer.Now()
	tx, err := from.GenNewPerpOrderTx(marketId, side, price, quantity, tifGTC)
	if err != nil {
		log.Printf("Failed to generate perp order tx: error=%v, marketId=%d, side=%d, role=%d, price=%s, quantity=%s",
			err, marketId, side, role, price.String(), quantity.String())
		return
	}
	_, err = from.SendTx(cli, tx)
	elapsed := boomer.Now() - start
	if err != nil {
		log.Printf("Failed to send perp order tx: error=%v, marketId=%d, side=%d, role=%d, price=%s, quantity=%s\n",
			err, marketId, side, role, price.String(), quantity.String())
		boomer.RecordFailure("http", "SendHalfFillPerpOrderTx"+" to "+endPoint, elapsed, err.Error())
		return
	}

	boomer.RecordSuccess("http", "SendHalfFillPerpOrderTx"+" to "+endPoint, elapsed, int64(10))
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}

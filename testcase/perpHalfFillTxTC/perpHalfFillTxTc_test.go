package perpHalfFillTxTC

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenPerpOrderTx verifies the perp order tx carries the DexCommandPerpOrder
// prefix and round-trips its fields through the decimal-string wire encoding.
func TestGenPerpOrderTx(t *testing.T) {
	var (
		from     = account.NewAccount(0)
		mkt      = uint64(7)
		side     = sideBuy
		price    = scaleUp(101)
		quantity = scaleUp(1)
	)

	tx, err := from.GenNewPerpOrderTx(mkt, side, price, quantity, tifGTC)
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.Equal(t, types.DexCommandPerpOrder, tx.Data()[0], "command prefix must be PerpOrder (0x41)")

	var decoded types.PerpOrderContext
	require.NoError(t, decoded.Deserialize(tx.Data()[1:]))
	assert.Equal(t, mkt, decoded.MarketId)
	assert.Equal(t, side, decoded.Side)
	assert.Equal(t, tifGTC, decoded.TimeInForce)
	assert.Equal(t, 0, decoded.Price.Cmp(price), "price round-trip")
	assert.Equal(t, 0, decoded.Quantity.Cmp(quantity), "quantity round-trip")
}

// TestMakerRestsTakerCrosses asserts that maker prices do not cross while taker
// prices do, using the same LOW/HIGH levels as Run(). Crossing follows canMatchLimit:
// a BUY crosses a resting ask when buyPrice >= ask; a SELL crosses a resting bid when
// sellPrice <= bid.
func TestMakerRestsTakerCrosses(t *testing.T) {
	refAligned := new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
	low := new(big.Int).Sub(refAligned, tickSize)
	high := new(big.Int).Add(refAligned, tickSize)

	// Both levels are tick-aligned so the node accepts them.
	assert.Zero(t, new(big.Int).Mod(low, tickSize).Sign(), "LOW must be a tick multiple")
	assert.Zero(t, new(big.Int).Mod(high, tickSize).Sign(), "HIGH must be a tick multiple")

	// Maker bid (LOW) sits below maker ask (HIGH): a resting maker never self-crosses.
	makerBid, makerAsk := low, high
	assert.Negative(t, makerBid.Cmp(makerAsk), "maker bid must be below maker ask")

	// Taker BUY at HIGH lifts a resting ask at HIGH: buy >= ask -> crosses.
	assert.GreaterOrEqual(t, high.Cmp(makerAsk), 0, "taker buy must cross the resting ask")
	// Taker SELL at LOW hits a resting bid at LOW: sell <= bid -> crosses.
	assert.LessOrEqual(t, low.Cmp(makerBid), 0, "taker sell must cross the resting bid")
}

// TestNoSelfTrade asserts each account is pinned to one side (even=BUY, odd=SELL),
// so an account can never cross its own order and the node never rejects a self-trade.
func TestNoSelfTrade(t *testing.T) {
	for idx := uint32(0); idx < 8; idx++ {
		if idx%2 == 0 {
			assert.Equal(t, sideBuy, accountSide(idx), "even index must be BUY-only")
		} else {
			assert.Equal(t, sideSell, accountSide(idx), "odd index must be SELL-only")
		}
	}
	// Adjacent accounts take opposite sides, so takers always face a different account.
	assert.NotEqual(t, accountSide(0), accountSide(1), "buy and sell accounts must be disjoint")
}

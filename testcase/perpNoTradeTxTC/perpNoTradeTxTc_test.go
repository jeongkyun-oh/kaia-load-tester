package perpNoTradeTxTC

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenPerpOrderTx verifies a perp order tx carries the DexCommandPerpOrder
// prefix and that its JSON payload round-trips back into a PerpOrderContext with
// the same fields (price/quantity survive the decimal-string wire encoding).
func TestGenPerpOrderTx(t *testing.T) {
	var (
		from     = account.NewAccount(0)
		mkt      = uint64(7)
		side     = sideBuy
		price    = scaleUp(99)
		quantity = scaleUp(1)
	)

	tx, err := from.GenNewPerpOrderTx(mkt, side, price, quantity, tifPost)
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.True(t, len(tx.Data()) > 0)

	require.Equal(t, types.DexCommandPerpOrder, tx.Data()[0], "command prefix must be PerpOrder (0x41)")

	var decoded types.PerpOrderContext
	require.NoError(t, decoded.Deserialize(tx.Data()[1:]))

	assert.Equal(t, from.GetAddress(), decoded.L1Owner)
	assert.Equal(t, mkt, decoded.MarketId)
	assert.Equal(t, side, decoded.Side)
	assert.Equal(t, tifPost, decoded.TimeInForce)
	assert.Equal(t, 0, decoded.Price.Cmp(price), "price round-trip: got %s want %s", decoded.Price, price)
	assert.Equal(t, 0, decoded.Quantity.Cmp(quantity), "quantity round-trip")
}

// TestGenPerpDepositTx verifies the perp-margin funding tx carries the
// DexCommandPerpDeposit prefix and round-trips token/amount.
func TestGenPerpDepositTx(t *testing.T) {
	var (
		from   = account.NewAccount(0)
		token  = "2"
		amount = new(big.Int).Mul(big.NewInt(1e9), big.NewInt(1e18))
	)

	tx, err := from.GenPerpDepositTx(token, amount)
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.True(t, len(tx.Data()) > 0)

	require.Equal(t, types.DexCommandPerpDeposit, tx.Data()[0], "command prefix must be PerpDeposit (0x12)")

	var decoded types.PerpDepositContext
	require.NoError(t, decoded.Deserialize(tx.Data()[1:]))

	assert.Equal(t, from.GetAddress(), decoded.L1Owner)
	assert.Equal(t, token, decoded.Token)
	assert.Equal(t, 0, decoded.Amount.Cmp(amount), "amount round-trip")
}

// TestNoTradePricesDoNotCross asserts the buy price rests strictly below the sell
// price (no cross) and that both are exact multiples of the tick size (accepted).
func TestNoTradePricesDoNotCross(t *testing.T) {
	refAligned := new(big.Int).Mul(new(big.Int).Div(refPrice, tickSize), tickSize)
	buy := new(big.Int).Sub(refAligned, tickSize)
	sell := new(big.Int).Add(refAligned, tickSize)

	assert.Negative(t, buy.Cmp(sell), "buy (%s) must be below sell (%s)", buy, sell)
	assert.Zero(t, new(big.Int).Mod(buy, tickSize).Sign(), "buy must be a tick multiple")
	assert.Zero(t, new(big.Int).Mod(sell, tickSize).Sign(), "sell must be a tick multiple")
}

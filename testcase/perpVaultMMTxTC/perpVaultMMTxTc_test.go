package perpVaultMMTxTC

import (
	"math/big"
	"testing"
)

func TestNonCrossingPrice(t *testing.T) {
	refPrice = scaleUp(100)
	tickSize = scaleUp(1)
	buy := makerPrice(sideBuy)
	sell := makerPrice(sideSell)
	if buy.Cmp(sell) >= 0 {
		t.Fatalf("buy %s must be < sell %s (non-crossing)", buy, sell)
	}
	// tick-aligned: (price % tick) == 0
	if new(big.Int).Mod(buy, tickSize).Sign() != 0 {
		t.Fatalf("buy price %s not tick-aligned", buy)
	}
}

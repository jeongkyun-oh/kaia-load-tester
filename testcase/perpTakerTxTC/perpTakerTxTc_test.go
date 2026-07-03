package perpTakerTxTC

import (
	"math/big"
	"testing"
)

func TestTakerPriceCrosses(t *testing.T) {
	refPrice = scaleUp(100)
	tickSize = scaleUp(1)
	buy := takerPrice(sideBuy)
	sell := takerPrice(sideSell)
	// Taker buy must be >= taker sell so it crosses resting liquidity between them.
	if buy.Cmp(sell) <= 0 {
		t.Fatalf("taker buy %s must be > sell %s (crossing direction)", buy, sell)
	}
	if new(big.Int).Mod(buy, tickSize).Sign() != 0 {
		t.Fatalf("buy price %s not tick-aligned", buy)
	}
	if new(big.Int).Mod(sell, tickSize).Sign() != 0 {
		t.Fatalf("sell price %s not tick-aligned", sell)
	}
}

// TestTakerPriceSellStaysPositive covers refPrice <= tickSize: refAligned-tickSize
// would be zero or negative, which the node rejects; the price must be clamped to
// the smallest valid tick level instead.
func TestTakerPriceSellStaysPositive(t *testing.T) {
	defer func() {
		refPrice = scaleUp(100)
		tickSize = scaleUp(1)
	}()
	tickSize = scaleUp(10)
	for _, ref := range []*big.Int{scaleUp(10), scaleUp(5)} {
		refPrice = ref
		sell := takerPrice(sideSell)
		if sell.Sign() <= 0 {
			t.Fatalf("sell price %s must stay positive when refPrice=%s tickSize=%s",
				sell, refPrice, tickSize)
		}
		if new(big.Int).Mod(sell, tickSize).Sign() != 0 {
			t.Fatalf("sell price %s not tick-aligned", sell)
		}
	}
}

// TestAccountSidePinsPerAccount asserts the perpHalfFillTxTC invariant: each account
// slot is pinned to one side (even = BUY, odd = SELL), so an account can never cross
// its own resting GTC remainder and self-trades are impossible.
func TestAccountSidePinsPerAccount(t *testing.T) {
	for idx := uint32(0); idx < 8; idx++ {
		side := accountSide(idx)
		want := sideBuy
		if idx%2 == 1 {
			want = sideSell
		}
		if side != want {
			t.Fatalf("accountSide(%d)=%d, want %d", idx, side, want)
		}
		if side != accountSide(idx) {
			t.Fatalf("accountSide(%d) must be deterministic", idx)
		}
	}
}

func TestTakerTifsAreTakerOnly(t *testing.T) {
	want := map[uint8]bool{0: true, 1: true, 3: true} // GTC, IOC, MARKET
	if len(takerTifs) != len(want) {
		t.Fatalf("takerTifs=%v, want exactly GTC/IOC/MARKET", takerTifs)
	}
	for _, tf := range takerTifs {
		if tf == 2 {
			t.Fatalf("takerTifs must not contain POST (2)")
		}
		if !want[tf] {
			t.Fatalf("takerTifs contains non-taker or duplicate tif %d", tf)
		}
		delete(want, tf) // a duplicate entry would hit the branch above
	}
	if len(want) != 0 {
		t.Fatalf("takerTifs is missing tifs: %v still expected", want)
	}
}

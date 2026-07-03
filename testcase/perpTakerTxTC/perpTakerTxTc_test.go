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

func TestTakerTifsAreTakerOnly(t *testing.T) {
	want := map[uint8]bool{0: true, 1: true, 3: true} // GTC, IOC, MARKET
	if len(takerTifs) != len(want) {
		t.Fatalf("takerTifs=%v, want exactly GTC/IOC/MARKET", takerTifs)
	}
	for _, tf := range takerTifs {
		if !want[tf] {
			t.Fatalf("takerTifs contains non-taker tif %d (POST=2 must be excluded)", tf)
		}
		if tf == 2 {
			t.Fatalf("takerTifs must not contain POST (2)")
		}
	}
}

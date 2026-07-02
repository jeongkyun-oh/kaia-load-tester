package config

import (
	"math/big"
	"testing"
)

func TestVaultDepositAmountScaling(t *testing.T) {
	cfg := &Config{vaultDepositAmount: 5} // human units
	got := cfg.GetVaultDepositAmount()
	want := new(big.Int).Mul(big.NewInt(5), big.NewInt(1e18))
	if got.Cmp(want) != 0 {
		t.Fatalf("GetVaultDepositAmount()=%s want %s", got, want)
	}
}

func TestVaultDefaults(t *testing.T) {
	cfg := &Config{vaultExecutorCount: 10, vaultTokenId: "2"}
	if cfg.GetVaultExecutorCount() != 10 {
		t.Fatalf("executor count=%d want 10", cfg.GetVaultExecutorCount())
	}
	if cfg.GetVaultTokenId() != "2" {
		t.Fatalf("tokenId=%q want \"2\"", cfg.GetVaultTokenId())
	}
}

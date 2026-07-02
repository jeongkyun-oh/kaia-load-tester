package account

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGenPerpOrderTxBySession(t *testing.T) {
	SetChainID(big.NewInt(2018)) // ensure package chainID set (see existing setter)

	owner, _ := crypto.GenerateKey()
	ownerAcc := &Account{privateKey: owner}

	sessionKey, _ := crypto.GenerateKey()
	l1owner := common.HexToAddress("0x000000000000000000000000000000000000dEaD")

	tx, err := ownerAcc.GenPerpOrderTxBySession(sessionKey, l1owner, 1, 0, big.NewInt(1e18), big.NewInt(1e18), 2)
	if err != nil {
		t.Fatalf("GenPerpOrderTxBySession err=%v", err)
	}
	// Target must be the dex precompile.
	if tx.To() == nil || *tx.To() != types.DexAddress {
		t.Fatalf("tx.To()=%v want DexAddress", tx.To())
	}
	// Outer signature must recover to the SESSION key, not the owner.
	signer := types.LatestSignerForChainID(big.NewInt(2018))
	sender, err := types.Sender(signer, tx)
	if err != nil {
		t.Fatalf("Sender err=%v", err)
	}
	if sender != crypto.PubkeyToAddress(sessionKey.PublicKey) {
		t.Fatalf("sender=%v want session key addr", sender)
	}
}

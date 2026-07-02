package vaultgen

import (
	"encoding/json"
	_ "embed"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:embed perpvaultproxy.json
var perpVaultProxyArtifact []byte

type proxyArtifact struct {
	ABI      json.RawMessage `json:"abi"`
	Bytecode string          `json:"bytecode"`
}

// DeployPerpVaultProxy deploys an ERC1967-style proxy pointing at `implementation`,
// invoking `data` (an ABI-encoded initialize(...) call) in the constructor.
func DeployPerpVaultProxy(auth *bind.TransactOpts, backend bind.ContractBackend, implementation common.Address, data []byte) (common.Address, *types.Transaction, error) {
	var art proxyArtifact
	if err := json.Unmarshal(perpVaultProxyArtifact, &art); err != nil {
		return common.Address{}, nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(string(art.ABI)))
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, tx, _, err := bind.DeployContract(auth, parsed, common.FromHex(art.Bytecode), backend, implementation, data)
	if err != nil {
		return common.Address{}, nil, err
	}
	return addr, tx, nil
}

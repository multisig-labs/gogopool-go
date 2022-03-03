package tokens

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

//
// Core ERC-20 functions
//

// Get fixed-supply GGP total supply
func GetFixedSupplyGGPTotalSupply(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return nil, err
	}
	return totalSupply(gogoTokenFixedSupplyGGP, "fixed-supply GGP", opts)
}

// Get fixed-supply GGP balance
func GetFixedSupplyGGPBalance(ggp *gogopool.GoGoPool, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return nil, err
	}
	return balanceOf(gogoTokenFixedSupplyGGP, "fixed-supply GGP", address, opts)
}

// Get fixed-supply GGP allowance
func GetFixedSupplyGGPAllowance(ggp *gogopool.GoGoPool, owner, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return nil, err
	}
	return allowance(gogoTokenFixedSupplyGGP, "fixed-supply GGP", owner, spender, opts)
}

// Estimate the gas of TransferFixedSupplyGGP
func EstimateTransferFixedSupplyGGPGas(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferGas(gogoTokenFixedSupplyGGP, "fixed-supply GGP", to, amount, opts)
}

// Transfer fixed-supply GGP
func TransferFixedSupplyGGP(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transfer(gogoTokenFixedSupplyGGP, "fixed-supply GGP", to, amount, opts)
}

// Estimate the gas of ApproveFixedSupplyGGP
func EstimateApproveFixedSupplyGGPGas(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateApproveGas(gogoTokenFixedSupplyGGP, "fixed-supply GGP", spender, amount, opts)
}

// Approve an fixed-supply GGP spender
func ApproveFixedSupplyGGP(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return approve(gogoTokenFixedSupplyGGP, "fixed-supply GGP", spender, amount, opts)
}

// Estimate the gas of TransferFromFixedSupplyGGP
func EstimateTransferFromFixedSupplyGGPGas(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferFromGas(gogoTokenFixedSupplyGGP, "fixed-supply GGP", from, to, amount, opts)
}

// Transfer fixed-supply GGP from a sender
func TransferFromFixedSupplyGGP(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenFixedSupplyGGP, err := getGoGoTokenGGPFixedSupply(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transferFrom(gogoTokenFixedSupplyGGP, "fixed-supply GGP", from, to, amount, opts)
}

//
// Contracts
//

// Get contracts
var gogoTokenFixedSupplyGGPLock sync.Mutex

func getGoGoTokenGGPFixedSupply(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoTokenFixedSupplyGGPLock.Lock()
	defer gogoTokenFixedSupplyGGPLock.Unlock()
	return ggp.GetContract("rocketTokenGGPFixedSupply")
}

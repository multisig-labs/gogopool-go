package tokens

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

//
// Core ERC-20 functions
//

// Get GGP total supply
func GetGGPTotalSupply(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return nil, err
	}
	return totalSupply(gogoTokenGGP, "GGP", opts)
}

// Get GGP balance
func GetGGPBalance(ggp *gogopool.GoGoPool, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return nil, err
	}
	return balanceOf(gogoTokenGGP, "GGP", address, opts)
}

// Get GGP allowance
func GetGGPAllowance(ggp *gogopool.GoGoPool, owner, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return nil, err
	}
	return allowance(gogoTokenGGP, "GGP", owner, spender, opts)
}

// Estimate the gas of TransferGGP
func EstimateTransferGGPGas(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferGas(gogoTokenGGP, "GGP", to, amount, opts)
}

// Transfer GGP
func TransferGGP(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transfer(gogoTokenGGP, "GGP", to, amount, opts)
}

// Estimate the gas of ApproveGGP
func EstimateApproveGGPGas(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateApproveGas(gogoTokenGGP, "GGP", spender, amount, opts)
}

// Approve an GGP spender
func ApproveGGP(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return approve(gogoTokenGGP, "GGP", spender, amount, opts)
}

// Estimate the gas of TransferFromGGP
func EstimateTransferFromGGPGas(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferFromGas(gogoTokenGGP, "GGP", from, to, amount, opts)
}

// Transfer GGP from a sender
func TransferFromGGP(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transferFrom(gogoTokenGGP, "GGP", from, to, amount, opts)
}

//
// GGP functions
//

// Estimate the gas of MintInflationGGP
func EstimateMintInflationGGPGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoTokenGGP.GetTransactionGasInfo(opts, "inflationMintTokens")
}

// Mint new GGP tokens from inflation
func MintInflationGGP(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoTokenGGP.Transact(opts, "inflationMintTokens")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not mint GGP tokens from inflation: %w", err)
	}
	return hash, nil
}

// Estimate the gas of SwapFixedSupplyGGPForGGP
func EstimateSwapFixedSupplyGGPForGGPGas(ggp *gogopool.GoGoPool, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoTokenGGP.GetTransactionGasInfo(opts, "swapTokens", amount)
}

// Swap fixed-supply GGP for new GGP tokens
func SwapFixedSupplyGGPForGGP(ggp *gogopool.GoGoPool, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoTokenGGP.Transact(opts, "swapTokens", amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not swap fixed-supply GGP for new GGP: %w", err)
	}
	return hash, nil
}

// Get the GGP inflation interval rate
func GetGGPInflationIntervalRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenGGP, err := getGoGoTokenGGP(ggp)
	if err != nil {
		return nil, err
	}
	rate := new(*big.Int)
	if err := gogoTokenGGP.Call(opts, rate, "getInflationIntervalRate"); err != nil {
		return nil, fmt.Errorf("Could not get GGP inflation interval rate: %w", err)
	}
	return *rate, nil
}

//
// Contracts
//

// Get contracts
var gogoTokenGGPLock sync.Mutex

func getGoGoTokenGGP(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoTokenGGPLock.Lock()
	defer gogoTokenGGPLock.Unlock()
	return ggp.GetContract("gogoTokenGGP")
}

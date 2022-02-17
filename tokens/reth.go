package tokens

import (
	"fmt"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/utils/avax"
)

//
// Core ERC-20 functions
//

// Get rETH total supply
func GetRETHTotalSupply(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	return totalSupply(gogoTokenRETH, "rETH", opts)
}

// Get rETH balance
func GetRETHBalance(ggp *gogopool.GoGoPool, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	return balanceOf(gogoTokenRETH, "rETH", address, opts)
}

// Get rETH allowance
func GetRETHAllowance(ggp *gogopool.GoGoPool, owner, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	return allowance(gogoTokenRETH, "rETH", owner, spender, opts)
}

// Estimate the gas of TransferRETH
func EstimateTransferRETHGas(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferGas(gogoTokenRETH, "rETH", to, amount, opts)
}

// Transfer rETH
func TransferRETH(ggp *gogopool.GoGoPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transfer(gogoTokenRETH, "rETH", to, amount, opts)
}

// Estimate the gas of ApproveRETH
func EstimateApproveRETHGas(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateApproveGas(gogoTokenRETH, "rETH", spender, amount, opts)
}

// Approve a rETH spender
func ApproveRETH(ggp *gogopool.GoGoPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return approve(gogoTokenRETH, "rETH", spender, amount, opts)
}

// Estimate the gas of TransferFromRETH
func EstimateTransferFromRETHGas(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateTransferFromGas(gogoTokenRETH, "rETH", from, to, amount, opts)
}

// Transfer rETH from a sender
func TransferFromRETH(ggp *gogopool.GoGoPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return transferFrom(gogoTokenRETH, "rETH", from, to, amount, opts)
}

//
// rETH functions
//

// Get the rETH contract ETH balance
func GetRETHContractETHBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	return contractETHBalance(ggp, gogoTokenRETH, opts)
}

// Get the ETH value of an amount of rETH
func GetETHValueOfRETH(ggp *gogopool.GoGoPool, rethAmount *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	GoGoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	ethValue := new(*big.Int)
	if err := GoGoTokenRETH.Call(opts, ethValue, "getEthValue", rethAmount); err != nil {
		return nil, fmt.Errorf("Could not get ETH value of rETH amount: %w", err)
	}
	return *ethValue, nil
}

// Get the rETH value of an amount of ETH
func GetRETHValueOfETH(ggp *gogopool.GoGoPool, ethAmount *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	rethValue := new(*big.Int)
	if err := gogoTokenRETH.Call(opts, rethValue, "getRethValue", ethAmount); err != nil {
		return nil, fmt.Errorf("Could not get rETH value of ETH amount: %w", err)
	}
	return *rethValue, nil
}

// Get the current ETH : rETH exchange rate
func GetRETHExchangeRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return 0, err
	}
	exchangeRate := new(*big.Int)
	if err := gogoTokenRETH.Call(opts, exchangeRate, "getExchangeRate"); err != nil {
		return 0, fmt.Errorf("Could not get rETH exchange rate: %w", err)
	}
	return avax.WeiToEth(*exchangeRate), nil
}

// Get the total amount of ETH collateral available for rETH trades
func GetRETHTotalCollateral(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return nil, err
	}
	totalCollateral := new(*big.Int)
	if err := gogoTokenRETH.Call(opts, totalCollateral, "getTotalCollateral"); err != nil {
		return nil, fmt.Errorf("Could not get rETH total collateral: %w", err)
	}
	return *totalCollateral, nil
}

// Get the rETH collateralization rate
func GetRETHCollateralRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return 0, err
	}
	collateralRate := new(*big.Int)
	if err := gogoTokenRETH.Call(opts, collateralRate, "getCollateralRate"); err != nil {
		return 0, fmt.Errorf("Could not get rETH collateral rate: %w", err)
	}
	return avax.WeiToEth(*collateralRate), nil
}

// Estimate the gas of BurnRETH
func EstimateBurnRETHGas(ggp *gogopool.GoGoPool, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoTokenRETH.GetTransactionGasInfo(opts, "burn", amount)
}

// Burn rETH for ETH
func BurnRETH(ggp *gogopool.GoGoPool, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoTokenRETH, err := getGoGoTokenRETH(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoTokenRETH.Transact(opts, "burn", amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not burn rETH: %w", err)
	}
	return hash, nil
}

//
// Contracts
//

// Get contracts
var gogoTokenRETHLock sync.Mutex

func getGoGoTokenRETH(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoTokenRETHLock.Lock()
	defer gogoTokenRETHLock.Unlock()
	return ggp.GetContract("gogoTokenRETH")
}

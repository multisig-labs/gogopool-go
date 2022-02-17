package tokens

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Token balances
type Balances struct {
	AVAX           *big.Int `json:"eth"`
	GAVAX          *big.Int `json:"reth"`
	GGP            *big.Int `json:"ggp"`
	FixedSupplyGGP *big.Int `json:"fixedSupplyGgp"`
}

// Get token balances of an address
func GetBalances(ggp *gogopool.GoGoPool, address common.Address, opts *bind.CallOpts) (Balances, error) {

	// Get call options block number
	var blockNumber *big.Int
	if opts != nil {
		blockNumber = opts.BlockNumber
	}

	// Data
	var wg errgroup.Group
	var avaxBalance *big.Int
	var gavaxBalance *big.Int
	var ggpBalance *big.Int
	var fixedSupplyGgpBalance *big.Int

	// Load data
	wg.Go(func() error {
		var err error
		avaxBalance, err = ggp.Client.BalanceAt(context.Background(), address, blockNumber)
		return err
	})
	wg.Go(func() error {
		var err error
		gavaxBalance, err = GetRETHBalance(ggp, address, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		ggpBalance, err = GetGGPBalance(ggp, address, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		fixedSupplyGgpBalance, err = GetFixedSupplyGGPBalance(ggp, address, opts)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return Balances{}, err
	}

	// Return
	return Balances{
		AVAX:           avaxBalance,
		GAVAX:          gavaxBalance,
		GGP:            ggpBalance,
		FixedSupplyGGP: fixedSupplyGgpBalance,
	}, nil

}

// Get a token contract's ETH balance
func contractETHBalance(ggp *gogopool.GoGoPool, tokenContract *gogopool.Contract, opts *bind.CallOpts) (*big.Int, error) {
	var blockNumber *big.Int
	if opts != nil {
		blockNumber = opts.BlockNumber
	}
	return ggp.Client.BalanceAt(context.Background(), *(tokenContract.Address), blockNumber)
}

// Get a token's total supply
func totalSupply(tokenContract *gogopool.Contract, tokenName string, opts *bind.CallOpts) (*big.Int, error) {
	totalSupply := new(*big.Int)
	if err := tokenContract.Call(opts, totalSupply, "totalSupply"); err != nil {
		return nil, fmt.Errorf("Could not get %s total supply: %w", tokenName, err)
	}
	return *totalSupply, nil
}

// Get a token balance
func balanceOf(tokenContract *gogopool.Contract, tokenName string, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	balance := new(*big.Int)
	if err := tokenContract.Call(opts, balance, "balanceOf", address); err != nil {
		return nil, fmt.Errorf("Could not get %s balance of %s: %w", tokenName, address.Hex(), err)
	}
	return *balance, nil
}

// Get a spender's allowance for an address
func allowance(tokenContract *gogopool.Contract, tokenName string, owner, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
	allowance := new(*big.Int)
	if err := tokenContract.Call(opts, allowance, "allowance", owner, spender); err != nil {
		return nil, fmt.Errorf("Could not get %s allowance of %s for %s: %w", tokenName, spender.Hex(), owner.Hex(), err)
	}
	return *allowance, nil
}

// Estimate the gas of transfer
func estimateTransferGas(tokenContract *gogopool.Contract, tokenName string, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return tokenContract.GetTransactionGasInfo(opts, "transfer", to, amount)
}

// Transfer tokens to an address
func transfer(tokenContract *gogopool.Contract, tokenName string, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := tokenContract.Transact(opts, "transfer", to, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not transfer %s to %s: %w", tokenName, to.Hex(), err)
	}
	return hash, nil
}

// Estimate the gas of approve
func estimateApproveGas(tokenContract *gogopool.Contract, tokenName string, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return tokenContract.GetTransactionGasInfo(opts, "approve", spender, amount)
}

// Approve a token allowance for a spender
func approve(tokenContract *gogopool.Contract, tokenName string, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := tokenContract.Transact(opts, "approve", spender, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not approve %s allowance for %s: %w", tokenName, spender.Hex(), err)
	}
	return hash, nil
}

// Estimate the gas of transferFrom
func estimateTransferFromGas(tokenContract *gogopool.Contract, tokenName string, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return tokenContract.GetTransactionGasInfo(opts, "transferFrom", from, to, amount)
}

// Transfer tokens from a sender to an address
func transferFrom(tokenContract *gogopool.Contract, tokenName string, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := tokenContract.Transact(opts, "transferFrom", from, to, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not transfer %s from %s to %s: %w", tokenName, from.Hex(), to.Hex(), err)
	}
	return hash, nil
}

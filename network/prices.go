package network

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Get the block number which network prices are current for
func GetPricesBlock(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return 0, err
	}
	pricesBlock := new(*big.Int)
	if err := gogoNetworkPrices.Call(opts, pricesBlock, "getPricesBlock"); err != nil {
		return 0, fmt.Errorf("Could not get network prices block: %w", err)
	}
	return (*pricesBlock).Uint64(), nil
}

// Get the current network GGP price in ETH
func GetGGPPrice(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return nil, err
	}
	ggpPrice := new(*big.Int)
	if err := gogoNetworkPrices.Call(opts, ggpPrice, "getGGPPrice"); err != nil {
		return nil, fmt.Errorf("Could not get network GGP price: %w", err)
	}
	return *ggpPrice, nil
}

// Estimate the gas of SubmitPrices
func EstimateSubmitPricesGas(ggp *gogopool.GoGoPool, block uint64, ggpPrice *big.Int, effectiveGgpStake *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoNetworkPrices.GetTransactionGasInfo(opts, "submitPrices", big.NewInt(int64(block)), ggpPrice, effectiveGgpStake)
}

// Submit network prices and total effective GGP stake for an epoch
func SubmitPrices(ggp *gogopool.GoGoPool, block uint64, ggpPrice, effectiveGgpStake *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoNetworkPrices.Transact(opts, "submitPrices", big.NewInt(int64(block)), ggpPrice, effectiveGgpStake)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not submit network prices: %w", err)
	}
	return hash, nil
}

// Check if the network is currently in consensus about the GGP price, or if it is still reaching consensus
func InConsensus(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return false, err
	}
	isInConsensus := new(bool)
	if err := gogoNetworkPrices.Call(opts, isInConsensus, "inConsensus"); err != nil {
		return false, fmt.Errorf("Could not get consensus status: %w", err)
	}
	return *isInConsensus, nil
}

// Returns the latest block number that oracles should be reporting prices for
func GetLatestReportablePricesBlock(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return nil, err
	}
	latestReportableBlock := new(*big.Int)
	if err := gogoNetworkPrices.Call(opts, latestReportableBlock, "getLatestReportableBlock"); err != nil {
		return nil, fmt.Errorf("Could not get latest reportable block: %w", err)
	}
	return *latestReportableBlock, nil
}

// Get contracts
var gogoNetworkPricesLock sync.Mutex

func getGoGoNetworkPrices(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNetworkPricesLock.Lock()
	defer gogoNetworkPricesLock.Unlock()
	return ggp.GetContract("gogoNetworkPrices")
}

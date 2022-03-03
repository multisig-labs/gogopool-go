package network

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Get the block number which network balances are current for
func GetBalancesBlock(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return 0, err
	}
	balancesBlock := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, balancesBlock, "getBalancesBlock"); err != nil {
		return 0, fmt.Errorf("Could not get network balances block: %w", err)
	}
	return (*balancesBlock).Uint64(), nil
}

// Get the current network total ETH balance
func GetTotalETHBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	totalEthBalance := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, totalEthBalance, "getTotalETHBalance"); err != nil {
		return nil, fmt.Errorf("Could not get network total ETH balance: %w", err)
	}
	return *totalEthBalance, nil
}

// Get the current network staking ETH balance
func GetStakingETHBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	stakingEthBalance := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, stakingEthBalance, "getStakingETHBalance"); err != nil {
		return nil, fmt.Errorf("Could not get network staking ETH balance: %w", err)
	}
	return *stakingEthBalance, nil
}

// Get the current network total rETH supply
func GetTotalRETHSupply(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	totalRethSupply := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, totalRethSupply, "getTotalRETHSupply"); err != nil {
		return nil, fmt.Errorf("Could not get network total rETH supply: %w", err)
	}
	return *totalRethSupply, nil
}

// Get the current network ETH utilization rate
func GetETHUtilizationRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return 0, err
	}
	ethUtilizationRate := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, ethUtilizationRate, "getETHUtilizationRate"); err != nil {
		return 0, fmt.Errorf("Could not get network ETH utilization rate: %w", err)
	}
	return avax.WeiToEth(*ethUtilizationRate), nil
}

// Estimate the gas of SubmitBalances
func EstimateSubmitBalancesGas(ggp *gogopool.GoGoPool, block uint64, totalEth, stakingEth, rethSupply *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoNetworkBalances.GetTransactionGasInfo(opts, "submitBalances", big.NewInt(int64(block)), totalEth, stakingEth, rethSupply)
}

// Submit network balances for an epoch
func SubmitBalances(ggp *gogopool.GoGoPool, block uint64, totalEth, stakingEth, rethSupply *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoNetworkBalances.Transact(opts, "submitBalances", big.NewInt(int64(block)), totalEth, stakingEth, rethSupply)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not submit network balances: %w", err)
	}
	return hash, nil
}

// Returns the latest block number that oracles should be reporting balances for
func GetLatestReportableBalancesBlock(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	latestReportableBlock := new(*big.Int)
	if err := gogoNetworkBalances.Call(opts, latestReportableBlock, "getLatestReportableBlock"); err != nil {
		return nil, fmt.Errorf("Could not get latest reportable block: %w", err)
	}
	return *latestReportableBlock, nil
}

// Get contracts
var gogoNetworkBalancesLock sync.Mutex

func getGoGoNetworkBalances(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNetworkBalancesLock.Lock()
	defer gogoNetworkBalancesLock.Unlock()
	return ggp.GetContract("rocketNetworkBalances")
}

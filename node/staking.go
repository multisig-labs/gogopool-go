package node

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Get the total GGP staked in the network
func GetTotalGGPStake(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	totalGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, totalGgpStake, "getTotalGGPStake"); err != nil {
		return nil, fmt.Errorf("Could not get total network GGP stake: %w", err)
	}
	return *totalGgpStake, nil
}

// Get the effective GGP staked in the network
func GetTotalEffectiveGGPStake(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	totalEffectiveGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, totalEffectiveGgpStake, "getTotalEffectiveGGPStake"); err != nil {
		return nil, fmt.Errorf("Could not get effective network GGP stake: %w", err)
	}
	return *totalEffectiveGgpStake, nil
}

// Get a node's GGP stake
func GetNodeGGPStake(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	nodeGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, nodeGgpStake, "getNodeGGPStake", nodeAddress); err != nil {
		return nil, fmt.Errorf("Could not get total node GGP stake: %w", err)
	}
	return *nodeGgpStake, nil
}

// Get a node's effective GGP stake
func GetNodeEffectiveGGPStake(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	nodeEffectiveGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, nodeEffectiveGgpStake, "getNodeEffectiveGGPStake", nodeAddress); err != nil {
		return nil, fmt.Errorf("Could not get effective node GGP stake: %w", err)
	}
	return *nodeEffectiveGgpStake, nil
}

// Get a node's minimum GGP stake to collateralize their minipools
func GetNodeMinimumGGPStake(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	nodeMinimumGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, nodeMinimumGgpStake, "getNodeMinimumGGPStake", nodeAddress); err != nil {
		return nil, fmt.Errorf("Could not get minimum node GGP stake: %w", err)
	}
	return *nodeMinimumGgpStake, nil
}

// Get a node's maximum GGP stake to collateralize their minipools
func GetNodeMaximumGGPStake(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	nodeMaximumGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, nodeMaximumGgpStake, "getNodeMaximumGGPStake", nodeAddress); err != nil {
		return nil, fmt.Errorf("Could not get maximum node GGP stake: %w", err)
	}
	return *nodeMaximumGgpStake, nil
}

// Get the time a node last staked GGP
func GetNodeGGPStakedTime(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return 0, err
	}
	nodeGgpStakedTime := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, nodeGgpStakedTime, "getNodeGGPStakedTime", nodeAddress); err != nil {
		return 0, fmt.Errorf("Could not get node GGP staked time: %w", err)
	}
	return (*nodeGgpStakedTime).Uint64(), nil
}

// Get a node's minipool limit based on GGP stake
func GetNodeMinipoolLimit(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return 0, err
	}
	minipoolLimit := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, minipoolLimit, "getNodeMinipoolLimit", nodeAddress); err != nil {
		return 0, fmt.Errorf("Could not get node minipool limit: %w", err)
	}
	return (*minipoolLimit).Uint64(), nil
}

// Estimate the gas of Stake
func EstimateStakeGas(ggp *gogopool.GoGoPool, ggpAmount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoNodeStaking.GetTransactionGasInfo(opts, "stakeGGP", ggpAmount)
}

// Stake GGP
func StakeGGP(ggp *gogopool.GoGoPool, ggpAmount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoNodeStaking.Transact(opts, "stakeGGP", ggpAmount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not stake GGP: %w", err)
	}
	return hash, nil
}

// Estimate the gas of WithdrawGGP
func EstimateWithdrawGGPGas(ggp *gogopool.GoGoPool, ggpAmount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoNodeStaking.GetTransactionGasInfo(opts, "withdrawGGP", ggpAmount)
}

// Withdraw staked GGP
func WithdrawGGP(ggp *gogopool.GoGoPool, ggpAmount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoNodeStaking.Transact(opts, "withdrawGGP", ggpAmount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not withdraw staked GGP: %w", err)
	}
	return hash, nil
}

// Calculate total effective GGP stake
func CalculateTotalEffectiveGGPStake(ggp *gogopool.GoGoPool, offset, limit, ggpPrice *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	gogoNodeStaking, err := getGoGoNodeStaking(ggp)
	if err != nil {
		return nil, err
	}
	totalEffectiveGgpStake := new(*big.Int)
	if err := gogoNodeStaking.Call(opts, totalEffectiveGgpStake, "calculateTotalEffectiveGGPStake", offset, limit, ggpPrice); err != nil {
		return nil, fmt.Errorf("Could not get total effective GGP stake: %w", err)
	}
	return *totalEffectiveGgpStake, nil
}

// Get contracts
var gogoNodeStakingLock sync.Mutex

func getGoGoNodeStaking(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNodeStakingLock.Lock()
	defer gogoNodeStakingLock.Unlock()
	return ggp.GetContract("rocketNodeStaking")
}

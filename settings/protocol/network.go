package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Config
const NetworkSettingsContractName = "rocketDAOProtocolSettingsNetwork"

// The threshold of trusted nodes that must reach consensus on oracle data to commit it
func GetNodeConsensusThreshold(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getNodeConsensusThreshold"); err != nil {
		return 0, fmt.Errorf("Could not get trusted node consensus threshold: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapNodeConsensusThreshold(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.consensus.threshold", avax.EthToWei(value), opts)
}

// Network balance submissions currently enabled
func GetSubmitBalancesEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := networkSettingsContract.Call(opts, value, "getSubmitBalancesEnabled"); err != nil {
		return false, fmt.Errorf("Could not get network balance submissions enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapSubmitBalancesEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, NetworkSettingsContractName, "network.submit.balances.enabled", value, opts)
}

// The frequency in blocks at which network balances should be submitted by trusted nodes
func GetSubmitBalancesFrequency(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getSubmitBalancesFrequency"); err != nil {
		return 0, fmt.Errorf("Could not get network balance submission frequency: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapSubmitBalancesFrequency(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.submit.balances.frequency", big.NewInt(int64(value)), opts)
}

// Network price submissions currently enabled
func GetSubmitPricesEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := networkSettingsContract.Call(opts, value, "getSubmitPricesEnabled"); err != nil {
		return false, fmt.Errorf("Could not get network price submissions enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapSubmitPricesEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, NetworkSettingsContractName, "network.submit.prices.enabled", value, opts)
}

// The frequency in blocks at which network prices should be submitted by trusted nodes
func GetSubmitPricesFrequency(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getSubmitPricesFrequency"); err != nil {
		return 0, fmt.Errorf("Could not get network price submission frequency: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapSubmitPricesFrequency(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.submit.prices.frequency", big.NewInt(int64(value)), opts)
}

// Minimum node commission rate
func GetMinimumNodeFee(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getMinimumNodeFee"); err != nil {
		return 0, fmt.Errorf("Could not get minimum node fee: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapMinimumNodeFee(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.node.fee.minimum", avax.EthToWei(value), opts)
}

// Target node commission rate
func GetTargetNodeFee(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getTargetNodeFee"); err != nil {
		return 0, fmt.Errorf("Could not get target node fee: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapTargetNodeFee(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.node.fee.target", avax.EthToWei(value), opts)
}

// Maximum node commission rate
func GetMaximumNodeFee(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getMaximumNodeFee"); err != nil {
		return 0, fmt.Errorf("Could not get maximum node fee: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapMaximumNodeFee(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.node.fee.maximum", avax.EthToWei(value), opts)
}

// The range of node demand values to base fee calculations on
func GetNodeFeeDemandRange(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getNodeFeeDemandRange"); err != nil {
		return nil, fmt.Errorf("Could not get node fee demand range: %w", err)
	}
	return *value, nil
}
func BootstrapNodeFeeDemandRange(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.node.fee.demand.range", value, opts)
}

// The target collateralization rate for the rETH contract as a fraction
func GetTargetRethCollateralRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	networkSettingsContract, err := getNetworkSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := networkSettingsContract.Call(opts, value, "getTargetRethCollateralRate"); err != nil {
		return 0, fmt.Errorf("Could not get target rETH contract collateralization rate: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapTargetRethCollateralRate(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NetworkSettingsContractName, "network.reth.collateral.target", avax.EthToWei(value), opts)
}

// Get contracts
var networkSettingsContractLock sync.Mutex

func getNetworkSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	networkSettingsContractLock.Lock()
	defer networkSettingsContractLock.Unlock()
	return ggp.GetContract(NetworkSettingsContractName)
}

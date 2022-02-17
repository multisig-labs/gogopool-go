package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// Config
const NodeSettingsContractName = "gogoDAOProtocolSettingsNode"

// Node registrations currently enabled
func GetNodeRegistrationEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	nodeSettingsContract, err := getNodeSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := nodeSettingsContract.Call(opts, value, "getRegistrationEnabled"); err != nil {
		return false, fmt.Errorf("Could not get node registrations enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapNodeRegistrationEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, NodeSettingsContractName, "node.registration.enabled", value, opts)
}

// Node deposits currently enabled
func GetNodeDepositEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	nodeSettingsContract, err := getNodeSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := nodeSettingsContract.Call(opts, value, "getDepositEnabled"); err != nil {
		return false, fmt.Errorf("Could not get node deposits enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapNodeDepositEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, NodeSettingsContractName, "node.deposit.enabled", value, opts)
}

// The minimum GGP stake per minipool as a fraction of assigned user ETH
func GetMinimumPerMinipoolStake(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	nodeSettingsContract, err := getNodeSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMinimumPerMinipoolStake"); err != nil {
		return 0, fmt.Errorf("Could not get minimum GGP stake per minipool: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapMinimumPerMinipoolStake(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NodeSettingsContractName, "node.per.minipool.stake.minimum", avax.EthToWei(value), opts)
}

// The maximum GGP stake per minipool as a fraction of assigned user ETH
func GetMaximumPerMinipoolStake(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	nodeSettingsContract, err := getNodeSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMaximumPerMinipoolStake"); err != nil {
		return 0, fmt.Errorf("Could not get maximum GGP stake per minipool: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapMaximumPerMinipoolStake(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, NodeSettingsContractName, "node.per.minipool.stake.maximum", avax.EthToWei(value), opts)
}

// Get contracts
var nodeSettingsContractLock sync.Mutex

func getNodeSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	nodeSettingsContractLock.Lock()
	defer nodeSettingsContractLock.Unlock()
	return ggp.GetContract(NodeSettingsContractName)
}

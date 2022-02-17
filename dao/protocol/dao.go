package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// Estimate the gas of BootstrapBool
func EstimateBootstrapBoolGas(ggp *gogopool.GoGoPool, contractName, settingPath string, value bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingBool", contractName, settingPath, value)
}

// Bootstrap a bool setting
func BootstrapBool(ggp *gogopool.GoGoPool, contractName, settingPath string, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAOProtocol.Transact(opts, "bootstrapSettingBool", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapUint
func EstimateBootstrapUintGas(ggp *gogopool.GoGoPool, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingUint", contractName, settingPath, value)
}

// Bootstrap a uint256 setting
func BootstrapUint(ggp *gogopool.GoGoPool, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAOProtocol.Transact(opts, "bootstrapSettingUint", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapAddress
func EstimateBootstrapAddressGas(ggp *gogopool.GoGoPool, contractName, settingPath string, value common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingAddress", contractName, settingPath, value)
}

// Bootstrap an address setting
func BootstrapAddress(ggp *gogopool.GoGoPool, contractName, settingPath string, value common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAOProtocol.Transact(opts, "bootstrapSettingAddress", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapClaimer
func EstimateBootstrapClaimerGas(ggp *gogopool.GoGoPool, contractName string, amount float64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingClaimer", contractName, avax.EthToWei(amount))
}

// Bootstrap a rewards claimer
func BootstrapClaimer(ggp *gogopool.GoGoPool, contractName string, amount float64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAOProtocol, err := getRocketDAOProtocol(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAOProtocol.Transact(opts, "bootstrapSettingClaimer", contractName, avax.EthToWei(amount))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol rewards claimer %s: %w", contractName, err)
	}
	return hash, nil
}

// Get contracts
var gogoDAOProtocolLock sync.Mutex

func getRocketDAOProtocol(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDAOProtocolLock.Lock()
	defer gogoDAOProtocolLock.Unlock()
	return ggp.GetContract("gogoDAOProtocol")
}

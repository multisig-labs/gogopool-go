package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Config
const DepositSettingsContractName = "gogoDAOProtocolSettingsDeposit"

// Deposits currently enabled
func GetDepositEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	depositSettingsContract, err := getDepositSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := depositSettingsContract.Call(opts, value, "getDepositEnabled"); err != nil {
		return false, fmt.Errorf("Could not get deposits enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapDepositEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, DepositSettingsContractName, "deposit.enabled", value, opts)
}

// Deposit assignments currently enabled
func GetAssignDepositsEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	depositSettingsContract, err := getDepositSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := depositSettingsContract.Call(opts, value, "getAssignDepositsEnabled"); err != nil {
		return false, fmt.Errorf("Could not get deposit assignments enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapAssignDepositsEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, DepositSettingsContractName, "deposit.assign.enabled", value, opts)
}

// Minimum deposit amount
func GetMinimumDeposit(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	depositSettingsContract, err := getDepositSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := depositSettingsContract.Call(opts, value, "getMinimumDeposit"); err != nil {
		return nil, fmt.Errorf("Could not get minimum deposit amount: %w", err)
	}
	return *value, nil
}
func BootstrapMinimumDeposit(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, DepositSettingsContractName, "deposit.minimum", value, opts)
}

// Maximum deposit pool size
func GetMaximumDepositPoolSize(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	depositSettingsContract, err := getDepositSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := depositSettingsContract.Call(opts, value, "getMaximumDepositPoolSize"); err != nil {
		return nil, fmt.Errorf("Could not get maximum deposit pool size: %w", err)
	}
	return *value, nil
}
func BootstrapMaximumDepositPoolSize(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, DepositSettingsContractName, "deposit.pool.maximum", value, opts)
}

// Maximum deposit assignments per transaction
func GetMaximumDepositAssignments(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	depositSettingsContract, err := getDepositSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := depositSettingsContract.Call(opts, value, "getMaximumDepositAssignments"); err != nil {
		return 0, fmt.Errorf("Could not get maximum deposit assignments: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapMaximumDepositAssignments(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, DepositSettingsContractName, "deposit.assign.maximum", big.NewInt(int64(value)), opts)
}

// Get contracts
var depositSettingsContractLock sync.Mutex

func getDepositSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	depositSettingsContractLock.Lock()
	defer depositSettingsContractLock.Unlock()
	return ggp.GetContract(DepositSettingsContractName)
}

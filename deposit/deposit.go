package deposit

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Get the deposit pool balance
func GetBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return nil, err
	}
	balance := new(*big.Int)
	if err := gogoDepositPool.Call(opts, balance, "getBalance"); err != nil {
		return nil, fmt.Errorf("Could not get deposit pool balance: %w", err)
	}
	return *balance, nil
}

// Get the excess deposit pool balance
func GetExcessBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return nil, err
	}
	excessBalance := new(*big.Int)
	if err := gogoDepositPool.Call(opts, excessBalance, "getExcessBalance"); err != nil {
		return nil, fmt.Errorf("Could not get deposit pool excess balance: %w", err)
	}
	return *excessBalance, nil
}

// Estimate the gas of Deposit
func EstimateDepositGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDepositPool.GetTransactionGasInfo(opts, "deposit")
}

// Make a deposit
func Deposit(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDepositPool.Transact(opts, "deposit")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not deposit: %w", err)
	}
	return hash, nil
}

// Estimate the gas of AssignDeposits
func EstimateAssignDepositsGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDepositPool.GetTransactionGasInfo(opts, "assignDeposits")
}

// Assign deposits
func AssignDeposits(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDepositPool, err := getGoGoDepositPool(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDepositPool.Transact(opts, "assignDeposits")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not assign deposits: %w", err)
	}
	return hash, nil
}

// Get contracts
var gogoDepositPoolLock sync.Mutex

func getGoGoDepositPool(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDepositPoolLock.Lock()
	defer gogoDepositPoolLock.Unlock()
	return ggp.GetContract("rocketDepositPool")
}

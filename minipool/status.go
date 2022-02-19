package minipool

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Estimate the gas of SubmitMinipoolWithdrawable
func EstimateSubmitMinipoolWithdrawableGas(ggp *gogopool.GoGoPool, minipoolAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoMinipoolStatus, err := getGoGoMinipoolStatus(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoMinipoolStatus.GetTransactionGasInfo(opts, "submitMinipoolWithdrawable", minipoolAddress)
}

// Submit a minipool withdrawable event
func SubmitMinipoolWithdrawable(ggp *gogopool.GoGoPool, minipoolAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoMinipoolStatus, err := getGoGoMinipoolStatus(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoMinipoolStatus.Transact(opts, "submitMinipoolWithdrawable", minipoolAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not submit minipool withdrawable event: %w", err)
	}
	return hash, nil
}

// Get contracts
var gogoMinipoolStatusLock sync.Mutex

func getGoGoMinipoolStatus(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoMinipoolStatusLock.Lock()
	defer gogoMinipoolStatusLock.Unlock()
	return ggp.GetContract("gogoMinipoolStatus")
}

package storage

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Get a node's withdrawal address
func GetNodeWithdrawalAddress(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (common.Address, error) {
	withdrawalAddress := new(common.Address)
	if err := ggp.GoGoStorageContract.Call(opts, withdrawalAddress, "getNodeWithdrawalAddress", nodeAddress); err != nil {
		return common.Address{}, fmt.Errorf("Could not get node %s withdrawal address: %w", nodeAddress.Hex(), err)
	}
	return *withdrawalAddress, nil
}

// Get a node's pending withdrawal address
func GetNodePendingWithdrawalAddress(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (common.Address, error) {
	withdrawalAddress := new(common.Address)
	if err := ggp.GoGoStorageContract.Call(opts, withdrawalAddress, "getNodePendingWithdrawalAddress", nodeAddress); err != nil {
		return common.Address{}, fmt.Errorf("Could not get node %s pending withdrawal address: %w", nodeAddress.Hex(), err)
	}
	return *withdrawalAddress, nil
}

// Estimate the gas of SetWithdrawalAddress
func EstimateSetWithdrawalAddressGas(ggp *gogopool.GoGoPool, nodeAddress common.Address, withdrawalAddress common.Address, confirm bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return ggp.GoGoStorageContract.GetTransactionGasInfo(opts, "setWithdrawalAddress", nodeAddress, withdrawalAddress, confirm)
}

// Set a node's withdrawal address
func SetWithdrawalAddress(ggp *gogopool.GoGoPool, nodeAddress common.Address, withdrawalAddress common.Address, confirm bool, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := ggp.GoGoStorageContract.Transact(opts, "setWithdrawalAddress", nodeAddress, withdrawalAddress, confirm)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not set node withdrawal address: %w", err)
	}
	return hash, nil
}

// Estimate the gas of ConfirmWithdrawalAddress
func EstimateConfirmWithdrawalAddressGas(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return ggp.GoGoStorageContract.GetTransactionGasInfo(opts, "confirmWithdrawalAddress", nodeAddress)
}

// Set a node's withdrawal address
func ConfirmWithdrawalAddress(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := ggp.GoGoStorageContract.Transact(opts, "confirmWithdrawalAddress", nodeAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not confirm node withdrawal address: %w", err)
	}
	return hash, nil
}

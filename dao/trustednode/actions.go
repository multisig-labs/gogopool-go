package trustednode

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Estimate the gas of Join
func EstimateJoinGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionJoin")
}

// Join the trusted node DAO
// Requires an executed invite proposal
func Join(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedActions.Transact(opts, "actionJoin")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not join the trusted node DAO: %w", err)
	}
	return hash, nil
}

// Estimate the gas of Leave
func EstimateLeaveGas(ggp *gogopool.GoGoPool, ggpBondRefundAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionLeave", ggpBondRefundAddress)
}

// Leave the trusted node DAO
// Requires an executed leave proposal
func Leave(ggp *gogopool.GoGoPool, ggpBondRefundAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedActions.Transact(opts, "actionLeave", ggpBondRefundAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not leave the trusted node DAO: %w", err)
	}
	return hash, nil
}

// Estimate the gas of MakeChallenge
func EstimateMakeChallengeGas(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionChallengeMake", memberAddress)
}

// Make a challenge against a node
func MakeChallenge(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedActions.Transact(opts, "actionChallengeMake", memberAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not challenge trusted node DAO member %s: %w", memberAddress.Hex(), err)
	}
	return hash, nil
}

// Estimate the gas of DecideChallenge
func EstimateDecideChallengeGas(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionChallengeDecide", memberAddress)
}

// Decide a challenge against a node
func DecideChallenge(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedActions.Transact(opts, "actionChallengeDecide", memberAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not decide the challenge against trusted node DAO member %s: %w", memberAddress.Hex(), err)
	}
	return hash, nil
}

// Get contracts
var gogoDAONodeTrustedActionsLock sync.Mutex

func getGoGoDAONodeTrustedActions(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDAONodeTrustedActionsLock.Lock()
	defer gogoDAONodeTrustedActionsLock.Unlock()
	return ggp.GetContract("rocketDAONodeTrustedActions")
}

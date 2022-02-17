package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Config
const (
	ProposalsSettingsContractName = "gogoDAONodeTrustedSettingsProposals"
	CooldownTimeSettingPath       = "proposal.cooldown.time"
	VoteTimeSettingPath           = "proposal.vote.time"
	VoteDelayTimeSettingPath      = "proposal.vote.delay.time"
	ExecuteTimeSettingPath        = "proposal.execute.time"
	ActionTimeSettingPath         = "proposal.action.time"
)

// The cooldown period a member must wait after making a proposal before making another in seconds
func GetProposalCooldownTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	proposalsSettingsContract, err := getProposalsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := proposalsSettingsContract.Call(opts, value, "getCooldownTime"); err != nil {
		return 0, fmt.Errorf("Could not get proposal cooldown period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapProposalCooldownTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, ProposalsSettingsContractName, CooldownTimeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeProposalCooldownTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", CooldownTimeSettingPath), ProposalsSettingsContractName, CooldownTimeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeProposalCooldownTimeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", CooldownTimeSettingPath), ProposalsSettingsContractName, CooldownTimeSettingPath, big.NewInt(int64(value)), opts)
}

// The period a proposal can be voted on for in seconds
func GetProposalVoteTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	proposalsSettingsContract, err := getProposalsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := proposalsSettingsContract.Call(opts, value, "getVoteTime"); err != nil {
		return 0, fmt.Errorf("Could not get proposal voting period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapProposalVoteTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, ProposalsSettingsContractName, VoteTimeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeProposalVoteTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", VoteTimeSettingPath), ProposalsSettingsContractName, VoteTimeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeProposalVoteTimeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", VoteTimeSettingPath), ProposalsSettingsContractName, VoteTimeSettingPath, big.NewInt(int64(value)), opts)
}

// The delay after creation before a proposal can be voted on in seconds
func GetProposalVoteDelayTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	proposalsSettingsContract, err := getProposalsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := proposalsSettingsContract.Call(opts, value, "getVoteDelayTime"); err != nil {
		return 0, fmt.Errorf("Could not get proposal voting delay: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapProposalVoteDelayTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, ProposalsSettingsContractName, VoteDelayTimeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeProposalVoteDelayTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", VoteDelayTimeSettingPath), ProposalsSettingsContractName, VoteDelayTimeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeProposalVoteDelayTimeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", VoteDelayTimeSettingPath), ProposalsSettingsContractName, VoteDelayTimeSettingPath, big.NewInt(int64(value)), opts)
}

// The period during which a passed proposal can be executed in time
func GetProposalExecuteTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	proposalsSettingsContract, err := getProposalsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := proposalsSettingsContract.Call(opts, value, "getExecuteTime"); err != nil {
		return 0, fmt.Errorf("Could not get proposal execution period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapProposalExecuteTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, ProposalsSettingsContractName, ExecuteTimeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeProposalExecuteTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ExecuteTimeSettingPath), ProposalsSettingsContractName, ExecuteTimeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeProposalExecuteTimeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ExecuteTimeSettingPath), ProposalsSettingsContractName, ExecuteTimeSettingPath, big.NewInt(int64(value)), opts)
}

// The period during which an action can be performed on an executed proposal in seconds
func GetProposalActionTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	proposalsSettingsContract, err := getProposalsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := proposalsSettingsContract.Call(opts, value, "getActionTime"); err != nil {
		return 0, fmt.Errorf("Could not get proposal action period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapProposalActionTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, ProposalsSettingsContractName, ActionTimeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeProposalActionTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ActionTimeSettingPath), ProposalsSettingsContractName, ActionTimeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeProposalActionTimeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ActionTimeSettingPath), ProposalsSettingsContractName, ActionTimeSettingPath, big.NewInt(int64(value)), opts)
}

// Get contracts
var proposalsSettingsContractLock sync.Mutex

func getProposalsSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	proposalsSettingsContractLock.Lock()
	defer proposalsSettingsContractLock.Unlock()
	return ggp.GetContract(ProposalsSettingsContractName)
}

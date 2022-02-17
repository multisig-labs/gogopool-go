package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/dao"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/strings"
)

// Estimate the gas of ProposeInviteMember
func EstimateProposeInviteMemberGas(ggp *gogopool.GoGoPool, message string, newMemberAddress common.Address, newMemberId, newMemberUrl string, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	newMemberUrl = strings.Sanitize(newMemberUrl)
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalInvite", newMemberId, newMemberUrl, newMemberAddress)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode invite member proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to invite a new member to the trusted node DAO
func ProposeInviteMember(ggp *gogopool.GoGoPool, message string, newMemberAddress common.Address, newMemberId, newMemberUrl string, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	newMemberUrl = strings.Sanitize(newMemberUrl)
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalInvite", newMemberId, newMemberUrl, newMemberAddress)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode invite member proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeMemberLeave
func EstimateProposeMemberLeaveGas(ggp *gogopool.GoGoPool, message string, memberAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalLeave", memberAddress)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode member leave proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal for a member to leave the trusted node DAO
func ProposeMemberLeave(ggp *gogopool.GoGoPool, message string, memberAddress common.Address, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalLeave", memberAddress)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode member leave proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeReplaceMember
func EstimateProposeReplaceMemberGas(ggp *gogopool.GoGoPool, message string, memberAddress, newMemberAddress common.Address, newMemberId, newMemberUrl string, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	newMemberUrl = strings.Sanitize(newMemberUrl)
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalReplace", memberAddress, newMemberId, newMemberUrl, newMemberAddress)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode replace member proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to replace a member in the trusted node DAO
func ProposeReplaceMember(ggp *gogopool.GoGoPool, message string, memberAddress, newMemberAddress common.Address, newMemberId, newMemberUrl string, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	newMemberUrl = strings.Sanitize(newMemberUrl)
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalReplace", memberAddress, newMemberId, newMemberUrl, newMemberAddress)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode replace member proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeKickMember
func EstimateProposeKickMemberGas(ggp *gogopool.GoGoPool, message string, memberAddress common.Address, ggpFineAmount *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalKick", memberAddress, ggpFineAmount)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode kick member proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to kick a member from the trusted node DAO
func ProposeKickMember(ggp *gogopool.GoGoPool, message string, memberAddress common.Address, ggpFineAmount *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalKick", memberAddress, ggpFineAmount)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode kick member proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeSetBool
func EstimateProposeSetBoolGas(ggp *gogopool.GoGoPool, message, contractName, settingPath string, value bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalSettingBool", contractName, settingPath, value)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode set bool setting proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to update a bool trusted node DAO setting
func ProposeSetBool(ggp *gogopool.GoGoPool, message, contractName, settingPath string, value bool, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalSettingBool", contractName, settingPath, value)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode set bool setting proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeSetUint
func EstimateProposeSetUintGas(ggp *gogopool.GoGoPool, message, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalSettingUint", contractName, settingPath, value)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode set uint setting proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to update a uint trusted node DAO setting
func ProposeSetUint(ggp *gogopool.GoGoPool, message, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalSettingUint", contractName, settingPath, value)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode set uint setting proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of ProposeUpgradeContract
func EstimateProposeUpgradeContractGas(ggp *gogopool.GoGoPool, message, upgradeType, contractName, contractAbi string, contractAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	compressedAbi, err := gogopool.EncodeAbiStr(contractAbi)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalUpgrade", upgradeType, contractName, compressedAbi, contractAddress)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not encode upgrade contract proposal payload: %w", err)
	}
	return EstimateProposalGas(ggp, message, payload, opts)
}

// Submit a proposal to upgrade a contract
func ProposeUpgradeContract(ggp *gogopool.GoGoPool, message, upgradeType, contractName, contractAbi string, contractAddress common.Address, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	compressedAbi, err := gogopool.EncodeAbiStr(contractAbi)
	if err != nil {
		return 0, common.Hash{}, err
	}
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	payload, err := gogoDAONodeTrustedProposals.ABI.Pack("proposalUpgrade", upgradeType, contractName, compressedAbi, contractAddress)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not encode upgrade contract proposal payload: %w", err)
	}
	return SubmitProposal(ggp, message, payload, opts)
}

// Estimate the gas of a proposal submission
func EstimateProposalGas(ggp *gogopool.GoGoPool, message string, payload []byte, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedProposals.GetTransactionGasInfo(opts, "propose", message, payload)
}

// Submit a trusted node DAO proposal
// Returns the ID of the new proposal
func SubmitProposal(ggp *gogopool.GoGoPool, message string, payload []byte, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	proposalCount, err := dao.GetProposalCount(ggp, nil)
	if err != nil {
		return 0, common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedProposals.Transact(opts, "propose", message, payload)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not submit trusted node DAO proposal: %w", err)
	}
	return proposalCount + 1, hash, nil
}

// Estimate the gas of CancelProposal
func EstimateCancelProposalGas(ggp *gogopool.GoGoPool, proposalId uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedProposals.GetTransactionGasInfo(opts, "cancel", big.NewInt(int64(proposalId)))
}

// Cancel a submitted proposal
func CancelProposal(ggp *gogopool.GoGoPool, proposalId uint64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedProposals.Transact(opts, "cancel", big.NewInt(int64(proposalId)))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not cancel trusted node DAO proposal %d: %w", proposalId, err)
	}
	return hash, nil
}

// Estimate the gas of VoteOnProposal
func EstimateVoteOnProposalGas(ggp *gogopool.GoGoPool, proposalId uint64, support bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedProposals.GetTransactionGasInfo(opts, "vote", big.NewInt(int64(proposalId)), support)
}

// Vote on a submitted proposal
func VoteOnProposal(ggp *gogopool.GoGoPool, proposalId uint64, support bool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedProposals.Transact(opts, "vote", big.NewInt(int64(proposalId)), support)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not vote on trusted node DAO proposal %d: %w", proposalId, err)
	}
	return hash, nil
}

// Estimate the gas of ExecuteProposal
func EstimateExecuteProposalGas(ggp *gogopool.GoGoPool, proposalId uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrustedProposals.GetTransactionGasInfo(opts, "execute", big.NewInt(int64(proposalId)))
}

// Execute a submitted proposal
func ExecuteProposal(ggp *gogopool.GoGoPool, proposalId uint64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrustedProposals, err := getRocketDAONodeTrustedProposals(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrustedProposals.Transact(opts, "execute", big.NewInt(int64(proposalId)))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not execute trusted node DAO proposal %d: %w", proposalId, err)
	}
	return hash, nil
}

// Get contracts
var gogoDAONodeTrustedProposalsLock sync.Mutex

func getRocketDAONodeTrustedProposals(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDAONodeTrustedProposalsLock.Lock()
	defer gogoDAONodeTrustedProposalsLock.Unlock()
	return ggp.GetContract("gogoDAONodeTrustedProposals")
}

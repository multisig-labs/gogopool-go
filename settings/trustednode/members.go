package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Config
const (
	MembersSettingsContractName       = "gogoDAONodeTrustedSettingsMembers"
	QuorumSettingPath                 = "members.quorum"
	GGPBondSettingPath                = "members.ggpbond"
	MinipoolUnbondedMaxSettingPath    = "members.minipool.unbonded.max"
	MinipoolUnbondedMinFeeSettingPath = "members.minipool.unbonded.min.fee"
	ChallengeCooldownSettingPath      = "members.challenge.cooldown"
	ChallengeWindowSettingPath        = "members.challenge.window"
	ChallengeCostSettingPath          = "members.challenge.cost"
)

// Member proposal quorum threshold
func GetQuorum(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getQuorum"); err != nil {
		return 0, fmt.Errorf("Could not get member quorum threshold: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapQuorum(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, QuorumSettingPath, avax.EthToWei(value), opts)
}
func ProposeQuorum(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", QuorumSettingPath), MembersSettingsContractName, QuorumSettingPath, avax.EthToWei(value), opts)
}
func EstimateProposeQuorumGas(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", QuorumSettingPath), MembersSettingsContractName, QuorumSettingPath, avax.EthToWei(value), opts)
}

// GGP bond required for a member
func GetGGPBond(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getGGPBond"); err != nil {
		return nil, fmt.Errorf("Could not get member GGP bond amount: %w", err)
	}
	return *value, nil
}
func BootstrapGGPBond(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, GGPBondSettingPath, value, opts)
}
func ProposeGGPBond(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", GGPBondSettingPath), MembersSettingsContractName, GGPBondSettingPath, value, opts)
}
func EstimateProposeGGPBondGas(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", GGPBondSettingPath), MembersSettingsContractName, GGPBondSettingPath, value, opts)
}

// The maximum number of unbonded minipools a member can run
func GetMinipoolUnbondedMax(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getMinipoolUnbondedMax"); err != nil {
		return 0, fmt.Errorf("Could not get member unbonded minipool limit: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapMinipoolUnbondedMax(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeMinipoolUnbondedMax(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", MinipoolUnbondedMaxSettingPath), MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeMinipoolUnbondedMaxGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", MinipoolUnbondedMaxSettingPath), MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}

// The minimum commission rate before unbonded minipools are allowed
func GetMinipoolUnbondedMinFee(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getMinipoolUnbondedMinFee"); err != nil {
		return 0, fmt.Errorf("Could not get member unbonded minipool minimum fee: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapMinipoolUnbondedMinFee(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeMinipoolUnbondedMinFee(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", MinipoolUnbondedMinFeeSettingPath), MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeMinipoolUnbondedMinFeeGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", MinipoolUnbondedMinFeeSettingPath), MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}

// The period a member must wait for before submitting another challenge, in blocks
func GetChallengeCooldown(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeCooldown"); err != nil {
		return 0, fmt.Errorf("Could not get member challenge cooldown period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapChallengeCooldown(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeChallengeCooldown(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ChallengeCooldownSettingPath), MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeChallengeCooldownGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ChallengeCooldownSettingPath), MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}

// The period during which a member can respond to a challenge, in blocks
func GetChallengeWindow(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeWindow"); err != nil {
		return 0, fmt.Errorf("Could not get member challenge window period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapChallengeWindow(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeChallengeWindow(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ChallengeWindowSettingPath), MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeChallengeWindowGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ChallengeWindowSettingPath), MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}

// The fee for a non-member to challenge a member, in wei
func GetChallengeCost(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	membersSettingsContract, err := getMembersSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeCost"); err != nil {
		return nil, fmt.Errorf("Could not get member challenge cost: %w", err)
	}
	return *value, nil
}
func BootstrapChallengeCost(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}
func ProposeChallengeCost(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ChallengeCostSettingPath), MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}
func EstimateProposeChallengeCostGas(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ChallengeCostSettingPath), MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}

// Get contracts
var membersSettingsContractLock sync.Mutex

func getMembersSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	membersSettingsContractLock.Lock()
	defer membersSettingsContractLock.Unlock()
	return ggp.GetContract(MembersSettingsContractName)
}

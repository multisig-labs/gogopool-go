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
	MinipoolSettingsContractName = "gogoDAONodeTrustedSettingsMinipool"
	ScrubPeriodPath              = "minipool.scrub.period"
	ScrubPenaltyEnabledPath      = "minipool.scrub.penalty.enabled"
)

// The cooldown period a member must wait after making a proposal before making another in seconds
func GetScrubPeriod(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getScrubPeriod"); err != nil {
		return 0, fmt.Errorf("Could not get scrub period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapScrubPeriod(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(ggp, MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}
func ProposeScrubPeriod(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(ggp, fmt.Sprintf("set %s", ScrubPeriodPath), MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeScrubPeriodGas(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(ggp, fmt.Sprintf("set %s", ScrubPeriodPath), MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}

// Whether or not the GGP slashing penalty is applied to scrubbed minipools
func GetScrubPenaltyEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := minipoolSettingsContract.Call(opts, value, "getScrubPenaltyEnabled"); err != nil {
		return false, fmt.Errorf("Could not get scrub penalty setting: %w", err)
	}
	return (*value), nil
}
func BootstrapScrubPenaltyEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapBool(ggp, MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}
func ProposeScrubPenaltyEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetBool(ggp, fmt.Sprintf("set %s", ScrubPenaltyEnabledPath), MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}
func EstimateProposeScrubPenaltyEnabledGas(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return trustednodedao.EstimateProposeSetBoolGas(ggp, fmt.Sprintf("set %s", ScrubPenaltyEnabledPath), MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}

// Get contracts
var minipoolSettingsContractLock sync.Mutex

func getMinipoolSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	minipoolSettingsContractLock.Lock()
	defer minipoolSettingsContractLock.Unlock()
	return ggp.GetContract(MinipoolSettingsContractName)
}

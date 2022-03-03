package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Config
const RewardsSettingsContractName = "rocketDAOProtocolSettingsRewards"

// The claim amount for a claimer as a fraction
func GetRewardsClaimerPerc(ggp *gogopool.GoGoPool, contractName string, opts *bind.CallOpts) (float64, error) {
	rewardsSettingsContract, err := getRewardsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := rewardsSettingsContract.Call(opts, value, "getRewardsClaimerPerc", contractName); err != nil {
		return 0, fmt.Errorf("Could not get rewards claimer percent: %w", err)
	}
	return avax.WeiToEth(*value), nil
}

// The time that a claimer's share was last updated
func GetRewardsClaimerPercTimeUpdated(ggp *gogopool.GoGoPool, contractName string, opts *bind.CallOpts) (uint64, error) {
	rewardsSettingsContract, err := getRewardsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := rewardsSettingsContract.Call(opts, value, "getRewardsClaimerPercTimeUpdated", contractName); err != nil {
		return 0, fmt.Errorf("Could not get rewards claimer updated time: %w", err)
	}
	return (*value).Uint64(), nil
}

// The total claim amount for all claimers as a fraction
func GetRewardsClaimersPercTotal(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	rewardsSettingsContract, err := getRewardsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := rewardsSettingsContract.Call(opts, value, "getRewardsClaimersPercTotal"); err != nil {
		return 0, fmt.Errorf("Could not get rewards claimers total percent: %w", err)
	}
	return avax.WeiToEth(*value), nil
}

// Rewards claim interval time
func GetRewardsClaimIntervalTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	rewardsSettingsContract, err := getRewardsSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := rewardsSettingsContract.Call(opts, value, "getRewardsClaimIntervalTime"); err != nil {
		return 0, fmt.Errorf("Could not get rewards claim interval: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapRewardsClaimIntervalTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, RewardsSettingsContractName, "ggp.rewards.claim.period.time", big.NewInt(int64(value)), opts)
}

// Get contracts
var rewardsSettingsContractLock sync.Mutex

func getRewardsSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	rewardsSettingsContractLock.Lock()
	defer rewardsSettingsContractLock.Unlock()
	return ggp.GetContract(RewardsSettingsContractName)
}

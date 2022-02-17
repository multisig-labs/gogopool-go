package rewards

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// Get whether a claims contract is enabled
func getEnabled(claimsContract *gogopool.Contract, claimsName string, opts *bind.CallOpts) (bool, error) {
	enabled := new(bool)
	if err := claimsContract.Call(opts, enabled, "getEnabled"); err != nil {
		return false, fmt.Errorf("Could not get %s claims contract enabled status: %w", claimsName, err)
	}
	return *enabled, nil
}

// Get whether a claimer can make a claim
// Use to check whether a claimer is able to make claims at all
func getClaimPossible(claimsContract *gogopool.Contract, claimsName string, claimerAddress common.Address, opts *bind.CallOpts) (bool, error) {
	claimPossible := new(bool)
	if err := claimsContract.Call(opts, claimPossible, "getClaimPossible", claimerAddress); err != nil {
		return false, fmt.Errorf("Could not get %s claim possible status for %s: %w", claimsName, claimerAddress.Hex(), err)
	}
	return *claimPossible, nil
}

// Get the percentage of rewards available to a claimer
func getClaimRewardsPerc(claimsContract *gogopool.Contract, claimsName string, claimerAddress common.Address, opts *bind.CallOpts) (float64, error) {
	claimRewardsPerc := new(*big.Int)
	if err := claimsContract.Call(opts, claimRewardsPerc, "getClaimRewardsPerc", claimerAddress); err != nil {
		return 0, fmt.Errorf("Could not get %s claim rewards percent for %s: %w", claimsName, claimerAddress.Hex(), err)
	}
	return avax.WeiToEth(*claimRewardsPerc), nil
}

// Get the total amount of rewards available to a claimer
// Use to check whether a claimer is able to make a claim for the current interval (returns zero if unable)
func getClaimRewardsAmount(claimsContract *gogopool.Contract, claimsName string, claimerAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	claimRewardsAmount := new(*big.Int)
	if err := claimsContract.Call(opts, claimRewardsAmount, "getClaimRewardsAmount", claimerAddress); err != nil {
		return nil, fmt.Errorf("Could not get %s claim rewards amount for %s: %w", claimsName, claimerAddress.Hex(), err)
	}
	return *claimRewardsAmount, nil
}

// Get the time that the user registered as a claimer
func getClaimingContractUserRegisteredTime(ggp *gogopool.GoGoPool, claimsContract string, claimerAddress common.Address, opts *bind.CallOpts) (time.Time, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return time.Time{}, err
	}
	claimTime := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, claimTime, "getClaimingContractUserRegisteredTime", claimsContract, claimerAddress); err != nil {
		return time.Time{}, fmt.Errorf("Could not get claims registration time on contract %s for %s: %w", claimsContract, claimerAddress.Hex(), err)
	}
	return time.Unix((*claimTime).Int64(), 0), nil
}

// Get the total amount claimed in the current interval by the given claiming contract
func getClaimingContractTotalClaimed(ggp *gogopool.GoGoPool, claimsContract string, opts *bind.CallOpts) (*big.Int, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return nil, err
	}
	totalClaimed := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, totalClaimed, "getClaimingContractTotalClaimed", claimsContract); err != nil {
		return nil, fmt.Errorf("Could not get total claimed for %s: %w", claimsContract, err)
	}
	return *totalClaimed, nil
}

// Estimate the gas of claim
func estimateClaimGas(claimsContract *gogopool.Contract, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	return claimsContract.GetTransactionGasInfo(opts, "claim")
}

// Claim rewards
func claim(claimsContract *gogopool.Contract, claimsName string, opts *bind.TransactOpts) (common.Hash, error) {
	hash, err := claimsContract.Transact(opts, "claim")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not claim %s rewards: %w", claimsName, err)
	}
	return hash, nil
}

// Get the timestamp that the current rewards interval started
func GetClaimIntervalTimeStart(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (time.Time, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return time.Time{}, err
	}
	unixTime := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, unixTime, "getClaimIntervalTimeStart"); err != nil {
		return time.Time{}, fmt.Errorf("Could not get claim interval time start: %w", err)
	}
	return time.Unix(int64((*unixTime).Uint64()), 0), nil
}

// Get the number of seconds in a claim interval
func GetClaimIntervalTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (time.Duration, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return 0, err
	}
	unixTime := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, unixTime, "getClaimIntervalTime"); err != nil {
		return 0, fmt.Errorf("Could not get claim interval time: %w", err)
	}
	return time.Duration((*unixTime).Int64()) * time.Second, nil
}

// Get the percent of checkpoint rewards that goes to node operators
func GetNodeOperatorRewardsPercent(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return 0, err
	}
	perc := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, perc, "getClaimingContractPerc", "gogoClaimNode"); err != nil {
		return 0, fmt.Errorf("Could not get node operator rewards percent: %w", err)
	}
	return avax.WeiToEth(*perc), nil
}

// Get the percent of checkpoint rewards that goes to ODAO members
func GetTrustedNodeOperatorRewardsPercent(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return 0, err
	}
	perc := new(*big.Int)
	if err := gogoRewardsPool.Call(opts, perc, "getClaimingContractPerc", "gogoClaimTrustedNode"); err != nil {
		return 0, fmt.Errorf("Could not get trusted node operator rewards percent: %w", err)
	}
	return avax.WeiToEth(*perc), nil
}

// Get contracts
var gogoRewardsPoolLock sync.Mutex

func getRocketRewardsPool(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoRewardsPoolLock.Lock()
	defer gogoRewardsPoolLock.Unlock()
	return ggp.GetContract("gogoRewardsPool")
}

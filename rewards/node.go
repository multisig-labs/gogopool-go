package rewards

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Get whether node reward claims are enabled
func GetNodeClaimsEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return false, err
	}
	return getEnabled(gogoClaimNode, "node", opts)
}

// Get whether a node rewards claimer can claim
func GetNodeClaimPossible(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (bool, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return false, err
	}
	return getClaimPossible(gogoClaimNode, "node", claimerAddress, opts)
}

// Get the percentage of rewards available for a node rewards claimer
func GetNodeClaimRewardsPerc(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (float64, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return 0, err
	}
	return getClaimRewardsPerc(gogoClaimNode, "node", claimerAddress, opts)
}

// Get the total amount of rewards available for a node rewards claimer
func GetNodeClaimRewardsAmount(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return nil, err
	}
	return getClaimRewardsAmount(gogoClaimNode, "node", claimerAddress, opts)
}

// Estimate the gas of ClaimNodeRewards
func EstimateClaimNodeRewardsGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateClaimGas(gogoClaimNode, opts)
}

// Claim node rewards
func ClaimNodeRewards(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return claim(gogoClaimNode, "node", opts)
}

// Filters through token claim events and sums the total amount claimed by claimerAddress
func CalculateLifetimeNodeRewards(ggp *gogopool.GoGoPool, claimerAddress common.Address, intervalSize *big.Int, startBlock *big.Int) (*big.Int, error) {
	// Get contracts
	gogoRewardsPool, err := getGoGoRewardsPool(ggp)
	if err != nil {
		return nil, err
	}
	gogoClaimNode, err := getGoGoClaimNode(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoRewardsPool.Address}
	// GGPTokensClaimed(address clamingContract, address claimingAddress, uint256 amount, uint256 time)
	topicFilter := [][]common.Hash{{gogoRewardsPool.ABI.Events["GGPTokensClaimed"].ID}, {gogoClaimNode.Address.Hash()}, {claimerAddress.Hash()}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, startBlock, nil, nil)
	if err != nil {
		return nil, err
	}

	// Iterate over the logs and sum the amount
	sum := big.NewInt(0)
	for _, log := range logs {
		values := make(map[string]interface{})
		// Decode the event
		if gogoRewardsPool.ABI.Events["GGPTokensClaimed"].Inputs.UnpackIntoMap(values, log.Data) != nil {
			return nil, err
		}
		// Add the amount argument to our sum
		amount := values["amount"].(*big.Int)
		sum.Add(sum, amount)
	}
	// Return the result
	return sum, nil
}

// Get the time that the user registered as a claimer
func GetNodeRegistrationTime(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (time.Time, error) {
	return getClaimingContractUserRegisteredTime(ggp, "rocketClaimNode", claimerAddress, opts)
}

// Get the total rewards claimed for this claiming contract this interval
func GetNodeTotalClaimed(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	return getClaimingContractTotalClaimed(ggp, "rocketClaimNode", opts)
}

// Get contracts
var gogoClaimNodeLock sync.Mutex

func getGoGoClaimNode(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoClaimNodeLock.Lock()
	defer gogoClaimNodeLock.Unlock()
	return ggp.GetContract("rocketClaimNode")
}

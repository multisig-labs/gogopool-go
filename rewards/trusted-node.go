package rewards

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// Get whether trusted node reward claims are enabled
func GetTrustedNodeClaimsEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return false, err
	}
	return getEnabled(gogoClaimTrustedNode, "trusted node", opts)
}

// Get whether a trusted node rewards claimer can claim
func GetTrustedNodeClaimPossible(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (bool, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return false, err
	}
	return getClaimPossible(gogoClaimTrustedNode, "trusted node", claimerAddress, opts)
}

// Get the percentage of rewards available for a trusted node rewards claimer
func GetTrustedNodeClaimRewardsPerc(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (float64, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return 0, err
	}
	return getClaimRewardsPerc(gogoClaimTrustedNode, "trusted node", claimerAddress, opts)
}

// Get the total amount of rewards available for a trusted node rewards claimer
func GetTrustedNodeClaimRewardsAmount(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return nil, err
	}
	return getClaimRewardsAmount(gogoClaimTrustedNode, "trusted node", claimerAddress, opts)
}

// Estimate the gas of ClaimTrustedNodeRewards
func EstimateClaimTrustedNodeRewardsGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return estimateClaimGas(gogoClaimTrustedNode, opts)
}

// Claim trusted node rewards
func ClaimTrustedNodeRewards(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	return claim(gogoClaimTrustedNode, "trusted node", opts)
}

// Filters through token claim events and sums the total amount claimed by claimerAddress
func CalculateLifetimeTrustedNodeRewards(ggp *gogopool.GoGoPool, claimerAddress common.Address, intervalSize *big.Int, startBlock *big.Int) (*big.Int, error) {
	// Get contracts
	gogoRewardsPool, err := getRocketRewardsPool(ggp)
	if err != nil {
		return nil, err
	}
	gogoClaimTrustedNode, err := getRocketClaimTrustedNode(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoRewardsPool.Address}
	// GGPTokensClaimed(address clamingContract, address clainingAddress, uint256 amount, uint256 time)
	topicFilter := [][]common.Hash{{gogoRewardsPool.ABI.Events["GGPTokensClaimed"].ID}, {gogoClaimTrustedNode.Address.Hash()}, {claimerAddress.Hash()}}

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
func GetTrustedNodeRegistrationTime(ggp *gogopool.GoGoPool, claimerAddress common.Address, opts *bind.CallOpts) (time.Time, error) {
	return getClaimingContractUserRegisteredTime(ggp, "gogoClaimTrustedNode", claimerAddress, opts)
}

// Get the total rewards claimed for this claiming contract this interval
func GetTrustedNodeTotalClaimed(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	return getClaimingContractTotalClaimed(ggp, "gogoClaimTrustedNode", opts)
}

// Get contracts
var gogoClaimTrustedNodeLock sync.Mutex

func getRocketClaimTrustedNode(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoClaimTrustedNodeLock.Lock()
	defer gogoClaimTrustedNodeLock.Unlock()
	return ggp.GetContract("gogoClaimTrustedNode")
}

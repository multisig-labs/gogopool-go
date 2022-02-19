package node

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"gonum.org/v1/gonum/mathext"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/storage"
	"github.com/multisig-labs/gogopool-go/utils/avax"
	"github.com/multisig-labs/gogopool-go/utils/strings"
)

// Settings
const (
	NodeAddressBatchSize = 50
	NodeDetailsBatchSize = 20
)

// Node details
type NodeDetails struct {
	Address                  common.Address `json:"address"`
	Exists                   bool           `json:"exists"`
	WithdrawalAddress        common.Address `json:"withdrawalAddress"`
	PendingWithdrawalAddress common.Address `json:"pendingWithdrawalAddress"`
	TimezoneLocation         string         `json:"timezoneLocation"`
}

// Count of nodes belonging to a timezone
type TimezoneCount struct {
	Timezone string   `abi:"timezone"`
	Count    *big.Int `abi:"count"`
}

// The results of the trusted node participation calculation
type TrustedNodeParticipation struct {
	StartBlock          uint64
	UpdateFrequency     uint64
	UpdateCount         uint64
	Probability         float64
	ExpectedSubmissions float64
	ActualSubmissions   map[common.Address]float64
	Participation       map[common.Address][]bool
}

// Get all node details
func GetNodes(ggp *gogopool.GoGoPool, opts *bind.CallOpts) ([]NodeDetails, error) {

	// Get node addresses
	nodeAddresses, err := GetNodeAddresses(ggp, opts)
	if err != nil {
		return []NodeDetails{}, err
	}

	// Load node details in batches
	details := make([]NodeDetails, len(nodeAddresses))
	for bsi := 0; bsi < len(nodeAddresses); bsi += NodeDetailsBatchSize {

		// Get batch start & end index
		nsi := bsi
		nei := bsi + NodeDetailsBatchSize
		if nei > len(nodeAddresses) {
			nei = len(nodeAddresses)
		}

		// Load details
		var wg errgroup.Group
		for ni := nsi; ni < nei; ni++ {
			ni := ni
			wg.Go(func() error {
				nodeAddress := nodeAddresses[ni]
				nodeDetails, err := GetNodeDetails(ggp, nodeAddress, opts)
				if err == nil {
					details[ni] = nodeDetails
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []NodeDetails{}, err
		}

	}

	// Return
	return details, nil

}

// Get all node addresses
func GetNodeAddresses(ggp *gogopool.GoGoPool, opts *bind.CallOpts) ([]common.Address, error) {

	// Get node count
	nodeCount, err := GetNodeCount(ggp, opts)
	if err != nil {
		return []common.Address{}, err
	}

	// Load node addresses in batches
	addresses := make([]common.Address, nodeCount)
	for bsi := uint64(0); bsi < nodeCount; bsi += NodeAddressBatchSize {

		// Get batch start & end index
		nsi := bsi
		nei := bsi + NodeAddressBatchSize
		if nei > nodeCount {
			nei = nodeCount
		}

		// Load addresses
		var wg errgroup.Group
		for ni := nsi; ni < nei; ni++ {
			ni := ni
			wg.Go(func() error {
				address, err := GetNodeAt(ggp, ni, opts)
				if err == nil {
					addresses[ni] = address
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []common.Address{}, err
		}

	}

	// Return
	return addresses, nil

}

// Get a node's details
func GetNodeDetails(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (NodeDetails, error) {

	// Data
	var wg errgroup.Group
	var exists bool
	var withdrawalAddress common.Address
	var pendingWithdrawalAddress common.Address
	var timezoneLocation string

	// Load data
	wg.Go(func() error {
		var err error
		exists, err = GetNodeExists(ggp, nodeAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		withdrawalAddress, err = storage.GetNodeWithdrawalAddress(ggp, nodeAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		pendingWithdrawalAddress, err = storage.GetNodePendingWithdrawalAddress(ggp, nodeAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		timezoneLocation, err = GetNodeTimezoneLocation(ggp, nodeAddress, opts)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return NodeDetails{}, err
	}

	// Return
	return NodeDetails{
		Address:                  nodeAddress,
		Exists:                   exists,
		WithdrawalAddress:        withdrawalAddress,
		PendingWithdrawalAddress: pendingWithdrawalAddress,
		TimezoneLocation:         timezoneLocation,
	}, nil

}

// Get the number of nodes in the network
func GetNodeCount(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return 0, err
	}
	nodeCount := new(*big.Int)
	if err := gogoNodeManager.Call(opts, nodeCount, "getNodeCount"); err != nil {
		return 0, fmt.Errorf("Could not get node count: %w", err)
	}
	return (*nodeCount).Uint64(), nil
}

// Get a breakdown of the number of nodes per timezone
func GetNodeCountPerTimezone(ggp *gogopool.GoGoPool, offset, limit *big.Int, opts *bind.CallOpts) ([]TimezoneCount, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return []TimezoneCount{}, err
	}
	timezoneCounts := new([]TimezoneCount)
	if err := gogoNodeManager.Call(opts, timezoneCounts, "getNodeCountPerTimezone", offset, limit); err != nil {
		return []TimezoneCount{}, fmt.Errorf("Could not get node count: %w", err)
	}
	return *timezoneCounts, nil
}

// Get a node address by index
func GetNodeAt(ggp *gogopool.GoGoPool, index uint64, opts *bind.CallOpts) (common.Address, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return common.Address{}, err
	}
	nodeAddress := new(common.Address)
	if err := gogoNodeManager.Call(opts, nodeAddress, "getNodeAt", big.NewInt(int64(index))); err != nil {
		return common.Address{}, fmt.Errorf("Could not get node %d address: %w", index, err)
	}
	return *nodeAddress, nil
}

// Check whether a node exists
func GetNodeExists(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (bool, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return false, err
	}
	exists := new(bool)
	if err := gogoNodeManager.Call(opts, exists, "getNodeExists", nodeAddress); err != nil {
		return false, fmt.Errorf("Could not get node %s exists status: %w", nodeAddress.Hex(), err)
	}
	return *exists, nil
}

// Get a node's timezone location
func GetNodeTimezoneLocation(ggp *gogopool.GoGoPool, nodeAddress common.Address, opts *bind.CallOpts) (string, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return "", err
	}
	timezoneLocation := new(string)
	if err := gogoNodeManager.Call(opts, timezoneLocation, "getNodeTimezoneLocation", nodeAddress); err != nil {
		return "", fmt.Errorf("Could not get node %s timezone location: %w", nodeAddress.Hex(), err)
	}
	return strings.Sanitize(*timezoneLocation), nil
}

// Estimate the gas of RegisterNode
func EstimateRegisterNodeGas(ggp *gogopool.GoGoPool, timezoneLocation string, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	_, err = time.LoadLocation(timezoneLocation)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not verify timezone [%s]: %w", timezoneLocation, err)
	}
	return gogoNodeManager.GetTransactionGasInfo(opts, "registerNode", timezoneLocation)
}

// Register a node
func RegisterNode(ggp *gogopool.GoGoPool, timezoneLocation string, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	_, err = time.LoadLocation(timezoneLocation)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not verify timezone [%s]: %w", timezoneLocation, err)
	}
	hash, err := gogoNodeManager.Transact(opts, "registerNode", timezoneLocation)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not register node: %w", err)
	}
	return hash, nil
}

// Estimate the gas of SetTimezoneLocation
func EstimateSetTimezoneLocationGas(ggp *gogopool.GoGoPool, timezoneLocation string, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	_, err = time.LoadLocation(timezoneLocation)
	if err != nil {
		return gogopool.GasInfo{}, fmt.Errorf("Could not verify timezone [%s]: %w", timezoneLocation, err)
	}
	return gogoNodeManager.GetTransactionGasInfo(opts, "setTimezoneLocation", timezoneLocation)
}

// Set a node's timezone location
func SetTimezoneLocation(ggp *gogopool.GoGoPool, timezoneLocation string, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNodeManager, err := getGoGoNodeManager(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	_, err = time.LoadLocation(timezoneLocation)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not verify timezone [%s]: %w", timezoneLocation, err)
	}
	hash, err := gogoNodeManager.Transact(opts, "setTimezoneLocation", timezoneLocation)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not set node timezone location: %w", err)
	}
	return hash, nil
}

// Returns an array of block numbers for prices submissions the given trusted node has submitted since fromBlock
func GetPricesSubmissions(ggp *gogopool.GoGoPool, nodeAddress common.Address, fromBlock uint64, intervalSize *big.Int) (*[]uint64, error) {
	// Get contracts
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoNetworkPrices.Address}
	topicFilter := [][]common.Hash{{gogoNetworkPrices.ABI.Events["PricesSubmitted"].ID}, {nodeAddress.Hash()}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, big.NewInt(int64(fromBlock)), nil, nil)
	if err != nil {
		return nil, err
	}
	timestamps := make([]uint64, len(logs))
	for i, log := range logs {
		values := make(map[string]interface{})
		// Decode the event
		if gogoNetworkPrices.ABI.Events["PricesSubmitted"].Inputs.UnpackIntoMap(values, log.Data) != nil {
			return nil, err
		}
		timestamps[i] = values["block"].(*big.Int).Uint64()
	}
	return &timestamps, nil
}

// Returns an array of block numbers for balances submissions the given trusted node has submitted since fromBlock
func GetBalancesSubmissions(ggp *gogopool.GoGoPool, nodeAddress common.Address, fromBlock uint64, intervalSize *big.Int) (*[]uint64, error) {
	// Get contracts
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoNetworkBalances.Address}
	topicFilter := [][]common.Hash{{gogoNetworkBalances.ABI.Events["BalancesSubmitted"].ID}, {nodeAddress.Hash()}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, big.NewInt(int64(fromBlock)), nil, nil)
	if err != nil {
		return nil, err
	}

	timestamps := make([]uint64, len(logs))
	for i, log := range logs {
		values := make(map[string]interface{})
		// Decode the event
		if gogoNetworkBalances.ABI.Events["BalancesSubmitted"].Inputs.UnpackIntoMap(values, log.Data) != nil {
			return nil, err
		}
		timestamps[i] = values["block"].(*big.Int).Uint64()
	}
	return &timestamps, nil
}

// Returns the most recent block number that the number of trusted nodes changed since fromBlock
func getLatestMemberCountChangedBlock(ggp *gogopool.GoGoPool, fromBlock uint64, intervalSize *big.Int) (uint64, error) {
	// Get contracts
	gogoDaoNodeTrustedActions, err := getGoGoDAONodeTrustedActions(ggp)
	if err != nil {
		return 0, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoDaoNodeTrustedActions.Address}
	topicFilter := [][]common.Hash{{gogoDaoNodeTrustedActions.ABI.Events["ActionJoined"].ID, gogoDaoNodeTrustedActions.ABI.Events["ActionLeave"].ID, gogoDaoNodeTrustedActions.ABI.Events["ActionKick"].ID, gogoDaoNodeTrustedActions.ABI.Events["ActionChallengeDecided"].ID}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, big.NewInt(int64(fromBlock)), nil, nil)
	if err != nil {
		return 0, err
	}

	for i := range logs {
		log := logs[len(logs)-i-1]
		if log.Topics[0] == gogoDaoNodeTrustedActions.ABI.Events["ActionChallengeDecided"].ID {
			values := make(map[string]interface{})
			// Decode the event
			if gogoDaoNodeTrustedActions.ABI.Events["ActionChallengeDecided"].Inputs.UnpackIntoMap(values, log.Data) != nil {
				return 0, err
			}
			if values["success"].(bool) {
				return log.BlockNumber, nil
			}
		} else {
			return log.BlockNumber, nil
		}
	}
	return fromBlock, nil
}

// Calculates the participation rate of every trusted node on price submission since the last block that member count changed
func CalculateTrustedNodePricesParticipation(ggp *gogopool.GoGoPool, intervalSize *big.Int, opts *bind.CallOpts) (*TrustedNodeParticipation, error) {
	// Get the update frequency
	updatePricesFrequency, err := protocol.GetSubmitPricesFrequency(ggp, opts)
	if err != nil {
		return nil, err
	}
	// Get the current block
	currentBlock, err := ggp.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	currentBlockNumber := currentBlock.Number.Uint64()
	// Get the block of the most recent member join (limiting to 50 intervals)
	minBlock := (currentBlockNumber/updatePricesFrequency - 50) * updatePricesFrequency
	latestMemberCountChangedBlock, err := getLatestMemberCountChangedBlock(ggp, minBlock, intervalSize)
	if err != nil {
		return nil, err
	}
	// Get the number of current members
	memberCount, err := trustednode.GetMemberCount(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Start block is the first interval after the latest join
	startBlock := (latestMemberCountChangedBlock/updatePricesFrequency + 1) * updatePricesFrequency
	// The number of members that have to submit each interval
	consensus := math.Floor(float64(memberCount)/2 + 1)
	// Check if any intervals have passed
	intervalsPassed := uint64(0)
	if currentBlockNumber > startBlock {
		// The number of intervals passed
		intervalsPassed = (currentBlockNumber-startBlock)/updatePricesFrequency + 1
	}
	// How many submissions would we expect per member given a random submission
	expected := float64(intervalsPassed) * consensus / float64(memberCount)
	// Get trusted members
	members, err := trustednode.GetMembers(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Construct the epoch map
	participationTable := make(map[common.Address][]bool)
	// Iterate members and sum chi-square
	submissions := make(map[common.Address]float64)
	chi := float64(0)
	for _, member := range members {
		participationTable[member.Address] = make([]bool, intervalsPassed)
		actual := 0
		if intervalsPassed > 0 {
			blocks, err := GetPricesSubmissions(ggp, member.Address, startBlock, intervalSize)
			if err != nil {
				return nil, err
			}
			actual = len(*blocks)
			delta := float64(actual) - expected
			chi += (delta * delta) / expected
			// Add to participation table
			for _, block := range *blocks {
				// Ignore out of step updates
				if block%updatePricesFrequency == 0 {
					index := block/updatePricesFrequency - startBlock/updatePricesFrequency
					participationTable[member.Address][index] = true
				}
			}
		}
		// Save actual submission
		submissions[member.Address] = float64(actual)
	}
	// Calculate inverse cumulative density function with members-1 DoF
	probability := float64(1)
	if intervalsPassed > 0 {
		probability = 1 - mathext.GammaIncReg(float64(len(members)-1)/2, chi/2)
	}
	// Construct return value
	participation := TrustedNodeParticipation{
		Probability:         probability,
		ExpectedSubmissions: expected,
		ActualSubmissions:   submissions,
		StartBlock:          startBlock,
		UpdateFrequency:     updatePricesFrequency,
		UpdateCount:         intervalsPassed,
		Participation:       participationTable,
	}
	return &participation, nil
}

// Calculates the participation rate of every trusted node on balance submission since the last block that member count changed
func CalculateTrustedNodeBalancesParticipation(ggp *gogopool.GoGoPool, intervalSize *big.Int, opts *bind.CallOpts) (*TrustedNodeParticipation, error) {
	// Get the update frequency
	updateBalancesFrequency, err := protocol.GetSubmitBalancesFrequency(ggp, opts)
	if err != nil {
		return nil, err
	}
	// Get the current block
	currentBlock, err := ggp.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	currentBlockNumber := currentBlock.Number.Uint64()
	// Get the block of the most recent member join (limiting to 50 intervals)
	minBlock := (currentBlockNumber/updateBalancesFrequency - 50) * updateBalancesFrequency
	latestMemberCountChangedBlock, err := getLatestMemberCountChangedBlock(ggp, minBlock, intervalSize)
	if err != nil {
		return nil, err
	}
	// Get the number of current members
	memberCount, err := trustednode.GetMemberCount(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Start block is the first interval after the latest join
	startBlock := (latestMemberCountChangedBlock/updateBalancesFrequency + 1) * updateBalancesFrequency
	// The number of members that have to submit each interval
	consensus := math.Floor(float64(memberCount)/2 + 1)
	// Check if any intervals have passed
	intervalsPassed := uint64(0)
	if currentBlockNumber > startBlock {
		// The number of intervals passed
		intervalsPassed = (currentBlockNumber-startBlock)/updateBalancesFrequency + 1
	}
	// How many submissions would we expect per member given a random submission
	expected := float64(intervalsPassed) * consensus / float64(memberCount)
	// Get trusted members
	members, err := trustednode.GetMembers(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Construct the epoch map
	participationTable := make(map[common.Address][]bool)
	// Iterate members and sum chi-square
	submissions := make(map[common.Address]float64)
	chi := float64(0)
	for _, member := range members {
		participationTable[member.Address] = make([]bool, intervalsPassed)
		actual := 0
		if intervalsPassed > 0 {
			blocks, err := GetBalancesSubmissions(ggp, member.Address, startBlock, intervalSize)
			if err != nil {
				return nil, err
			}
			actual = len(*blocks)
			delta := float64(actual) - expected
			chi += (delta * delta) / expected
			// Add to participation table
			for _, block := range *blocks {
				// Ignore out of step updates
				if block%updateBalancesFrequency == 0 {
					index := block/updateBalancesFrequency - startBlock/updateBalancesFrequency
					participationTable[member.Address][index] = true
				}
			}
		}
		// Save actual submission
		submissions[member.Address] = float64(actual)
	}
	// Calculate inverse cumulative density function with members-1 DoF
	probability := float64(1)
	if intervalsPassed > 0 {
		probability = 1 - mathext.GammaIncReg(float64(len(members)-1)/2, chi/2)
	}
	// Construct return value
	participation := TrustedNodeParticipation{
		Probability:         probability,
		ExpectedSubmissions: expected,
		ActualSubmissions:   submissions,
		StartBlock:          startBlock,
		UpdateFrequency:     updateBalancesFrequency,
		UpdateCount:         intervalsPassed,
		Participation:       participationTable,
	}
	return &participation, nil
}

// Returns an array of members who submitted a balance since fromBlock
func GetLatestBalancesSubmissions(ggp *gogopool.GoGoPool, fromBlock uint64, intervalSize *big.Int) ([]common.Address, error) {
	// Get contracts
	gogoNetworkBalances, err := getGoGoNetworkBalances(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoNetworkBalances.Address}
	topicFilter := [][]common.Hash{{gogoNetworkBalances.ABI.Events["BalancesSubmitted"].ID}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, big.NewInt(int64(fromBlock)), nil, nil)
	if err != nil {
		return nil, err
	}

	results := make([]common.Address, len(logs))
	for i, log := range logs {
		// Topic 0 is the event, topic 1 is the "from" address
		address := common.BytesToAddress(log.Topics[1].Bytes())
		results[i] = address
	}
	return results, nil
}

// Returns a mapping of members and whether they have submitted balances this interval or not
func GetTrustedNodeLatestBalancesParticipation(ggp *gogopool.GoGoPool, intervalSize *big.Int, opts *bind.CallOpts) (map[common.Address]bool, error) {
	// Get the update frequency
	updateBalancesFrequency, err := protocol.GetSubmitBalancesFrequency(ggp, opts)
	if err != nil {
		return nil, err
	}
	// Get the current block
	currentBlock, err := ggp.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	currentBlockNumber := currentBlock.Number.Uint64()
	// Get trusted members
	members, err := trustednode.GetMembers(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Get submission within the current interval
	fromBlock := currentBlockNumber / updateBalancesFrequency * updateBalancesFrequency
	submissions, err := GetLatestBalancesSubmissions(ggp, fromBlock, intervalSize)
	if err != nil {
		return nil, err
	}
	// Build and return result table
	participationTable := make(map[common.Address]bool)
	for _, member := range members {
		participationTable[member.Address] = false
	}
	for _, submission := range submissions {
		participationTable[submission] = true
	}
	return participationTable, nil
}

// Returns an array of members who submitted prices since fromBlock
func GetLatestPricesSubmissions(ggp *gogopool.GoGoPool, fromBlock uint64, intervalSize *big.Int) ([]common.Address, error) {
	// Get contracts
	gogoNetworkPrices, err := getGoGoNetworkPrices(ggp)
	if err != nil {
		return nil, err
	}
	// Construct a filter query for relevant logs
	addressFilter := []common.Address{*gogoNetworkPrices.Address}
	topicFilter := [][]common.Hash{{gogoNetworkPrices.ABI.Events["PricesSubmitted"].ID}}

	// Get the event logs
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, big.NewInt(int64(fromBlock)), nil, nil)
	if err != nil {
		return nil, err
	}

	results := make([]common.Address, len(logs))
	for i, log := range logs {
		// Topic 0 is the event, topic 1 is the "from" address
		address := common.BytesToAddress(log.Topics[1].Bytes())
		results[i] = address
	}
	return results, nil
}

// Returns a mapping of members and whether they have submitted prices this interval or not
func GetTrustedNodeLatestPricesParticipation(ggp *gogopool.GoGoPool, intervalSize *big.Int, opts *bind.CallOpts) (map[common.Address]bool, error) {
	// Get the update frequency
	updatePricesFrequency, err := protocol.GetSubmitPricesFrequency(ggp, opts)
	if err != nil {
		return nil, err
	}
	// Get the current block
	currentBlock, err := ggp.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	currentBlockNumber := currentBlock.Number.Uint64()
	// Get trusted members
	members, err := trustednode.GetMembers(ggp, nil)
	if err != nil {
		return nil, err
	}
	// Get submission within the current interval
	fromBlock := currentBlockNumber / updatePricesFrequency * updatePricesFrequency
	submissions, err := GetLatestPricesSubmissions(ggp, fromBlock, intervalSize)
	if err != nil {
		return nil, err
	}
	// Build and return result table
	participationTable := make(map[common.Address]bool)
	for _, member := range members {
		participationTable[member.Address] = false
	}
	for _, submission := range submissions {
		participationTable[submission] = true
	}
	return participationTable, nil
}

// Get contracts
var gogoNodeManagerLock sync.Mutex

func getGoGoNodeManager(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNodeManagerLock.Lock()
	defer gogoNodeManagerLock.Unlock()
	return ggp.GetContract("gogoNodeManager")
}

var gogoNetworkPricesLock sync.Mutex

func getGoGoNetworkPrices(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNetworkPricesLock.Lock()
	defer gogoNetworkPricesLock.Unlock()
	return ggp.GetContract("gogoNetworkPrices")
}

var gogoNetworkBalancesLock sync.Mutex

func getGoGoNetworkBalances(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNetworkBalancesLock.Lock()
	defer gogoNetworkBalancesLock.Unlock()
	return ggp.GetContract("gogoNetworkBalances")
}

var gogoDAONodeTrustedActionsLock sync.Mutex

func getGoGoDAONodeTrustedActions(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDAONodeTrustedActionsLock.Lock()
	defer gogoDAONodeTrustedActionsLock.Unlock()
	return ggp.GetContract("gogoDAONodeTrustedActions")
}

package avax

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/multisig-labs/gogopool-go/gogopool"
)

type FilterQuery struct {
	BlockHash *common.Hash
	FromBlock *big.Int
	ToBlock   *big.Int
	Topics    [][]common.Hash
}

func FilterContractLogs(ggp *gogopool.GoGoPool, contractName string, q FilterQuery, intervalSize *big.Int) ([]types.Log, error) {
	gogoDaoNodeTrustedUpgrade, err := ggp.GetContract("gogoDAONodeTrustedUpgrade")
	if err != nil {
		return nil, err
	}
	// Get all the addresses this contract has ever been deployed at
	addresses := make([]common.Address, 0)
	// Construct a filter to query ContractUpgraded event
	addressFilter := []common.Address{*gogoDaoNodeTrustedUpgrade.Address}
	topicFilter := [][]common.Hash{{gogoDaoNodeTrustedUpgrade.ABI.Events["ContractUpgraded"].ID}, {crypto.Keccak256Hash([]byte(contractName))}}
	logs, err := GetLogs(ggp, addressFilter, topicFilter, intervalSize, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	// Interate the logs and store every past contract address
	for _, log := range logs {
		addresses = append(addresses, common.HexToAddress(log.Topics[2].Hex()))
	}
	// Append current address
	currentAddress, err := ggp.GetAddress(contractName)
	if err != nil {
		return nil, err
	}
	addresses = append(addresses, *currentAddress)
	// Perform the desired getLogs call and return results
	return GetLogs(ggp, addresses, q.Topics, intervalSize, q.FromBlock, q.ToBlock, q.BlockHash)
}

// Gets the logs for a particular log request, breaking the calls into batches if necessary
func GetLogs(ggp *gogopool.GoGoPool, addressFilter []common.Address, topicFilter [][]common.Hash, intervalSize, fromBlock, toBlock *big.Int, blockHash *common.Hash) ([]types.Log, error) {
	var logs []types.Log

	// Get the block that Rocket Pool was deployed on as the lower bound if one wasn't specified
	if fromBlock == nil {
		var err error
		deployBlockHash := crypto.Keccak256Hash([]byte("deploy.block"))
		fromBlock, err = ggp.GoGoStorage.GetUint(nil, deployBlockHash)
		if err != nil {
			return nil, err
		}
	}

	if intervalSize == nil {
		// Handle unlimited intervals with a single call
		logs, err := ggp.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
			Addresses: addressFilter,
			Topics:    topicFilter,
			FromBlock: fromBlock,
			ToBlock:   toBlock,
			BlockHash: blockHash,
		})
		if err != nil {
			return nil, err
		}
		return logs, nil
	} else {
		// Get the latest block
		if toBlock == nil {
			latestBlock, err := ggp.Client.BlockNumber(context.Background())
			if err != nil {
				return nil, err
			}
			toBlock = big.NewInt(0)
			toBlock.SetUint64(latestBlock)
		}

		// Set the start and end, clamping on the latest block
		intervalSize.Sub(intervalSize, big.NewInt(1))
		start := fromBlock
		end := big.NewInt(0).Add(start, intervalSize)
		if end.Cmp(toBlock) == 1 {
			end = toBlock
		}
		for {
			// Get the logs using the current interval
			newLogs, err := ggp.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
				Addresses: addressFilter,
				Topics:    topicFilter,
				FromBlock: start,
				ToBlock:   end,
				BlockHash: blockHash,
			})
			if err != nil {
				return nil, err
			}

			// Append the logs to the total list
			logs = append(logs, newLogs...)

			// Return once we've finished iterating
			if end.Cmp(toBlock) == 0 {
				return logs, nil
			}

			// Update to the next interval (end+1 : that + interval - 1)
			start.Add(end, big.NewInt(1))
			end.Add(start, intervalSize)
			if end.Cmp(toBlock) == 1 {
				end = toBlock
			}
		}
	}
}

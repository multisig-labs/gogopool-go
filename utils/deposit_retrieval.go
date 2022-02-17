package utils

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/multisig-labs/gogopool-go/gogopool"
	ggptypes "github.com/multisig-labs/gogopool-go/types"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// BeaconDepositEvent represents a DepositEvent event raised by the BeaconDeposit contract.
type BeaconDepositEvent struct {
	Pubkey                []byte    `abi:"pubkey"`
	WithdrawalCredentials []byte    `abi:"withdrawal_credentials"`
	Amount                []byte    `abi:"amount"`
	Signature             []byte    `abi:"signature"`
	Index                 []byte    `abi:"index"`
	Raw                   types.Log // Blockchain specific contextual infos
}

// Formatted Beacon deposit event data
type DepositData struct {
	Pubkey                ggptypes.ValidatorPubkey    `json:"pubkey"`
	WithdrawalCredentials common.Hash                 `json:"withdrawalCredentials"`
	Amount                uint64                      `json:"amount"`
	Signature             ggptypes.ValidatorSignature `json:"signature"`
	TxHash                common.Hash                 `json:"txHash"`
	BlockNumber           uint64                      `json:"blockNumber"`
	TxIndex               uint                        `json:"txIndex"`
}

// Gets all of the deposit contract's deposit events for the provided pubkeys
func GetDeposits(ggp *gogopool.GoGoPool, pubkeys map[ggptypes.ValidatorPubkey]bool, startBlock *big.Int, intervalSize *big.Int, opts *bind.CallOpts) (map[ggptypes.ValidatorPubkey][]DepositData, error) {

	// Get the deposit contract wrapper
	casperDeposit, err := getCasperDeposit(ggp)
	if err != nil {
		return nil, err
	}

	// Create the initial map and pubkey lookup
	depositMap := make(map[ggptypes.ValidatorPubkey][]DepositData, len(pubkeys))

	// Get the deposit events
	addressFilter := []common.Address{*casperDeposit.Address}
	topicFilter := [][]common.Hash{{casperDeposit.ABI.Events["DepositEvent"].ID}}
	logs, err := avax.GetLogs(ggp, addressFilter, topicFilter, intervalSize, startBlock, nil, nil)
	if err != nil {
		return nil, err
	}

	// Process each event
	for _, log := range logs {
		depositEvent := new(BeaconDepositEvent)
		err = casperDeposit.Contract.UnpackLog(depositEvent, "DepositEvent", log)
		if err != nil {
			return nil, err
		}

		// Check if this is a deposit for one of the pubkeys we're looking for
		pubkey := ggptypes.BytesToValidatorPubkey(depositEvent.Pubkey)
		_, exists := pubkeys[pubkey]
		if exists {
			// Convert the deposit amount from little-endian binary to a uint64
			var amount uint64
			buf := bytes.NewReader(depositEvent.Amount)
			err = binary.Read(buf, binary.LittleEndian, &amount)
			if err != nil {
				return nil, err
			}

			// Create the deposit data wrapper and add it to this pubkey's collection
			depositData := DepositData{
				Pubkey:                pubkey,
				WithdrawalCredentials: common.BytesToHash(depositEvent.WithdrawalCredentials),
				Amount:                amount,
				Signature:             ggptypes.BytesToValidatorSignature(depositEvent.Signature),
				TxHash:                log.TxHash,
				BlockNumber:           log.BlockNumber,
				TxIndex:               log.TxIndex,
			}
			depositMap[pubkey] = append(depositMap[pubkey], depositData)
		}
	}

	// Sort deposits by time
	for _, deposits := range depositMap {
		if len(deposits) > 1 {
			sortDepositData(deposits)
		}
	}

	return depositMap, nil
}

// Sorts a slice of deposit data entries - lower blocks come first, and if multiple transactions occur
// in the same block, lower transaction indices come first
func sortDepositData(data []DepositData) {
	sort.Slice(data, func(i int, j int) bool {
		first := data[i]
		second := data[j]
		if first.BlockNumber == second.BlockNumber {
			return first.TxIndex < second.TxIndex
		}
		return first.BlockNumber < second.BlockNumber
	})
}

// Get contracts
var casperDepositLock sync.Mutex

func getCasperDeposit(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	casperDepositLock.Lock()
	defer casperDepositLock.Unlock()
	return ggp.GetContract("casperDeposit")
}

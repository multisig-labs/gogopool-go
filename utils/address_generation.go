package utils

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/minipool"
	ggptypes "github.com/multisig-labs/gogopool-go/types"
)

// Combine a node's address and a salt to retreive a new salt compatible with depositing
func GetNodeSalt(nodeAddress common.Address, salt *big.Int) common.Hash {
	// Create a new salt by hashing the original and the node address
	saltBytes := [32]byte{}
	salt.FillBytes(saltBytes[:])
	saltHash := crypto.Keccak256Hash(nodeAddress.Bytes(), saltBytes[:])
	return saltHash
}

// Precompute the address of a minipool based on the node wallet, deposit type, and unique salt
// If you set minipoolBytecode to nil, this will retrieve it from the contracts using minipool.GetMinipoolBytecode().
func GenerateAddress(ggp *gogopool.GoGoPool, nodeAddress common.Address, depositType ggptypes.MinipoolDeposit, salt *big.Int, minipoolBytecode []byte) (common.Address, error) {

	// Get dependencies
	gogoMinipoolManager, err := getGoGoMinipoolManager(ggp)
	if err != nil {
		return common.Address{}, err
	}
	minipoolAbi, err := ggp.GetABI("rocketMinipool")
	if err != nil {
		return common.Address{}, err
	}

	if len(minipoolBytecode) == 0 {
		minipoolBytecode, err = minipool.GetMinipoolBytecode(ggp, nil)
		if err != nil {
			return common.Address{}, fmt.Errorf("Error getting minipool bytecode: %w", err)
		}
	}

	// Create the hash of the minipool constructor call
	depositTypeBytes := [32]byte{}
	depositTypeBytes[0] = byte(depositType)
	packedConstructorArgs, err := minipoolAbi.Pack("", ggp.GoGoStorageContract.Address, nodeAddress, depositType)
	if err != nil {
		return common.Address{}, fmt.Errorf("Error creating minipool constructor args: %w", err)
	}

	// Get the node salt and initialization data
	nodeSalt := GetNodeSalt(nodeAddress, salt)
	initData := append(minipoolBytecode, packedConstructorArgs...)
	initHash := crypto.Keccak256(initData)

	address := crypto.CreateAddress2(*gogoMinipoolManager.Address, nodeSalt, initHash)
	return address, nil

}

// Get contracts
var gogoMinipoolManagerLock sync.Mutex

func getGoGoMinipoolManager(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoMinipoolManagerLock.Lock()
	defer gogoMinipoolManagerLock.Unlock()
	return ggp.GetContract("rocketMinipoolManager")
}

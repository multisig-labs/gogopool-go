package gogopool

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/sync/errgroup"

	"github.com/multisig-labs/gogopool-go/contracts"
	"github.com/multisig-labs/gogopool-go/utils/client"
)

// Cache settings
const CacheTTL = 300 // 5 minutes

// Cached data types
type cachedAddress struct {
	address *common.Address
	time    int64
}
type cachedABI struct {
	abi  *abi.ABI
	time int64
}
type cachedContract struct {
	contract *Contract
	time     int64
}

// Rocket Pool contract manager
type GoGoPool struct {
	Client              *client.EthClientProxy
	GoGoStorage         *contracts.GoGoStorage
	GoGoStorageContract *Contract
	addresses           map[string]cachedAddress
	abis                map[string]cachedABI
	contracts           map[string]cachedContract
	addressesLock       sync.RWMutex
	abisLock            sync.RWMutex
	contractsLock       sync.RWMutex
}

// Create new contract manager
func NewGoGoPool(client *client.EthClientProxy, gogoStorageAddress common.Address) (*GoGoPool, error) {

	// Initialize GoGoStorage contract
	gogoStorage, err := contracts.NewGoGoStorage(gogoStorageAddress, client)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize Rocket Pool storage contract: %w", err)
	}

	// Create a Contract for it
	rsAbi, err := abi.JSON(strings.NewReader(contracts.GoGoStorageABI))
	if err != nil {
		return nil, err
	}
	contract := &Contract{
		Contract: bind.NewBoundContract(gogoStorageAddress, rsAbi, client, client, client),
		Address:  &gogoStorageAddress,
		ABI:      &rsAbi,
		Client:   client,
	}

	// Create and return
	return &GoGoPool{
		Client:              client,
		GoGoStorage:         gogoStorage,
		GoGoStorageContract: contract,
		addresses:           make(map[string]cachedAddress),
		abis:                make(map[string]cachedABI),
		contracts:           make(map[string]cachedContract),
	}, nil

}

// Load Rocket Pool contract addresses
func (ggp *GoGoPool) GetAddress(contractName string) (*common.Address, error) {

	// Check for cached address
	if cached, ok := ggp.getCachedAddress(contractName); ok {
		if time.Now().Unix()-cached.time <= CacheTTL {
			return cached.address, nil
		} else {
			ggp.deleteCachedAddress(contractName)
		}
	}

	// Get address
	address, err := ggp.GoGoStorage.GetAddress(nil, crypto.Keccak256Hash([]byte("contract.address"), []byte(contractName)))
	if err != nil {
		return nil, fmt.Errorf("Could not load contract %s address: %w", contractName, err)
	}

	// Cache address
	ggp.setCachedAddress(contractName, cachedAddress{
		address: &address,
		time:    time.Now().Unix(),
	})

	// Return
	return &address, nil

}
func (ggp *GoGoPool) GetAddresses(contractNames ...string) ([]*common.Address, error) {

	// Data
	var wg errgroup.Group
	addresses := make([]*common.Address, len(contractNames))

	// Load addresses
	for ci, contractName := range contractNames {
		ci, contractName := ci, contractName
		wg.Go(func() error {
			address, err := ggp.GetAddress(contractName)
			if err == nil {
				addresses[ci] = address
			}
			return err
		})
	}

	// Wait for data
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	// Return
	return addresses, nil

}

// Load Rocket Pool contract ABIs
func (ggp *GoGoPool) GetABI(contractName string) (*abi.ABI, error) {

	// Check for cached ABI
	if cached, ok := ggp.getCachedABI(contractName); ok {
		if time.Now().Unix()-cached.time <= CacheTTL {
			return cached.abi, nil
		} else {
			ggp.deleteCachedABI(contractName)
		}
	}

	// Get ABI
	abiEncoded, err := ggp.GoGoStorage.GetString(nil, crypto.Keccak256Hash([]byte("contract.abi"), []byte(contractName)))
	if err != nil {
		return nil, fmt.Errorf("Could not load contract %s ABI: %w", contractName, err)
	}

	// Decode ABI
	abi, err := DecodeAbi(abiEncoded)
	if err != nil {
		return nil, fmt.Errorf("Could not decode contract %s ABI: %w", contractName, err)
	}

	// Cache ABI
	ggp.setCachedABI(contractName, cachedABI{
		abi:  abi,
		time: time.Now().Unix(),
	})

	// Return
	return abi, nil

}
func (ggp *GoGoPool) GetABIs(contractNames ...string) ([]*abi.ABI, error) {

	// Data
	var wg errgroup.Group
	abis := make([]*abi.ABI, len(contractNames))

	// Load ABIs
	for ci, contractName := range contractNames {
		ci, contractName := ci, contractName
		wg.Go(func() error {
			abi, err := ggp.GetABI(contractName)
			if err == nil {
				abis[ci] = abi
			}
			return err
		})
	}

	// Wait for data
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	// Return
	return abis, nil

}

// Load Rocket Pool contracts
func (ggp *GoGoPool) GetContract(contractName string) (*Contract, error) {

	// Check for cached contract
	if cached, ok := ggp.getCachedContract(contractName); ok {
		if time.Now().Unix()-cached.time <= CacheTTL {
			return cached.contract, nil
		} else {
			ggp.deleteCachedContract(contractName)
		}
	}

	// Data
	var wg errgroup.Group
	var address *common.Address
	var abi *abi.ABI

	// Load data
	wg.Go(func() error {
		var err error
		address, err = ggp.GetAddress(contractName)
		return err
	})
	wg.Go(func() error {
		var err error
		abi, err = ggp.GetABI(contractName)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	// Create contract
	contract := &Contract{
		Contract: bind.NewBoundContract(*address, *abi, ggp.Client, ggp.Client, ggp.Client),
		Address:  address,
		ABI:      abi,
		Client:   ggp.Client,
	}

	// Cache contract
	ggp.setCachedContract(contractName, cachedContract{
		contract: contract,
		time:     time.Now().Unix(),
	})

	// Return
	return contract, nil

}
func (ggp *GoGoPool) GetContracts(contractNames ...string) ([]*Contract, error) {

	// Data
	var wg errgroup.Group
	contracts := make([]*Contract, len(contractNames))

	// Load contracts
	for ci, contractName := range contractNames {
		ci, contractName := ci, contractName
		wg.Go(func() error {
			contract, err := ggp.GetContract(contractName)
			if err == nil {
				contracts[ci] = contract
			}
			return err
		})
	}

	// Wait for data
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	// Return
	return contracts, nil

}

// Create a Rocket Pool contract instance
func (ggp *GoGoPool) MakeContract(contractName string, address common.Address) (*Contract, error) {

	// Load ABI
	abi, err := ggp.GetABI(contractName)
	if err != nil {
		return nil, err
	}

	// Create and return
	return &Contract{
		Contract: bind.NewBoundContract(address, *abi, ggp.Client, ggp.Client, ggp.Client),
		Address:  &address,
		ABI:      abi,
		Client:   ggp.Client,
	}, nil

}

// Address cache control
func (ggp *GoGoPool) getCachedAddress(contractName string) (cachedAddress, bool) {
	ggp.addressesLock.RLock()
	defer ggp.addressesLock.RUnlock()
	value, ok := ggp.addresses[contractName]
	return value, ok
}
func (ggp *GoGoPool) setCachedAddress(contractName string, value cachedAddress) {
	ggp.addressesLock.Lock()
	defer ggp.addressesLock.Unlock()
	ggp.addresses[contractName] = value
}
func (ggp *GoGoPool) deleteCachedAddress(contractName string) {
	ggp.addressesLock.Lock()
	defer ggp.addressesLock.Unlock()
	delete(ggp.addresses, contractName)
}

// ABI cache control
func (ggp *GoGoPool) getCachedABI(contractName string) (cachedABI, bool) {
	ggp.abisLock.RLock()
	defer ggp.abisLock.RUnlock()
	value, ok := ggp.abis[contractName]
	return value, ok
}
func (ggp *GoGoPool) setCachedABI(contractName string, value cachedABI) {
	ggp.abisLock.Lock()
	defer ggp.abisLock.Unlock()
	ggp.abis[contractName] = value
}
func (ggp *GoGoPool) deleteCachedABI(contractName string) {
	ggp.abisLock.Lock()
	defer ggp.abisLock.Unlock()
	delete(ggp.abis, contractName)
}

// Contract cache control
func (ggp *GoGoPool) getCachedContract(contractName string) (cachedContract, bool) {
	ggp.contractsLock.RLock()
	defer ggp.contractsLock.RUnlock()
	value, ok := ggp.contracts[contractName]
	return value, ok
}
func (ggp *GoGoPool) setCachedContract(contractName string, value cachedContract) {
	ggp.contractsLock.Lock()
	defer ggp.contractsLock.Unlock()
	ggp.contracts[contractName] = value
}
func (ggp *GoGoPool) deleteCachedContract(contractName string) {
	ggp.contractsLock.Lock()
	defer ggp.contractsLock.Unlock()
	delete(ggp.contracts, contractName)
}

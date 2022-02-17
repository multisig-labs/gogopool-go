// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GoGoStorageABI is the input ABI used to generate the binding from.
const GoGoStorageABI = "[{\"inputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"constructor\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"oldGuardian\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"newGuardian\",\"type\": \"address\"}],\"name\": \"GuardianChanged\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"node\",\"type\": \"address\"},{\"indexed\": true,\"internalType\": \"address\",\"name\": \"withdrawalAddress\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"time\",\"type\": \"uint256\"}],\"name\": \"NodeWithdrawalAddressSet\",\"type\": \"event\"},{\"inputs\": [],\"name\": \"getGuardian\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_newAddress\",\"type\": \"address\"}],\"name\": \"setGuardian\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"confirmGuardian\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getDeployedStatus\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [],\"name\": \"setDeployedStatus\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_nodeAddress\",\"type\": \"address\"}],\"name\": \"getNodeWithdrawalAddress\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_nodeAddress\",\"type\": \"address\"}],\"name\": \"getNodePendingWithdrawalAddress\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_nodeAddress\",\"type\": \"address\"},{\"internalType\": \"address\",\"name\": \"_newWithdrawalAddress\",\"type\": \"address\"},{\"internalType\": \"bool\",\"name\": \"_confirm\",\"type\": \"bool\"}],\"name\": \"setWithdrawalAddress\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_nodeAddress\",\"type\": \"address\"}],\"name\": \"confirmWithdrawalAddress\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getAddress\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"r\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getUint\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"r\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getString\",\"outputs\": [{\"internalType\": \"string\",\"name\": \"\",\"type\": \"string\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getBytes\",\"outputs\": [{\"internalType\": \"bytes\",\"name\": \"\",\"type\": \"bytes\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getBool\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"r\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getInt\",\"outputs\": [{\"internalType\": \"int256\",\"name\": \"r\",\"type\": \"int256\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"getBytes32\",\"outputs\": [{\"internalType\": \"bytes32\",\"name\": \"r\",\"type\": \"bytes32\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"address\",\"name\": \"_value\",\"type\": \"address\"}],\"name\": \"setAddress\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"uint256\",\"name\": \"_value\",\"type\": \"uint256\"}],\"name\": \"setUint\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"string\",\"name\": \"_value\",\"type\": \"string\"}],\"name\": \"setString\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"bytes\",\"name\": \"_value\",\"type\": \"bytes\"}],\"name\": \"setBytes\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"bool\",\"name\": \"_value\",\"type\": \"bool\"}],\"name\": \"setBool\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"int256\",\"name\": \"_value\",\"type\": \"int256\"}],\"name\": \"setInt\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"bytes32\",\"name\": \"_value\",\"type\": \"bytes32\"}],\"name\": \"setBytes32\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteAddress\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteUint\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteString\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteBytes\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteBool\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteInt\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"}],\"name\": \"deleteBytes32\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"uint256\",\"name\": \"_amount\",\"type\": \"uint256\"}],\"name\": \"addUint\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_key\",\"type\": \"bytes32\"},{\"internalType\": \"uint256\",\"name\": \"_amount\",\"type\": \"uint256\"}],\"name\": \"subUint\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]\r\n"

// GoGoStorage is an auto generated Go binding around an Ethereum contract.
type GoGoStorage struct {
	GoGoStorageCaller     // Read-only binding to the contract
	GoGoStorageTransactor // Write-only binding to the contract
	GoGoStorageFilterer   // Log filterer for contract events
}

// GoGoStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoGoStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoGoStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoGoStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoGoStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoGoStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoGoStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoGoStorageSession struct {
	Contract     *GoGoStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoGoStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoGoStorageCallerSession struct {
	Contract *GoGoStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GoGoStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoGoStorageTransactorSession struct {
	Contract     *GoGoStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GoGoStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoGoStorageRaw struct {
	Contract *GoGoStorage // Generic contract binding to access the raw methods on
}

// GoGoStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoGoStorageCallerRaw struct {
	Contract *GoGoStorageCaller // Generic read-only contract binding to access the raw methods on
}

// GoGoStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoGoStorageTransactorRaw struct {
	Contract *GoGoStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoGoStorage creates a new instance of GoGoStorage, bound to a specific deployed contract.
func NewGoGoStorage(address common.Address, backend bind.ContractBackend) (*GoGoStorage, error) {
	contract, err := bindGoGoStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoGoStorage{GoGoStorageCaller: GoGoStorageCaller{contract: contract}, GoGoStorageTransactor: GoGoStorageTransactor{contract: contract}, GoGoStorageFilterer: GoGoStorageFilterer{contract: contract}}, nil
}

// NewGoGoStorageCaller creates a new read-only instance of GoGoStorage, bound to a specific deployed contract.
func NewGoGoStorageCaller(address common.Address, caller bind.ContractCaller) (*GoGoStorageCaller, error) {
	contract, err := bindGoGoStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoGoStorageCaller{contract: contract}, nil
}

// NewGoGoStorageTransactor creates a new write-only instance of GoGoStorage, bound to a specific deployed contract.
func NewGoGoStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*GoGoStorageTransactor, error) {
	contract, err := bindGoGoStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoGoStorageTransactor{contract: contract}, nil
}

// NewGoGoStorageFilterer creates a new log filterer instance of GoGoStorage, bound to a specific deployed contract.
func NewGoGoStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*GoGoStorageFilterer, error) {
	contract, err := bindGoGoStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoGoStorageFilterer{contract: contract}, nil
}

// bindGoGoStorage binds a generic wrapper to an already deployed contract.
func bindGoGoStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoGoStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoGoStorage *GoGoStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoGoStorage.Contract.GoGoStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoGoStorage *GoGoStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoGoStorage.Contract.GoGoStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoGoStorage *GoGoStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoGoStorage.Contract.GoGoStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoGoStorage *GoGoStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoGoStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoGoStorage *GoGoStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoGoStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoGoStorage *GoGoStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoGoStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_GoGoStorage *GoGoStorageCaller) GetAddress(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_GoGoStorage *GoGoStorageSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _GoGoStorage.Contract.GetAddress(&_GoGoStorage.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_GoGoStorage *GoGoStorageCallerSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _GoGoStorage.Contract.GetAddress(&_GoGoStorage.CallOpts, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _key) view returns(bool)
func (_GoGoStorage *GoGoStorageCaller) GetBool(opts *bind.CallOpts, _key [32]byte) (bool, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getBool", _key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _key) view returns(bool)
func (_GoGoStorage *GoGoStorageSession) GetBool(_key [32]byte) (bool, error) {
	return _GoGoStorage.Contract.GetBool(&_GoGoStorage.CallOpts, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _key) view returns(bool)
func (_GoGoStorage *GoGoStorageCallerSession) GetBool(_key [32]byte) (bool, error) {
	return _GoGoStorage.Contract.GetBool(&_GoGoStorage.CallOpts, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (_GoGoStorage *GoGoStorageCaller) GetBytes(opts *bind.CallOpts, _key [32]byte) ([]byte, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getBytes", _key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (_GoGoStorage *GoGoStorageSession) GetBytes(_key [32]byte) ([]byte, error) {
	return _GoGoStorage.Contract.GetBytes(&_GoGoStorage.CallOpts, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (_GoGoStorage *GoGoStorageCallerSession) GetBytes(_key [32]byte) ([]byte, error) {
	return _GoGoStorage.Contract.GetBytes(&_GoGoStorage.CallOpts, _key)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32)
func (_GoGoStorage *GoGoStorageCaller) GetBytes32(opts *bind.CallOpts, _key [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getBytes32", _key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32)
func (_GoGoStorage *GoGoStorageSession) GetBytes32(_key [32]byte) ([32]byte, error) {
	return _GoGoStorage.Contract.GetBytes32(&_GoGoStorage.CallOpts, _key)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32)
func (_GoGoStorage *GoGoStorageCallerSession) GetBytes32(_key [32]byte) ([32]byte, error) {
	return _GoGoStorage.Contract.GetBytes32(&_GoGoStorage.CallOpts, _key)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 _key) view returns(int256)
func (_GoGoStorage *GoGoStorageCaller) GetInt(opts *bind.CallOpts, _key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getInt", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 _key) view returns(int256)
func (_GoGoStorage *GoGoStorageSession) GetInt(_key [32]byte) (*big.Int, error) {
	return _GoGoStorage.Contract.GetInt(&_GoGoStorage.CallOpts, _key)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 _key) view returns(int256)
func (_GoGoStorage *GoGoStorageCallerSession) GetInt(_key [32]byte) (*big.Int, error) {
	return _GoGoStorage.Contract.GetInt(&_GoGoStorage.CallOpts, _key)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (_GoGoStorage *GoGoStorageCaller) GetString(opts *bind.CallOpts, _key [32]byte) (string, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getString", _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (_GoGoStorage *GoGoStorageSession) GetString(_key [32]byte) (string, error) {
	return _GoGoStorage.Contract.GetString(&_GoGoStorage.CallOpts, _key)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (_GoGoStorage *GoGoStorageCallerSession) GetString(_key [32]byte) (string, error) {
	return _GoGoStorage.Contract.GetString(&_GoGoStorage.CallOpts, _key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256)
func (_GoGoStorage *GoGoStorageCaller) GetUint(opts *bind.CallOpts, _key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GoGoStorage.contract.Call(opts, &out, "getUint", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256)
func (_GoGoStorage *GoGoStorageSession) GetUint(_key [32]byte) (*big.Int, error) {
	return _GoGoStorage.Contract.GetUint(&_GoGoStorage.CallOpts, _key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256)
func (_GoGoStorage *GoGoStorageCallerSession) GetUint(_key [32]byte) (*big.Int, error) {
	return _GoGoStorage.Contract.GetUint(&_GoGoStorage.CallOpts, _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteAddress(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteAddress", _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteAddress(&_GoGoStorage.TransactOpts, _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteAddress(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteBool(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteBool", _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBool(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBool(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteBytes(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteBytes", _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBytes(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBytes(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteBytes32(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteBytes32", _key)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteBytes32(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBytes32(&_GoGoStorage.TransactOpts, _key)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteBytes32(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteBytes32(&_GoGoStorage.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteInt(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteInt", _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteInt(&_GoGoStorage.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteInt(&_GoGoStorage.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteString(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteString", _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteString(&_GoGoStorage.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteString(&_GoGoStorage.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactor) DeleteUint(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "deleteUint", _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteUint(&_GoGoStorage.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(bytes32 _key) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.DeleteUint(&_GoGoStorage.TransactOpts, _key)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetAddress(opts *bind.TransactOpts, _key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setAddress", _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetAddress(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetAddress(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _key, bool _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetBool(opts *bind.TransactOpts, _key [32]byte, _value bool) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setBool", _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _key, bool _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBool(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _key, bool _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBool(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 _key, bytes _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetBytes(opts *bind.TransactOpts, _key [32]byte, _value []byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setBytes", _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 _key, bytes _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBytes(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 _key, bytes _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBytes(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _key, bytes32 _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetBytes32(opts *bind.TransactOpts, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setBytes32", _key, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _key, bytes32 _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetBytes32(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBytes32(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _key, bytes32 _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetBytes32(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetBytes32(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 _key, int256 _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetInt(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setInt", _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 _key, int256 _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetInt(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 _key, int256 _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetInt(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 _key, string _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetString(opts *bind.TransactOpts, _key [32]byte, _value string) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setString", _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 _key, string _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetString(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 _key, string _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetString(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _key, uint256 _value) returns()
func (_GoGoStorage *GoGoStorageTransactor) SetUint(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.contract.Transact(opts, "setUint", _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _key, uint256 _value) returns()
func (_GoGoStorage *GoGoStorageSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetUint(&_GoGoStorage.TransactOpts, _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _key, uint256 _value) returns()
func (_GoGoStorage *GoGoStorageTransactorSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GoGoStorage.Contract.SetUint(&_GoGoStorage.TransactOpts, _key, _value)
}

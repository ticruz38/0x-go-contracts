// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IAssetProxyABI is the input ABI used to generate the binding from.
const IAssetProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAuthorizedAddressAtIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetData\",\"type\":\"bytes\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorizedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAssetProxy is an auto generated Go binding around an Ethereum contract.
type IAssetProxy struct {
	IAssetProxyCaller     // Read-only binding to the contract
	IAssetProxyTransactor // Write-only binding to the contract
	IAssetProxyFilterer   // Log filterer for contract events
}

// IAssetProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssetProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssetProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssetProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssetProxySession struct {
	Contract     *IAssetProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAssetProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssetProxyCallerSession struct {
	Contract *IAssetProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IAssetProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssetProxyTransactorSession struct {
	Contract     *IAssetProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IAssetProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssetProxyRaw struct {
	Contract *IAssetProxy // Generic contract binding to access the raw methods on
}

// IAssetProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssetProxyCallerRaw struct {
	Contract *IAssetProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IAssetProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssetProxyTransactorRaw struct {
	Contract *IAssetProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssetProxy creates a new instance of IAssetProxy, bound to a specific deployed contract.
func NewIAssetProxy(address common.Address, backend bind.ContractBackend) (*IAssetProxy, error) {
	contract, err := bindIAssetProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssetProxy{IAssetProxyCaller: IAssetProxyCaller{contract: contract}, IAssetProxyTransactor: IAssetProxyTransactor{contract: contract}, IAssetProxyFilterer: IAssetProxyFilterer{contract: contract}}, nil
}

// NewIAssetProxyCaller creates a new read-only instance of IAssetProxy, bound to a specific deployed contract.
func NewIAssetProxyCaller(address common.Address, caller bind.ContractCaller) (*IAssetProxyCaller, error) {
	contract, err := bindIAssetProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetProxyCaller{contract: contract}, nil
}

// NewIAssetProxyTransactor creates a new write-only instance of IAssetProxy, bound to a specific deployed contract.
func NewIAssetProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssetProxyTransactor, error) {
	contract, err := bindIAssetProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetProxyTransactor{contract: contract}, nil
}

// NewIAssetProxyFilterer creates a new log filterer instance of IAssetProxy, bound to a specific deployed contract.
func NewIAssetProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssetProxyFilterer, error) {
	contract, err := bindIAssetProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssetProxyFilterer{contract: contract}, nil
}

// bindIAssetProxy binds a generic wrapper to an already deployed contract.
func bindIAssetProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAssetProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetProxy *IAssetProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAssetProxy.Contract.IAssetProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetProxy *IAssetProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetProxy.Contract.IAssetProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetProxy *IAssetProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetProxy.Contract.IAssetProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetProxy *IAssetProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAssetProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetProxy *IAssetProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetProxy *IAssetProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetProxy.Contract.contract.Transact(opts, method, params...)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_IAssetProxy *IAssetProxyCaller) GetAuthorizedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IAssetProxy.contract.Call(opts, out, "getAuthorizedAddresses")
	return *ret0, err
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_IAssetProxy *IAssetProxySession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _IAssetProxy.Contract.GetAuthorizedAddresses(&_IAssetProxy.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_IAssetProxy *IAssetProxyCallerSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _IAssetProxy.Contract.GetAuthorizedAddresses(&_IAssetProxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_IAssetProxy *IAssetProxyCaller) GetProxyId(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IAssetProxy.contract.Call(opts, out, "getProxyId")
	return *ret0, err
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_IAssetProxy *IAssetProxySession) GetProxyId() ([4]byte, error) {
	return _IAssetProxy.Contract.GetProxyId(&_IAssetProxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_IAssetProxy *IAssetProxyCallerSession) GetProxyId() ([4]byte, error) {
	return _IAssetProxy.Contract.GetProxyId(&_IAssetProxy.CallOpts)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxyTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.contract.Transact(opts, "addAuthorizedAddress", target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxySession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.AddAuthorizedAddress(&_IAssetProxy.TransactOpts, target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxyTransactorSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.AddAuthorizedAddress(&_IAssetProxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxyTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.contract.Transact(opts, "removeAuthorizedAddress", target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxySession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.RemoveAuthorizedAddress(&_IAssetProxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_IAssetProxy *IAssetProxyTransactorSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.RemoveAuthorizedAddress(&_IAssetProxy.TransactOpts, target)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_IAssetProxy *IAssetProxyTransactor) RemoveAuthorizedAddressAtIndex(opts *bind.TransactOpts, target common.Address, index *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.contract.Transact(opts, "removeAuthorizedAddressAtIndex", target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_IAssetProxy *IAssetProxySession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.Contract.RemoveAuthorizedAddressAtIndex(&_IAssetProxy.TransactOpts, target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_IAssetProxy *IAssetProxyTransactorSession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.Contract.RemoveAuthorizedAddressAtIndex(&_IAssetProxy.TransactOpts, target, index)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xa85e59e4.
//
// Solidity: function transferFrom(bytes assetData, address from, address to, uint256 amount) returns()
func (_IAssetProxy *IAssetProxyTransactor) TransferFrom(opts *bind.TransactOpts, assetData []byte, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.contract.Transact(opts, "transferFrom", assetData, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xa85e59e4.
//
// Solidity: function transferFrom(bytes assetData, address from, address to, uint256 amount) returns()
func (_IAssetProxy *IAssetProxySession) TransferFrom(assetData []byte, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.Contract.TransferFrom(&_IAssetProxy.TransactOpts, assetData, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xa85e59e4.
//
// Solidity: function transferFrom(bytes assetData, address from, address to, uint256 amount) returns()
func (_IAssetProxy *IAssetProxyTransactorSession) TransferFrom(assetData []byte, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAssetProxy.Contract.TransferFrom(&_IAssetProxy.TransactOpts, assetData, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssetProxy *IAssetProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IAssetProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssetProxy *IAssetProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.TransferOwnership(&_IAssetProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssetProxy *IAssetProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IAssetProxy.Contract.TransferOwnership(&_IAssetProxy.TransactOpts, newOwner)
}

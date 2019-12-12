// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// MultiAssetProxyABI is the input ABI used to generate the binding from.
const MultiAssetProxyABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"assetProxies\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetProxyId\",\"type\":\"bytes4\"}],\"name\":\"getAssetProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAuthorizedAddressAtIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetProxy\",\"type\":\"address\"}],\"name\":\"registerAssetProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorizedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes4\"},{\"indexed\":false,\"name\":\"assetProxy\",\"type\":\"address\"}],\"name\":\"AssetProxyRegistered\",\"type\":\"event\"}]"

// MultiAssetProxy is an auto generated Go binding around an Ethereum contract.
type MultiAssetProxy struct {
	MultiAssetProxyCaller     // Read-only binding to the contract
	MultiAssetProxyTransactor // Write-only binding to the contract
	MultiAssetProxyFilterer   // Log filterer for contract events
}

// MultiAssetProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiAssetProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAssetProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiAssetProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAssetProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiAssetProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAssetProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiAssetProxySession struct {
	Contract     *MultiAssetProxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiAssetProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiAssetProxyCallerSession struct {
	Contract *MultiAssetProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MultiAssetProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiAssetProxyTransactorSession struct {
	Contract     *MultiAssetProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MultiAssetProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiAssetProxyRaw struct {
	Contract *MultiAssetProxy // Generic contract binding to access the raw methods on
}

// MultiAssetProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiAssetProxyCallerRaw struct {
	Contract *MultiAssetProxyCaller // Generic read-only contract binding to access the raw methods on
}

// MultiAssetProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiAssetProxyTransactorRaw struct {
	Contract *MultiAssetProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiAssetProxy creates a new instance of MultiAssetProxy, bound to a specific deployed contract.
func NewMultiAssetProxy(address common.Address, backend bind.ContractBackend) (*MultiAssetProxy, error) {
	contract, err := bindMultiAssetProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxy{MultiAssetProxyCaller: MultiAssetProxyCaller{contract: contract}, MultiAssetProxyTransactor: MultiAssetProxyTransactor{contract: contract}, MultiAssetProxyFilterer: MultiAssetProxyFilterer{contract: contract}}, nil
}

// NewMultiAssetProxyCaller creates a new read-only instance of MultiAssetProxy, bound to a specific deployed contract.
func NewMultiAssetProxyCaller(address common.Address, caller bind.ContractCaller) (*MultiAssetProxyCaller, error) {
	contract, err := bindMultiAssetProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyCaller{contract: contract}, nil
}

// NewMultiAssetProxyTransactor creates a new write-only instance of MultiAssetProxy, bound to a specific deployed contract.
func NewMultiAssetProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiAssetProxyTransactor, error) {
	contract, err := bindMultiAssetProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyTransactor{contract: contract}, nil
}

// NewMultiAssetProxyFilterer creates a new log filterer instance of MultiAssetProxy, bound to a specific deployed contract.
func NewMultiAssetProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiAssetProxyFilterer, error) {
	contract, err := bindMultiAssetProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyFilterer{contract: contract}, nil
}

// bindMultiAssetProxy binds a generic wrapper to an already deployed contract.
func bindMultiAssetProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiAssetProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiAssetProxy *MultiAssetProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiAssetProxy.Contract.MultiAssetProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiAssetProxy *MultiAssetProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.MultiAssetProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiAssetProxy *MultiAssetProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.MultiAssetProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiAssetProxy *MultiAssetProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiAssetProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiAssetProxy *MultiAssetProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiAssetProxy *MultiAssetProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetProxies is a free data retrieval call binding the contract method 0x3fd3c997.
//
// Solidity: function assetProxies(bytes4 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCaller) AssetProxies(opts *bind.CallOpts, arg0 [4]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "assetProxies", arg0)
	return *ret0, err
}

// AssetProxies is a free data retrieval call binding the contract method 0x3fd3c997.
//
// Solidity: function assetProxies(bytes4 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxySession) AssetProxies(arg0 [4]byte) (common.Address, error) {
	return _MultiAssetProxy.Contract.AssetProxies(&_MultiAssetProxy.CallOpts, arg0)
}

// AssetProxies is a free data retrieval call binding the contract method 0x3fd3c997.
//
// Solidity: function assetProxies(bytes4 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) AssetProxies(arg0 [4]byte) (common.Address, error) {
	return _MultiAssetProxy.Contract.AssetProxies(&_MultiAssetProxy.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxySession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _MultiAssetProxy.Contract.Authorities(&_MultiAssetProxy.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _MultiAssetProxy.Contract.Authorities(&_MultiAssetProxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_MultiAssetProxy *MultiAssetProxyCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "authorized", arg0)
	return *ret0, err
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_MultiAssetProxy *MultiAssetProxySession) Authorized(arg0 common.Address) (bool, error) {
	return _MultiAssetProxy.Contract.Authorized(&_MultiAssetProxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _MultiAssetProxy.Contract.Authorized(&_MultiAssetProxy.CallOpts, arg0)
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCaller) GetAssetProxy(opts *bind.CallOpts, assetProxyId [4]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "getAssetProxy", assetProxyId)
	return *ret0, err
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxySession) GetAssetProxy(assetProxyId [4]byte) (common.Address, error) {
	return _MultiAssetProxy.Contract.GetAssetProxy(&_MultiAssetProxy.CallOpts, assetProxyId)
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) GetAssetProxy(assetProxyId [4]byte) (common.Address, error) {
	return _MultiAssetProxy.Contract.GetAssetProxy(&_MultiAssetProxy.CallOpts, assetProxyId)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_MultiAssetProxy *MultiAssetProxyCaller) GetAuthorizedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "getAuthorizedAddresses")
	return *ret0, err
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_MultiAssetProxy *MultiAssetProxySession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _MultiAssetProxy.Contract.GetAuthorizedAddresses(&_MultiAssetProxy.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_MultiAssetProxy *MultiAssetProxyCallerSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _MultiAssetProxy.Contract.GetAuthorizedAddresses(&_MultiAssetProxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_MultiAssetProxy *MultiAssetProxyCaller) GetProxyId(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "getProxyId")
	return *ret0, err
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_MultiAssetProxy *MultiAssetProxySession) GetProxyId() ([4]byte, error) {
	return _MultiAssetProxy.Contract.GetProxyId(&_MultiAssetProxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) GetProxyId() ([4]byte, error) {
	return _MultiAssetProxy.Contract.GetProxyId(&_MultiAssetProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAssetProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MultiAssetProxy *MultiAssetProxySession) Owner() (common.Address, error) {
	return _MultiAssetProxy.Contract.Owner(&_MultiAssetProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MultiAssetProxy *MultiAssetProxyCallerSession) Owner() (common.Address, error) {
	return _MultiAssetProxy.Contract.Owner(&_MultiAssetProxy.CallOpts)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.contract.Transact(opts, "addAuthorizedAddress", target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxySession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.AddAuthorizedAddress(&_MultiAssetProxy.TransactOpts, target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactorSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.AddAuthorizedAddress(&_MultiAssetProxy.TransactOpts, target)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactor) RegisterAssetProxy(opts *bind.TransactOpts, assetProxy common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.contract.Transact(opts, "registerAssetProxy", assetProxy)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_MultiAssetProxy *MultiAssetProxySession) RegisterAssetProxy(assetProxy common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RegisterAssetProxy(&_MultiAssetProxy.TransactOpts, assetProxy)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactorSession) RegisterAssetProxy(assetProxy common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RegisterAssetProxy(&_MultiAssetProxy.TransactOpts, assetProxy)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.contract.Transact(opts, "removeAuthorizedAddress", target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxySession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RemoveAuthorizedAddress(&_MultiAssetProxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactorSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RemoveAuthorizedAddress(&_MultiAssetProxy.TransactOpts, target)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactor) RemoveAuthorizedAddressAtIndex(opts *bind.TransactOpts, target common.Address, index *big.Int) (*types.Transaction, error) {
	return _MultiAssetProxy.contract.Transact(opts, "removeAuthorizedAddressAtIndex", target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_MultiAssetProxy *MultiAssetProxySession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RemoveAuthorizedAddressAtIndex(&_MultiAssetProxy.TransactOpts, target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactorSession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.RemoveAuthorizedAddressAtIndex(&_MultiAssetProxy.TransactOpts, target, index)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiAssetProxy *MultiAssetProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.TransferOwnership(&_MultiAssetProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiAssetProxy *MultiAssetProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiAssetProxy.Contract.TransferOwnership(&_MultiAssetProxy.TransactOpts, newOwner)
}

// MultiAssetProxyAssetProxyRegisteredIterator is returned from FilterAssetProxyRegistered and is used to iterate over the raw logs and unpacked data for AssetProxyRegistered events raised by the MultiAssetProxy contract.
type MultiAssetProxyAssetProxyRegisteredIterator struct {
	Event *MultiAssetProxyAssetProxyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiAssetProxyAssetProxyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAssetProxyAssetProxyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiAssetProxyAssetProxyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiAssetProxyAssetProxyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAssetProxyAssetProxyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAssetProxyAssetProxyRegistered represents a AssetProxyRegistered event raised by the MultiAssetProxy contract.
type MultiAssetProxyAssetProxyRegistered struct {
	Id         [4]byte
	AssetProxy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetProxyRegistered is a free log retrieval operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_MultiAssetProxy *MultiAssetProxyFilterer) FilterAssetProxyRegistered(opts *bind.FilterOpts) (*MultiAssetProxyAssetProxyRegisteredIterator, error) {

	logs, sub, err := _MultiAssetProxy.contract.FilterLogs(opts, "AssetProxyRegistered")
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyAssetProxyRegisteredIterator{contract: _MultiAssetProxy.contract, event: "AssetProxyRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetProxyRegistered is a free log subscription operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_MultiAssetProxy *MultiAssetProxyFilterer) WatchAssetProxyRegistered(opts *bind.WatchOpts, sink chan<- *MultiAssetProxyAssetProxyRegistered) (event.Subscription, error) {

	logs, sub, err := _MultiAssetProxy.contract.WatchLogs(opts, "AssetProxyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAssetProxyAssetProxyRegistered)
				if err := _MultiAssetProxy.contract.UnpackLog(event, "AssetProxyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetProxyRegistered is a log parse operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_MultiAssetProxy *MultiAssetProxyFilterer) ParseAssetProxyRegistered(log types.Log) (*MultiAssetProxyAssetProxyRegistered, error) {
	event := new(MultiAssetProxyAssetProxyRegistered)
	if err := _MultiAssetProxy.contract.UnpackLog(event, "AssetProxyRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiAssetProxyAuthorizedAddressAddedIterator is returned from FilterAuthorizedAddressAdded and is used to iterate over the raw logs and unpacked data for AuthorizedAddressAdded events raised by the MultiAssetProxy contract.
type MultiAssetProxyAuthorizedAddressAddedIterator struct {
	Event *MultiAssetProxyAuthorizedAddressAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiAssetProxyAuthorizedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAssetProxyAuthorizedAddressAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiAssetProxyAuthorizedAddressAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiAssetProxyAuthorizedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAssetProxyAuthorizedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAssetProxyAuthorizedAddressAdded represents a AuthorizedAddressAdded event raised by the MultiAssetProxy contract.
type MultiAssetProxyAuthorizedAddressAdded struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressAdded is a free log retrieval operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) FilterAuthorizedAddressAdded(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*MultiAssetProxyAuthorizedAddressAddedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MultiAssetProxy.contract.FilterLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyAuthorizedAddressAddedIterator{contract: _MultiAssetProxy.contract, event: "AuthorizedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressAdded is a free log subscription operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) WatchAuthorizedAddressAdded(opts *bind.WatchOpts, sink chan<- *MultiAssetProxyAuthorizedAddressAdded, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MultiAssetProxy.contract.WatchLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAssetProxyAuthorizedAddressAdded)
				if err := _MultiAssetProxy.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedAddressAdded is a log parse operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) ParseAuthorizedAddressAdded(log types.Log) (*MultiAssetProxyAuthorizedAddressAdded, error) {
	event := new(MultiAssetProxyAuthorizedAddressAdded)
	if err := _MultiAssetProxy.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiAssetProxyAuthorizedAddressRemovedIterator is returned from FilterAuthorizedAddressRemoved and is used to iterate over the raw logs and unpacked data for AuthorizedAddressRemoved events raised by the MultiAssetProxy contract.
type MultiAssetProxyAuthorizedAddressRemovedIterator struct {
	Event *MultiAssetProxyAuthorizedAddressRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiAssetProxyAuthorizedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAssetProxyAuthorizedAddressRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiAssetProxyAuthorizedAddressRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiAssetProxyAuthorizedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAssetProxyAuthorizedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAssetProxyAuthorizedAddressRemoved represents a AuthorizedAddressRemoved event raised by the MultiAssetProxy contract.
type MultiAssetProxyAuthorizedAddressRemoved struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressRemoved is a free log retrieval operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) FilterAuthorizedAddressRemoved(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*MultiAssetProxyAuthorizedAddressRemovedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MultiAssetProxy.contract.FilterLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MultiAssetProxyAuthorizedAddressRemovedIterator{contract: _MultiAssetProxy.contract, event: "AuthorizedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressRemoved is a free log subscription operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) WatchAuthorizedAddressRemoved(opts *bind.WatchOpts, sink chan<- *MultiAssetProxyAuthorizedAddressRemoved, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MultiAssetProxy.contract.WatchLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAssetProxyAuthorizedAddressRemoved)
				if err := _MultiAssetProxy.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedAddressRemoved is a log parse operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_MultiAssetProxy *MultiAssetProxyFilterer) ParseAuthorizedAddressRemoved(log types.Log) (*MultiAssetProxyAuthorizedAddressRemoved, error) {
	event := new(MultiAssetProxyAuthorizedAddressRemoved)
	if err := _MultiAssetProxy.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

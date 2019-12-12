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

// ERC20ProxyABI is the input ABI used to generate the binding from.
const ERC20ProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAuthorizedAddressAtIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorizedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressRemoved\",\"type\":\"event\"}]"

// ERC20Proxy is an auto generated Go binding around an Ethereum contract.
type ERC20Proxy struct {
	ERC20ProxyCaller     // Read-only binding to the contract
	ERC20ProxyTransactor // Write-only binding to the contract
	ERC20ProxyFilterer   // Log filterer for contract events
}

// ERC20ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ProxySession struct {
	Contract     *ERC20Proxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ProxyCallerSession struct {
	Contract *ERC20ProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ProxyTransactorSession struct {
	Contract     *ERC20ProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ProxyRaw struct {
	Contract *ERC20Proxy // Generic contract binding to access the raw methods on
}

// ERC20ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ProxyCallerRaw struct {
	Contract *ERC20ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ProxyTransactorRaw struct {
	Contract *ERC20ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Proxy creates a new instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20Proxy(address common.Address, backend bind.ContractBackend) (*ERC20Proxy, error) {
	contract, err := bindERC20Proxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Proxy{ERC20ProxyCaller: ERC20ProxyCaller{contract: contract}, ERC20ProxyTransactor: ERC20ProxyTransactor{contract: contract}, ERC20ProxyFilterer: ERC20ProxyFilterer{contract: contract}}, nil
}

// NewERC20ProxyCaller creates a new read-only instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyCaller(address common.Address, caller bind.ContractCaller) (*ERC20ProxyCaller, error) {
	contract, err := bindERC20Proxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyCaller{contract: contract}, nil
}

// NewERC20ProxyTransactor creates a new write-only instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ProxyTransactor, error) {
	contract, err := bindERC20Proxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyTransactor{contract: contract}, nil
}

// NewERC20ProxyFilterer creates a new log filterer instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ProxyFilterer, error) {
	contract, err := bindERC20Proxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyFilterer{contract: contract}, nil
}

// bindERC20Proxy binds a generic wrapper to an already deployed contract.
func bindERC20Proxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Proxy *ERC20ProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Proxy.Contract.ERC20ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Proxy *ERC20ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ERC20ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Proxy *ERC20ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ERC20ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Proxy *ERC20ProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Proxy *ERC20ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Proxy *ERC20ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.contract.Transact(opts, method, params...)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_ERC20Proxy *ERC20ProxyCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_ERC20Proxy *ERC20ProxySession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ERC20Proxy.Contract.Authorities(&_ERC20Proxy.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) constant returns(address)
func (_ERC20Proxy *ERC20ProxyCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ERC20Proxy.Contract.Authorities(&_ERC20Proxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_ERC20Proxy *ERC20ProxyCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "authorized", arg0)
	return *ret0, err
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_ERC20Proxy *ERC20ProxySession) Authorized(arg0 common.Address) (bool, error) {
	return _ERC20Proxy.Contract.Authorized(&_ERC20Proxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) constant returns(bool)
func (_ERC20Proxy *ERC20ProxyCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _ERC20Proxy.Contract.Authorized(&_ERC20Proxy.CallOpts, arg0)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_ERC20Proxy *ERC20ProxyCaller) GetAuthorizedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "getAuthorizedAddresses")
	return *ret0, err
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_ERC20Proxy *ERC20ProxySession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _ERC20Proxy.Contract.GetAuthorizedAddresses(&_ERC20Proxy.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_ERC20Proxy *ERC20ProxyCallerSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _ERC20Proxy.Contract.GetAuthorizedAddresses(&_ERC20Proxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_ERC20Proxy *ERC20ProxyCaller) GetProxyId(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "getProxyId")
	return *ret0, err
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_ERC20Proxy *ERC20ProxySession) GetProxyId() ([4]byte, error) {
	return _ERC20Proxy.Contract.GetProxyId(&_ERC20Proxy.CallOpts)
}

// GetProxyId is a free data retrieval call binding the contract method 0xae25532e.
//
// Solidity: function getProxyId() constant returns(bytes4)
func (_ERC20Proxy *ERC20ProxyCallerSession) GetProxyId() ([4]byte, error) {
	return _ERC20Proxy.Contract.GetProxyId(&_ERC20Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Proxy *ERC20ProxySession) Owner() (common.Address, error) {
	return _ERC20Proxy.Contract.Owner(&_ERC20Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCallerSession) Owner() (common.Address, error) {
	return _ERC20Proxy.Contract.Owner(&_ERC20Proxy.CallOpts)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "addAuthorizedAddress", target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxySession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.AddAuthorizedAddress(&_ERC20Proxy.TransactOpts, target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.AddAuthorizedAddress(&_ERC20Proxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "removeAuthorizedAddress", target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxySession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RemoveAuthorizedAddress(&_ERC20Proxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RemoveAuthorizedAddress(&_ERC20Proxy.TransactOpts, target)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) RemoveAuthorizedAddressAtIndex(opts *bind.TransactOpts, target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "removeAuthorizedAddressAtIndex", target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ERC20Proxy *ERC20ProxySession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RemoveAuthorizedAddressAtIndex(&_ERC20Proxy.TransactOpts, target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RemoveAuthorizedAddressAtIndex(&_ERC20Proxy.TransactOpts, target, index)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Proxy *ERC20ProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.TransferOwnership(&_ERC20Proxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.TransferOwnership(&_ERC20Proxy.TransactOpts, newOwner)
}

// ERC20ProxyAuthorizedAddressAddedIterator is returned from FilterAuthorizedAddressAdded and is used to iterate over the raw logs and unpacked data for AuthorizedAddressAdded events raised by the ERC20Proxy contract.
type ERC20ProxyAuthorizedAddressAddedIterator struct {
	Event *ERC20ProxyAuthorizedAddressAdded // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyAuthorizedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyAuthorizedAddressAdded)
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
		it.Event = new(ERC20ProxyAuthorizedAddressAdded)
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
func (it *ERC20ProxyAuthorizedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyAuthorizedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyAuthorizedAddressAdded represents a AuthorizedAddressAdded event raised by the ERC20Proxy contract.
type ERC20ProxyAuthorizedAddressAdded struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressAdded is a free log retrieval operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterAuthorizedAddressAdded(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*ERC20ProxyAuthorizedAddressAddedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyAuthorizedAddressAddedIterator{contract: _ERC20Proxy.contract, event: "AuthorizedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressAdded is a free log subscription operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchAuthorizedAddressAdded(opts *bind.WatchOpts, sink chan<- *ERC20ProxyAuthorizedAddressAdded, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyAuthorizedAddressAdded)
				if err := _ERC20Proxy.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
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
func (_ERC20Proxy *ERC20ProxyFilterer) ParseAuthorizedAddressAdded(log types.Log) (*ERC20ProxyAuthorizedAddressAdded, error) {
	event := new(ERC20ProxyAuthorizedAddressAdded)
	if err := _ERC20Proxy.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20ProxyAuthorizedAddressRemovedIterator is returned from FilterAuthorizedAddressRemoved and is used to iterate over the raw logs and unpacked data for AuthorizedAddressRemoved events raised by the ERC20Proxy contract.
type ERC20ProxyAuthorizedAddressRemovedIterator struct {
	Event *ERC20ProxyAuthorizedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyAuthorizedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyAuthorizedAddressRemoved)
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
		it.Event = new(ERC20ProxyAuthorizedAddressRemoved)
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
func (it *ERC20ProxyAuthorizedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyAuthorizedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyAuthorizedAddressRemoved represents a AuthorizedAddressRemoved event raised by the ERC20Proxy contract.
type ERC20ProxyAuthorizedAddressRemoved struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressRemoved is a free log retrieval operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterAuthorizedAddressRemoved(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*ERC20ProxyAuthorizedAddressRemovedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyAuthorizedAddressRemovedIterator{contract: _ERC20Proxy.contract, event: "AuthorizedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressRemoved is a free log subscription operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchAuthorizedAddressRemoved(opts *bind.WatchOpts, sink chan<- *ERC20ProxyAuthorizedAddressRemoved, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyAuthorizedAddressRemoved)
				if err := _ERC20Proxy.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
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
func (_ERC20Proxy *ERC20ProxyFilterer) ParseAuthorizedAddressRemoved(log types.Log) (*ERC20ProxyAuthorizedAddressRemoved, error) {
	event := new(ERC20ProxyAuthorizedAddressRemoved)
	if err := _ERC20Proxy.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

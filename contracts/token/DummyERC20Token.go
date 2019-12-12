// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// DummyERC20TokenABI is the input ABI used to generate the binding from.
const DummyERC20TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_MINT_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// DummyERC20Token is an auto generated Go binding around an Ethereum contract.
type DummyERC20Token struct {
	DummyERC20TokenCaller     // Read-only binding to the contract
	DummyERC20TokenTransactor // Write-only binding to the contract
	DummyERC20TokenFilterer   // Log filterer for contract events
}

// DummyERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyERC20TokenSession struct {
	Contract     *DummyERC20Token  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyERC20TokenCallerSession struct {
	Contract *DummyERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DummyERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyERC20TokenTransactorSession struct {
	Contract     *DummyERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DummyERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyERC20TokenRaw struct {
	Contract *DummyERC20Token // Generic contract binding to access the raw methods on
}

// DummyERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyERC20TokenCallerRaw struct {
	Contract *DummyERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// DummyERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyERC20TokenTransactorRaw struct {
	Contract *DummyERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyERC20Token creates a new instance of DummyERC20Token, bound to a specific deployed contract.
func NewDummyERC20Token(address common.Address, backend bind.ContractBackend) (*DummyERC20Token, error) {
	contract, err := bindDummyERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyERC20Token{DummyERC20TokenCaller: DummyERC20TokenCaller{contract: contract}, DummyERC20TokenTransactor: DummyERC20TokenTransactor{contract: contract}, DummyERC20TokenFilterer: DummyERC20TokenFilterer{contract: contract}}, nil
}

// NewDummyERC20TokenCaller creates a new read-only instance of DummyERC20Token, bound to a specific deployed contract.
func NewDummyERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*DummyERC20TokenCaller, error) {
	contract, err := bindDummyERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyERC20TokenCaller{contract: contract}, nil
}

// NewDummyERC20TokenTransactor creates a new write-only instance of DummyERC20Token, bound to a specific deployed contract.
func NewDummyERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyERC20TokenTransactor, error) {
	contract, err := bindDummyERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyERC20TokenTransactor{contract: contract}, nil
}

// NewDummyERC20TokenFilterer creates a new log filterer instance of DummyERC20Token, bound to a specific deployed contract.
func NewDummyERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyERC20TokenFilterer, error) {
	contract, err := bindDummyERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyERC20TokenFilterer{contract: contract}, nil
}

// bindDummyERC20Token binds a generic wrapper to an already deployed contract.
func bindDummyERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyERC20Token *DummyERC20TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyERC20Token.Contract.DummyERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyERC20Token *DummyERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.DummyERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyERC20Token *DummyERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.DummyERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyERC20Token *DummyERC20TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyERC20Token *DummyERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyERC20Token *DummyERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.contract.Transact(opts, method, params...)
}

// MAXMINTAMOUNT is a free data retrieval call binding the contract method 0xfa9b7018.
//
// Solidity: function MAX_MINT_AMOUNT() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCaller) MAXMINTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "MAX_MINT_AMOUNT")
	return *ret0, err
}

// MAXMINTAMOUNT is a free data retrieval call binding the contract method 0xfa9b7018.
//
// Solidity: function MAX_MINT_AMOUNT() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenSession) MAXMINTAMOUNT() (*big.Int, error) {
	return _DummyERC20Token.Contract.MAXMINTAMOUNT(&_DummyERC20Token.CallOpts)
}

// MAXMINTAMOUNT is a free data retrieval call binding the contract method 0xfa9b7018.
//
// Solidity: function MAX_MINT_AMOUNT() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCallerSession) MAXMINTAMOUNT() (*big.Int, error) {
	return _DummyERC20Token.Contract.MAXMINTAMOUNT(&_DummyERC20Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DummyERC20Token.Contract.Allowance(&_DummyERC20Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DummyERC20Token.Contract.Allowance(&_DummyERC20Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DummyERC20Token.Contract.BalanceOf(&_DummyERC20Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DummyERC20Token.Contract.BalanceOf(&_DummyERC20Token.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenSession) Decimals() (*big.Int, error) {
	return _DummyERC20Token.Contract.Decimals(&_DummyERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCallerSession) Decimals() (*big.Int, error) {
	return _DummyERC20Token.Contract.Decimals(&_DummyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenSession) Name() (string, error) {
	return _DummyERC20Token.Contract.Name(&_DummyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenCallerSession) Name() (string, error) {
	return _DummyERC20Token.Contract.Name(&_DummyERC20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DummyERC20Token *DummyERC20TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DummyERC20Token *DummyERC20TokenSession) Owner() (common.Address, error) {
	return _DummyERC20Token.Contract.Owner(&_DummyERC20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DummyERC20Token *DummyERC20TokenCallerSession) Owner() (common.Address, error) {
	return _DummyERC20Token.Contract.Owner(&_DummyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenSession) Symbol() (string, error) {
	return _DummyERC20Token.Contract.Symbol(&_DummyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DummyERC20Token *DummyERC20TokenCallerSession) Symbol() (string, error) {
	return _DummyERC20Token.Contract.Symbol(&_DummyERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DummyERC20Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _DummyERC20Token.Contract.TotalSupply(&_DummyERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DummyERC20Token *DummyERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DummyERC20Token.Contract.TotalSupply(&_DummyERC20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Approve(&_DummyERC20Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Approve(&_DummyERC20Token.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenTransactor) Mint(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "mint", _value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenSession) Mint(_value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Mint(&_DummyERC20Token.TransactOpts, _value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenTransactorSession) Mint(_value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Mint(&_DummyERC20Token.TransactOpts, _value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address _target, uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenTransactor) SetBalance(opts *bind.TransactOpts, _target common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "setBalance", _target, _value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address _target, uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenSession) SetBalance(_target common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.SetBalance(&_DummyERC20Token.TransactOpts, _target, _value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address _target, uint256 _value) returns()
func (_DummyERC20Token *DummyERC20TokenTransactorSession) SetBalance(_target common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.SetBalance(&_DummyERC20Token.TransactOpts, _target, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Transfer(&_DummyERC20Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.Transfer(&_DummyERC20Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.TransferFrom(&_DummyERC20Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DummyERC20Token *DummyERC20TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.TransferFrom(&_DummyERC20Token.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DummyERC20Token *DummyERC20TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DummyERC20Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DummyERC20Token *DummyERC20TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.TransferOwnership(&_DummyERC20Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DummyERC20Token *DummyERC20TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DummyERC20Token.Contract.TransferOwnership(&_DummyERC20Token.TransactOpts, newOwner)
}

// DummyERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DummyERC20Token contract.
type DummyERC20TokenApprovalIterator struct {
	Event *DummyERC20TokenApproval // Event containing the contract specifics and raw log

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
func (it *DummyERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyERC20TokenApproval)
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
		it.Event = new(DummyERC20TokenApproval)
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
func (it *DummyERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyERC20TokenApproval represents a Approval event raised by the DummyERC20Token contract.
type DummyERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*DummyERC20TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DummyERC20Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &DummyERC20TokenApprovalIterator{contract: _DummyERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DummyERC20TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DummyERC20Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyERC20TokenApproval)
				if err := _DummyERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) ParseApproval(log types.Log) (*DummyERC20TokenApproval, error) {
	event := new(DummyERC20TokenApproval)
	if err := _DummyERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DummyERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DummyERC20Token contract.
type DummyERC20TokenTransferIterator struct {
	Event *DummyERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *DummyERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyERC20TokenTransfer)
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
		it.Event = new(DummyERC20TokenTransfer)
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
func (it *DummyERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyERC20TokenTransfer represents a Transfer event raised by the DummyERC20Token contract.
type DummyERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DummyERC20TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DummyERC20Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DummyERC20TokenTransferIterator{contract: _DummyERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DummyERC20TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DummyERC20Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyERC20TokenTransfer)
				if err := _DummyERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DummyERC20Token *DummyERC20TokenFilterer) ParseTransfer(log types.Log) (*DummyERC20TokenTransfer, error) {
	event := new(DummyERC20TokenTransfer)
	if err := _DummyERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

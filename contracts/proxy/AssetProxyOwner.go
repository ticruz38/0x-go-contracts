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

// AssetProxyOwnerABI is the input ABI used to generate the binding from.
const AssetProxyOwnerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeRemoveAuthorizedAddressAtIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsTimeLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetProxyContract\",\"type\":\"address\"},{\"name\":\"isRegistered\",\"type\":\"bool\"}],\"name\":\"registerAssetProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_secondsTimeLocked\",\"type\":\"uint256\"}],\"name\":\"changeTimeLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAssetProxyRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmationTimes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_assetProxyContracts\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_secondsTimeLocked\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"assetProxyContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isRegistered\",\"type\":\"bool\"}],\"name\":\"AssetProxyRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmationTime\",\"type\":\"uint256\"}],\"name\":\"ConfirmationTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"secondsTimeLocked\",\"type\":\"uint256\"}],\"name\":\"TimeLockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]"

// AssetProxyOwner is an auto generated Go binding around an Ethereum contract.
type AssetProxyOwner struct {
	AssetProxyOwnerCaller     // Read-only binding to the contract
	AssetProxyOwnerTransactor // Write-only binding to the contract
	AssetProxyOwnerFilterer   // Log filterer for contract events
}

// AssetProxyOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetProxyOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetProxyOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetProxyOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetProxyOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetProxyOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetProxyOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetProxyOwnerSession struct {
	Contract     *AssetProxyOwner  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetProxyOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetProxyOwnerCallerSession struct {
	Contract *AssetProxyOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetProxyOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetProxyOwnerTransactorSession struct {
	Contract     *AssetProxyOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetProxyOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetProxyOwnerRaw struct {
	Contract *AssetProxyOwner // Generic contract binding to access the raw methods on
}

// AssetProxyOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetProxyOwnerCallerRaw struct {
	Contract *AssetProxyOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// AssetProxyOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetProxyOwnerTransactorRaw struct {
	Contract *AssetProxyOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetProxyOwner creates a new instance of AssetProxyOwner, bound to a specific deployed contract.
func NewAssetProxyOwner(address common.Address, backend bind.ContractBackend) (*AssetProxyOwner, error) {
	contract, err := bindAssetProxyOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwner{AssetProxyOwnerCaller: AssetProxyOwnerCaller{contract: contract}, AssetProxyOwnerTransactor: AssetProxyOwnerTransactor{contract: contract}, AssetProxyOwnerFilterer: AssetProxyOwnerFilterer{contract: contract}}, nil
}

// NewAssetProxyOwnerCaller creates a new read-only instance of AssetProxyOwner, bound to a specific deployed contract.
func NewAssetProxyOwnerCaller(address common.Address, caller bind.ContractCaller) (*AssetProxyOwnerCaller, error) {
	contract, err := bindAssetProxyOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerCaller{contract: contract}, nil
}

// NewAssetProxyOwnerTransactor creates a new write-only instance of AssetProxyOwner, bound to a specific deployed contract.
func NewAssetProxyOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetProxyOwnerTransactor, error) {
	contract, err := bindAssetProxyOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerTransactor{contract: contract}, nil
}

// NewAssetProxyOwnerFilterer creates a new log filterer instance of AssetProxyOwner, bound to a specific deployed contract.
func NewAssetProxyOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetProxyOwnerFilterer, error) {
	contract, err := bindAssetProxyOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerFilterer{contract: contract}, nil
}

// bindAssetProxyOwner binds a generic wrapper to an already deployed contract.
func bindAssetProxyOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetProxyOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetProxyOwner *AssetProxyOwnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetProxyOwner.Contract.AssetProxyOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetProxyOwner *AssetProxyOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.AssetProxyOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetProxyOwner *AssetProxyOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.AssetProxyOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetProxyOwner *AssetProxyOwnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetProxyOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetProxyOwner *AssetProxyOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetProxyOwner *AssetProxyOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _AssetProxyOwner.Contract.MAXOWNERCOUNT(&_AssetProxyOwner.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _AssetProxyOwner.Contract.MAXOWNERCOUNT(&_AssetProxyOwner.CallOpts)
}

// ConfirmationTimes is a free data retrieval call binding the contract method 0xd38f2d82.
//
// Solidity: function confirmationTimes(uint256 ) constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCaller) ConfirmationTimes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "confirmationTimes", arg0)
	return *ret0, err
}

// ConfirmationTimes is a free data retrieval call binding the contract method 0xd38f2d82.
//
// Solidity: function confirmationTimes(uint256 ) constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerSession) ConfirmationTimes(arg0 *big.Int) (*big.Int, error) {
	return _AssetProxyOwner.Contract.ConfirmationTimes(&_AssetProxyOwner.CallOpts, arg0)
}

// ConfirmationTimes is a free data retrieval call binding the contract method 0xd38f2d82.
//
// Solidity: function confirmationTimes(uint256 ) constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) ConfirmationTimes(arg0 *big.Int) (*big.Int, error) {
	return _AssetProxyOwner.Contract.ConfirmationTimes(&_AssetProxyOwner.CallOpts, arg0)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.Confirmations(&_AssetProxyOwner.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.Confirmations(&_AssetProxyOwner.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _AssetProxyOwner.Contract.GetConfirmationCount(&_AssetProxyOwner.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _AssetProxyOwner.Contract.GetConfirmationCount(&_AssetProxyOwner.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_AssetProxyOwner *AssetProxyOwnerCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_AssetProxyOwner *AssetProxyOwnerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _AssetProxyOwner.Contract.GetConfirmations(&_AssetProxyOwner.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _AssetProxyOwner.Contract.GetConfirmations(&_AssetProxyOwner.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_AssetProxyOwner *AssetProxyOwnerCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_AssetProxyOwner *AssetProxyOwnerSession) GetOwners() ([]common.Address, error) {
	return _AssetProxyOwner.Contract.GetOwners(&_AssetProxyOwner.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) GetOwners() ([]common.Address, error) {
	return _AssetProxyOwner.Contract.GetOwners(&_AssetProxyOwner.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _AssetProxyOwner.Contract.GetTransactionCount(&_AssetProxyOwner.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _AssetProxyOwner.Contract.GetTransactionCount(&_AssetProxyOwner.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_AssetProxyOwner *AssetProxyOwnerCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_AssetProxyOwner *AssetProxyOwnerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _AssetProxyOwner.Contract.GetTransactionIds(&_AssetProxyOwner.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _AssetProxyOwner.Contract.GetTransactionIds(&_AssetProxyOwner.CallOpts, from, to, pending, executed)
}

// IsAssetProxyRegistered is a free data retrieval call binding the contract method 0x83250f79.
//
// Solidity: function isAssetProxyRegistered(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCaller) IsAssetProxyRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "isAssetProxyRegistered", arg0)
	return *ret0, err
}

// IsAssetProxyRegistered is a free data retrieval call binding the contract method 0x83250f79.
//
// Solidity: function isAssetProxyRegistered(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerSession) IsAssetProxyRegistered(arg0 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.IsAssetProxyRegistered(&_AssetProxyOwner.CallOpts, arg0)
}

// IsAssetProxyRegistered is a free data retrieval call binding the contract method 0x83250f79.
//
// Solidity: function isAssetProxyRegistered(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) IsAssetProxyRegistered(arg0 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.IsAssetProxyRegistered(&_AssetProxyOwner.CallOpts, arg0)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _AssetProxyOwner.Contract.IsConfirmed(&_AssetProxyOwner.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _AssetProxyOwner.Contract.IsConfirmed(&_AssetProxyOwner.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.IsOwner(&_AssetProxyOwner.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _AssetProxyOwner.Contract.IsOwner(&_AssetProxyOwner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_AssetProxyOwner *AssetProxyOwnerCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_AssetProxyOwner *AssetProxyOwnerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _AssetProxyOwner.Contract.Owners(&_AssetProxyOwner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _AssetProxyOwner.Contract.Owners(&_AssetProxyOwner.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerSession) Required() (*big.Int, error) {
	return _AssetProxyOwner.Contract.Required(&_AssetProxyOwner.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) Required() (*big.Int, error) {
	return _AssetProxyOwner.Contract.Required(&_AssetProxyOwner.CallOpts)
}

// SecondsTimeLocked is a free data retrieval call binding the contract method 0x37bd78a0.
//
// Solidity: function secondsTimeLocked() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCaller) SecondsTimeLocked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "secondsTimeLocked")
	return *ret0, err
}

// SecondsTimeLocked is a free data retrieval call binding the contract method 0x37bd78a0.
//
// Solidity: function secondsTimeLocked() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerSession) SecondsTimeLocked() (*big.Int, error) {
	return _AssetProxyOwner.Contract.SecondsTimeLocked(&_AssetProxyOwner.CallOpts)
}

// SecondsTimeLocked is a free data retrieval call binding the contract method 0x37bd78a0.
//
// Solidity: function secondsTimeLocked() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) SecondsTimeLocked() (*big.Int, error) {
	return _AssetProxyOwner.Contract.SecondsTimeLocked(&_AssetProxyOwner.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AssetProxyOwner.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerSession) TransactionCount() (*big.Int, error) {
	return _AssetProxyOwner.Contract.TransactionCount(&_AssetProxyOwner.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) TransactionCount() (*big.Int, error) {
	return _AssetProxyOwner.Contract.TransactionCount(&_AssetProxyOwner.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_AssetProxyOwner *AssetProxyOwnerCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _AssetProxyOwner.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_AssetProxyOwner *AssetProxyOwnerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _AssetProxyOwner.Contract.Transactions(&_AssetProxyOwner.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_AssetProxyOwner *AssetProxyOwnerCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _AssetProxyOwner.Contract.Transactions(&_AssetProxyOwner.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.AddOwner(&_AssetProxyOwner.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.AddOwner(&_AssetProxyOwner.TransactOpts, owner)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ChangeRequirement(&_AssetProxyOwner.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ChangeRequirement(&_AssetProxyOwner.TransactOpts, _required)
}

// ChangeTimeLock is a paid mutator transaction binding the contract method 0x7ad28c51.
//
// Solidity: function changeTimeLock(uint256 _secondsTimeLocked) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ChangeTimeLock(opts *bind.TransactOpts, _secondsTimeLocked *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "changeTimeLock", _secondsTimeLocked)
}

// ChangeTimeLock is a paid mutator transaction binding the contract method 0x7ad28c51.
//
// Solidity: function changeTimeLock(uint256 _secondsTimeLocked) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ChangeTimeLock(_secondsTimeLocked *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ChangeTimeLock(&_AssetProxyOwner.TransactOpts, _secondsTimeLocked)
}

// ChangeTimeLock is a paid mutator transaction binding the contract method 0x7ad28c51.
//
// Solidity: function changeTimeLock(uint256 _secondsTimeLocked) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ChangeTimeLock(_secondsTimeLocked *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ChangeTimeLock(&_AssetProxyOwner.TransactOpts, _secondsTimeLocked)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ConfirmTransaction(&_AssetProxyOwner.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ConfirmTransaction(&_AssetProxyOwner.TransactOpts, transactionId)
}

// ExecuteRemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x3589b35c.
//
// Solidity: function executeRemoveAuthorizedAddressAtIndex(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ExecuteRemoveAuthorizedAddressAtIndex(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "executeRemoveAuthorizedAddressAtIndex", transactionId)
}

// ExecuteRemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x3589b35c.
//
// Solidity: function executeRemoveAuthorizedAddressAtIndex(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ExecuteRemoveAuthorizedAddressAtIndex(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ExecuteRemoveAuthorizedAddressAtIndex(&_AssetProxyOwner.TransactOpts, transactionId)
}

// ExecuteRemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x3589b35c.
//
// Solidity: function executeRemoveAuthorizedAddressAtIndex(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ExecuteRemoveAuthorizedAddressAtIndex(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ExecuteRemoveAuthorizedAddressAtIndex(&_AssetProxyOwner.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ExecuteTransaction(&_AssetProxyOwner.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ExecuteTransaction(&_AssetProxyOwner.TransactOpts, transactionId)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0x5a1a66af.
//
// Solidity: function registerAssetProxy(address assetProxyContract, bool isRegistered) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) RegisterAssetProxy(opts *bind.TransactOpts, assetProxyContract common.Address, isRegistered bool) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "registerAssetProxy", assetProxyContract, isRegistered)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0x5a1a66af.
//
// Solidity: function registerAssetProxy(address assetProxyContract, bool isRegistered) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) RegisterAssetProxy(assetProxyContract common.Address, isRegistered bool) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RegisterAssetProxy(&_AssetProxyOwner.TransactOpts, assetProxyContract, isRegistered)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0x5a1a66af.
//
// Solidity: function registerAssetProxy(address assetProxyContract, bool isRegistered) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) RegisterAssetProxy(assetProxyContract common.Address, isRegistered bool) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RegisterAssetProxy(&_AssetProxyOwner.TransactOpts, assetProxyContract, isRegistered)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RemoveOwner(&_AssetProxyOwner.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RemoveOwner(&_AssetProxyOwner.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ReplaceOwner(&_AssetProxyOwner.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.ReplaceOwner(&_AssetProxyOwner.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RevokeConfirmation(&_AssetProxyOwner.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.RevokeConfirmation(&_AssetProxyOwner.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_AssetProxyOwner *AssetProxyOwnerTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AssetProxyOwner.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_AssetProxyOwner *AssetProxyOwnerSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.SubmitTransaction(&_AssetProxyOwner.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_AssetProxyOwner *AssetProxyOwnerTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AssetProxyOwner.Contract.SubmitTransaction(&_AssetProxyOwner.TransactOpts, destination, value, data)
}

// AssetProxyOwnerAssetProxyRegistrationIterator is returned from FilterAssetProxyRegistration and is used to iterate over the raw logs and unpacked data for AssetProxyRegistration events raised by the AssetProxyOwner contract.
type AssetProxyOwnerAssetProxyRegistrationIterator struct {
	Event *AssetProxyOwnerAssetProxyRegistration // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerAssetProxyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerAssetProxyRegistration)
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
		it.Event = new(AssetProxyOwnerAssetProxyRegistration)
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
func (it *AssetProxyOwnerAssetProxyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerAssetProxyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerAssetProxyRegistration represents a AssetProxyRegistration event raised by the AssetProxyOwner contract.
type AssetProxyOwnerAssetProxyRegistration struct {
	AssetProxyContract common.Address
	IsRegistered       bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAssetProxyRegistration is a free log retrieval operation binding the contract event 0xdaef8ff7dc66c5e34eb9c338aab679d9f427f89868d9228494455a4d982eb2b0.
//
// Solidity: event AssetProxyRegistration(address assetProxyContract, bool isRegistered)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterAssetProxyRegistration(opts *bind.FilterOpts) (*AssetProxyOwnerAssetProxyRegistrationIterator, error) {

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "AssetProxyRegistration")
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerAssetProxyRegistrationIterator{contract: _AssetProxyOwner.contract, event: "AssetProxyRegistration", logs: logs, sub: sub}, nil
}

// WatchAssetProxyRegistration is a free log subscription operation binding the contract event 0xdaef8ff7dc66c5e34eb9c338aab679d9f427f89868d9228494455a4d982eb2b0.
//
// Solidity: event AssetProxyRegistration(address assetProxyContract, bool isRegistered)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchAssetProxyRegistration(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerAssetProxyRegistration) (event.Subscription, error) {

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "AssetProxyRegistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerAssetProxyRegistration)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "AssetProxyRegistration", log); err != nil {
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

// ParseAssetProxyRegistration is a log parse operation binding the contract event 0xdaef8ff7dc66c5e34eb9c338aab679d9f427f89868d9228494455a4d982eb2b0.
//
// Solidity: event AssetProxyRegistration(address assetProxyContract, bool isRegistered)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseAssetProxyRegistration(log types.Log) (*AssetProxyOwnerAssetProxyRegistration, error) {
	event := new(AssetProxyOwnerAssetProxyRegistration)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "AssetProxyRegistration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the AssetProxyOwner contract.
type AssetProxyOwnerConfirmationIterator struct {
	Event *AssetProxyOwnerConfirmation // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerConfirmation)
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
		it.Event = new(AssetProxyOwnerConfirmation)
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
func (it *AssetProxyOwnerConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerConfirmation represents a Confirmation event raised by the AssetProxyOwner contract.
type AssetProxyOwnerConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*AssetProxyOwnerConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerConfirmationIterator{contract: _AssetProxyOwner.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerConfirmation)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseConfirmation(log types.Log) (*AssetProxyOwnerConfirmation, error) {
	event := new(AssetProxyOwnerConfirmation)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerConfirmationTimeSetIterator is returned from FilterConfirmationTimeSet and is used to iterate over the raw logs and unpacked data for ConfirmationTimeSet events raised by the AssetProxyOwner contract.
type AssetProxyOwnerConfirmationTimeSetIterator struct {
	Event *AssetProxyOwnerConfirmationTimeSet // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerConfirmationTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerConfirmationTimeSet)
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
		it.Event = new(AssetProxyOwnerConfirmationTimeSet)
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
func (it *AssetProxyOwnerConfirmationTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerConfirmationTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerConfirmationTimeSet represents a ConfirmationTimeSet event raised by the AssetProxyOwner contract.
type AssetProxyOwnerConfirmationTimeSet struct {
	TransactionId    *big.Int
	ConfirmationTime *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConfirmationTimeSet is a free log retrieval operation binding the contract event 0x0b237afe65f1514fd7ea3f923ea4fe792bdd07000a912b6cd1602a8e7f573c8d.
//
// Solidity: event ConfirmationTimeSet(uint256 indexed transactionId, uint256 confirmationTime)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterConfirmationTimeSet(opts *bind.FilterOpts, transactionId []*big.Int) (*AssetProxyOwnerConfirmationTimeSetIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "ConfirmationTimeSet", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerConfirmationTimeSetIterator{contract: _AssetProxyOwner.contract, event: "ConfirmationTimeSet", logs: logs, sub: sub}, nil
}

// WatchConfirmationTimeSet is a free log subscription operation binding the contract event 0x0b237afe65f1514fd7ea3f923ea4fe792bdd07000a912b6cd1602a8e7f573c8d.
//
// Solidity: event ConfirmationTimeSet(uint256 indexed transactionId, uint256 confirmationTime)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchConfirmationTimeSet(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerConfirmationTimeSet, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "ConfirmationTimeSet", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerConfirmationTimeSet)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "ConfirmationTimeSet", log); err != nil {
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

// ParseConfirmationTimeSet is a log parse operation binding the contract event 0x0b237afe65f1514fd7ea3f923ea4fe792bdd07000a912b6cd1602a8e7f573c8d.
//
// Solidity: event ConfirmationTimeSet(uint256 indexed transactionId, uint256 confirmationTime)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseConfirmationTimeSet(log types.Log) (*AssetProxyOwnerConfirmationTimeSet, error) {
	event := new(AssetProxyOwnerConfirmationTimeSet)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "ConfirmationTimeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AssetProxyOwner contract.
type AssetProxyOwnerDepositIterator struct {
	Event *AssetProxyOwnerDeposit // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerDeposit)
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
		it.Event = new(AssetProxyOwnerDeposit)
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
func (it *AssetProxyOwnerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerDeposit represents a Deposit event raised by the AssetProxyOwner contract.
type AssetProxyOwnerDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*AssetProxyOwnerDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerDepositIterator{contract: _AssetProxyOwner.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerDeposit)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseDeposit(log types.Log) (*AssetProxyOwnerDeposit, error) {
	event := new(AssetProxyOwnerDeposit)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the AssetProxyOwner contract.
type AssetProxyOwnerExecutionIterator struct {
	Event *AssetProxyOwnerExecution // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerExecution)
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
		it.Event = new(AssetProxyOwnerExecution)
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
func (it *AssetProxyOwnerExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerExecution represents a Execution event raised by the AssetProxyOwner contract.
type AssetProxyOwnerExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*AssetProxyOwnerExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerExecutionIterator{contract: _AssetProxyOwner.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerExecution)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "Execution", log); err != nil {
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

// ParseExecution is a log parse operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseExecution(log types.Log) (*AssetProxyOwnerExecution, error) {
	event := new(AssetProxyOwnerExecution)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the AssetProxyOwner contract.
type AssetProxyOwnerExecutionFailureIterator struct {
	Event *AssetProxyOwnerExecutionFailure // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerExecutionFailure)
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
		it.Event = new(AssetProxyOwnerExecutionFailure)
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
func (it *AssetProxyOwnerExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerExecutionFailure represents a ExecutionFailure event raised by the AssetProxyOwner contract.
type AssetProxyOwnerExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*AssetProxyOwnerExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerExecutionFailureIterator{contract: _AssetProxyOwner.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerExecutionFailure)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseExecutionFailure(log types.Log) (*AssetProxyOwnerExecutionFailure, error) {
	event := new(AssetProxyOwnerExecutionFailure)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the AssetProxyOwner contract.
type AssetProxyOwnerOwnerAdditionIterator struct {
	Event *AssetProxyOwnerOwnerAddition // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerOwnerAddition)
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
		it.Event = new(AssetProxyOwnerOwnerAddition)
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
func (it *AssetProxyOwnerOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerOwnerAddition represents a OwnerAddition event raised by the AssetProxyOwner contract.
type AssetProxyOwnerOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*AssetProxyOwnerOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerOwnerAdditionIterator{contract: _AssetProxyOwner.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerOwnerAddition)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseOwnerAddition(log types.Log) (*AssetProxyOwnerOwnerAddition, error) {
	event := new(AssetProxyOwnerOwnerAddition)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the AssetProxyOwner contract.
type AssetProxyOwnerOwnerRemovalIterator struct {
	Event *AssetProxyOwnerOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerOwnerRemoval)
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
		it.Event = new(AssetProxyOwnerOwnerRemoval)
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
func (it *AssetProxyOwnerOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerOwnerRemoval represents a OwnerRemoval event raised by the AssetProxyOwner contract.
type AssetProxyOwnerOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*AssetProxyOwnerOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerOwnerRemovalIterator{contract: _AssetProxyOwner.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerOwnerRemoval)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseOwnerRemoval(log types.Log) (*AssetProxyOwnerOwnerRemoval, error) {
	event := new(AssetProxyOwnerOwnerRemoval)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the AssetProxyOwner contract.
type AssetProxyOwnerRequirementChangeIterator struct {
	Event *AssetProxyOwnerRequirementChange // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerRequirementChange)
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
		it.Event = new(AssetProxyOwnerRequirementChange)
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
func (it *AssetProxyOwnerRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerRequirementChange represents a RequirementChange event raised by the AssetProxyOwner contract.
type AssetProxyOwnerRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*AssetProxyOwnerRequirementChangeIterator, error) {

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerRequirementChangeIterator{contract: _AssetProxyOwner.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerRequirementChange) (event.Subscription, error) {

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerRequirementChange)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseRequirementChange(log types.Log) (*AssetProxyOwnerRequirementChange, error) {
	event := new(AssetProxyOwnerRequirementChange)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the AssetProxyOwner contract.
type AssetProxyOwnerRevocationIterator struct {
	Event *AssetProxyOwnerRevocation // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerRevocation)
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
		it.Event = new(AssetProxyOwnerRevocation)
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
func (it *AssetProxyOwnerRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerRevocation represents a Revocation event raised by the AssetProxyOwner contract.
type AssetProxyOwnerRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*AssetProxyOwnerRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerRevocationIterator{contract: _AssetProxyOwner.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerRevocation)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseRevocation(log types.Log) (*AssetProxyOwnerRevocation, error) {
	event := new(AssetProxyOwnerRevocation)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the AssetProxyOwner contract.
type AssetProxyOwnerSubmissionIterator struct {
	Event *AssetProxyOwnerSubmission // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerSubmission)
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
		it.Event = new(AssetProxyOwnerSubmission)
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
func (it *AssetProxyOwnerSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerSubmission represents a Submission event raised by the AssetProxyOwner contract.
type AssetProxyOwnerSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*AssetProxyOwnerSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerSubmissionIterator{contract: _AssetProxyOwner.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerSubmission)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "Submission", log); err != nil {
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

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseSubmission(log types.Log) (*AssetProxyOwnerSubmission, error) {
	event := new(AssetProxyOwnerSubmission)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetProxyOwnerTimeLockChangeIterator is returned from FilterTimeLockChange and is used to iterate over the raw logs and unpacked data for TimeLockChange events raised by the AssetProxyOwner contract.
type AssetProxyOwnerTimeLockChangeIterator struct {
	Event *AssetProxyOwnerTimeLockChange // Event containing the contract specifics and raw log

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
func (it *AssetProxyOwnerTimeLockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetProxyOwnerTimeLockChange)
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
		it.Event = new(AssetProxyOwnerTimeLockChange)
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
func (it *AssetProxyOwnerTimeLockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetProxyOwnerTimeLockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetProxyOwnerTimeLockChange represents a TimeLockChange event raised by the AssetProxyOwner contract.
type AssetProxyOwnerTimeLockChange struct {
	SecondsTimeLocked *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTimeLockChange is a free log retrieval operation binding the contract event 0xd1c9101a34feff75cccef14a28785a0279cb0b49c1f321f21f5f422e746b4377.
//
// Solidity: event TimeLockChange(uint256 secondsTimeLocked)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) FilterTimeLockChange(opts *bind.FilterOpts) (*AssetProxyOwnerTimeLockChangeIterator, error) {

	logs, sub, err := _AssetProxyOwner.contract.FilterLogs(opts, "TimeLockChange")
	if err != nil {
		return nil, err
	}
	return &AssetProxyOwnerTimeLockChangeIterator{contract: _AssetProxyOwner.contract, event: "TimeLockChange", logs: logs, sub: sub}, nil
}

// WatchTimeLockChange is a free log subscription operation binding the contract event 0xd1c9101a34feff75cccef14a28785a0279cb0b49c1f321f21f5f422e746b4377.
//
// Solidity: event TimeLockChange(uint256 secondsTimeLocked)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) WatchTimeLockChange(opts *bind.WatchOpts, sink chan<- *AssetProxyOwnerTimeLockChange) (event.Subscription, error) {

	logs, sub, err := _AssetProxyOwner.contract.WatchLogs(opts, "TimeLockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetProxyOwnerTimeLockChange)
				if err := _AssetProxyOwner.contract.UnpackLog(event, "TimeLockChange", log); err != nil {
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

// ParseTimeLockChange is a log parse operation binding the contract event 0xd1c9101a34feff75cccef14a28785a0279cb0b49c1f321f21f5f422e746b4377.
//
// Solidity: event TimeLockChange(uint256 secondsTimeLocked)
func (_AssetProxyOwner *AssetProxyOwnerFilterer) ParseTimeLockChange(log types.Log) (*AssetProxyOwnerTimeLockChange, error) {
	event := new(AssetProxyOwnerTimeLockChange)
	if err := _AssetProxyOwner.contract.UnpackLog(event, "TimeLockChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

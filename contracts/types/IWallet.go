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

// IWalletABI is the input ABI used to generate the binding from.
const IWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IWallet is an auto generated Go binding around an Ethereum contract.
type IWallet struct {
	IWalletCaller     // Read-only binding to the contract
	IWalletTransactor // Write-only binding to the contract
	IWalletFilterer   // Log filterer for contract events
}

// IWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWalletSession struct {
	Contract     *IWallet          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWalletCallerSession struct {
	Contract *IWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWalletTransactorSession struct {
	Contract     *IWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWalletRaw struct {
	Contract *IWallet // Generic contract binding to access the raw methods on
}

// IWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWalletCallerRaw struct {
	Contract *IWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWalletTransactorRaw struct {
	Contract *IWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWallet creates a new instance of IWallet, bound to a specific deployed contract.
func NewIWallet(address common.Address, backend bind.ContractBackend) (*IWallet, error) {
	contract, err := bindIWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWallet{IWalletCaller: IWalletCaller{contract: contract}, IWalletTransactor: IWalletTransactor{contract: contract}, IWalletFilterer: IWalletFilterer{contract: contract}}, nil
}

// NewIWalletCaller creates a new read-only instance of IWallet, bound to a specific deployed contract.
func NewIWalletCaller(address common.Address, caller bind.ContractCaller) (*IWalletCaller, error) {
	contract, err := bindIWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWalletCaller{contract: contract}, nil
}

// NewIWalletTransactor creates a new write-only instance of IWallet, bound to a specific deployed contract.
func NewIWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IWalletTransactor, error) {
	contract, err := bindIWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWalletTransactor{contract: contract}, nil
}

// NewIWalletFilterer creates a new log filterer instance of IWallet, bound to a specific deployed contract.
func NewIWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IWalletFilterer, error) {
	contract, err := bindIWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWalletFilterer{contract: contract}, nil
}

// bindIWallet binds a generic wrapper to an already deployed contract.
func bindIWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWallet *IWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWallet.Contract.IWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWallet *IWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWallet.Contract.IWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWallet *IWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWallet.Contract.IWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWallet *IWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWallet *IWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWallet *IWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWallet.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) constant returns(bytes4)
func (_IWallet *IWalletCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IWallet.contract.Call(opts, out, "isValidSignature", hash, signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) constant returns(bytes4)
func (_IWallet *IWalletSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _IWallet.Contract.IsValidSignature(&_IWallet.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) constant returns(bytes4)
func (_IWallet *IWalletCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _IWallet.Contract.IsValidSignature(&_IWallet.CallOpts, hash, signature)
}

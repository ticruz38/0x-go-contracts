// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dutchAuction

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

// DutchAuctionABI is the input ABI used to generate the binding from.
const DutchAuctionABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"makerAddress\",\"type\":\"address\"},{\"name\":\"takerAddress\",\"type\":\"address\"},{\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"name\":\"senderAddress\",\"type\":\"address\"},{\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"makerFee\",\"type\":\"uint256\"},{\"name\":\"takerFee\",\"type\":\"uint256\"},{\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"},{\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getAuctionDetails\",\"outputs\":[{\"components\":[{\"name\":\"beginTimeSeconds\",\"type\":\"uint256\"},{\"name\":\"endTimeSeconds\",\"type\":\"uint256\"},{\"name\":\"beginAmount\",\"type\":\"uint256\"},{\"name\":\"endAmount\",\"type\":\"uint256\"},{\"name\":\"currentAmount\",\"type\":\"uint256\"},{\"name\":\"currentTimeSeconds\",\"type\":\"uint256\"}],\"name\":\"auctionDetails\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"makerAddress\",\"type\":\"address\"},{\"name\":\"takerAddress\",\"type\":\"address\"},{\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"name\":\"senderAddress\",\"type\":\"address\"},{\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"makerFee\",\"type\":\"uint256\"},{\"name\":\"takerFee\",\"type\":\"uint256\"},{\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"},{\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"makerAddress\",\"type\":\"address\"},{\"name\":\"takerAddress\",\"type\":\"address\"},{\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"name\":\"senderAddress\",\"type\":\"address\"},{\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"name\":\"makerFee\",\"type\":\"uint256\"},{\"name\":\"takerFee\",\"type\":\"uint256\"},{\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"},{\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"name\":\"buySignature\",\"type\":\"bytes\"},{\"name\":\"sellSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"name\":\"right\",\"type\":\"tuple\"},{\"name\":\"leftMakerAssetSpreadAmount\",\"type\":\"uint256\"}],\"name\":\"matchedFillResults\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_exchange\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DutchAuction is an auto generated Go binding around an Ethereum contract.
type DutchAuction struct {
	DutchAuctionCaller     // Read-only binding to the contract
	DutchAuctionTransactor // Write-only binding to the contract
	DutchAuctionFilterer   // Log filterer for contract events
}

// DutchAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DutchAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DutchAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DutchAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DutchAuctionSession struct {
	Contract     *DutchAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DutchAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DutchAuctionCallerSession struct {
	Contract *DutchAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DutchAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DutchAuctionTransactorSession struct {
	Contract     *DutchAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DutchAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DutchAuctionRaw struct {
	Contract *DutchAuction // Generic contract binding to access the raw methods on
}

// DutchAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DutchAuctionCallerRaw struct {
	Contract *DutchAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// DutchAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DutchAuctionTransactorRaw struct {
	Contract *DutchAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDutchAuction creates a new instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuction(address common.Address, backend bind.ContractBackend) (*DutchAuction, error) {
	contract, err := bindDutchAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DutchAuction{DutchAuctionCaller: DutchAuctionCaller{contract: contract}, DutchAuctionTransactor: DutchAuctionTransactor{contract: contract}, DutchAuctionFilterer: DutchAuctionFilterer{contract: contract}}, nil
}

// NewDutchAuctionCaller creates a new read-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionCaller(address common.Address, caller bind.ContractCaller) (*DutchAuctionCaller, error) {
	contract, err := bindDutchAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionCaller{contract: contract}, nil
}

// NewDutchAuctionTransactor creates a new write-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*DutchAuctionTransactor, error) {
	contract, err := bindDutchAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionTransactor{contract: contract}, nil
}

// NewDutchAuctionFilterer creates a new log filterer instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*DutchAuctionFilterer, error) {
	contract, err := bindDutchAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionFilterer{contract: contract}, nil
}

// bindDutchAuction binds a generic wrapper to an already deployed contract.
func bindDutchAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DutchAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.DutchAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transact(opts, method, params...)
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	Left                       Struct2
	Right                      Struct2
	LeftMakerAssetSpreadAmount *big.Int
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	MakerAddress          common.Address
	TakerAddress          common.Address
	FeeRecipientAddress   common.Address
	SenderAddress         common.Address
	MakerAssetAmount      *big.Int
	TakerAssetAmount      *big.Int
	MakerFee              *big.Int
	TakerFee              *big.Int
	ExpirationTimeSeconds *big.Int
	Salt                  *big.Int
	MakerAssetData        []byte
	TakerAssetData        []byte
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	BeginTimeSeconds   *big.Int
	EndTimeSeconds     *big.Int
	BeginAmount        *big.Int
	EndAmount          *big.Int
	CurrentAmount      *big.Int
	CurrentTimeSeconds *big.Int
}

// GetAuctionDetails is a paid mutator transaction binding the contract method 0x2e9cd033.
//
// Solidity: function getAuctionDetails(Struct0 order) returns(Struct1 auctionDetails)
func (_DutchAuction *DutchAuctionTransactor) GetAuctionDetails(opts *bind.TransactOpts, order Struct0) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "getAuctionDetails", order)
}

// GetAuctionDetails is a paid mutator transaction binding the contract method 0x2e9cd033.
//
// Solidity: function getAuctionDetails(Struct0 order) returns(Struct1 auctionDetails)
func (_DutchAuction *DutchAuctionSession) GetAuctionDetails(order Struct0) (*types.Transaction, error) {
	return _DutchAuction.Contract.GetAuctionDetails(&_DutchAuction.TransactOpts, order)
}

// GetAuctionDetails is a paid mutator transaction binding the contract method 0x2e9cd033.
//
// Solidity: function getAuctionDetails(Struct0 order) returns(Struct1 auctionDetails)
func (_DutchAuction *DutchAuctionTransactorSession) GetAuctionDetails(order Struct0) (*types.Transaction, error) {
	return _DutchAuction.Contract.GetAuctionDetails(&_DutchAuction.TransactOpts, order)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c28d861.
//
// Solidity: function matchOrders(Struct0 buyOrder, Struct0 sellOrder, bytes buySignature, bytes sellSignature) returns(Struct3 matchedFillResults)
func (_DutchAuction *DutchAuctionTransactor) MatchOrders(opts *bind.TransactOpts, buyOrder Struct0, sellOrder Struct0, buySignature []byte, sellSignature []byte) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "matchOrders", buyOrder, sellOrder, buySignature, sellSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c28d861.
//
// Solidity: function matchOrders(Struct0 buyOrder, Struct0 sellOrder, bytes buySignature, bytes sellSignature) returns(Struct3 matchedFillResults)
func (_DutchAuction *DutchAuctionSession) MatchOrders(buyOrder Struct0, sellOrder Struct0, buySignature []byte, sellSignature []byte) (*types.Transaction, error) {
	return _DutchAuction.Contract.MatchOrders(&_DutchAuction.TransactOpts, buyOrder, sellOrder, buySignature, sellSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c28d861.
//
// Solidity: function matchOrders(Struct0 buyOrder, Struct0 sellOrder, bytes buySignature, bytes sellSignature) returns(Struct3 matchedFillResults)
func (_DutchAuction *DutchAuctionTransactorSession) MatchOrders(buyOrder Struct0, sellOrder Struct0, buySignature []byte, sellSignature []byte) (*types.Transaction, error) {
	return _DutchAuction.Contract.MatchOrders(&_DutchAuction.TransactOpts, buyOrder, sellOrder, buySignature, sellSignature)
}

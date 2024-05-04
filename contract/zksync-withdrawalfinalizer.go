// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WithdrawalFinalizerRequestFinalizeWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalFinalizerRequestFinalizeWithdrawal struct {
	L2BlockNumber     *big.Int
	L2MessageIndex    *big.Int
	L2TxNumberInBlock uint16
	Message           []byte
	MerkleProof       [][32]byte
	IsEth             bool
	Gas               *big.Int
}

// WithdrawalFinalizerResult is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalFinalizerResult struct {
	L2BlockNumber  *big.Int
	L2MessageIndex *big.Int
	Gas            *big.Int
	Success        bool
}

// WithdrawalFinalizerMetaData contains all meta data concerning the WithdrawalFinalizer contract.
var WithdrawalFinalizerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"_isEth\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawalFinalizer.RequestFinalizeWithdrawal[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"name\":\"finalizeWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structWithdrawalFinalizer.Result[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WithdrawalFinalizerABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalFinalizerMetaData.ABI instead.
var WithdrawalFinalizerABI = WithdrawalFinalizerMetaData.ABI

// WithdrawalFinalizer is an auto generated Go binding around an Ethereum contract.
type WithdrawalFinalizer struct {
	WithdrawalFinalizerCaller     // Read-only binding to the contract
	WithdrawalFinalizerTransactor // Write-only binding to the contract
	WithdrawalFinalizerFilterer   // Log filterer for contract events
}

// WithdrawalFinalizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalFinalizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalFinalizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalFinalizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalFinalizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalFinalizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalFinalizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalFinalizerSession struct {
	Contract     *WithdrawalFinalizer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WithdrawalFinalizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalFinalizerCallerSession struct {
	Contract *WithdrawalFinalizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// WithdrawalFinalizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalFinalizerTransactorSession struct {
	Contract     *WithdrawalFinalizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// WithdrawalFinalizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalFinalizerRaw struct {
	Contract *WithdrawalFinalizer // Generic contract binding to access the raw methods on
}

// WithdrawalFinalizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalFinalizerCallerRaw struct {
	Contract *WithdrawalFinalizerCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalFinalizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalFinalizerTransactorRaw struct {
	Contract *WithdrawalFinalizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawalFinalizer creates a new instance of WithdrawalFinalizer, bound to a specific deployed contract.
func NewWithdrawalFinalizer(address common.Address, backend bind.ContractBackend) (*WithdrawalFinalizer, error) {
	contract, err := bindWithdrawalFinalizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawalFinalizer{WithdrawalFinalizerCaller: WithdrawalFinalizerCaller{contract: contract}, WithdrawalFinalizerTransactor: WithdrawalFinalizerTransactor{contract: contract}, WithdrawalFinalizerFilterer: WithdrawalFinalizerFilterer{contract: contract}}, nil
}

// NewWithdrawalFinalizerCaller creates a new read-only instance of WithdrawalFinalizer, bound to a specific deployed contract.
func NewWithdrawalFinalizerCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalFinalizerCaller, error) {
	contract, err := bindWithdrawalFinalizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalFinalizerCaller{contract: contract}, nil
}

// NewWithdrawalFinalizerTransactor creates a new write-only instance of WithdrawalFinalizer, bound to a specific deployed contract.
func NewWithdrawalFinalizerTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalFinalizerTransactor, error) {
	contract, err := bindWithdrawalFinalizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalFinalizerTransactor{contract: contract}, nil
}

// NewWithdrawalFinalizerFilterer creates a new log filterer instance of WithdrawalFinalizer, bound to a specific deployed contract.
func NewWithdrawalFinalizerFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalFinalizerFilterer, error) {
	contract, err := bindWithdrawalFinalizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalFinalizerFilterer{contract: contract}, nil
}

// bindWithdrawalFinalizer binds a generic wrapper to an already deployed contract.
func bindWithdrawalFinalizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawalFinalizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalFinalizer *WithdrawalFinalizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalFinalizer.Contract.WithdrawalFinalizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalFinalizer *WithdrawalFinalizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.WithdrawalFinalizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalFinalizer *WithdrawalFinalizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.WithdrawalFinalizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalFinalizer *WithdrawalFinalizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalFinalizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalFinalizer *WithdrawalFinalizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalFinalizer *WithdrawalFinalizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.contract.Transact(opts, method, params...)
}

// FinalizeWithdrawals is a paid mutator transaction binding the contract method 0x32bfc64d.
//
// Solidity: function finalizeWithdrawals((uint256,uint256,uint16,bytes,bytes32[],bool,uint256)[] requests) returns((uint256,uint256,uint256,bool)[])
func (_WithdrawalFinalizer *WithdrawalFinalizerTransactor) FinalizeWithdrawals(opts *bind.TransactOpts, requests []WithdrawalFinalizerRequestFinalizeWithdrawal) (*types.Transaction, error) {
	return _WithdrawalFinalizer.contract.Transact(opts, "finalizeWithdrawals", requests)
}

// FinalizeWithdrawals is a paid mutator transaction binding the contract method 0x32bfc64d.
//
// Solidity: function finalizeWithdrawals((uint256,uint256,uint16,bytes,bytes32[],bool,uint256)[] requests) returns((uint256,uint256,uint256,bool)[])
func (_WithdrawalFinalizer *WithdrawalFinalizerSession) FinalizeWithdrawals(requests []WithdrawalFinalizerRequestFinalizeWithdrawal) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.FinalizeWithdrawals(&_WithdrawalFinalizer.TransactOpts, requests)
}

// FinalizeWithdrawals is a paid mutator transaction binding the contract method 0x32bfc64d.
//
// Solidity: function finalizeWithdrawals((uint256,uint256,uint16,bytes,bytes32[],bool,uint256)[] requests) returns((uint256,uint256,uint256,bool)[])
func (_WithdrawalFinalizer *WithdrawalFinalizerTransactorSession) FinalizeWithdrawals(requests []WithdrawalFinalizerRequestFinalizeWithdrawal) (*types.Transaction, error) {
	return _WithdrawalFinalizer.Contract.FinalizeWithdrawals(&_WithdrawalFinalizer.TransactOpts, requests)
}

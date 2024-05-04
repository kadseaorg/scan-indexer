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

// IScrollChainMetaData contains all meta data concerning the IScrollChain contract.
var IScrollChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggrProof\",\"type\":\"bytes\"}],\"name\":\"finalizeBatchWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IScrollChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IScrollChainMetaData.ABI instead.
var IScrollChainABI = IScrollChainMetaData.ABI

// IScrollChain is an auto generated Go binding around an Ethereum contract.
type IScrollChain struct {
	IScrollChainCaller     // Read-only binding to the contract
	IScrollChainTransactor // Write-only binding to the contract
	IScrollChainFilterer   // Log filterer for contract events
}

// IScrollChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScrollChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScrollChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScrollChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScrollChainSession struct {
	Contract     *IScrollChain     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IScrollChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScrollChainCallerSession struct {
	Contract *IScrollChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IScrollChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScrollChainTransactorSession struct {
	Contract     *IScrollChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IScrollChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScrollChainRaw struct {
	Contract *IScrollChain // Generic contract binding to access the raw methods on
}

// IScrollChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScrollChainCallerRaw struct {
	Contract *IScrollChainCaller // Generic read-only contract binding to access the raw methods on
}

// IScrollChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScrollChainTransactorRaw struct {
	Contract *IScrollChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScrollChain creates a new instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChain(address common.Address, backend bind.ContractBackend) (*IScrollChain, error) {
	contract, err := bindIScrollChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScrollChain{IScrollChainCaller: IScrollChainCaller{contract: contract}, IScrollChainTransactor: IScrollChainTransactor{contract: contract}, IScrollChainFilterer: IScrollChainFilterer{contract: contract}}, nil
}

// NewIScrollChainCaller creates a new read-only instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainCaller(address common.Address, caller bind.ContractCaller) (*IScrollChainCaller, error) {
	contract, err := bindIScrollChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollChainCaller{contract: contract}, nil
}

// NewIScrollChainTransactor creates a new write-only instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IScrollChainTransactor, error) {
	contract, err := bindIScrollChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollChainTransactor{contract: contract}, nil
}

// NewIScrollChainFilterer creates a new log filterer instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IScrollChainFilterer, error) {
	contract, err := bindIScrollChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScrollChainFilterer{contract: contract}, nil
}

// bindIScrollChain binds a generic wrapper to an already deployed contract.
func bindIScrollChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IScrollChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollChain *IScrollChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollChain.Contract.IScrollChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollChain *IScrollChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollChain.Contract.IScrollChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollChain *IScrollChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollChain.Contract.IScrollChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollChain *IScrollChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollChain *IScrollChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollChain *IScrollChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollChain.Contract.contract.Transact(opts, method, params...)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "committedBatches", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.CommittedBatches(&_IScrollChain.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.CommittedBatches(&_IScrollChain.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) FinalizedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "finalizedStateRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.FinalizedStateRoots(&_IScrollChain.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.FinalizedStateRoots(&_IScrollChain.CallOpts, batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainCaller) IsBatchFinalized(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "isBatchFinalized", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainSession) IsBatchFinalized(batchIndex *big.Int) (bool, error) {
	return _IScrollChain.Contract.IsBatchFinalized(&_IScrollChain.CallOpts, batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainCallerSession) IsBatchFinalized(batchIndex *big.Int) (bool, error) {
	return _IScrollChain.Contract.IsBatchFinalized(&_IScrollChain.CallOpts, batchIndex)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) WithdrawRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "withdrawRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) WithdrawRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.WithdrawRoots(&_IScrollChain.CallOpts, batchIndex)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) WithdrawRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.WithdrawRoots(&_IScrollChain.CallOpts, batchIndex)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainTransactor) CommitBatch(opts *bind.TransactOpts, version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "commitBatch", version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainSession) CommitBatch(version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.CommitBatch(&_IScrollChain.TransactOpts, version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainTransactorSession) CommitBatch(version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.CommitBatch(&_IScrollChain.TransactOpts, version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactor) FinalizeBatchWithProof(opts *bind.TransactOpts, batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "finalizeBatchWithProof", batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainSession) FinalizeBatchWithProof(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactorSession) FinalizeBatchWithProof(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainTransactor) RevertBatch(opts *bind.TransactOpts, batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "revertBatch", batchHeader, count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainSession) RevertBatch(batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.Contract.RevertBatch(&_IScrollChain.TransactOpts, batchHeader, count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainTransactorSession) RevertBatch(batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.Contract.RevertBatch(&_IScrollChain.TransactOpts, batchHeader, count)
}

// IScrollChainCommitBatchIterator is returned from FilterCommitBatch and is used to iterate over the raw logs and unpacked data for CommitBatch events raised by the IScrollChain contract.
type IScrollChainCommitBatchIterator struct {
	Event *IScrollChainCommitBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainCommitBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainCommitBatch)
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
		it.Event = new(IScrollChainCommitBatch)
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
func (it *IScrollChainCommitBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainCommitBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainCommitBatch represents a CommitBatch event raised by the IScrollChain contract.
type IScrollChainCommitBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitBatch is a free log retrieval operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainCommitBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainCommitBatchIterator{contract: _IScrollChain.contract, event: "CommitBatch", logs: logs, sub: sub}, nil
}

// WatchCommitBatch is a free log subscription operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) WatchCommitBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainCommitBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainCommitBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
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

// ParseCommitBatch is a log parse operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) ParseCommitBatch(log types.Log) (*IScrollChainCommitBatch, error) {
	event := new(IScrollChainCommitBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainFinalizeBatchIterator is returned from FilterFinalizeBatch and is used to iterate over the raw logs and unpacked data for FinalizeBatch events raised by the IScrollChain contract.
type IScrollChainFinalizeBatchIterator struct {
	Event *IScrollChainFinalizeBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainFinalizeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainFinalizeBatch)
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
		it.Event = new(IScrollChainFinalizeBatch)
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
func (it *IScrollChainFinalizeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainFinalizeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainFinalizeBatch represents a FinalizeBatch event raised by the IScrollChain contract.
type IScrollChainFinalizeBatch struct {
	BatchIndex   *big.Int
	BatchHash    [32]byte
	StateRoot    [32]byte
	WithdrawRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatch is a free log retrieval operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainFinalizeBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainFinalizeBatchIterator{contract: _IScrollChain.contract, event: "FinalizeBatch", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatch is a free log subscription operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) WatchFinalizeBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainFinalizeBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainFinalizeBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
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

// ParseFinalizeBatch is a log parse operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) ParseFinalizeBatch(log types.Log) (*IScrollChainFinalizeBatch, error) {
	event := new(IScrollChainFinalizeBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainRevertBatchIterator is returned from FilterRevertBatch and is used to iterate over the raw logs and unpacked data for RevertBatch events raised by the IScrollChain contract.
type IScrollChainRevertBatchIterator struct {
	Event *IScrollChainRevertBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainRevertBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainRevertBatch)
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
		it.Event = new(IScrollChainRevertBatch)
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
func (it *IScrollChainRevertBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainRevertBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainRevertBatch represents a RevertBatch event raised by the IScrollChain contract.
type IScrollChainRevertBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevertBatch is a free log retrieval operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainRevertBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainRevertBatchIterator{contract: _IScrollChain.contract, event: "RevertBatch", logs: logs, sub: sub}, nil
}

// WatchRevertBatch is a free log subscription operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) WatchRevertBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainRevertBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainRevertBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
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

// ParseRevertBatch is a log parse operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) ParseRevertBatch(log types.Log) (*IScrollChainRevertBatch, error) {
	event := new(IScrollChainRevertBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

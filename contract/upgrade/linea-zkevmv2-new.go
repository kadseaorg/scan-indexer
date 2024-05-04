// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgrade

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
	_ = abi.ConvertType
)

// IZkEvmV2BlockData is an auto generated low-level Go binding around an user-defined struct.
type IZkEvmV2BlockData struct {
	BlockRootHash         [32]byte
	L2BlockTimestamp      uint32
	Transactions          [][]byte
	L2ToL1MsgHashes       [][32]byte
	FromAddresses         []byte
	BatchReceptionIndices []uint16
}

// LineaMetaData contains all meta data concerning the Linea contract.
var LineaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlockTimestampError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeePaymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"L1L2MessageNotSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inde\",\"type\":\"uint256\"}],\"name\":\"MemoryOutOfBounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageAlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadySent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageDoesNotExistOrHasAlreadyBeenClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"MessageSendingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartingRootHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownTransactionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueSentTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueShouldBeGreaterThanFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resettingAddress\",\"type\":\"address\"}],\"name\":\"AmountUsedInPeriodReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastBlockFinalized\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"startingRootHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"finalRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withProof\",\"type\":\"bool\"}],\"name\":\"DataFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastBlockFinalized\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"startingRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"finalRootHash\",\"type\":\"bytes32\"}],\"name\":\"BlocksVerificationDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageHashes\",\"type\":\"bytes32[]\"}],\"name\":\"L1L2MessagesReceivedOnL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"L2L1MessageHashAddedToInbox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"amountChangeBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"amountUsedLoweredToLimit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"usedAmountResetToZero\",\"type\":\"bool\"}],\"name\":\"LimitAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierSetBy\",\"type\":\"address\"}],\"name\":\"VerifierAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERAL_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_L2_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_L1_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_SENT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVING_SYSTEM_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodAmountInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2ToL1MsgHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"fromAddresses\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"batchReceptionIndices\",\"type\":\"uint16[]\"}],\"internalType\":\"structIZkEvmV2.BlockData[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_parentStateRootHash\",\"type\":\"bytes32\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2ToL1MsgHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"fromAddresses\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"batchReceptionIndices\",\"type\":\"uint16[]\"}],\"internalType\":\"structIZkEvmV2.BlockData[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"}],\"name\":\"finalizeBlocksWithoutProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inboxL2L1MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_initialStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_initialL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_defaultVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitPeriodInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitAmountInWei\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"outboxL1L2MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"pauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pauseTypeStatuses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetAmountUsedInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resetRateLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"unPauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LineaABI is the input ABI used to generate the binding from.
// Deprecated: Use LineaMetaData.ABI instead.
var LineaABI = LineaMetaData.ABI

// Linea is an auto generated Go binding around an Ethereum contract.
type Linea struct {
	LineaCaller     // Read-only binding to the contract
	LineaTransactor // Write-only binding to the contract
	LineaFilterer   // Log filterer for contract events
}

// LineaCaller is an auto generated read-only Go binding around an Ethereum contract.
type LineaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LineaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LineaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LineaSession struct {
	Contract     *Linea            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LineaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LineaCallerSession struct {
	Contract *LineaCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LineaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LineaTransactorSession struct {
	Contract     *LineaTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LineaRaw is an auto generated low-level Go binding around an Ethereum contract.
type LineaRaw struct {
	Contract *Linea // Generic contract binding to access the raw methods on
}

// LineaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LineaCallerRaw struct {
	Contract *LineaCaller // Generic read-only contract binding to access the raw methods on
}

// LineaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LineaTransactorRaw struct {
	Contract *LineaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinea creates a new instance of Linea, bound to a specific deployed contract.
func NewLinea(address common.Address, backend bind.ContractBackend) (*Linea, error) {
	contract, err := bindLinea(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Linea{LineaCaller: LineaCaller{contract: contract}, LineaTransactor: LineaTransactor{contract: contract}, LineaFilterer: LineaFilterer{contract: contract}}, nil
}

// NewLineaCaller creates a new read-only instance of Linea, bound to a specific deployed contract.
func NewLineaCaller(address common.Address, caller bind.ContractCaller) (*LineaCaller, error) {
	contract, err := bindLinea(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LineaCaller{contract: contract}, nil
}

// NewLineaTransactor creates a new write-only instance of Linea, bound to a specific deployed contract.
func NewLineaTransactor(address common.Address, transactor bind.ContractTransactor) (*LineaTransactor, error) {
	contract, err := bindLinea(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LineaTransactor{contract: contract}, nil
}

// NewLineaFilterer creates a new log filterer instance of Linea, bound to a specific deployed contract.
func NewLineaFilterer(address common.Address, filterer bind.ContractFilterer) (*LineaFilterer, error) {
	contract, err := bindLinea(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LineaFilterer{contract: contract}, nil
}

// bindLinea binds a generic wrapper to an already deployed contract.
func bindLinea(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LineaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Linea *LineaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Linea.Contract.LineaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Linea *LineaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linea.Contract.LineaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Linea *LineaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Linea.Contract.LineaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Linea *LineaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Linea.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Linea *LineaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linea.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Linea *LineaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Linea.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Linea *LineaCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Linea *LineaSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Linea.Contract.DEFAULTADMINROLE(&_Linea.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Linea *LineaCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Linea.Contract.DEFAULTADMINROLE(&_Linea.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCaller) GENERALPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "GENERAL_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaSession) GENERALPAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.GENERALPAUSETYPE(&_Linea.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCallerSession) GENERALPAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.GENERALPAUSETYPE(&_Linea.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaCaller) INBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "INBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _Linea.Contract.INBOXSTATUSRECEIVED(&_Linea.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaCallerSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _Linea.Contract.INBOXSTATUSRECEIVED(&_Linea.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaCaller) INBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "INBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _Linea.Contract.INBOXSTATUSUNKNOWN(&_Linea.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaCallerSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _Linea.Contract.INBOXSTATUSUNKNOWN(&_Linea.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCaller) L1L2PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "L1_L2_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaSession) L1L2PAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.L1L2PAUSETYPE(&_Linea.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCallerSession) L1L2PAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.L1L2PAUSETYPE(&_Linea.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCaller) L2L1PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "L2_L1_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaSession) L2L1PAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.L2L1PAUSETYPE(&_Linea.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCallerSession) L2L1PAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.L2L1PAUSETYPE(&_Linea.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Linea *LineaCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Linea *LineaSession) OPERATORROLE() ([32]byte, error) {
	return _Linea.Contract.OPERATORROLE(&_Linea.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Linea *LineaCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Linea.Contract.OPERATORROLE(&_Linea.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaCaller) OUTBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "OUTBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaSession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSRECEIVED(&_Linea.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_Linea *LineaCallerSession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSRECEIVED(&_Linea.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_Linea *LineaCaller) OUTBOXSTATUSSENT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "OUTBOX_STATUS_SENT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_Linea *LineaSession) OUTBOXSTATUSSENT() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSSENT(&_Linea.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_Linea *LineaCallerSession) OUTBOXSTATUSSENT() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSSENT(&_Linea.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaCaller) OUTBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "OUTBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaSession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSUNKNOWN(&_Linea.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_Linea *LineaCallerSession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _Linea.Contract.OUTBOXSTATUSUNKNOWN(&_Linea.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_Linea *LineaCaller) PAUSEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "PAUSE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_Linea *LineaSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _Linea.Contract.PAUSEMANAGERROLE(&_Linea.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_Linea *LineaCallerSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _Linea.Contract.PAUSEMANAGERROLE(&_Linea.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCaller) PROVINGSYSTEMPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "PROVING_SYSTEM_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaSession) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.PROVINGSYSTEMPAUSETYPE(&_Linea.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_Linea *LineaCallerSession) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _Linea.Contract.PROVINGSYSTEMPAUSETYPE(&_Linea.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Linea *LineaCaller) RATELIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "RATE_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Linea *LineaSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _Linea.Contract.RATELIMITSETTERROLE(&_Linea.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Linea *LineaCallerSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _Linea.Contract.RATELIMITSETTERROLE(&_Linea.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_Linea *LineaCaller) CurrentL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "currentL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_Linea *LineaSession) CurrentL2BlockNumber() (*big.Int, error) {
	return _Linea.Contract.CurrentL2BlockNumber(&_Linea.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_Linea *LineaCallerSession) CurrentL2BlockNumber() (*big.Int, error) {
	return _Linea.Contract.CurrentL2BlockNumber(&_Linea.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_Linea *LineaCaller) CurrentPeriodAmountInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "currentPeriodAmountInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_Linea *LineaSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _Linea.Contract.CurrentPeriodAmountInWei(&_Linea.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_Linea *LineaCallerSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _Linea.Contract.CurrentPeriodAmountInWei(&_Linea.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_Linea *LineaCaller) CurrentPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "currentPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_Linea *LineaSession) CurrentPeriodEnd() (*big.Int, error) {
	return _Linea.Contract.CurrentPeriodEnd(&_Linea.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_Linea *LineaCallerSession) CurrentPeriodEnd() (*big.Int, error) {
	return _Linea.Contract.CurrentPeriodEnd(&_Linea.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_Linea *LineaCaller) CurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "currentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_Linea *LineaSession) CurrentTimestamp() (*big.Int, error) {
	return _Linea.Contract.CurrentTimestamp(&_Linea.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_Linea *LineaCallerSession) CurrentTimestamp() (*big.Int, error) {
	return _Linea.Contract.CurrentTimestamp(&_Linea.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Linea *LineaCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Linea *LineaSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Linea.Contract.GetRoleAdmin(&_Linea.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Linea *LineaCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Linea.Contract.GetRoleAdmin(&_Linea.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Linea *LineaCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Linea *LineaSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Linea.Contract.HasRole(&_Linea.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Linea *LineaCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Linea.Contract.HasRole(&_Linea.CallOpts, role, account)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaCaller) InboxL2L1MessageStatus(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "inboxL2L1MessageStatus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaSession) InboxL2L1MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _Linea.Contract.InboxL2L1MessageStatus(&_Linea.CallOpts, arg0)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaCallerSession) InboxL2L1MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _Linea.Contract.InboxL2L1MessageStatus(&_Linea.CallOpts, arg0)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_Linea *LineaCaller) LimitInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "limitInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_Linea *LineaSession) LimitInWei() (*big.Int, error) {
	return _Linea.Contract.LimitInWei(&_Linea.CallOpts)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_Linea *LineaCallerSession) LimitInWei() (*big.Int, error) {
	return _Linea.Contract.LimitInWei(&_Linea.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_Linea *LineaCaller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_Linea *LineaSession) NextMessageNumber() (*big.Int, error) {
	return _Linea.Contract.NextMessageNumber(&_Linea.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_Linea *LineaCallerSession) NextMessageNumber() (*big.Int, error) {
	return _Linea.Contract.NextMessageNumber(&_Linea.CallOpts)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaCaller) OutboxL1L2MessageStatus(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "outboxL1L2MessageStatus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaSession) OutboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _Linea.Contract.OutboxL1L2MessageStatus(&_Linea.CallOpts, arg0)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_Linea *LineaCallerSession) OutboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _Linea.Contract.OutboxL1L2MessageStatus(&_Linea.CallOpts, arg0)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_Linea *LineaCaller) PauseTypeStatuses(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "pauseTypeStatuses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_Linea *LineaSession) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _Linea.Contract.PauseTypeStatuses(&_Linea.CallOpts, arg0)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_Linea *LineaCallerSession) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _Linea.Contract.PauseTypeStatuses(&_Linea.CallOpts, arg0)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_Linea *LineaCaller) PeriodInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "periodInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_Linea *LineaSession) PeriodInSeconds() (*big.Int, error) {
	return _Linea.Contract.PeriodInSeconds(&_Linea.CallOpts)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_Linea *LineaCallerSession) PeriodInSeconds() (*big.Int, error) {
	return _Linea.Contract.PeriodInSeconds(&_Linea.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_Linea *LineaCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_Linea *LineaSession) Sender() (common.Address, error) {
	return _Linea.Contract.Sender(&_Linea.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_Linea *LineaCallerSession) Sender() (common.Address, error) {
	return _Linea.Contract.Sender(&_Linea.CallOpts)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_Linea *LineaCaller) StateRootHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "stateRootHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_Linea *LineaSession) StateRootHashes(arg0 *big.Int) ([32]byte, error) {
	return _Linea.Contract.StateRootHashes(&_Linea.CallOpts, arg0)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_Linea *LineaCallerSession) StateRootHashes(arg0 *big.Int) ([32]byte, error) {
	return _Linea.Contract.StateRootHashes(&_Linea.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Linea *LineaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Linea *LineaSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Linea.Contract.SupportsInterface(&_Linea.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Linea *LineaCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Linea.Contract.SupportsInterface(&_Linea.CallOpts, interfaceId)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_Linea *LineaCaller) Verifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Linea.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_Linea *LineaSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _Linea.Contract.Verifiers(&_Linea.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_Linea *LineaCallerSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _Linea.Contract.Verifiers(&_Linea.CallOpts, arg0)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_Linea *LineaTransactor) ClaimMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "claimMessage", _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_Linea *LineaSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.ClaimMessage(&_Linea.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_Linea *LineaTransactorSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.ClaimMessage(&_Linea.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_Linea *LineaTransactor) FinalizeBlocks(opts *bind.TransactOpts, _blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "finalizeBlocks", _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_Linea *LineaSession) FinalizeBlocks(_blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.FinalizeBlocks(&_Linea.TransactOpts, _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_Linea *LineaTransactorSession) FinalizeBlocks(_blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.FinalizeBlocks(&_Linea.TransactOpts, _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_Linea *LineaTransactor) FinalizeBlocksWithoutProof(opts *bind.TransactOpts, _blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "finalizeBlocksWithoutProof", _blocksData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_Linea *LineaSession) FinalizeBlocksWithoutProof(_blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _Linea.Contract.FinalizeBlocksWithoutProof(&_Linea.TransactOpts, _blocksData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_Linea *LineaTransactorSession) FinalizeBlocksWithoutProof(_blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _Linea.Contract.FinalizeBlocksWithoutProof(&_Linea.TransactOpts, _blocksData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Linea *LineaSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.GrantRole(&_Linea.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.GrantRole(&_Linea.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_Linea *LineaTransactor) Initialize(opts *bind.TransactOpts, _initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "initialize", _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_Linea *LineaSession) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.Initialize(&_Linea.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_Linea *LineaTransactorSession) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.Initialize(&_Linea.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaTransactor) PauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "pauseByType", _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaSession) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.PauseByType(&_Linea.TransactOpts, _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaTransactorSession) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.PauseByType(&_Linea.TransactOpts, _pauseType)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Linea *LineaSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.RenounceRole(&_Linea.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.RenounceRole(&_Linea.TransactOpts, role, account)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_Linea *LineaTransactor) ResetAmountUsedInPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "resetAmountUsedInPeriod")
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_Linea *LineaSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _Linea.Contract.ResetAmountUsedInPeriod(&_Linea.TransactOpts)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_Linea *LineaTransactorSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _Linea.Contract.ResetAmountUsedInPeriod(&_Linea.TransactOpts)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_Linea *LineaTransactor) ResetRateLimitAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "resetRateLimitAmount", _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_Linea *LineaSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.ResetRateLimitAmount(&_Linea.TransactOpts, _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_Linea *LineaTransactorSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.ResetRateLimitAmount(&_Linea.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Linea *LineaSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.RevokeRole(&_Linea.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Linea *LineaTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Linea.Contract.RevokeRole(&_Linea.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_Linea *LineaTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "sendMessage", _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_Linea *LineaSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Linea.Contract.SendMessage(&_Linea.TransactOpts, _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_Linea *LineaTransactorSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Linea.Contract.SendMessage(&_Linea.TransactOpts, _to, _fee, _calldata)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_Linea *LineaTransactor) SetVerifierAddress(opts *bind.TransactOpts, _newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "setVerifierAddress", _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_Linea *LineaSession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.SetVerifierAddress(&_Linea.TransactOpts, _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_Linea *LineaTransactorSession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _Linea.Contract.SetVerifierAddress(&_Linea.TransactOpts, _newVerifierAddress, _proofType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaTransactor) UnPauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.contract.Transact(opts, "unPauseByType", _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaSession) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.UnPauseByType(&_Linea.TransactOpts, _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_Linea *LineaTransactorSession) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _Linea.Contract.UnPauseByType(&_Linea.TransactOpts, _pauseType)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Linea *LineaTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linea.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Linea *LineaSession) Receive() (*types.Transaction, error) {
	return _Linea.Contract.Receive(&_Linea.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Linea *LineaTransactorSession) Receive() (*types.Transaction, error) {
	return _Linea.Contract.Receive(&_Linea.TransactOpts)
}

// LineaAmountUsedInPeriodResetIterator is returned from FilterAmountUsedInPeriodReset and is used to iterate over the raw logs and unpacked data for AmountUsedInPeriodReset events raised by the Linea contract.
type LineaAmountUsedInPeriodResetIterator struct {
	Event *LineaAmountUsedInPeriodReset // Event containing the contract specifics and raw log

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
func (it *LineaAmountUsedInPeriodResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaAmountUsedInPeriodReset)
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
		it.Event = new(LineaAmountUsedInPeriodReset)
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
func (it *LineaAmountUsedInPeriodResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaAmountUsedInPeriodResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaAmountUsedInPeriodReset represents a AmountUsedInPeriodReset event raised by the Linea contract.
type LineaAmountUsedInPeriodReset struct {
	ResettingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAmountUsedInPeriodReset is a free log retrieval operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_Linea *LineaFilterer) FilterAmountUsedInPeriodReset(opts *bind.FilterOpts, resettingAddress []common.Address) (*LineaAmountUsedInPeriodResetIterator, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return &LineaAmountUsedInPeriodResetIterator{contract: _Linea.contract, event: "AmountUsedInPeriodReset", logs: logs, sub: sub}, nil
}

// WatchAmountUsedInPeriodReset is a free log subscription operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_Linea *LineaFilterer) WatchAmountUsedInPeriodReset(opts *bind.WatchOpts, sink chan<- *LineaAmountUsedInPeriodReset, resettingAddress []common.Address) (event.Subscription, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaAmountUsedInPeriodReset)
				if err := _Linea.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
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

// ParseAmountUsedInPeriodReset is a log parse operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_Linea *LineaFilterer) ParseAmountUsedInPeriodReset(log types.Log) (*LineaAmountUsedInPeriodReset, error) {
	event := new(LineaAmountUsedInPeriodReset)
	if err := _Linea.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaBlocksVerificationDoneIterator is returned from FilterBlocksVerificationDone and is used to iterate over the raw logs and unpacked data for BlocksVerificationDone events raised by the Linea contract.
type LineaBlocksVerificationDoneIterator struct {
	Event *LineaBlocksVerificationDone // Event containing the contract specifics and raw log

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
func (it *LineaBlocksVerificationDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaBlocksVerificationDone)
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
		it.Event = new(LineaBlocksVerificationDone)
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
func (it *LineaBlocksVerificationDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaBlocksVerificationDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaBlocksVerificationDone represents a BlocksVerificationDone event raised by the Linea contract.
type LineaBlocksVerificationDone struct {
	LastBlockFinalized *big.Int
	StartingRootHash   [32]byte
	FinalRootHash      [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerificationDone is a free log retrieval operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_Linea *LineaFilterer) FilterBlocksVerificationDone(opts *bind.FilterOpts, lastBlockFinalized []*big.Int) (*LineaBlocksVerificationDoneIterator, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return &LineaBlocksVerificationDoneIterator{contract: _Linea.contract, event: "BlocksVerificationDone", logs: logs, sub: sub}, nil
}

// WatchBlocksVerificationDone is a free log subscription operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_Linea *LineaFilterer) WatchBlocksVerificationDone(opts *bind.WatchOpts, sink chan<- *LineaBlocksVerificationDone, lastBlockFinalized []*big.Int) (event.Subscription, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaBlocksVerificationDone)
				if err := _Linea.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
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

// ParseBlocksVerificationDone is a log parse operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_Linea *LineaFilterer) ParseBlocksVerificationDone(log types.Log) (*LineaBlocksVerificationDone, error) {
	event := new(LineaBlocksVerificationDone)
	if err := _Linea.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaDataFinalizedIterator is returned from FilterDataFinalized and is used to iterate over the raw logs and unpacked data for DataFinalized events raised by the Linea contract.
type LineaDataFinalizedIterator struct {
	Event *LineaDataFinalized // Event containing the contract specifics and raw log

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
func (it *LineaDataFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaDataFinalized)
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
		it.Event = new(LineaDataFinalized)
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
func (it *LineaDataFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaDataFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaDataFinalized represents a DataFinalized event raised by the Linea contract.
type LineaDataFinalized struct {
	LastBlockFinalized *big.Int
	StartingRootHash   [32]byte
	FinalRootHash      [32]byte
	WithProof          bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDataFinalized is a free log retrieval operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_Linea *LineaFilterer) FilterDataFinalized(opts *bind.FilterOpts, lastBlockFinalized []*big.Int, startingRootHash [][32]byte, finalRootHash [][32]byte) (*LineaDataFinalizedIterator, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}
	var startingRootHashRule []interface{}
	for _, startingRootHashItem := range startingRootHash {
		startingRootHashRule = append(startingRootHashRule, startingRootHashItem)
	}
	var finalRootHashRule []interface{}
	for _, finalRootHashItem := range finalRootHash {
		finalRootHashRule = append(finalRootHashRule, finalRootHashItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "DataFinalized", lastBlockFinalizedRule, startingRootHashRule, finalRootHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaDataFinalizedIterator{contract: _Linea.contract, event: "DataFinalized", logs: logs, sub: sub}, nil
}

// WatchDataFinalized is a free log subscription operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_Linea *LineaFilterer) WatchDataFinalized(opts *bind.WatchOpts, sink chan<- *LineaDataFinalized, lastBlockFinalized []*big.Int, startingRootHash [][32]byte, finalRootHash [][32]byte) (event.Subscription, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}
	var startingRootHashRule []interface{}
	for _, startingRootHashItem := range startingRootHash {
		startingRootHashRule = append(startingRootHashRule, startingRootHashItem)
	}
	var finalRootHashRule []interface{}
	for _, finalRootHashItem := range finalRootHash {
		finalRootHashRule = append(finalRootHashRule, finalRootHashItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "DataFinalized", lastBlockFinalizedRule, startingRootHashRule, finalRootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaDataFinalized)
				if err := _Linea.contract.UnpackLog(event, "DataFinalized", log); err != nil {
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

// ParseDataFinalized is a log parse operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_Linea *LineaFilterer) ParseDataFinalized(log types.Log) (*LineaDataFinalized, error) {
	event := new(LineaDataFinalized)
	if err := _Linea.contract.UnpackLog(event, "DataFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Linea contract.
type LineaInitializedIterator struct {
	Event *LineaInitialized // Event containing the contract specifics and raw log

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
func (it *LineaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaInitialized)
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
		it.Event = new(LineaInitialized)
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
func (it *LineaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaInitialized represents a Initialized event raised by the Linea contract.
type LineaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Linea *LineaFilterer) FilterInitialized(opts *bind.FilterOpts) (*LineaInitializedIterator, error) {

	logs, sub, err := _Linea.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LineaInitializedIterator{contract: _Linea.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Linea *LineaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LineaInitialized) (event.Subscription, error) {

	logs, sub, err := _Linea.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaInitialized)
				if err := _Linea.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Linea *LineaFilterer) ParseInitialized(log types.Log) (*LineaInitialized, error) {
	event := new(LineaInitialized)
	if err := _Linea.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaL1L2MessagesReceivedOnL2Iterator is returned from FilterL1L2MessagesReceivedOnL2 and is used to iterate over the raw logs and unpacked data for L1L2MessagesReceivedOnL2 events raised by the Linea contract.
type LineaL1L2MessagesReceivedOnL2Iterator struct {
	Event *LineaL1L2MessagesReceivedOnL2 // Event containing the contract specifics and raw log

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
func (it *LineaL1L2MessagesReceivedOnL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaL1L2MessagesReceivedOnL2)
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
		it.Event = new(LineaL1L2MessagesReceivedOnL2)
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
func (it *LineaL1L2MessagesReceivedOnL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaL1L2MessagesReceivedOnL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaL1L2MessagesReceivedOnL2 represents a L1L2MessagesReceivedOnL2 event raised by the Linea contract.
type LineaL1L2MessagesReceivedOnL2 struct {
	MessageHashes [][32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL1L2MessagesReceivedOnL2 is a free log retrieval operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_Linea *LineaFilterer) FilterL1L2MessagesReceivedOnL2(opts *bind.FilterOpts) (*LineaL1L2MessagesReceivedOnL2Iterator, error) {

	logs, sub, err := _Linea.contract.FilterLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return &LineaL1L2MessagesReceivedOnL2Iterator{contract: _Linea.contract, event: "L1L2MessagesReceivedOnL2", logs: logs, sub: sub}, nil
}

// WatchL1L2MessagesReceivedOnL2 is a free log subscription operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_Linea *LineaFilterer) WatchL1L2MessagesReceivedOnL2(opts *bind.WatchOpts, sink chan<- *LineaL1L2MessagesReceivedOnL2) (event.Subscription, error) {

	logs, sub, err := _Linea.contract.WatchLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaL1L2MessagesReceivedOnL2)
				if err := _Linea.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
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

// ParseL1L2MessagesReceivedOnL2 is a log parse operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_Linea *LineaFilterer) ParseL1L2MessagesReceivedOnL2(log types.Log) (*LineaL1L2MessagesReceivedOnL2, error) {
	event := new(LineaL1L2MessagesReceivedOnL2)
	if err := _Linea.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaL2L1MessageHashAddedToInboxIterator is returned from FilterL2L1MessageHashAddedToInbox and is used to iterate over the raw logs and unpacked data for L2L1MessageHashAddedToInbox events raised by the Linea contract.
type LineaL2L1MessageHashAddedToInboxIterator struct {
	Event *LineaL2L1MessageHashAddedToInbox // Event containing the contract specifics and raw log

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
func (it *LineaL2L1MessageHashAddedToInboxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaL2L1MessageHashAddedToInbox)
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
		it.Event = new(LineaL2L1MessageHashAddedToInbox)
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
func (it *LineaL2L1MessageHashAddedToInboxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaL2L1MessageHashAddedToInboxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaL2L1MessageHashAddedToInbox represents a L2L1MessageHashAddedToInbox event raised by the Linea contract.
type LineaL2L1MessageHashAddedToInbox struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterL2L1MessageHashAddedToInbox is a free log retrieval operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_Linea *LineaFilterer) FilterL2L1MessageHashAddedToInbox(opts *bind.FilterOpts, messageHash [][32]byte) (*LineaL2L1MessageHashAddedToInboxIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaL2L1MessageHashAddedToInboxIterator{contract: _Linea.contract, event: "L2L1MessageHashAddedToInbox", logs: logs, sub: sub}, nil
}

// WatchL2L1MessageHashAddedToInbox is a free log subscription operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_Linea *LineaFilterer) WatchL2L1MessageHashAddedToInbox(opts *bind.WatchOpts, sink chan<- *LineaL2L1MessageHashAddedToInbox, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaL2L1MessageHashAddedToInbox)
				if err := _Linea.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
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

// ParseL2L1MessageHashAddedToInbox is a log parse operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_Linea *LineaFilterer) ParseL2L1MessageHashAddedToInbox(log types.Log) (*LineaL2L1MessageHashAddedToInbox, error) {
	event := new(LineaL2L1MessageHashAddedToInbox)
	if err := _Linea.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaLimitAmountChangedIterator is returned from FilterLimitAmountChanged and is used to iterate over the raw logs and unpacked data for LimitAmountChanged events raised by the Linea contract.
type LineaLimitAmountChangedIterator struct {
	Event *LineaLimitAmountChanged // Event containing the contract specifics and raw log

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
func (it *LineaLimitAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaLimitAmountChanged)
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
		it.Event = new(LineaLimitAmountChanged)
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
func (it *LineaLimitAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaLimitAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaLimitAmountChanged represents a LimitAmountChanged event raised by the Linea contract.
type LineaLimitAmountChanged struct {
	AmountChangeBy           common.Address
	Amount                   *big.Int
	AmountUsedLoweredToLimit bool
	UsedAmountResetToZero    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLimitAmountChanged is a free log retrieval operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_Linea *LineaFilterer) FilterLimitAmountChanged(opts *bind.FilterOpts, amountChangeBy []common.Address) (*LineaLimitAmountChangedIterator, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return &LineaLimitAmountChangedIterator{contract: _Linea.contract, event: "LimitAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLimitAmountChanged is a free log subscription operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_Linea *LineaFilterer) WatchLimitAmountChanged(opts *bind.WatchOpts, sink chan<- *LineaLimitAmountChanged, amountChangeBy []common.Address) (event.Subscription, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaLimitAmountChanged)
				if err := _Linea.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
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

// ParseLimitAmountChanged is a log parse operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_Linea *LineaFilterer) ParseLimitAmountChanged(log types.Log) (*LineaLimitAmountChanged, error) {
	event := new(LineaLimitAmountChanged)
	if err := _Linea.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaMessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the Linea contract.
type LineaMessageClaimedIterator struct {
	Event *LineaMessageClaimed // Event containing the contract specifics and raw log

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
func (it *LineaMessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaMessageClaimed)
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
		it.Event = new(LineaMessageClaimed)
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
func (it *LineaMessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaMessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaMessageClaimed represents a MessageClaimed event raised by the Linea contract.
type LineaMessageClaimed struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*LineaMessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaMessageClaimedIterator{contract: _Linea.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *LineaMessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaMessageClaimed)
				if err := _Linea.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) ParseMessageClaimed(log types.Log) (*LineaMessageClaimed, error) {
	event := new(LineaMessageClaimed)
	if err := _Linea.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Linea contract.
type LineaMessageSentIterator struct {
	Event *LineaMessageSent // Event containing the contract specifics and raw log

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
func (it *LineaMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaMessageSent)
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
		it.Event = new(LineaMessageSent)
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
func (it *LineaMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaMessageSent represents a MessageSent event raised by the Linea contract.
type LineaMessageSent struct {
	From        common.Address
	To          common.Address
	Fee         *big.Int
	Value       *big.Int
	Nonce       *big.Int
	Calldata    []byte
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*LineaMessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaMessageSentIterator{contract: _Linea.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *LineaMessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaMessageSent)
				if err := _Linea.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_Linea *LineaFilterer) ParseMessageSent(log types.Log) (*LineaMessageSent, error) {
	event := new(LineaMessageSent)
	if err := _Linea.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Linea contract.
type LineaPausedIterator struct {
	Event *LineaPaused // Event containing the contract specifics and raw log

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
func (it *LineaPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaPaused)
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
		it.Event = new(LineaPaused)
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
func (it *LineaPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaPaused represents a Paused event raised by the Linea contract.
type LineaPaused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) FilterPaused(opts *bind.FilterOpts) (*LineaPausedIterator, error) {

	logs, sub, err := _Linea.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LineaPausedIterator{contract: _Linea.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LineaPaused) (event.Subscription, error) {

	logs, sub, err := _Linea.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaPaused)
				if err := _Linea.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) ParsePaused(log types.Log) (*LineaPaused, error) {
	event := new(LineaPaused)
	if err := _Linea.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Linea contract.
type LineaRoleAdminChangedIterator struct {
	Event *LineaRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LineaRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaRoleAdminChanged)
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
		it.Event = new(LineaRoleAdminChanged)
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
func (it *LineaRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaRoleAdminChanged represents a RoleAdminChanged event raised by the Linea contract.
type LineaRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Linea *LineaFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LineaRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LineaRoleAdminChangedIterator{contract: _Linea.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Linea *LineaFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LineaRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaRoleAdminChanged)
				if err := _Linea.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Linea *LineaFilterer) ParseRoleAdminChanged(log types.Log) (*LineaRoleAdminChanged, error) {
	event := new(LineaRoleAdminChanged)
	if err := _Linea.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Linea contract.
type LineaRoleGrantedIterator struct {
	Event *LineaRoleGranted // Event containing the contract specifics and raw log

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
func (it *LineaRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaRoleGranted)
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
		it.Event = new(LineaRoleGranted)
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
func (it *LineaRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaRoleGranted represents a RoleGranted event raised by the Linea contract.
type LineaRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LineaRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LineaRoleGrantedIterator{contract: _Linea.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LineaRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaRoleGranted)
				if err := _Linea.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) ParseRoleGranted(log types.Log) (*LineaRoleGranted, error) {
	event := new(LineaRoleGranted)
	if err := _Linea.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Linea contract.
type LineaRoleRevokedIterator struct {
	Event *LineaRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LineaRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaRoleRevoked)
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
		it.Event = new(LineaRoleRevoked)
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
func (it *LineaRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaRoleRevoked represents a RoleRevoked event raised by the Linea contract.
type LineaRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LineaRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LineaRoleRevokedIterator{contract: _Linea.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LineaRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaRoleRevoked)
				if err := _Linea.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Linea *LineaFilterer) ParseRoleRevoked(log types.Log) (*LineaRoleRevoked, error) {
	event := new(LineaRoleRevoked)
	if err := _Linea.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the Linea contract.
type LineaUnPausedIterator struct {
	Event *LineaUnPaused // Event containing the contract specifics and raw log

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
func (it *LineaUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaUnPaused)
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
		it.Event = new(LineaUnPaused)
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
func (it *LineaUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaUnPaused represents a UnPaused event raised by the Linea contract.
type LineaUnPaused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) FilterUnPaused(opts *bind.FilterOpts) (*LineaUnPausedIterator, error) {

	logs, sub, err := _Linea.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &LineaUnPausedIterator{contract: _Linea.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *LineaUnPaused) (event.Subscription, error) {

	logs, sub, err := _Linea.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaUnPaused)
				if err := _Linea.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_Linea *LineaFilterer) ParseUnPaused(log types.Log) (*LineaUnPaused, error) {
	event := new(LineaUnPaused)
	if err := _Linea.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaVerifierAddressChangedIterator is returned from FilterVerifierAddressChanged and is used to iterate over the raw logs and unpacked data for VerifierAddressChanged events raised by the Linea contract.
type LineaVerifierAddressChangedIterator struct {
	Event *LineaVerifierAddressChanged // Event containing the contract specifics and raw log

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
func (it *LineaVerifierAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaVerifierAddressChanged)
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
		it.Event = new(LineaVerifierAddressChanged)
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
func (it *LineaVerifierAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaVerifierAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaVerifierAddressChanged represents a VerifierAddressChanged event raised by the Linea contract.
type LineaVerifierAddressChanged struct {
	VerifierAddress common.Address
	ProofType       *big.Int
	VerifierSetBy   common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVerifierAddressChanged is a free log retrieval operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_Linea *LineaFilterer) FilterVerifierAddressChanged(opts *bind.FilterOpts, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (*LineaVerifierAddressChangedIterator, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _Linea.contract.FilterLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return &LineaVerifierAddressChangedIterator{contract: _Linea.contract, event: "VerifierAddressChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierAddressChanged is a free log subscription operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_Linea *LineaFilterer) WatchVerifierAddressChanged(opts *bind.WatchOpts, sink chan<- *LineaVerifierAddressChanged, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (event.Subscription, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _Linea.contract.WatchLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaVerifierAddressChanged)
				if err := _Linea.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
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

// ParseVerifierAddressChanged is a log parse operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_Linea *LineaFilterer) ParseVerifierAddressChanged(log types.Log) (*LineaVerifierAddressChanged, error) {
	event := new(LineaVerifierAddressChanged)
	if err := _Linea.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

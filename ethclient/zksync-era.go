package ethclient

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	ZKSNewPriorityRequestTopic0 = "0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382"
)

var (
	ErrWithdrawalFinalizationRequestNotFound = errors.New("withdrawal finalization request not found")
)

func (ec Client) GetZKSL1BatchNumber(ctx context.Context) (uint64, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "zks_L1BatchNumber")
	return result.ToInt().Uint64(), err
}

type RpcL1BatchDetails struct {
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAA  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	CommitTxHash   *common.Hash `json:"commitTxHash"`
	CommittedAt    *time.Time   `json:"committedAt"`
	ExecuteTxHash  *common.Hash `json:"executeTxHash"`
	ExecutedAt     *time.Time   `json:"executedAt"`
	L1GasPrice     uint64       `json:"l1GasPrice"`
	L1TxCount      uint64       `json:"l1TxCount"`
	L2FairGasPrice uint64       `json:"l2FairGasPrice"`
	L2TxCount      uint64       `json:"l2TxCount"`
	Number         uint64       `json:"number"`
	ProveTxHash    *common.Hash `json:"proveTxHash"`
	ProvenAt       *time.Time   `json:"provenAt"`
	RootHash       *common.Hash `json:"rootHash"`
	Status         string       `json:"status"`
	Timestamp      int64        `json:"timestamp"`
}

func (c *Client) GetZKSL1BatchDetails(ctx context.Context, batchNumber uint64) (*RpcL1BatchDetails, error) {
	var result RpcL1BatchDetails
	err := c.rpcClient.CallContext(ctx, &result, "zks_getL1BatchDetails", batchNumber)
	return &result, err
}

type BlockRange struct {
	Start hexutil.Big
	End   hexutil.Big
}

func (ec Client) GetZKSL1BatchBlockRange(ctx context.Context, batchNumber uint64) (BlockRange, error) {
	var result [2]hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "zks_getL1BatchBlockRange", batchNumber)
	if err != nil {
		return BlockRange{}, err
	}
	return BlockRange{Start: result[0], End: result[1]}, nil
}

// For Bridge History
type ZKSDepositEvent struct {
	*contract.MailboxFacetNewPriorityRequest
	BlockTimestamp uint64
}

// GetZKSL1DepositEvents returns the NewPriorityRequest events of the ZKSDiamondProxy contract.
func (ec Client) GetZKSL1DepositEvents(ctx context.Context, ZKSDiamondProxy string, startBlock uint64, endBlock uint64) ([]ZKSDepositEvent, error) {
	mailbox, err := contract.NewMailboxFacet(common.HexToAddress(ZKSDiamondProxy), ec)
	if err != nil {
		return nil, err
	}

	// Define the filter query
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlock),
		ToBlock:   new(big.Int).SetUint64(endBlock),
		Addresses: []common.Address{common.HexToAddress(ZKSDiamondProxy)},
		Topics:    [][]common.Hash{{common.HexToHash(ZKSNewPriorityRequestTopic0)}},
	}

	// Fetch the logs
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	var result []ZKSDepositEvent
	// Iterate over the logs and parse them
	for _, vLog := range logs {
		// Parse the log into the specific event struct
		event, err := mailbox.ParseNewPriorityRequest(vLog)
		if err != nil {
			return nil, err
		}

		blockTime, err := ec.GetBlockTimestampByTxHash(ctx, vLog.TxHash)
		if err != nil {
			return nil, err
		}

		result = append(result, ZKSDepositEvent{
			MailboxFacetNewPriorityRequest: event,
			BlockTimestamp:                 blockTime,
		})

	}

	return result, nil
}

// ParseWithdrawalFinalizationInput parses the input data of the finalizeWithdrawals method of the WithdrawalFinalizer contract.
func ParseWithdrawalFinalizationInput(inputData string) ([]contract.WithdrawalFinalizerRequestFinalizeWithdrawal, error) {
	inputBytes, err := hex.DecodeString(strings.TrimPrefix(inputData, "0x"))
	if err != nil {
		return nil, err
	}

	contractABI, err := contract.WithdrawalFinalizerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	method, ok := contractABI.Methods["finalizeWithdrawals"]
	if !ok {
		return nil, ErrWithdrawalFinalizationRequestNotFound
	}

	methodID := method.ID
	if !strings.HasPrefix(strings.TrimPrefix(inputData, "0x"), hex.EncodeToString(methodID[:])) {
		return nil, errors.New("input data does not match the finalizeWithdrawals method")
	}

	unpacked, err := method.Inputs.Unpack(inputBytes[4:])
	if err != nil {
		return nil, err
	}

	var params []contract.WithdrawalFinalizerRequestFinalizeWithdrawal
	for _, result := range unpacked {
		switch v := result.(type) {
		case []struct {
			L2BlockNumber     *big.Int    `json:"_l2BlockNumber"` // actually this means l1BatchNumber
			L2MessageIndex    *big.Int    `json:"_l2MessageIndex"`
			L2TxNumberInBlock uint16      `json:"_l2TxNumberInBlock"` // actually this means l1BatchTxIndex
			Message           []uint8     `json:"_message"`
			MerkleProof       [][32]uint8 `json:"_merkleProof"`
			IsEth             bool        `json:"_isEth"`
			Gas               *big.Int    `json:"_gas"`
		}:
			for _, item := range v {
				params = append(params, contract.WithdrawalFinalizerRequestFinalizeWithdrawal{
					L2BlockNumber:     item.L2BlockNumber,
					L2MessageIndex:    item.L2MessageIndex,
					L2TxNumberInBlock: item.L2TxNumberInBlock,
					Message:           item.Message,
					MerkleProof:       item.MerkleProof,
					IsEth:             item.IsEth,
					Gas:               item.Gas,
				})
			}
		default:
			log.Fatal().Msgf("Unexpected type: %T\n", v)
		}
	}

	return params, nil
}

type ETHWithdrawalFinalizerData struct {
	L2BlockNumber     *big.Int
	L2MessageIndex    *big.Int
	L2TxNumberInBlock uint16
	Message           []byte
	MerkleProof       [][32]byte
}

// ParseWithdrawalFinalizationInput parses the input data of the FinalizeEthWithdrawal method of the ZksyncDiamondProxy contract.
func ParseETHWithdrawalFinalizationInput(inputData string) (ETHWithdrawalFinalizerData, error) {
	inputBytes, err := hex.DecodeString(strings.TrimPrefix(inputData, "0x"))
	if err != nil {
		return ETHWithdrawalFinalizerData{}, err
	}

	contractABI, err := contract.MailboxFacetMetaData.GetAbi()
	if err != nil {
		return ETHWithdrawalFinalizerData{}, err
	}

	method, ok := contractABI.Methods["finalizeEthWithdrawal"]
	if !ok {
		return ETHWithdrawalFinalizerData{}, ErrWithdrawalFinalizationRequestNotFound
	}

	methodID := method.ID
	if !strings.HasPrefix(strings.TrimPrefix(inputData, "0x"), hex.EncodeToString(methodID[:])) {
		return ETHWithdrawalFinalizerData{}, nil
	}

	unpacked, err := method.Inputs.Unpack(inputBytes[4:])
	if err != nil {
		return ETHWithdrawalFinalizerData{}, err
	}

	if len(unpacked) != 5 {
		return ETHWithdrawalFinalizerData{}, errors.New("unexpected number of parameters in unpacked data")
	}
	params := ETHWithdrawalFinalizerData{}
	params.L2BlockNumber, _ = unpacked[0].(*big.Int)
	params.L2MessageIndex, _ = unpacked[1].(*big.Int)
	params.L2TxNumberInBlock, _ = unpacked[2].(uint16)
	params.Message, _ = unpacked[3].([]byte)
	params.MerkleProof, _ = unpacked[4].([][32]byte)

	return params, nil
}

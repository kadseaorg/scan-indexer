package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	SequencerInbox                         = "0x1c479675ad559DC151F6Ec7ed3FbF8ceE79582B6"
	OldSequencer                           = "0x4c6f947ae67f572afa4ae0730947de7c874f95ef"
	SequencerBatchDeliveredTopic           = "0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7"
	SequencerBatchDeliveredFromOriginTopic = "0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682"
	NodeInterfaceAddr                      = "0x00000000000000000000000000000000000000C8"
)

var (
	uint64T, _ = abi.NewType("uint64", "", nil)

	findBatchContainingBlockMethod = abi.NewMethod(
		"findBatchContainingBlock",
		"findBatchContainingBlock",
		abi.Function,
		"",
		false,
		true,
		abi.Arguments{
			{Name: "l2BlockNum", Type: uint64T},
		},
		abi.Arguments{
			{Name: "l1BatchNum", Type: uint64T},
		},
	)
)

type ArbOneSequencerBatchDeliveredEvent struct {
	BatchSequenceNumber   uint64
	L1BlockNumber         uint64
	L1TxHash              common.Hash
	Timestamp             uint64
	FinalizedBlockNumbers []*big.Int
	Number                uint64
	OutputRootHash        common.Hash
	L2BlockNumber         uint64
}

func (ec Client) GetArbOneBatchContainingBlock(ctx context.Context, L2BlockNumber uint64) (uint64, error) {
	input, err := findBatchContainingBlockMethod.Inputs.Pack(L2BlockNumber)
	if err != nil {
		return 0, err
	}

	req := map[string]any{
		"from": common.HexToAddress("0x"),
		"to":   common.HexToAddress("0x00000000000000000000000000000000000000C8"),
		"data": hexutil.Encode(append(findBatchContainingBlockMethod.ID, input...)),
	}

	var out any
	if err := ec.rpcClient.Call(&out, "eth_call", &req, "latest"); err != nil {
		return 0, err
	}

	hex, ok := out.(string)
	if !ok {
		return 0, err
	}
	data, err := hexutil.Decode(hex)
	if err != nil {
		return 0, err
	}

	nums, err := findBatchContainingBlockMethod.Outputs.Unpack(data)
	if err != nil {
		return 0, err
	}

	return nums[0].(uint64), err
}

// Old SequencerBatchDeliveredFromOriginTopic is used for L1 blocks from 0 - 15447147
// After that, the new SequencerBatchDeliveredTopic is used
func (ec Client) GetArbOneSequencerBatchDeliveredEventLogs(ctx context.Context, fromBlock, toBlock uint64) ([]*ArbOneSequencerBatchDeliveredEvent, error) {
	seqInboxContract, err := contract.NewArbOne(common.HexToAddress(SequencerInbox), ec)
	if err != nil {
		return nil, err
	}

	oldquery := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(OldSequencer)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(SequencerBatchDeliveredFromOriginTopic),
			},
		},
	}
	oldlogs, err := ec.FilterLogs(ctx, oldquery)
	if err != nil {
		return nil, err
	}

	newquery := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(SequencerInbox)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(SequencerBatchDeliveredTopic),
			},
		},
	}
	newlogs, err := ec.FilterLogs(ctx, newquery)
	if err != nil {
		return nil, err
	}

	alllogs := append(oldlogs, newlogs...)

	txBatchMap := make(map[common.Hash]*ArbOneSequencerBatchDeliveredEvent)
	for _, txlog := range alllogs {
		if len(txlog.Topics) == 0 {
			continue
		}

		if txlog.Topics[0] == common.HexToHash(SequencerBatchDeliveredTopic) {
			deliveredBatch, err := seqInboxContract.ParseSequencerBatchDelivered(txlog)
			if err != nil {
				return nil, err
			}

			txBatchMap[txlog.TxHash] = &ArbOneSequencerBatchDeliveredEvent{
				Number:              deliveredBatch.BatchSequenceNumber.Uint64(),
				BatchSequenceNumber: deliveredBatch.BatchSequenceNumber.Uint64(),
				L1BlockNumber:       txlog.BlockNumber,
				L1TxHash:            txlog.TxHash,
				OutputRootHash:      txlog.TxHash,
			}
		} else if txlog.Topics[0] == common.HexToHash(SequencerBatchDeliveredFromOriginTopic) {
			deliveredBatch, err := seqInboxContract.ParseSequencerBatchDeliveredFromOrigin(txlog)
			if err != nil {
				return nil, err
			}

			txBatchMap[txlog.TxHash] = &ArbOneSequencerBatchDeliveredEvent{
				Number:              deliveredBatch.SeqBatchIndex.Uint64(),
				BatchSequenceNumber: deliveredBatch.SeqBatchIndex.Uint64(),
				L1BlockNumber:       txlog.BlockNumber,
				L1TxHash:            txlog.TxHash,
				OutputRootHash:      txlog.TxHash,
			}
		}
	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*ArbOneSequencerBatchDeliveredEvent
	for _, txBatch := range txBatchMap {
		if _, ok := blockTimeMap[txBatch.L1BlockNumber]; !ok {
			block, err := ec.HeaderByNumber(ctx, big.NewInt(int64(txBatch.L1BlockNumber)))
			if err != nil {
				return nil, err
			}
			blockTimeMap[txBatch.L1BlockNumber] = block.Time
		}
		txBatch.Timestamp = blockTimeMap[txBatch.L1BlockNumber]
		txBatches = append(txBatches, txBatch)
	}

	return txBatches, nil
}

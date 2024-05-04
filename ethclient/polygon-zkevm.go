package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	PolygonZkEVMSequencerInbox = "0x5132A183E9F3CB7C848b0AAC5Ae0c4f0491B7aB2"

	PolygonZkEVMSequenceBatchesTopic = "0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe"

	BsquaredSequencerInbox = "0x67d269191c92Caf3cD7723F116c85e6E9bf55933"

	OKX2SequencerInbox = "0x6662621411A8DACC3cA7049C8BddABaa9a999ce3"

	OKX1SequenceBatchesTopic = "0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe"
)

type PolygonZkEVMBatchEvent struct {
	L1BlockNumber uint64
	L1TxHash      common.Hash
	L2InitBatch   uint64
	L2FinalBatch  uint64
	L2BatchBlocks []*big.Int
	Timestamp     uint64
	RootHash      *common.Hash
}

func (ec Client) GetPolygonZkEVMSequenceBatchesEventLogs(ctx context.Context, rollupAddr string, fromBlock, toBlock uint64) ([]*PolygonZkEVMBatchEvent, error) {
	seqInboxContract, err := contract.NewPolygonzkevm(common.HexToAddress(rollupAddr), ec)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(rollupAddr)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(PolygonZkEVMSequenceBatchesTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	txBatchMap := make(map[common.Hash]*PolygonZkEVMBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0] != common.HexToHash(PolygonZkEVMSequenceBatchesTopic) {
			continue
		}

		data, err := seqInboxContract.ParseVerifyBatchesTrustedAggregator(log)
		if err != nil {
			return nil, err
		}

		txBatchMap[log.TxHash] = &PolygonZkEVMBatchEvent{
			L2FinalBatch:  data.NumBatch,
			L1TxHash:      log.TxHash,
			L1BlockNumber: log.BlockNumber,
			RootHash:      &data.Raw.TxHash,
			L2BatchBlocks: make([]*big.Int, 0),
		}
	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*PolygonZkEVMBatchEvent
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

func (ec Client) GetOKX1SequenceBatchesEventLogs(ctx context.Context, rollupAddr string, fromBlock, toBlock uint64) ([]*PolygonZkEVMBatchEvent, error) {
	seqInboxContract, err := contract.NewPolygonzkevm(common.HexToAddress(rollupAddr), ec)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(rollupAddr)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(OKX1SequenceBatchesTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	txBatchMap := make(map[common.Hash]*PolygonZkEVMBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0] != common.HexToHash(PolygonZkEVMSequenceBatchesTopic) {
			continue
		}

		data, err := seqInboxContract.ParseVerifyBatchesTrustedAggregator(log)
		if err != nil {
			return nil, err
		}

		txBatchMap[log.TxHash] = &PolygonZkEVMBatchEvent{
			L2FinalBatch:  data.NumBatch,
			L1TxHash:      log.TxHash,
			L1BlockNumber: log.BlockNumber,
			RootHash:      &data.Raw.TxHash,
			L2BatchBlocks: make([]*big.Int, 0),
		}
	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*PolygonZkEVMBatchEvent
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

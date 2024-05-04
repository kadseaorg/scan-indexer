package ethclient

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	OKX1SepoliaConsensusSmartContract = "0x6662621411A8DACC3cA7049C8BddABaa9a999ce3"
	OKX1SepoliaSequenceBatchesTopic   = "0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe"
	OKX1L1InitialBatchHeight          = 4653166
)

type OKX1SepoliaBatchEvent struct {
	L1BlockNumber uint64
	L1TxHash      common.Hash
	L2InitBatch   uint64
	L2FinalBatch  uint64
	L2BatchBlocks []*big.Int
	Timestamp     uint64
	RootHash      *common.Hash
}

func (ec Client) GetOKX1SepoliaSequenceBatchesEventLogs(ctx context.Context, rollupAddr string, fromBlock, toBlock uint64) ([]*OKX1SepoliaBatchEvent, error) {
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
				common.HexToHash(OKX1SepoliaSequenceBatchesTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	txBatchMap := make(map[common.Hash]*OKX1SepoliaBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0] != common.HexToHash(OKX1SepoliaSequenceBatchesTopic) {
			continue
		}

		data, err := seqInboxContract.ParseVerifyBatchesTrustedAggregator(log)
		if err != nil {
			return nil, err
		}
		txBatchMap[log.TxHash] = &OKX1SepoliaBatchEvent{
			L2FinalBatch:  data.NumBatch,
			L1TxHash:      log.TxHash,
			L1BlockNumber: log.BlockNumber,
			RootHash:      &data.Raw.TxHash,
			L2BatchBlocks: make([]*big.Int, 0),
		}
	}
	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*OKX1SepoliaBatchEvent
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

	SortOKX1BatchEvents(txBatches)

	return txBatches, nil
}

func SortOKX1BatchEvents(events []*OKX1SepoliaBatchEvent) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].L2FinalBatch < events[j].L2FinalBatch
	})
}

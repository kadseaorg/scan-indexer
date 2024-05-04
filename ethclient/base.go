package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	BaseL2OutputOracleProxy = "0x56315b90c40730925ec5485cf004d835058518A0"
	OutputProposedTopic     = "0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2"
)

type BaseBatchEvent struct {
	Number         uint64 // L2OutputIndex
	OutputRootHash common.Hash
	L2BlockNumber  uint64
	L1TxHash       common.Hash
	L1BlockNumber  uint64
	Timestamp      uint64
}

func (ec Client) GetBaseOutputProposedEventLogs(ctx context.Context, fromBlock, toBlock uint64) ([]*BaseBatchEvent, error) {
	outputOracleContract, err := contract.NewL2OutputOracle(common.HexToAddress(BaseL2OutputOracleProxy), ec)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(BaseL2OutputOracleProxy)},
		Topics: [][]common.Hash{
			{
				// OutputProposed
				common.HexToHash(OutputProposedTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	txBatchMap := make(map[common.Hash]*BaseBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] != common.HexToHash(OutputProposedTopic) {
			continue
		}

		// OutputProposed
		outputProposed, err := outputOracleContract.ParseOutputProposed(log)
		if err != nil {
			return nil, err
		}

		// one tx one batch
		txBatchMap[log.TxHash] = &BaseBatchEvent{
			OutputRootHash: outputProposed.OutputRoot,
			Number:         outputProposed.L2OutputIndex.Uint64(),
			L2BlockNumber:  outputProposed.L2BlockNumber.Uint64(),
			L1TxHash:       log.TxHash,
			L1BlockNumber:  log.BlockNumber,
		}

	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*BaseBatchEvent
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
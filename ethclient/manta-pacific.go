package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	MantaPacificL2OutputAddr = "0x30c789674ad3B458886BBC9abf42EEe19EA05C1D"
	MPL2OutputProposedTopic  = "0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2"
)

type MantaPacificBatchEvent struct {
	Number                uint64
	FinalRootHash         common.Hash
	Timestamp             uint64
	FinalizedBlockNumbers []*big.Int
	OutputRootHash        common.Hash
	L2BlockNumber         uint64
	L1Timestamp           uint64
	L1TxHash              common.Hash
	L1BlockNumber         uint64
}

func (ec Client) GetMantaPacificOutputProposedEventLogs(ctx context.Context, fromBlock, toBlock uint64) ([]*MantaPacificBatchEvent, error) {

	MantaPacificL2Contract, err := contract.NewMantaPacificL2(common.HexToAddress(MantaPacificL2OutputAddr), ec)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(MantaPacificL2OutputAddr)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(MPL2OutputProposedTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	txBatchMap := make(map[common.Hash]*MantaPacificBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] != common.HexToHash(MPL2OutputProposedTopic) {
			continue
		}

		// OutputProposed
		outputProposed, err := MantaPacificL2Contract.ParseOutputProposed(log)
		if err != nil {
			return nil, err
		}

		// event OutputProposed(
		// 	bytes32 indexed outputRoot,
		// 	uint256 indexed l2OutputIndex,
		// 	uint256 indexed l2BlockNumber,
		// 	uint256 l1Timestamp
		// );

		txBatchMap[log.TxHash] = &MantaPacificBatchEvent{
			OutputRootHash: outputProposed.OutputRoot,
			Number:         outputProposed.L2OutputIndex.Uint64(),
			L2BlockNumber:  outputProposed.L2BlockNumber.Uint64(),
			L1TxHash:       log.TxHash,
			L1BlockNumber:  log.BlockNumber,
		}

		// Each batch contains 120 L2 blocks.
		// For the first batch, it is special and contains 121 blocks including the block 0
		// which is gensis.
		// The L2BlockNumber in outputProposed indicates the number of the last block.
		L2BlockNumberEnd := outputProposed.L2BlockNumber.Uint64()
		L2BlockNumberBeg := L2BlockNumberEnd - 120
		if L2BlockNumberBeg != 0 {
			L2BlockNumberBeg = L2BlockNumberBeg + 1
		}

		for i := L2BlockNumberBeg; i <= L2BlockNumberEnd; i++ {
			t := big.NewInt(int64(i))
			txBatchMap[log.TxHash].FinalizedBlockNumbers = append(txBatchMap[log.TxHash].FinalizedBlockNumbers, t)
		}
	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*MantaPacificBatchEvent
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

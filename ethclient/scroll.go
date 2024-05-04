package ethclient

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unifra20/l2scan-indexer/contract"
)

const (
	BlockContextLength       = 60
	ScrollCommitBatchTopic   = "0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f"
	ScrollFinalizeBatchTopic = "0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d"
)

type ScrollBatchStatus string

const (
	ScrollBatchStatusCommitted ScrollBatchStatus = "committed"
	ScrollBatchStatusProven    ScrollBatchStatus = "finalized"
)

type ScrollBatchEvent struct {
	Number          uint64
	BlockNumber     uint64
	TxHash          common.Hash
	CommittedBlocks []uint64
	RootHash        *common.Hash
	Status          ScrollBatchStatus
	Timestamp       uint64
}

func (ec Client) GetScrollBatches(ctx context.Context, rollupAddr string, fromBlock, toBlock uint64) ([]*ScrollBatchEvent, error) {
	scrollRollupContract, err := contract.NewIScrollChain(common.HexToAddress(rollupAddr), ec)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(rollupAddr)},
		Topics: [][]common.Hash{
			{
				// CommitBatch
				common.HexToHash(ScrollCommitBatchTopic),
				// FinalizeBatch
				common.HexToHash(ScrollFinalizeBatchTopic),
			},
		},
	}
	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	// log.Debug().Msgf("GetScrollBatches: batchNumber=%d, fromBlock=%d, toBlock=%d, logs=%d", batchNumber, fromBlock, toBlock, len(logs))

	txBatchMap := make(map[common.Hash]*ScrollBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] == common.HexToHash(ScrollCommitBatchTopic) {
			// CommitBatch
			commitBatch, err := scrollRollupContract.ParseCommitBatch(log)
			if err != nil {
				return nil, err
			}

			// get input data from transaction
			tx, _, err := ec.TransactionByHash(ctx, log.TxHash)
			if err != nil {
				return nil, err
			}

			// parse committed blocks from input data
			committedBlocks, err := ParseCommittedBlocks(tx.Data())
			if err != nil {
				return nil, err
			}

			if _, ok := txBatchMap[log.TxHash]; !ok {
				txBatchMap[log.TxHash] = &ScrollBatchEvent{
					Number:          commitBatch.BatchIndex.Uint64(),
					BlockNumber:     log.BlockNumber,
					TxHash:          log.TxHash,
					CommittedBlocks: committedBlocks,
					Status:          ScrollBatchStatusCommitted,
				}
			}
		} else if log.Topics[0] == common.HexToHash(ScrollFinalizeBatchTopic) {
			// FinalizeBatch
			finalizBatch, err := scrollRollupContract.ParseFinalizeBatch(log)
			if err != nil {
				return nil, err
			}
			rootHash := common.BytesToHash(finalizBatch.StateRoot[:])
			if _, ok := txBatchMap[log.TxHash]; !ok {
				txBatchMap[log.TxHash] = &ScrollBatchEvent{
					Number:      finalizBatch.BatchIndex.Uint64(),
					BlockNumber: log.BlockNumber,
					TxHash:      log.TxHash,
					RootHash:    &rootHash,
					Status:      ScrollBatchStatusProven,
				}
			}
		}
	}

	// get block time
	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*ScrollBatchEvent
	for _, txBatch := range txBatchMap {
		if _, ok := blockTimeMap[txBatch.BlockNumber]; !ok {
			block, err := ec.HeaderByNumber(ctx, big.NewInt(int64(txBatch.BlockNumber)))
			if err != nil {
				return nil, err
			}
			blockTimeMap[txBatch.BlockNumber] = block.Time
		}
		txBatch.Timestamp = blockTimeMap[txBatch.BlockNumber]
		txBatches = append(txBatches, txBatch)
	}

	return txBatches, nil
}

type BlockContext struct {
	BlockNumber     uint64
	Timestamp       uint64
	BaseFee         *big.Int
	GasLimit        uint64
	NumTransactions uint16
	NumL1Messages   uint16
}

type Chunk struct {
	Blocks         []BlockContext
	L2Transactions []byte
}

// Function: commitBatch(uint8 _version,bytes _parentBatchHeader,bytes[] _chunks,bytes _skippedL1MessageBitmap)
type commitBlocksInput struct {
	Version                uint8
	ParentBatchHeader      []byte
	Chunks                 [][]byte
	SkippedL1MessageBitmap []byte
}

func ParseCommittedBlocks(input []byte) ([]uint64, error) {
	abi, err := contract.IScrollChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	// parse input data
	commitBlocksInput := new(commitBlocksInput)
	unpacked, err := abi.Methods["commitBatch"].Inputs.Unpack(input[4:])
	if err != nil {
		return nil, err
	}
	err = abi.Methods["commitBatch"].Inputs.Copy(commitBlocksInput, unpacked)
	if err != nil {
		return nil, err
	}

	// parse chunk
	chunks := make([]Chunk, len(commitBlocksInput.Chunks))
	for i, chunkData := range commitBlocksInput.Chunks {
		chunk, err := ParseChunk(chunkData)
		if err != nil {
			return nil, err
		}
		chunks[i] = *chunk
	}

	// parse committed blocks
	var committedBlocks []uint64
	for _, chunk := range chunks {
		for _, block := range chunk.Blocks {
			committedBlocks = append(committedBlocks, block.BlockNumber)
		}
	}

	return committedBlocks, nil
}

func ParseChunk(chunkData []byte) (*Chunk, error) {
	chunkLength := len(chunkData)
	numBlocks := int(chunkData[0])

	if numBlocks == 0 {
		return nil, fmt.Errorf("no block in chunk")
	}

	if chunkLength < 1+numBlocks*BlockContextLength {
		return nil, fmt.Errorf("invalid chunk length")
	}

	blocks := make([]BlockContext, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blockStart := 1 + i*BlockContextLength
		block := BlockContext{
			BlockNumber:     binary.BigEndian.Uint64(chunkData[blockStart : blockStart+8]),
			Timestamp:       binary.BigEndian.Uint64(chunkData[blockStart+8 : blockStart+16]),
			BaseFee:         new(big.Int).SetBytes(chunkData[blockStart+16 : blockStart+48]),
			GasLimit:        binary.BigEndian.Uint64(chunkData[blockStart+48 : blockStart+56]),
			NumTransactions: binary.BigEndian.Uint16(chunkData[blockStart+56 : blockStart+58]),
			NumL1Messages:   binary.BigEndian.Uint16(chunkData[blockStart+58 : blockStart+60]),
		}
		blocks[i] = block
	}

	l2TxStart := 1 + numBlocks*BlockContextLength
	l2Transactions := chunkData[l2TxStart:]

	return &Chunk{
		Blocks:         blocks,
		L2Transactions: l2Transactions,
	}, nil
}

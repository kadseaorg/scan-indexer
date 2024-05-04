package ethclient

import (
	"context"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/contract"
	contract_upgrade "github.com/unifra20/l2scan-indexer/contract/upgrade"
)

const (
	LineaZkevmV2Address         = "0xd19d4B5d358258f05D7B411E21A1460D11B0876F"
	BlockFinalizedTopic         = "0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432"
	DataFinalizedTopic          = "0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c"
	BlocksVerificationDoneTopic = "0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f"

	// After block height 19217957 on l1, BlockFinalized has been replaced by DataFinalized.
	BlockFinalizeUpgradeBlockHeightOnL1 = 19217957
	BlockFinalizeUpgradeBlockHeightOnL2 = 2238199
)

var (
	lineaContract         *LineaContract
	initLineaContractOnce sync.Once
)

type LineaBatchEvent struct {
	Number                uint64 // ignore on linea
	FinalRootHash         common.Hash
	TransactionHash       common.Hash
	Timestamp             uint64
	FinalizedBlockNumbers []uint64
	BlockNumber           uint64
}

type LineaContract struct {
	oldLinea *contract.Linea
	newLinea *contract_upgrade.Linea
}

func (c LineaContract) ParseEthLog(logs types.Log, last_l2_block_number *uint64, txBatchMap map[common.Hash]*LineaBatchEvent) error {
	blockHeight := logs.BlockNumber
	if blockHeight <= BlockFinalizeUpgradeBlockHeightOnL1 {
		if logs.Topics[0] == common.HexToHash(BlockFinalizedTopic) {
			blockFinalized, err := c.oldLinea.ParseBlockFinalized(logs)
			if err != nil {
				return err
			}
			txBatchMap[logs.TxHash].FinalizedBlockNumbers = append(txBatchMap[logs.TxHash].FinalizedBlockNumbers, blockFinalized.BlockNumber.Uint64())
		} else if logs.Topics[0] == common.HexToHash(BlocksVerificationDoneTopic) {
			verificationDoneEvent, err := c.oldLinea.ParseBlocksVerificationDone(logs)
			if err != nil {
				return err
			}
			txBatchMap[logs.TxHash].FinalRootHash = verificationDoneEvent.FinalRootHash
		}
	} else {
		if logs.Topics[0] == common.HexToHash(DataFinalizedTopic) {
			dataFinalized, err := c.newLinea.ParseDataFinalized(logs)
			if err != nil {
				return err
			}

			syncedFinalizedBlockNumber := dataFinalized.LastBlockFinalized.Uint64()
			finalizedBlockNumbers := make([]uint64, syncedFinalizedBlockNumber-*last_l2_block_number)

			for i := uint64(0); i < syncedFinalizedBlockNumber-*last_l2_block_number; i++ {
				finalizedBlockNumbers[i] = *last_l2_block_number + 1 + i
			}
			txBatchMap[logs.TxHash].FinalizedBlockNumbers = finalizedBlockNumbers
			txBatchMap[logs.TxHash].FinalRootHash = dataFinalized.FinalRootHash

			// Update last_l2_block_number, and then we will know all the finalized blocks of the current transaction
			// by LastBlockFinalized in DataFinalize
			*last_l2_block_number = syncedFinalizedBlockNumber
		} else if logs.Topics[0] == common.HexToHash(BlocksVerificationDoneTopic) {
			verificationDoneEvent, err := c.oldLinea.ParseBlocksVerificationDone(logs)
			if err != nil {
				return err
			}

			syncedFinalizedBlockNumber := verificationDoneEvent.LastBlockFinalized.Uint64()

			finalizedBlockNumbers := make([]uint64, syncedFinalizedBlockNumber-*last_l2_block_number)

			for i := uint64(0); i < syncedFinalizedBlockNumber-*last_l2_block_number; i++ {
				finalizedBlockNumbers[i] = *last_l2_block_number + 1 + i
			}

			txBatchMap[logs.TxHash].FinalizedBlockNumbers = finalizedBlockNumbers
			txBatchMap[logs.TxHash].FinalRootHash = verificationDoneEvent.FinalRootHash
			*last_l2_block_number = syncedFinalizedBlockNumber
		}
	}

	return nil
}

func (ec Client) initLineaContract() {
	initLineaContractOnce.Do(func() {
		oldLinea, err := contract.NewLinea(common.HexToAddress(LineaZkevmV2Address), ec)
		if err != nil {
			log.Error().Msgf("Error crate old linea contract failed: %s", err)
			return
		}
		newLinea, err := contract_upgrade.NewLinea(common.HexToAddress(LineaZkevmV2Address), ec)
		if err != nil {
			log.Error().Msgf("Error crate old linea contract failed: %s", err)
			return
		}

		lineaContract = &LineaContract{
			oldLinea,
			newLinea,
		}
	})
}

func (ec Client) GetLineaBatchEventsFromL1(ctx context.Context, last_l2_block_number *uint64, fromBlock, toBlock uint64) ([]*LineaBatchEvent, error) {
	ec.initLineaContract()

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(LineaZkevmV2Address)},
		Topics: [][]common.Hash{
			{
				common.HexToHash(BlockFinalizedTopic),
				common.HexToHash(DataFinalizedTopic),
				common.HexToHash(BlocksVerificationDoneTopic),
			},
		},
	}

	logs, err := ec.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	// log.Debug().Msgf("GetLineaFinalizedEventLogs: batchNumber=%d, fromBlock=%d, toBlock=%d, logs=%d", batchNumber, fromBlock, toBlock, len(logs))

	txBatchMap := make(map[common.Hash]*LineaBatchEvent)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}

		if _, ok := txBatchMap[log.TxHash]; !ok {
			txBatchMap[log.TxHash] = &LineaBatchEvent{
				TransactionHash: log.TxHash,
				BlockNumber:     log.BlockNumber,
			}
		}

		if err := lineaContract.ParseEthLog(log, last_l2_block_number, txBatchMap); err != nil {
			return nil, err
		}
	}

	blockTimeMap := make(map[uint64]uint64)
	var txBatches []*LineaBatchEvent
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

	SortLineaBatchEvents(txBatches)

	return txBatches, nil
}

func SortLineaBatchEvents(events []*LineaBatchEvent) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].BlockNumber < events[j].BlockNumber
	})
}

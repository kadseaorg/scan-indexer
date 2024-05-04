package indexer

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (i *Indexer) StartL1IndexerForPolygonZkEVM(worker int, forceStartBlock *uint64) {
	log.Info().Msg("🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢 Starting L1 Indexer 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢")
	var lastSyncedBlockNumber uint64
	var err error

	if forceStartBlock != nil {
		log.Warn().Msgf("Force starting L1 indexer from block %d", *forceStartBlock)
		lastSyncedBlockNumber = *forceStartBlock
	} else {
		lastSyncedBlockNumber, err = i.dbClient.GetLatestL1SyncedBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 block number: %s", err)
			return
		}

	}

	var eventStep uint64
	switch i.chain.Name {
	case BsquaredTestnet.Name:
		eventStep = 500

	case PolygonZkEVM.Name:
		eventStep = 300

	default:
		eventStep = 50
	}

	for {
		startBlock := lastSyncedBlockNumber + 1
		for {
			latestBlockNumber, err := i.l1client.BlockNumber(i.ctx)
			if err != nil {
				// exit and let the process restart
				log.Error().Msgf("Error getting latest block number: %s", err)
				os.Exit(-1)
			}

			//It means that it has been synchronized to the latest height,
			//so wait a little longer to reduce the request pressure on the L1 node.
			if startBlock-eventStep-1 >= latestBlockNumber {
				time.Sleep(time.Second * 30)
			}

			// reset startBlock if it's greater than latestBlockNumber
			if startBlock > latestBlockNumber {
				startBlock = latestBlockNumber
			}
			endBlock := startBlock + eventStep
			if endBlock > latestBlockNumber {
				endBlock = latestBlockNumber
			}

			batchNumbers, batchesLen, err := i.indexL1BatchesForPolygonZkEVM(startBlock, endBlock)
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
					// has indexed, skip this block range
					log.Info().Msgf("🤷 [L1 Indexer] Already indexed range: [%d, %d]", startBlock, endBlock)
					startBlock = endBlock + 1
					continue
				}
				log.Error().Msgf("Error indexing L1 batches for range: [%d, %d]: %s", startBlock, endBlock, err)
				break
			}
			if batchesLen == 0 {
				log.Info().Msgf("🔍 [L1 Indexer] No new batches for range: [%d, %d], skipping", startBlock, endBlock)
				startBlock = endBlock + 1
				continue
			}

			log.Info().Msgf("🟢 [L1 Indexer] Indexed batch num: %d, range: [%d, %d], len: %d", batchNumbers, startBlock, endBlock, batchesLen)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 Indexer] Have indexed all batches up to block %d", latestBlockNumber)
				break
			}

			startBlock = endBlock + 1
		}
		lastSyncedBlockNumber, err = i.dbClient.GetLatestL1SyncedBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 block number in loop: %s", err)
		}
		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
	}
}

func (i *Indexer) indexL1BatchesForPolygonZkEVM(from uint64, to uint64) ([]uint64, int, error) {
	batchEvents, err := i.l1client.GetPolygonZkEVMSequenceBatchesEventLogs(i.ctx, i.chain.RollupContractAddress, from, to)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error getting L1 batch events: %w", err)
	}

	if len(batchEvents) == 0 {
		return []uint64{0}, 0, nil
	}

	L2InitBatch, err := i.dbClient.GetLatestBatchNumber(i.ctx)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error getting L2 Init batch information: %w", err)
	}

	newBatchEvents := make([]*ethclient.PolygonZkEVMBatchEvent, 0)
	for _, batch := range batchEvents {
		for batchNum := L2InitBatch + 1; batchNum <= batch.L2FinalBatch; batchNum++ {
			newBatchEvent := &ethclient.PolygonZkEVMBatchEvent{
				L1BlockNumber: batch.L1BlockNumber,
				L1TxHash:      batch.L1TxHash,
				L2InitBatch:   batch.L2InitBatch,
				L2FinalBatch:  batchNum,
				L2BatchBlocks: batch.L2BatchBlocks,
				Timestamp:     batch.Timestamp,
				RootHash:      batch.RootHash,
			}
			newBatchEvents = append(newBatchEvents, newBatchEvent)
		}
		L2InitBatch = batch.L2FinalBatch
	}

	for _, batch := range newBatchEvents {
		batchInfo, err := i.l2client.GetBatchByNumber(i.ctx, batch.L2FinalBatch)
		if err != nil {
			return []uint64{0}, 0, fmt.Errorf("error getting L2 batch information: %w", err)
		}
		for _, blockHash := range batchInfo.Blocks {
			blockInfo, err := i.l2client.GetBlockByHash(i.ctx, blockHash)
			if err != nil {
				return []uint64{0}, 0, fmt.Errorf("error getting L2 block information: %w", err)
			}
			batch.L2BatchBlocks = append(batch.L2BatchBlocks, blockInfo.Number.ToInt())
		}
	}

	// check if we need to wait for the block to be finalized
	if err := i.WaitingWriteBatchForPolygonZkEVM(newBatchEvents); err != nil {
		return []uint64{0}, 0, fmt.Errorf("error waiting for block to be finalized: %w", err)
	}

	batchNumbers, err := i.dbClient.UpsertL1BatchForPolygonZkevmAll(i.ctx, newBatchEvents)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error upserting L1 batch events: %w", err)
	}

	err = i.dbClient.UpsertL1SyncProgress(i.ctx, to)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error upserting L1 sync progress: %w", err)
	}

	return *batchNumbers, len(newBatchEvents), nil
}

func (i *Indexer) WaitingWriteBatchForPolygonZkEVM(batchEvents []*ethclient.PolygonZkEVMBatchEvent) error {
	latestIndexedBlockNumber, err := i.dbClient.GetLatestBlockNumber(i.ctx)
	if err != nil {
		return fmt.Errorf("error getting latest indexed block number: %w", err)
	}

	for _, event := range batchEvents {
		for _, l2block := range event.L2BatchBlocks {
			if l2block.Uint64() <= latestIndexedBlockNumber {
				continue
			}
			for {
				time.Sleep(10 * time.Second)
				latestIndexedBlockNumber, err = i.dbClient.GetLatestBlockNumber(i.ctx)
				if err != nil {
					return fmt.Errorf("error getting latest indexed block number: %w", err)
				}
				if l2block.Uint64() <= latestIndexedBlockNumber {
					break
				}
				log.Debug().Msgf("🚧 Waiting for l2 block %d to be indexed, so we can write batch info", l2block)
			}
		}
	}

	log.Info().Msgf("🟢 All batch events can be written to the database")
	return nil
}

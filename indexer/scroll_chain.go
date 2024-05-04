package indexer

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (i *Indexer) StartL1IndexerForScroll(worker int, forceStartBlock *uint64) {
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

	const eventStep = 50
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

			batchNumber, batchesLen, err := i.indexL1BatchesForScroll(startBlock, endBlock)
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
				// log.Info().Msgf("🔍 [L1 Indexer] No new batches for range: [%d, %d], skipping", startBlock, endBlock)
				startBlock = endBlock + 1
				continue
			}

			log.Info().Msgf("🟢 [L1 Indexer] Indexed batch num: %d, range: [%d, %d], len: %d", batchNumber, startBlock, endBlock, batchesLen)

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

func (i *Indexer) WaitingWriteBatchForScroll(batchEvents []*ethclient.ScrollBatchEvent) error {
	latestIndexedBlockNumber, err := i.dbClient.GetLatestBlockNumber(i.ctx)
	if err != nil {
		return fmt.Errorf("error getting latest indexed block number: %w", err)
	}

	for _, event := range batchEvents {
		for _, l2block := range event.CommittedBlocks {
			if l2block <= latestIndexedBlockNumber {
				continue
			}
			for {
				time.Sleep(10 * time.Second)
				latestIndexedBlockNumber, err = i.dbClient.GetLatestBlockNumber(i.ctx)
				if err != nil {
					return fmt.Errorf("error getting latest indexed block number: %w", err)
				}
				if l2block <= latestIndexedBlockNumber {
					break
				}
				log.Debug().Msgf("🚧 Waiting for l2 block %d to be indexed, so we can write batch info", l2block)
			}
		}
	}

	log.Info().Msgf("🟢 All batch events can be written to the database")
	return nil
}

func (i *Indexer) indexL1BatchesForScroll(from uint64, to uint64) (uint64, int, error) {
	batchEvents, err := i.l1client.GetScrollBatches(i.ctx, i.chain.RollupContractAddress, from, to)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting L1 batch events: %w", err)
	}

	if len(batchEvents) == 0 {
		return 0, 0, nil
	}

	// check if we need to wait for the block to be finalized
	if err := i.WaitingWriteBatchForScroll(batchEvents); err != nil {
		return 0, 0, fmt.Errorf("error waiting for block to be finalized: %w", err)
	}

	batchNumber, err := i.dbClient.UpsertL1BatchForScrollAll(i.ctx, batchEvents)
	if err != nil {
		return 0, 0, fmt.Errorf("error upserting L1 batch events: %w", err)
	}

	err = i.dbClient.UpsertL1SyncProgress(i.ctx, to)
	if err != nil {
		return 0, 0, fmt.Errorf("error upserting L1 sync progress: %w", err)
	}

	return *batchNumber, len(batchEvents), nil
}

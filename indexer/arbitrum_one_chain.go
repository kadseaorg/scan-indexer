package indexer

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/unifra20/l2scan-indexer/ethclient"

	"github.com/rs/zerolog/log"
)

const (
	NitroGensisBlockNum         = 22207817
	L1BlockNumberForNitroGensis = 15447728
)

func (i *Indexer) StartL1IndexerForArbitrumOne(worker int, forceStartBlock *uint64) {
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

	if i.chain == ArbitrumOne {
		if lastSyncedBlockNumber <= L1BlockNumberForNitroGensis {
			lastSyncedBlockNumber = L1BlockNumberForNitroGensis - 1
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

			batchNumber, batchesLen, err := i.indexL1BatchesForArbitrumOne(startBlock, endBlock)
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
			if batchesLen == 0 && endBlock < latestBlockNumber {
				log.Info().Msgf("🔍 [L1 Indexer] No new batches for range: [%d, %d], skipping", startBlock, endBlock)
				startBlock = endBlock + 1
				time.Sleep(time.Millisecond * 500)
				continue
			}

			log.Info().Msgf("🟢 [L1 Indexer] Indexed batch num: %d, range: [%d, %d], len: %d", batchNumber, startBlock, endBlock, batchesLen)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 Indexer] Have indexed all batches up to block %d", latestBlockNumber)
				time.Sleep(time.Second)
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

func (i *Indexer) indexL1BatchesForArbitrumOne(from uint64, to uint64) (uint64, int, error) {
	batchEvents, err := i.l1client.GetArbOneSequencerBatchDeliveredEventLogs(i.ctx, from, to)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting L1 batch events: %w", err)
	}

	if len(batchEvents) == 0 {
		return 0, 0, nil
	}

	batchNumber, err := i.dbClient.UpsertL1BatchForArbOneAll(i.ctx, batchEvents)
	if err != nil {
		return 0, 0, fmt.Errorf("error upserting L1 batch events: %w", err)
	}

	err = i.dbClient.UpsertL1SyncProgress(i.ctx, to)
	if err != nil {
		return 0, 0, fmt.Errorf("error upserting L1 sync progress: %w", err)
	}

	return *batchNumber, len(batchEvents), nil
}

func (i *Indexer) indexL1BatchesForArbitrumOneForRecovery(from uint64, to uint64) {
	batchEvents, err := i.l1client.GetArbOneSequencerBatchDeliveredEventLogs(i.ctx, from, to)
	if err != nil {
		log.Error().Msgf("error getting Arb L1 batch events: %s", err)
		return
	}

	if len(batchEvents) == 0 {
		return
	}

	batchNumber, err := i.dbClient.UpsertL1BatchForArbOneAllForRecovery(i.ctx, batchEvents)
	if err != nil {
		log.Error().Msgf("error upserting Arb L1 batch events: %v", err)
		return
	}

	err = i.dbClient.UpsertRecoveryL1BlockHeightProgress(i.ctx, to)
	if err != nil {
		log.Error().Msgf("error upserting recovery L1 block height progress: %v", err)
		return
	}

	log.Info().Msgf("👻 [L1 Recovery] Recovery Arb missing batches %d", *batchNumber)
}

func (i *Indexer) AddBatchNumberToL2blocksArbOne(blockNumber *big.Int) error {
	if i.chain != ArbitrumOne {
		return nil
	}

	if blockNumber.Uint64() <= NitroGensisBlockNum {
		return nil
	}

	log.Info().Msgf("🔵 [Add batch info] Start to add batch info for block: %d", blockNumber)

	var batchNumber uint64
	var latestSyncedBatchNumber uint64
	var err error

	// Get the batch number which contains current block number.
	// try several times and bail out if keeping fail.
	for n := 0; n < 10; n++ {
		batchNumber, err = i.l2client.GetArbOneBatchContainingBlock(i.ctx, blockNumber.Uint64())
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	if err != nil {
		log.Error().Msgf("Error querying batch info for L2 block %d - batchNumber %d: %s", blockNumber.Uint64(), batchNumber, err)
		return err
	}

	// Get the latest batch number in our local DB, compare it with batch number containing current block.
	// This is to ensure our batch number is indexed and saved in our local DB.
	for n := 0; n < 10; n++ {
		latestSyncedBatchNumber, err = i.dbClient.GetLatestBatchNumber(i.ctx)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if batchNumber > latestSyncedBatchNumber {
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}

	if err != nil {
		log.Error().Msgf("Error getting the latest batch number for L2 block %d: %s", blockNumber.Uint64(), err)
		return err
	}

	// The current L2 block is indexed, the tx in it are indexed, its batch number is quried,
	// the batch itself is indexed as well.
	// Now let us update our DB to add the batch info for blocks and transactions.
	err = i.dbClient.UpdateBatchInfoForArbOne(i.ctx, blockNumber.Uint64(), batchNumber)
	if err != nil {
		log.Error().Msgf("Error updating batch info for L2 block %d: %s", blockNumber.Uint64(), err)
		return err
	}

	log.Info().Msgf("🌈 [Add batch info] Finished to add batch[%d] info for L2 block[%d]", batchNumber, blockNumber.Uint64())
	return nil
}

// TODO: our RPC server doesn't support this one yet.
func (i *Indexer) indexArbOneClassicInternalTransaction(tx *ethclient.RpcTransaction, block *ethclient.RpcBlock) error {
	return nil
}

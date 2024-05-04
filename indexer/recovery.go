package indexer

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/unifra20/l2scan-indexer/ethclient"

	"github.com/alitto/pond"
	"github.com/rs/zerolog/log"
)

const (
	RetryLimit        = 20
	ErrRetryGap       = time.Second * 10
	CycleGap          = time.Minute * 10
	SlowQueryRetryGap = time.Hour * 1
	RecoveryLimit     = 10000
)

func (i *Indexer) L2RecoveryExecutor(worker int, checkMisMatchedBlocks bool) {
	var isFirstRun bool = true

	// Define the runner function that encapsulates the logic to be executed periodically.
	runner := func() {
		var missingBlocks []uint64
		var err error
		limit := RecoveryLimit

		// If it's the first run, set limit to 0
		if isFirstRun {
			limit = 0
			isFirstRun = false
		}

		if checkMisMatchedBlocks {
			missingBlocks, err = i.dbClient.GetMismatchedTransactionBlocks(i.ctx, limit)
		} else {
			missingBlocks, err = i.dbClient.GetMissingBlocks(i.ctx, limit)
		}

		if err != nil {
			log.Error().Err(err).Msg("Error getting missing blocks from db(Recovery)")
			return
		}

		log.Info().Msgf("🔄 RecoveryExecutor: missingBlocksLen: %d", len(missingBlocks))

		i.RecoverMissingBlocks(worker, checkMisMatchedBlocks, missingBlocks)
	}

	// Call forever with the context, interval, and the runner function.
	forever(i.ctx, time.Minute*1, runner)
}

func (i *Indexer) RecoverMissingBlocks(worker int, checkMisMatchedBlocks bool, missingBlocks []uint64) {
	log.Info().Msgf("[🔄 RecoveryExecutor] recovering missing blocks (missingBlocksLen: %d) (checkMisMatchedBlocks: %v)",
		len(missingBlocks), checkMisMatchedBlocks)

	missingBlocksLen := len(missingBlocks)
	pool := pond.New(worker, 2*worker)
	for idx, block := range missingBlocks {
		block := block
		idx := idx
		pool.Submit(func() {
			log.Info().Msgf("🔄 Recovering block %d (%d/%d)", block, idx+1, missingBlocksLen)
			i.handleBlock(big.NewInt(int64(block)))

		})
	}
	pool.StopAndWait()

	log.Info().Msgf("[🔄 RecoveryExecutor] finished recovering missing blocks (missingBlocksLen: %d) (checkMisMatchedBlocks: %v)", len(missingBlocks), checkMisMatchedBlocks)
}

func (i *Indexer) RecoverZKSyncMissingBatches(worker int) {
	if i.chain.Name != ZkSyncEra.Name {
		return
	}

	log.Info().Msg("🔄 Recovering ZKS missing batches")

	startBatch := uint64(1)
	endBatch, err := i.l2client.GetZKSL1BatchNumber(i.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting latest batch number")
		return
	}

	// get missing batches from db
	log.Info().Uint64("start_batch_number", startBatch).Uint64("end_batch_number", endBatch).Msg("🔄 Getting missing batches")
	missingBatches, err := i.dbClient.GetMissingBatches(i.ctx, startBatch, endBatch)
	if err != nil {
		log.Error().Err(err).Msg("Error getting missing batches from db")
		return
	}

	batchLen := len(missingBatches)
	log.Info().Msgf("🔄 Recovering %d missing batches", batchLen)
	pool := pond.New(worker, 2*worker)
	for idx, batch := range missingBatches {
		batch := batch
		pool.Submit(func() {
			log.Info().Msgf("🔄 Recovering batch %d (%d/%d)", batch, idx+1, batchLen)
			time.Sleep(time.Millisecond * 300)
			i.indexZKSEraL1Batch(batch)
		})
	}
	pool.StopAndWait()

	log.Info().Msg("🔄 Finished recovering missing batches")
}

func (i *Indexer) RecoverL1BatchesForArbitrumOne(worker int, recoveryStartBlock *uint64) {
	if i.chain != ArbitrumOne {
		return
	}

	log.Info().Msg("🔄 Recovering Arb missing batches")
	var lastSyncedBlockNumber uint64
	var err error

	if recoveryStartBlock != nil {
		log.Warn().Msgf(" starting Recovery L1 batcher from block %d", *recoveryStartBlock)
		lastSyncedBlockNumber = *recoveryStartBlock
	} else {
		lastSyncedBlockNumber, err = i.dbClient.GetRecoveryL1BlockHeightProgress(i.ctx)
		if err != nil {
			log.Error().Msg("🔄 Err Get Recovery L1 Block Height Progress")
			return
		}
	}

	if lastSyncedBlockNumber <= L1BlockNumberForNitroGensis {
		lastSyncedBlockNumber = L1BlockNumberForNitroGensis - 1
	}

	//Check every 24 hours
	const eventStep = 50
	for j := 0; j < 10000; j++ {
		log.Info().Msgf("🔄Recovering Arb missing batches, Run round %d ", j)
		if j > 0 {
			log.Info().Msg("🔄 Continue to Recovering the missing batches of the arb after 24 hours")
			time.Sleep(time.Hour * 24)
		}

		for {
			startBlock := lastSyncedBlockNumber + 1
			latestBlockNumber, err := i.l1client.BlockNumber(i.ctx)
			if err != nil {
				// exit and let the process restart
				log.Error().Msgf("Error getting latest block number: %s", err)
				os.Exit(-1)
			}
			// reset startBlock if it's greater than latestBlockNumber
			if startBlock > latestBlockNumber {
				startBlock = latestBlockNumber
			}
			endBlock := startBlock + eventStep
			if endBlock > latestBlockNumber {
				endBlock = latestBlockNumber
			}

			i.indexL1BatchesForArbitrumOneForRecovery(startBlock, endBlock)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 Recovery] Have indexed all batches up to block %d", latestBlockNumber)
				lastSyncedBlockNumber = endBlock
				break
			}

			lastSyncedBlockNumber = endBlock
		}

		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
		lastSyncedBlockNumber, err = i.dbClient.GetRecoveryL1BlockHeightProgress(i.ctx)
		if err != nil {
			log.Error().Msg("🔄 Err Get Recovery L1 Block Height Progress")
			return
		}
	}
}

// okx1
func (i *Indexer) StartL1RecoveryForOKX1Sepolia(worker int) {
	log.Info().Msg("🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢 Starting L1 Recovery 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢")
	var lastSyncedBlockNumber uint64
	var err error

	lastSyncedBlockNumber, err = i.dbClient.GetRecoveryL1BlockHeightProgress(i.ctx)
	if err != nil {
		log.Error().Msgf("Error getting recovery L1 block number: %s", err)
		return
	}

	var eventStep uint64
	switch i.chain.Name {
	case BsquaredTestnet.Name:
		eventStep = 1000

	case PolygonZkEVM.Name:
		eventStep = 100

	case OKX1Sepolia.Name:
		if lastSyncedBlockNumber <= ethclient.OKX1L1InitialBatchHeight {
			lastSyncedBlockNumber = ethclient.OKX1L1InitialBatchHeight - 1
		}
		eventStep = 2000

	default:
		//eventStep = 50
		return
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

			batchNumbers, batchesLen, err := i.recoveryL1BatchesForOKX1Sepolia(startBlock, endBlock)
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
					// has indexed, skip this block range
					log.Info().Msgf("🤷 [L1 recovery] Already indexed range: [%d, %d]", startBlock, endBlock)
					startBlock = endBlock + 1
					continue
				}
				log.Error().Msgf("Error recovery L1 batches for range: [%d, %d]: %s", startBlock, endBlock, err)
				break
			}
			if batchesLen == 0 {
				log.Info().Msgf("🔍 [L1 recovery] No new batches for range: [%d, %d], skipping", startBlock, endBlock)
				startBlock = endBlock + 1
				continue
			}

			log.Info().Msgf("🟢 [L1 recovery] Indexed batch num: %d, range: [%d, %d], len: %d", batchNumbers, startBlock, endBlock, batchesLen)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 recovery] Have indexed all batches up to block %d", latestBlockNumber)
				break
			}

			startBlock = endBlock + 1
		}
		lastSyncedBlockNumber, err = i.dbClient.GetRecoveryL1BlockHeightProgress(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 block number in loop: %s", err)
		}

		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
	}
}

func (i *Indexer) recoveryL1BatchesForOKX1Sepolia(from uint64, to uint64) ([]uint64, int, error) {
	batchEvents, err := i.l1client.GetOKX1SepoliaSequenceBatchesEventLogs(i.ctx, i.chain.RollupContractAddress, from, to)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error getting recovery L1 batch events: %w", err)
	}

	if len(batchEvents) == 0 {
		return []uint64{0}, 0, nil
	}

	L2InitBatch, err := i.dbClient.GetLastRecoverybatchNumber(i.ctx)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error getting recovery L2 Init batch information: %w", err)
	}

	newBatchEvents := make([]*ethclient.OKX1SepoliaBatchEvent, 0)
	for _, batch := range batchEvents {
		for batchNum := L2InitBatch + 1; batchNum <= batch.L2FinalBatch; batchNum++ {
			newBatchEvent := &ethclient.OKX1SepoliaBatchEvent{
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
			return []uint64{0}, 0, fmt.Errorf("error getting recovery L2 batch information: %w", err)
		}
		for _, blockHash := range batchInfo.Blocks {
			blockInfo, err := i.l2client.GetBlockByHash(i.ctx, blockHash)
			if err != nil {
				return []uint64{0}, 0, fmt.Errorf("error getting recovery L2 block information: %w", err)
			}
			batch.L2BatchBlocks = append(batch.L2BatchBlocks, blockInfo.Number.ToInt())
		}
	}

	// check if we need to wait for the block to be finalized
	if err := i.WaitingWriteBatchForOKX1Sepolia(newBatchEvents); err != nil {
		return []uint64{0}, 0, fmt.Errorf("error waiting for block to be finalized: %w", err)
	}

	batchNumbers, err := i.dbClient.UpsertL1BatchForOKX1SepoliaAll(i.ctx, newBatchEvents)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error upserting recovery L1 batch events: %w", err)
	}

	err = i.dbClient.UpsertLastRecoverybatchNumber(i.ctx, batchEvents[len(batchEvents)-1].L2FinalBatch)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error upserting recovery L1 sync progress: %w", err)
	}

	err = i.dbClient.UpsertRecoveryL1BlockHeightProgress(i.ctx, to)
	if err != nil {
		return []uint64{0}, 0, fmt.Errorf("error upserting recovery L1 sync progress: %w", err)
	}

	return *batchNumbers, len(newBatchEvents), nil
}

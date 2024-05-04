package indexer

import (
	"math/big"
	"strings"
	"time"

	"github.com/unifra20/l2scan-indexer/ethclient"

	"github.com/alitto/pond"
	"github.com/rs/zerolog/log"
)

const scheduleInterval = 10 * time.Minute
const unfinalizedScheduleInterval = 1 * time.Minute

// CheckAndUpdateZKSL1Status checks if unfinalized L1 blocks are finalized and updates the L1 status
func (i *Indexer) CheckAndUpdateZKSL1Status(worker int) {
	log.Info().Msg("🎯 Checking and updating ZKS L1 status")

	// get unfinalized L1 batches from db
	log.Info().Msg("🎯 Getting unfinalized ZKS L1 batches")
	unfinalizedL1Batches, err := i.dbClient.GetUnfinalizedL1Batches(i.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting unfinalized ZKS L1 batches from db")
		return
	}

	log.Info().Msgf("🎯 Checking %d unfinalized ZKS L1 batches", len(unfinalizedL1Batches))

	pool := pond.New(worker, 2*worker)
	for _, batch := range unfinalizedL1Batches {
		batch := batch
		pool.Submit(func() {
			log.Info().Msgf("🎯 Checking ZKS L1 batch %d", batch)
			time.Sleep(time.Millisecond * 300)
			err := i.indexZKSEraL1Batch(batch)
			if err != nil {
				log.Error().Err(err).Msgf("Error checking ZKS L1 batch %d", batch)
				return
			}
		})
	}
	pool.StopAndWait()

	log.Info().Msgf("🎯 Finished checking and updating %d ZKS L1 batches", len(unfinalizedL1Batches))
}

func (i *Indexer) CheckAndUpdateUnfinalizedL2Blocks(worker int) {
	log.Info().Msgf("🔄 Checking and updating ZKS unfinalized L2 blocks")

	unfinalizedL2Blocks, err := i.dbClient.GetUnfinalizedL2Blocks(i.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting unfinalized blocks from db")
		return
	}

	unfinalizedL2BlocksLen := len(unfinalizedL2Blocks)
	log.Info().Msgf("🔄 Checking and updating ZKS %d unfinalized blocks", unfinalizedL2BlocksLen)
	pool := pond.New(worker, 2*worker)
	for idx, block := range unfinalizedL2Blocks {
		block := block
		idx := idx
		pool.Submit(func() {
			log.Info().Msgf("🔄 Updating ZKS block %d (%d/%d)", block, idx+1, unfinalizedL2BlocksLen)
			i.handleBlock(big.NewInt(int64(block)))
		})
	}
	pool.StopAndWait()

	log.Info().Msg("🔄 Finished checking and updating ZKS missing blocks")
}

// ScheduleZKSL1StatusUpdate schedules a job to call CheckAndUpdateZKSL1Status
func (i *Indexer) ScheduleZKSL1StatusUpdate(worker int) {
	if i.chain.Name != ZkSyncEra.Name && i.chain.Name != ZkSyncEraSepolia.Name {
		return
	}

	log.Info().Msgf("🎯 Scheduling ZKS L1 status update every %s", scheduleInterval)

	ticker := time.NewTicker(scheduleInterval)

	go func() {
		for range ticker.C {
			i.CheckAndUpdateZKSL1Status(worker)
		}
	}()

	tickerUnfinalized := time.NewTicker(unfinalizedScheduleInterval)

	go func() {
		for range tickerUnfinalized.C {
			i.CheckAndUpdateUnfinalizedL2Blocks(worker)
		}
	}()
}

func (i *Indexer) ScheduleArbL1StatusUpdate(worker int, forceStartBlock *uint64) {
	if i.chain.Name != ArbitrumOne.Name {
		return
	}

	log.Info().Msgf("🎯 Scheduling Arb L2 status update every %s", scheduleInterval)

	i.CheckAndUpdateArbL2UnfinalizedBlock(worker / 2)
	i.CheckAndUpdateArbL2MissingFinalizedBlock(worker)

	tickerMissing := time.NewTicker(time.Minute * 10)
	tickerUnfinalized := time.NewTicker(time.Minute * 4)

	go func() {
		for {
			select {
			case <-tickerMissing.C:
				i.CheckAndUpdateArbL2MissingFinalizedBlock(worker / 2)

			case <-tickerUnfinalized.C:
				i.CheckAndUpdateArbL2UnfinalizedBlock(worker / 2)
			}
		}
	}()
}

func (i *Indexer) CheckAndUpdateArbL2UnfinalizedBlock(worker int) {
	if i.latestL2BlockNumber < NitroGensisBlockNum {
		return
	}

	startBlock := i.latestL2BlockNumber - 2000
	endBlock := i.latestL2BlockNumber
	log.Info().Msgf("🎯 Checking and updating Arb L2 unfinalized block [%d - %d]", startBlock, endBlock)

	if startBlock <= NitroGensisBlockNum {
		startBlock = NitroGensisBlockNum - 1
	}

	unfinalizedL2Blocks, err := i.dbClient.SelectNumbersForNullBatches(i.ctx, startBlock, endBlock)
	if err != nil {
		log.Error().Err(err).Msg("Error getting unfinalized Arb L2 blocks from db")
		return
	}

	unfinalizedL2BlocksLen := len(unfinalizedL2Blocks)
	log.Info().Msgf("🔄 Checking and updating Arb len[%d] unfinalized blocks [%d - %d]", unfinalizedL2BlocksLen, startBlock, endBlock)
	pool := pond.New(worker, 2*worker)
	for idx, block := range unfinalizedL2Blocks {
		block := block
		idx := idx
		pool.Submit(func() {
			log.Info().Msgf("🔄 Updating Arb unfinalized block %d (%d/%d)", block, idx+1, unfinalizedL2BlocksLen)
			err := i.AddBatchNumberToL2blocksArbOne(big.NewInt(int64(block)))
			if err != nil {
				log.Error().Msgf("🔄 Error updating Arb unfinalized block %d :%s", block, err)
				return
			}
		})
	}
	pool.StopAndWait()

	log.Info().Msg("🔄 Finished checking and updating Arb unfinalized blocks")
}

func (i *Indexer) CheckAndUpdateArbL2MissingFinalizedBlock(worker int) {
	log.Info().Msg("🎯 Checking and updating Arb L2 missing finalized block")

	// get unfinalized L1 batches from db
	startBlock, err := i.dbClient.GetRecoveryL1L2MappingStatusProgress(i.ctx)
	if err != nil {
		log.Error().Msg("🔄 Err Get Recovery l1 l2 mapping Status Progress")
		return
	}

	if startBlock <= NitroGensisBlockNum {
		startBlock = NitroGensisBlockNum - 1
	}

	for {
		endBlock := startBlock + 5000
		if i.latestL2BlockNumber <= endBlock {
			endBlock = i.latestL2BlockNumber
		}

		unfinalizedL2Blocks, err := i.dbClient.SelectNumbersForNullBatches(i.ctx, startBlock, endBlock)
		if err != nil {
			log.Error().Err(err).Msg("Error getting missing finalized Arb L2 blocks from db")
			return
		}

		unfinalizedL2BlocksLen := len(unfinalizedL2Blocks)
		log.Info().Msgf("🔄 Checking and updating Arb len[%d] missing finalized blocks [%d - %d] - %d", unfinalizedL2BlocksLen, startBlock, endBlock, i.latestL2BlockNumber)
		pool := pond.New(worker, 2*worker)
		for idx, block := range unfinalizedL2Blocks {
			block := block
			idx := idx
			pool.Submit(func() {
				log.Info().Msgf("🔄 Updating Arb missing finalized block %d (%d/%d)", block, idx+1, unfinalizedL2BlocksLen)
				err := i.AddBatchNumberToL2blocksArbOne(big.NewInt(int64(block)))
				if err == nil {
					err = i.dbClient.UpsertRecoveryL1L2MappingStatusProgress(i.ctx, block)
					if err != nil {
						log.Error().Msgf("Error Upsert recovery l1 l2 mapping status progress, block %d: %s", block, err)
						return
					}
				}
			})
		}
		pool.StopAndWait()

		if i.latestL2BlockNumber == endBlock {
			for {
				if i.latestL2BlockNumber > endBlock {
					break
				} else {
					time.Sleep(time.Millisecond * 1000)
				}
			}
		}
		startBlock = endBlock
		time.Sleep(time.Millisecond * 1000)
	}

	log.Info().Msg("🔄 Finished checking and updating Arb missing finalized blocks")
}

func (i *Indexer) CheckAndUpdateLineaUnfinalizedL2Blocks() {
	if i.chain.Name != Linea.Name {
		return
	}

	log.Info().Msgf("🔄 Checking and updating Linea unfinalized L2 blocks")

	latestL1BlockNumber, err := i.l1client.BlockNumber(i.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting l1 blocks")
		return
	}

	recheckL1LastFinalizedBlock, err := i.dbClient.GetRecheckLastFinalizedBlock(i.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting l1 blocks")
		return
	}

	recheckLastL2BlockNumber := uint64(ethclient.BlockFinalizeUpgradeBlockHeightOnL2)

	log.Info().Msgf("🔄 Checking and updating Linea unfinalized blocks srart l1 %d l2 %d latestL1BlockNumber %d",
		recheckL1LastFinalizedBlock, recheckLastL2BlockNumber, latestL1BlockNumber)

	startBlock := recheckL1LastFinalizedBlock
	for {
		endBlock := startBlock + 50

		if startBlock >= latestL1BlockNumber {
			startBlock = latestL1BlockNumber
		}

		if endBlock >= latestL1BlockNumber {
			endBlock = latestL1BlockNumber
		}

		batchNumber, batchesLen, err := i.indexL1BatchesForLinea(&recheckLastL2BlockNumber, startBlock, endBlock)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				// has indexed, skip this block range
				log.Info().Msgf("🤷 [L1 Indexer] Checking already indexed range: [%d, %d]", startBlock, endBlock)
				startBlock = endBlock + 1
				continue
			}
			log.Error().Msgf("Error Checking indexing L1 batches for range: [%d, %d]: %s", startBlock, endBlock, err)
			break
		}

		if batchesLen == 0 && endBlock < latestL1BlockNumber {
			log.Info().Msgf("🔍 [L1 Indexer] Linea no new batches for range: [%d, %d], skipping", startBlock, endBlock)
			startBlock = endBlock + 1
			continue
		}

		if endBlock >= latestL1BlockNumber {
			err = i.dbClient.UpsertRecheckLastFinalizedBlock(i.ctx, endBlock)
			if err != nil {
				log.Error().Msgf("Error upsert recheck last finalized block range: %d: %s", recheckLastL2BlockNumber, err)
				continue
			}

			log.Info().Msgf("👻 [L1 Indexer] Have Checking indexed all batches up to block %d", latestL1BlockNumber)
			break
		}

		log.Info().Msgf("🟢 [L1 Indexer] Checking indexed batch num: %d, range: [%d, %d], len: %d", batchNumber, startBlock, endBlock, batchesLen)

		err = i.dbClient.UpsertRecheckLastFinalizedBlock(i.ctx, endBlock)
		if err != nil {
			log.Error().Msgf("Error upsert recheck last finalized block range: %d: %s", recheckLastL2BlockNumber, err)
			continue
		}
		startBlock = endBlock + 1
	}

	log.Info().Msg("🔄 Finished checking and updating Linea missing blocks")
}

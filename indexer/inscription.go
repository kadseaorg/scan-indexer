package indexer

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func (i *Indexer) StartInscriptionIndexer(worker int, forceStartBlock *uint64) {
	if i.chain != ZkSyncEra && i.chain != Scroll && i.chain != MantaPacific && i.chain != Base && i.chain != Linea {
		return
	}

	log.Info().Msg("🪽🪽🪽🪽🪽🪽🪽🪽🪽🪽 Starting Inscription Indexer 🪽🪽🪽🪽🪽🪽🪽🪽🪽🪽")
	var lastSyncedBlockNumber uint64
	var err error

	if forceStartBlock != nil {
		log.Warn().Msgf("Force starting inscription indexer from block %d", *forceStartBlock)
		lastSyncedBlockNumber = *forceStartBlock
	} else {
		lastSyncedBlockNumber, err = i.dbClient.GetLatestSyncedInscriptionBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last synced inscription block number: %s", err)
			return
		}
	}

	const extractStep = 10000
	for {
		startBlock := lastSyncedBlockNumber + 1
		for {
			latestBlockNumber, err := i.dbClient.GetLatestBlockNumber(i.ctx)
			if err != nil {
				log.Error().Msgf("Error getting latest block number: %s", err)
				os.Exit(-1)
			}

			// reset startBlock if it's greater than latestBlockNumber
			if startBlock > latestBlockNumber {
				// we should wait for 3 mins to ensure that the latest block transactions are indexed into the database
				time.Sleep(3 * time.Minute)
				startBlock = latestBlockNumber
			}
			endBlock := startBlock + extractStep
			if endBlock > latestBlockNumber {
				endBlock = latestBlockNumber
			}

			inscriptionsLen, err := i.IndexInscriptions(startBlock, endBlock)
			if err != nil {
				log.Error().Msgf("Error extracting inscriptions: %s", err)
				break
			}

			if inscriptionsLen == 0 {
				log.Info().Msgf("🪽 [Inscription Indexer] No inscriptions found from block %d to %d", startBlock, endBlock)

				startBlock = endBlock + 1
				time.Sleep(5 * time.Second)
				continue
			}

			log.Info().Msgf("🪽 [Inscription Indexer] Extracted inscriptions from block %d to %d, len: %d", startBlock, endBlock, inscriptionsLen)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [Inscription Indexer] Have extracted all inscriptions up to block %d", latestBlockNumber)
				break
			}

			startBlock = endBlock + 1
		}
		lastSyncedBlockNumber, err = i.dbClient.GetLatestSyncedInscriptionBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last synced inscription block number: %s", err)
		}
		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
	}
}

func (i *Indexer) IndexInscriptions(startBlock uint64, endBlock uint64) (int, error) {
	nums, err := i.dbClient.ExtractInscriptions(startBlock, endBlock)
	if err != nil {
		return 0, err
	}
	return nums, i.dbClient.UpsertLastetSyncedInscriptionBlockNumber(i.ctx, endBlock)
}

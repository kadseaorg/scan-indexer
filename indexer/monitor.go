package indexer

import (
	"github.com/alitto/pond"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

func (i *Indexer) L2MonitorExecutor(worker int) {
	var (
		lastCacheBlockNumber   uint64
		lastIndexedBlockNumber uint64
		err                    error
	)

	lastCacheBlockNumber, err = i.dbClient.GetLatestCache(i.ctx)
	if err != nil {
		log.Error().Msgf("Error getting last cache block number: %s", err)
		return
	}

	pool := pond.New(worker, worker*2)

	for {
		lastIndexedBlockNumber, err = i.dbClient.GetLatestBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last indexed block number: %s", err)
			return
		}

		log.Info().Msgf("L2 Monitor: lastIndexedBlockNumber %d latestCacheBlockNumber %d", lastIndexedBlockNumber, lastCacheBlockNumber)

		if lastIndexedBlockNumber > lastCacheBlockNumber {
			for n := lastCacheBlockNumber; n < lastIndexedBlockNumber; n++ {
				blockNumber := big.NewInt(int64(n)) // Create a new variable to avoid data race
				if ok, err := i.dbClient.HasBlocksByBlockNumber(i.ctx, blockNumber); err != nil {
					log.Error().Msgf("Error checking block %d: %s", blockNumber, err)
					return
				} else if !ok {
					pool.Submit(func() {
						i.handleBlock(blockNumber)
					})
				}
			}

			lastCacheBlockNumber = lastIndexedBlockNumber

			if err := i.dbClient.UpdateCache(i.ctx, big.NewInt(int64(lastCacheBlockNumber))); err != nil {
				log.Error().Msgf("Error UpdateCache block number: %s", err)
				return
			}
		}

		time.Sleep(3 * time.Second) // wait for 5 seconds
	}

}

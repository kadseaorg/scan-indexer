package indexer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/alitto/pond"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

const (
	TraceRecoveryInterval = 5 * time.Second
	TraceRecoveryLimit    = 10000
)

func (i *Indexer) RecoverL2MissingTraces(ctx context.Context, workers int) {
	log.Info().Msgf("🔄 Start watching trace recovery transactions every %s", TraceRecoveryInterval)

	var isFirstRun bool = true

	forever(ctx, TraceRecoveryInterval, func() {
		limit := TraceRecoveryLimit

		// If it's the first run, set limit to 0 to recover all missed traces
		if isFirstRun {
			limit = 0
			isFirstRun = false
		}

		if err := i.watch(ctx, workers, limit); err != nil {
			log.Error().Msgf("watch trace recovery transactions error: %s, retry in %s", err, TraceRecoveryInterval)
		}
	})
}

func (i *Indexer) watch(ctx context.Context, workers int, limit int) (err error) {
	missedBlockNumbers, err := i.dbClient.GetMissedTraceBlockNumbers(ctx, limit)
	if err != nil {
		err = fmt.Errorf("query new blcok number error: %w", err)
		return
	}

	if len(missedBlockNumbers) == 0 {
		return
	}

	log.Info().Msgf("🔄 Received %d missed trace block numbers", len(missedBlockNumbers))

	pool := pond.New(workers, 2*workers)
	lenMissedBlockNumbers := len(missedBlockNumbers)
	for index, blockNumber := range missedBlockNumbers {
		n := blockNumber
		pool.Submit(func() {
			start := time.Now()

			traceCount, err := i.indexBlockTrace(n)

			elapsedTime := time.Since(start).Milliseconds()

			if err != nil {
				log.Error().Msgf("Error indexing traced block %d (%d/%d) in %d ms: %s", n, index+1, lenMissedBlockNumbers, elapsedTime, err)
				return
			}

			log.Info().Msgf("✅ Indexed trace block %d (%d/%d) with %d internal transactions in %d ms", n, index+1, lenMissedBlockNumbers, traceCount, elapsedTime)

		})
	}

	return
}

func (i *Indexer) indexBlockTrace(height uint64) (internalTxnCount int, err error) {
	blockNumber := big.NewInt(int64(height))

	block, err := i.l2client.GetBlockByNumber(i.ctx, blockNumber)
	if err != nil {
		return
	}

	pgxBatch := &pgx.Batch{}
	for _, tx := range block.Transactions {
		if i.chain == ArbitrumOne && block.Number.ToInt().Uint64() <= 22207815 {
			err = i.indexArbOneClassicInternalTransaction(tx, block)
		} else {
			internalTxnCount, err = i.indexInternalTransaction(pgxBatch, tx, block, 0)
		}

		if err != nil {
			return 0, err
		}
	}

	i.dbClient.UpdateBlockTraceChecked(pgxBatch, blockNumber, internalTxnCount)

	// Execute the batch

	err = i.dbClient.ExecBatch(pgxBatch)
	if err != nil {
		// if bulk insert fails, try to rollback the block
		if err := i.dbClient.DeleteBlock(i.ctx, blockNumber.Uint64()); err != nil {
			return 0, fmt.Errorf("error rolling back block %d: %w", blockNumber, err)
		}
		return 0, err
	}

	return internalTxnCount, nil
}

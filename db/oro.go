package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (db *Client) UpsertL1BatchForOroAll(ctx context.Context, batches []*ethclient.BaseBatchEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error
	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForOro(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}
func (db *Client) UpsertL1BatchForOro(ctx context.Context, batch *ethclient.BaseBatchEvent) (*uint64, error) {
	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (number) DO UPDATE SET
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6,
		l2_block_number = $7
	`
	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err := db.pool.Exec(ctx, sqlStmt, batch.Number, batch.L1TxHash.Hex(), provenAt, batch.OutputRootHash.Hex(), batch.Timestamp, "finalized", batch.L2BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("insert batch failed: %v", err)
	}

	// Get the l2 block number of the current batch
	previousL2BlockNumber := 0
	currentL2BlockNumber := batch.L2BlockNumber

	// Get the l2 block number of the previous batch
	if batch.Number > 0 {
		err = db.pool.QueryRow(ctx, "SELECT l2_block_number FROM l1_batches WHERE number = $1 - 1", batch.Number).Scan(&previousL2BlockNumber)
		if err != nil {
			return nil, fmt.Errorf("QueryRow failed: %v", err)
		}
	}

	// Set l1_batch_number for all blocks in the range to the current batch number
	_, err = db.pool.Exec(ctx, "UPDATE blocks SET l1_batch_number = $1 WHERE number > $2 AND number <= $3", batch.Number, previousL2BlockNumber, currentL2BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("update blocks failed: %v", err)
	}

	// Set l1_batch_number for all transactions in the range to the current batch number
	_, err = db.pool.Exec(ctx, "UPDATE transactions SET l1_batch_number = $1 WHERE block_number > $2 AND block_number <= $3", batch.Number, previousL2BlockNumber, currentL2BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("update transactions failed: %v", err)
	}

	return &batch.Number, nil
}

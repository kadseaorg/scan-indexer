package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (db *Client) UpsertL1BatchForArbOneAll(ctx context.Context, batches []*ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error
	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForArbOne(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpsertL1BatchForArbOne(ctx context.Context, batch *ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number, l1_prove_block_number )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT (number) DO UPDATE SET
	    number = $1,
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6,
		l2_block_number = $7,
	    l1_prove_block_number = $8                
	`

	l2BlockNumber := big.NewInt(0)
	if len(batch.FinalizedBlockNumbers) > 0 {
		l2BlockNumber = batch.FinalizedBlockNumbers[len(batch.FinalizedBlockNumbers)-1]
	} else {
		l2BlockNumber = nil
	}

	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err := db.pool.Exec(ctx, sqlStmt,
		batch.Number,
		batch.L1TxHash.Hex(),
		provenAt,
		batch.OutputRootHash.Hex(),
		batch.Timestamp,
		"finalized",
		l2BlockNumber,
		batch.L1BlockNumber)

	if err != nil {
		return nil, fmt.Errorf("insert batch failed: %v", err)
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.FinalizedBlockNumbers), batch.Number)

	return &batch.Number, nil
}

func (db *Client) UpsertL1BatchForArbOneAllForRecovery(ctx context.Context, batches []*ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error

	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForArbOneIfNotExist(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpsertL1BatchForArbOneIfNotExist(ctx context.Context, batch *ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	sqlStmtCheck := `
		SELECT COUNT(*) FROM l1_batches WHERE number = $1
		`
	var count int
	err := db.pool.QueryRow(ctx, sqlStmtCheck, batch.Number).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error checking if batch with number %d exists: %v", batch.Number, err)
	}

	if count > 0 {
		log.Debug().Msgf("Batch with number %d already exists, no need to insert.", batch.Number)
		return &batch.Number, nil
	}

	l2BlockNumber := big.NewInt(0)
	if len(batch.FinalizedBlockNumbers) > 0 {
		l2BlockNumber = batch.FinalizedBlockNumbers[len(batch.FinalizedBlockNumbers)-1]
	} else {
		l2BlockNumber = nil
	}

	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number, l1_prove_block_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT (number) DO UPDATE SET
	    number = $1,
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6,
		l2_block_number = $7
	    l1_prove_block_number = $8
	`
	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err = db.pool.Exec(ctx, sqlStmt,
		batch.Number,
		batch.L1TxHash.Hex(),
		provenAt,
		batch.OutputRootHash.Hex(),
		batch.Timestamp,
		"finalized",
		l2BlockNumber,
		batch.L1BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("insert batch failed: %v", err)
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.FinalizedBlockNumbers), batch.Number)

	return &batch.Number, nil
}

func (db *Client) UpdateL1BatchForArbOneAllForRecovery(ctx context.Context, batches []*ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error

	for i, batch := range batches {
		batchNumber, err = db.UpdateL1BlockNumberForArbOne(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpdateL1BlockNumberForArbOne(ctx context.Context, batch *ethclient.ArbOneSequencerBatchDeliveredEvent) (*uint64, error) {
	sqlStmt := `
  UPDATE l1_batches
  SET l2_block_number = $1,
      l1_prove_block_number = $2
  WHERE number = $3
  `

	l2BlockNumber := big.NewInt(0)
	if len(batch.FinalizedBlockNumbers) > 0 {
		l2BlockNumber = batch.FinalizedBlockNumbers[len(batch.FinalizedBlockNumbers)-1]
	} else {
		l2BlockNumber = nil
	}

	_, err := db.pool.Exec(ctx, sqlStmt, l2BlockNumber, batch.L1BlockNumber, batch.Number)
	if err != nil {
		return nil, fmt.Errorf("update batch failed: %v", err)
	}

	log.Debug().Msgf("🆙 Updated l1_batches table for number %d", batch.Number)
	return &batch.Number, nil
}

func (db *Client) UpdateBatchInfoForArbOne(ctx context.Context, blockNum uint64, batchNum uint64) error {

	// Get l1_batch_timestamp from l1_batches table according to the batchNum
	var batchTimestamp sql.NullInt64
	sqlStmtCheck := `
	SELECT timestamp FROM l1_batches WHERE number = $1
	`
	err := db.pool.QueryRow(ctx, sqlStmtCheck, batchNum).Scan(&batchTimestamp)
	if err != nil {
		return fmt.Errorf("[batch] error querying timestamp for batch[%d]: %v", batchNum, err)
	}

	// Update batch info in batches table with batch number and batch timestamp
	sqlStmtUpdateBatch := `
	UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
	`
	_, err = db.pool.Exec(ctx, sqlStmtUpdateBatch, batchNum, batchTimestamp, blockNum)
	if err != nil {
		return fmt.Errorf("error updating batch[%d] info for block[%d]: %v", batchNum, blockNum, err)
	}

	// Update batch info in transactions table with batch number
	sqlStmtUpdateTX := `
	UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
	`
	_, err = db.pool.Exec(ctx, sqlStmtUpdateTX, batchNum, blockNum)
	if err != nil {
		return fmt.Errorf("error updating batch[%d] info for transactions: %v", batchNum, err)
	}

	return nil
}

func (db *Client) SelectNumbersForNullBatches(ctx context.Context, startNumber, endNumber uint64) ([]uint64, error) {
	var numbers []uint64

	sqlStmtSelect := `
    SELECT number 
    FROM public.blocks 
    WHERE l1_batch_number IS NULL AND number >= $1 AND number < $2
    `

	rows, err := db.pool.Query(ctx, sqlStmtSelect, startNumber, endNumber)
	if err != nil {
		return nil, fmt.Errorf("[batch] error querying numbers for blocks with null batch: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var blockNum uint64
		if err := rows.Scan(&blockNum); err != nil {
			return nil, fmt.Errorf("error scanning block number: %v", err)
		}
		numbers = append(numbers, blockNum)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %v", err)
	}

	return numbers, nil
}

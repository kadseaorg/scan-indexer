package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

type Client struct {
	pool  *pgxpool.Pool
	mutex sync.Mutex
}

//func NewClient(url string) (*Client, error) {
//	pool, err := pgxpool.New(context.Background(), url)
//	if err != nil {
//		log.Error().Msgf("Error connecting to Postgres database: %s", err)
//		return nil, err
//	}
//	return &Client{pool: pool, mutex: sync.Mutex{}}, nil
//}

func NewClient(url string) (*Client, error) {
	connConfig, err := pgxpool.ParseConfig(url)
	connConfig.MaxConns = 200
	connConfig.MinConns = 50
	if err != nil {
		log.Error().Msgf("Error parse config to Postgres database: %s", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), connConfig)
	if err != nil {
		log.Error().Msgf("Error new with config to Postgres database: %s", err)
		return nil, err
	}
	return &Client{pool: pool, mutex: sync.Mutex{}}, nil
}

func (db *Client) ExecBatch(batch *pgx.Batch) error {
	ctx := context.Background()
	results := db.pool.SendBatch(ctx, batch)
	defer results.Close()

	// Check if there are any errors in the batch
	for i := 0; i < batch.Len(); i++ {
		if _, err := results.Exec(); err != nil {
			return err
		}
	}

	return nil
}

func (db *Client) CreateFieldsL1BlockNumberForL1Batches() error {
	sql := `DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'l1_batches' AND column_name = 'l1_commit_block_number') THEN
        ALTER TABLE public.l1_batches ADD COLUMN l1_commit_block_number int8;
    END IF;

	IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'l1_batches' AND column_name = 'l1_execute_block_number') THEN
        ALTER TABLE public.l1_batches ADD COLUMN l1_execute_block_number int8;
    END IF;

	IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'l1_batches' AND column_name = 'l1_prove_block_number') THEN
        ALTER TABLE public.l1_batches ADD COLUMN l1_prove_block_number int8;
    END IF;
END $$;`

	ctx := context.Background()
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	_ = tx.Commit(ctx)
	return nil
}

func (db *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	var blockNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT MAX(number) FROM blocks").Scan(&blockNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		log.Error().Msgf("Error getting latest block number: %s", err)
		return 0, err
	}
	if !blockNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(blockNumber.Int64), nil
}

func (db *Client) GetLatestBatchNumber(ctx context.Context) (uint64, error) {
	var batchNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT MAX(number) FROM l1_batches").Scan(&batchNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	if !batchNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(batchNumber.Int64), nil
}

func (db *Client) GetLatestBatchOfLastL2FinalizedBlock(ctx context.Context) (uint64, error) {
	var batchNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT l2_block_number FROM l1_batches WHERE number = (SELECT MAX(number) FROM l1_batches)").Scan(&batchNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	if !batchNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(batchNumber.Int64), nil
}

func (db *Client) GetLatestBatchOfLastL1CommitBlock(ctx context.Context) (uint64, error) {
	var batchNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT l1_commit_block_number FROM l1_batches WHERE number = (SELECT MAX(number) FROM l1_batches)").Scan(&batchNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	if !batchNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(batchNumber.Int64), nil
}

func (db *Client) GetLatestBatchOfLastL1ExecuteBlock(ctx context.Context) (uint64, error) {
	var batchNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT l1_execute_block_number FROM l1_batches WHERE number = (SELECT MAX(number) FROM l1_batches)").Scan(&batchNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	if !batchNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(batchNumber.Int64), nil
}

func (db *Client) GetLatestBatchOfLastL1ProveBlock(ctx context.Context) (uint64, error) {
	var batchNumber sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT l1_prove_block_number FROM l1_batches WHERE number = (SELECT MAX(number) FROM l1_batches)").Scan(&batchNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	if !batchNumber.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}
	return uint64(batchNumber.Int64), nil
}

func (db *Client) GetLatestL1SyncedBlockNumber(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "l1_synced_block_number")
}

func (db *Client) GetRecoveryL1BlockHeightProgress(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l1_synced_block_number")
}

func (db *Client) GetRecoveryL1L2MappingStatusProgress(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l2_batch_block_number")
}

func (db *Client) GetRecoveryL2SyncProgress(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l2_synced_block_number")
}

func (db *Client) GetRecoveryL2SyncTarget(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l2_synced_target_block_number")
}

func (db *Client) GetRecheckLastFinalizedBlock(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recheck_l1_last_finalized_block")
}

func (db *Client) GetLastRecoveryInternalTransactionBlock(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l2_internal_transaction_block_number")
}

func (db *Client) GetLastRecoverybatchNumber(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, "recovery_l1_batch_number")
}

// l1_bridge_deposit_synced_block_number
// l1_bridge_withdraw_synced_block_number
func (db *Client) GetSyncProgress(ctx context.Context, key string) (uint64, error) {
	var value sql.NullInt64
	err := db.pool.QueryRow(ctx, "SELECT value FROM sync_progress WHERE key = $1", key).Scan(&value)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	if !value.Valid {
		// handle the case where the number is NULL in the DB
		return 0, nil
	}

	return uint64(value.Int64), nil
}

func (db *Client) UpsertBlock(pgxBatch *pgx.Batch, block *ethclient.RpcBlock, traceChecked bool, internalTransactionCount int) {
	sqlStmt := `
		INSERT INTO blocks
		(number, hash, transaction_count, validator, difficulty, total_difficulty, size, nonce, gas_used, gas_limit, extra_data, parent_hash, sha3_uncle, timestamp, l1_batch_number, l1_batch_timestamp, trace_checked, internal_transaction_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) ON CONFLICT(number) DO UPDATE SET
		hash = $2, transaction_count = $3, validator = $4, difficulty = $5, total_difficulty = $6, size = $7, nonce = $8, gas_used = $9, gas_limit = $10, extra_data = $11, parent_hash = $12, sha3_uncle = $13, timestamp = $14, l1_batch_number = $15, l1_batch_timestamp = $16, trace_checked = $17, internal_transaction_count = $18
	`
	params := []interface{}{
		block.Number.ToInt(),           // $1
		block.Hash.Hex(),               // $2
		len(block.Transactions),        // $3
		block.Miner.Hex(),              // $4
		block.Difficulty.ToInt(),       // $5
		block.TotalDifficulty.ToInt(),  // $6
		block.Size.ToInt(),             // $7
		block.Nonce.String(),           // $8
		block.GasUsed.ToInt(),          // $9
		block.GasLimit.ToInt(),         // $10
		block.ExtraData.String(),       // $11
		block.ParentHash.Hex(),         // $12
		block.Sha3Uncles.Hex(),         // $13
		block.Timestamp.ToInt(),        // $14
		block.L1BatchNumber.ToInt(),    // $15
		block.L1BatchTimestamp.ToInt(), // $16
		traceChecked,                   // $17
		internalTransactionCount,       // $18
	}

	pgxBatch.Queue(sqlStmt, params...)
}

func (db *Client) UpsertBlockInternalTransactionCount(ctx context.Context, block *big.Int, internalTransactionCount int) error {
	stmt := `SELECT internal_transaction_count FROM blocks WHERE number = $1`

	row := db.pool.QueryRow(ctx, stmt, block.String())

	var existingInternalTransactionCount *int
	err := row.Scan(&existingInternalTransactionCount)

	if err != nil {
		return err
	}

	if existingInternalTransactionCount != nil && *existingInternalTransactionCount > 0 {
		return nil
	}

	sqlStmt := `
             UPDATE blocks SET internal_transaction_count = $1 WHERE number = $2
             `
	_, err = db.pool.Exec(ctx, sqlStmt, internalTransactionCount, block)
	if err != nil {
		return err
	}

	return nil
}

func (db *Client) DeleteBlockInternalTransactionCount(ctx context.Context, block *big.Int) error {
	sqlStmt := `
             UPDATE blocks SET internal_transaction_count = $1 WHERE number = $2
             `
	_, err := db.pool.Exec(ctx, sqlStmt, nil, block)
	if err != nil {
		return err
	}

	return nil
}

func (db *Client) AddBlockInternalTransactionCount(ctx context.Context, block *big.Int, internalTransactionCount int) error {
	stmt := `SELECT internal_transaction_count FROM blocks WHERE number = $1`

	row := db.pool.QueryRow(ctx, stmt, block.String())

	var existingInternalTransactionCount *int
	err := row.Scan(&existingInternalTransactionCount)

	if err != nil {
		return err
	}

	var icr int
	if existingInternalTransactionCount == nil {
		icr = internalTransactionCount
	} else {
		icr = *existingInternalTransactionCount + internalTransactionCount
	}

	sqlStmt := `
             UPDATE blocks SET internal_transaction_count = $1 WHERE number = $2
             `
	_, err = db.pool.Exec(ctx, sqlStmt, icr, block)
	if err != nil {
		return err
	}

	return nil
}

func (db *Client) InsertTransaction(pgxBatch *pgx.Batch, block *ethclient.RpcBlock, tx *ethclient.RpcTransaction) {
	sqlStmt := `
        INSERT INTO transactions 
		(hash, block_hash, block_number, from_address, to_address, value, gas_price, gas_limit, method_id, input, nonce, transaction_type, l1_batch_number, l1_batch_tx_index, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT(hash) DO UPDATE SET 
		block_hash = $2, block_number = $3, from_address = $4, to_address = $5, value = $6, gas_price = $7, gas_limit = $8, method_id = $9, input = $10, nonce = $11, transaction_type = $12, l1_batch_number = $13, l1_batch_tx_index = $14, timestamp = $15
    `

	var methodId []byte
	if len(tx.Input) > 4 {
		methodId = tx.Input[:4]
	}
	var to *string
	if tx.To != nil {
		toStr := strings.ToLower(tx.To.Hex())
		to = &toStr
	}
	// Preparing the parameters for the SQL statement.
	params := []interface{}{
		tx.Hash.Hex(),                  // $1
		block.Hash.Hex(),               // $2
		block.Number.ToInt(),           // $3
		strings.ToLower(tx.From.Hex()), // $4
		to,                             // $5
		tx.Value.ToInt(),               // $6
		tx.GasPrice.ToInt(),            // $7
		tx.Gas.ToInt(),                 // $8
		hexutil.Encode(methodId),       // $9
		tx.Input.String(),              // $10
		tx.Nonce.ToInt(),               // $11
		tx.Type,                        // $12
		tx.L1BatchNumber.ToInt(),       // $13
		tx.L1BatchTxIndex.ToInt(),      // $14
		block.Timestamp.ToInt(),        // $15
	}

	// Adding the command to the batch.
	pgxBatch.Queue(sqlStmt, params...)
}

func (db *Client) UpsertL1Batch(ctx context.Context, batch *ethclient.RpcL1BatchDetails) error {
	sqlStmt := `
	INSERT INTO l1_batches (number, commit_tx_hash, committed_at, execute_tx_hash, executed_at, prove_tx_hash, proven_at, root_hash, status, l1_gas_price, l1_tx_count, l2_fair_gas_price, l2_tx_count, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	ON CONFLICT (number) DO UPDATE SET
		commit_tx_hash = $2,
		committed_at = $3,
		execute_tx_hash = $4,
		executed_at = $5,
		prove_tx_hash = $6,
		proven_at = $7,
		root_hash = $8,
		status = $9,
		l1_gas_price = $10,
		l1_tx_count = $11,
		l2_fair_gas_price = $12,
		l2_tx_count = $13,
		timestamp = $14
	`
	commitTxHash := sql.NullString{}
	if batch.CommitTxHash != nil {
		commitTxHash.String = batch.CommitTxHash.Hex()
		commitTxHash.Valid = true
	}

	committedAt := pq.NullTime{}
	if batch.CommittedAt != nil {
		committedAt.Time = *batch.CommittedAt
		committedAt.Valid = true
	}

	executeTxHash := sql.NullString{}
	if batch.ExecuteTxHash != nil {
		executeTxHash.String = batch.ExecuteTxHash.Hex()
		executeTxHash.Valid = true
	}

	executedAt := pq.NullTime{}
	if batch.ExecutedAt != nil {
		executedAt.Time = *batch.ExecutedAt
		executedAt.Valid = true
	}

	proveTxHash := sql.NullString{}
	if batch.ProveTxHash != nil {
		proveTxHash.String = batch.ProveTxHash.Hex()
		proveTxHash.Valid = true
	}

	provenAt := pq.NullTime{}
	if batch.ProvenAt != nil {
		provenAt.Time = *batch.ProvenAt
		provenAt.Valid = true
	}

	rootHash := sql.NullString{}
	if batch.RootHash != nil {
		rootHash.String = batch.RootHash.Hex()
		rootHash.Valid = true
	}

	_, err := db.pool.Exec(ctx, sqlStmt,
		batch.Number,
		commitTxHash,
		committedAt,
		executeTxHash,
		executedAt,
		proveTxHash,
		provenAt,
		rootHash,
		batch.Status,
		batch.L1GasPrice,
		batch.L1TxCount,
		batch.L2FairGasPrice,
		batch.L2TxCount,
		batch.Timestamp,
	)

	if err != nil {
		return err
	}
	return nil
}

func (db *Client) IsHashRootExist(ctx context.Context, batchRootHash *common.Hash) (bool, error) {
	sqlStmt := `SELECT COUNT(*) FROM l1_batches WHERE root_hash = $1`

	rootHash := sql.NullString{}
	if batchRootHash != nil {
		rootHash.String = batchRootHash.Hex()
		rootHash.Valid = true
	}

	count := 0

	err := db.pool.QueryRow(ctx, sqlStmt, rootHash).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *Client) UpsertL1BatchForMantaPacificAll(ctx context.Context, batches []*ethclient.MantaPacificBatchEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error
	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForMantaPacific(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpsertL1BatchForPolygonZkevmAll(ctx context.Context, batches []*ethclient.PolygonZkEVMBatchEvent) (*[]uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumbers []uint64
	for i, batch := range batches {
		batchNumbersTemp, err := db.UpsertL1BatchForPolygonZkevm(ctx, batch)
		if err != nil {
			return nil, err
		}
		batchNumbers = append(batchNumbers, *batchNumbersTemp)
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), batchNumbers)
	}
	return &batchNumbers, nil
}

func (db *Client) UpsertL1BatchForOKX1SepoliaAll(ctx context.Context, batches []*ethclient.OKX1SepoliaBatchEvent) (*[]uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumbers []uint64
	for i, batch := range batches {
		batchNumbersTemp, err := db.UpsertL1BatchForOKX1Sepolia(ctx, batch)
		if err != nil {
			return nil, err
		}
		batchNumbers = append(batchNumbers, *batchNumbersTemp)
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), batch.L2FinalBatch)
	}

	log.Debug().Msgf("🚧 Finish upserting new batchs %d [%d]", len(batches), batchNumbers)
	return &batchNumbers, nil
}

func (db *Client) UpsertL1BatchForLineaAll(ctx context.Context, batches []*ethclient.LineaBatchEvent) (*int64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *int64
	var err error
	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForLinea(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpsertL1BatchForScrollAll(ctx context.Context, batches []*ethclient.ScrollBatchEvent) (*uint64, error) {
	if len(batches) == 0 {
		return nil, nil
	}

	var batchNumber *uint64
	var err error
	for i, batch := range batches {
		batchNumber, err = db.UpsertL1BatchForScroll(ctx, batch)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("🚧 Upserting new batch %d/%d [%d]", i+1, len(batches), *batchNumber)
	}
	return batchNumber, nil
}

func (db *Client) UpsertL1SyncProgress(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "l1_synced_block_number", blockNumber)
}

func (db *Client) UpsertRecoveryL1BlockHeightProgress(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l1_synced_block_number", blockNumber)
}

func (db *Client) UpsertRecoveryL1L2MappingStatusProgress(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l2_batch_block_number", blockNumber)
}

func (db *Client) UpsertRecoveryL2SyncProgress(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l2_synced_block_number", blockNumber)
}

func (db *Client) UpsertRecoveryL2SyncTarget(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l2_synced_target_block_number", blockNumber)
}

func (db *Client) UpsertRecheckLastFinalizedBlock(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recheck_l1_last_finalized_block", blockNumber)
}

func (db *Client) UpsertLastRecoveryInternalTransactionBlock(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l2_internal_transaction_block_number", blockNumber)
}

func (db *Client) UpsertLastRecoverybatchNumber(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, "recovery_l1_batch_number", blockNumber)
}

func (db *Client) UpdateSyncProgress(ctx context.Context, key string, value uint64) error {
	_, err := db.pool.Exec(ctx, `INSERT INTO sync_progress (key, value) VALUES ($1, $2) 
        ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value`,
		key, value)
	return err
}

func (db *Client) UpsertL1BatchForMantaPacific(ctx context.Context, batch *ethclient.MantaPacificBatchEvent) (*uint64, error) {
	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (number) DO UPDATE SET
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6
	`
	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err := db.pool.Exec(ctx, sqlStmt, batch.Number, batch.L1TxHash.Hex(), provenAt, batch.OutputRootHash.Hex(), batch.Timestamp, "finalized")
	if err != nil {
		return nil, fmt.Errorf("insert batch failed: %v", err)
	}

	// Update all blocks batch info
	for _, blockNumber := range batch.FinalizedBlockNumbers {
		sqlStmt = `
		UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.Number, batch.Timestamp, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating block: %v", err)
		}
	}

	// Update all transactions batch info
	for _, blockNumber := range batch.FinalizedBlockNumbers {
		sqlStmt = `
		UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.Number, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating transaction: %v", err)
		}
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.FinalizedBlockNumbers), batch.Number)

	return &batch.Number, nil
}

func (db *Client) UpsertL1BatchForPolygonZkevm(ctx context.Context, batch *ethclient.PolygonZkEVMBatchEvent) (*uint64, error) {
	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number, l1_prove_block_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT (number) DO UPDATE SET
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6,
		l2_block_number = $7,
	    l1_prove_block_number = $8 
	`

	var l2BlockBumber *big.Int
	if len(batch.L2BatchBlocks) > 0 {
		l2BlockBumber = batch.L2BatchBlocks[0]
	} else {
		l2BlockBumber = nil
	}

	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err := db.pool.Exec(ctx, sqlStmt,
		batch.L2FinalBatch,
		batch.L1TxHash.Hex(),
		provenAt,
		batch.RootHash.Hex(),
		batch.Timestamp,
		"finalized",
		l2BlockBumber,
		batch.L1BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("[batches] error inserting batch: %v", err)
	}

	// Update all blocks batch info
	for _, blockNumber := range batch.L2BatchBlocks {
		sqlStmt = `
			UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
			`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.L2FinalBatch, batch.Timestamp, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating block: %v", err)
		}
	}

	// Update all transactions batch info
	for _, blockNumber := range batch.L2BatchBlocks {
		sqlStmt = `
		UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.L2FinalBatch, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating transaction: %v", err)
		}
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.L2BatchBlocks), batch.L2FinalBatch)

	var updatedBatch = batch.L2FinalBatch

	return &updatedBatch, nil
}

func (db *Client) UpsertL1BatchForOKX1Sepolia(ctx context.Context, batch *ethclient.OKX1SepoliaBatchEvent) (*uint64, error) {
	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (number, prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number, l1_prove_block_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT (number) DO UPDATE SET
		prove_tx_hash = $2,
		proven_at = $3,
		root_hash = $4,
		timestamp = $5,
		status = $6,
		l2_block_number = $7,
	    l1_prove_block_number = $8 
	`

	var l2BlockBumber *big.Int
	if len(batch.L2BatchBlocks) > 0 {
		l2BlockBumber = batch.L2BatchBlocks[0]
	} else {
		l2BlockBumber = nil
	}

	provenAt := time.Unix(int64(batch.Timestamp), 0)
	_, err := db.pool.Exec(ctx, sqlStmt,
		batch.L2FinalBatch,
		batch.L1TxHash.Hex(),
		provenAt,
		batch.RootHash.Hex(),
		batch.Timestamp,
		"finalized",
		l2BlockBumber,
		batch.L1BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("[batches] error inserting batch: %v", err)
	}

	// Update all blocks batch info
	for _, blockNumber := range batch.L2BatchBlocks {
		sqlStmt = `
			UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
			`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.L2FinalBatch, batch.Timestamp, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating block: %v", err)
		}
	}

	// Update all transactions batch info
	for _, blockNumber := range batch.L2BatchBlocks {
		sqlStmt = `
		UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batch.L2FinalBatch, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating transaction: %v", err)
		}
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.L2BatchBlocks), batch.L2FinalBatch)

	var updatedBatch = batch.L2FinalBatch

	return &updatedBatch, nil
}

func (db *Client) UpsertL1BatchForLinea(ctx context.Context, batch *ethclient.LineaBatchEvent) (*int64, error) {
	// to ensure that we do not encounter a skipped sequence number, we should verify the existence of the batch by checking the prove_tx_hash.
	sqlStmtCheck := `
	SELECT number FROM l1_batches WHERE prove_tx_hash = $1
	`
	var existedBatchNumber int64
	err := db.pool.QueryRow(ctx, sqlStmtCheck, batch.TransactionHash.Hex()).Scan(&existedBatchNumber)
	if err == nil {
		return &existedBatchNumber, nil
	}

	// Insert batch
	sqlStmt := `
	INSERT INTO l1_batches (prove_tx_hash, proven_at, root_hash, timestamp, status, l2_block_number, l1_prove_block_number)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (root_hash) DO UPDATE
	SET prove_tx_hash = $1,
    proven_at = $2,
	root_hash = $3,  
    timestamp = $4,
    status = $5,
	l2_block_number = $6, 
	l1_prove_block_number = $7
	RETURNING number
	`

	var finalizedBlockNumbers uint64
	var batchNumber int64

	if len(batch.FinalizedBlockNumbers) == 0 {
		return nil, errors.New("upsert l1 batch for Linea err: finalizedBlockNumbers is nil")
	} else {
		finalizedBlockNumbers = batch.FinalizedBlockNumbers[len(batch.FinalizedBlockNumbers)-1]
	}

	provenAt := time.Unix(int64(batch.Timestamp), 0)
	err = db.pool.QueryRow(ctx, sqlStmt,
		batch.TransactionHash.Hex(),
		provenAt,
		batch.FinalRootHash.Hex(),
		batch.Timestamp,
		"verified",
		finalizedBlockNumbers,
		batch.BlockNumber).Scan(&batchNumber)
	if err != nil {
		return nil, fmt.Errorf("[batches] error inserting batch: %v", err)
	}

	// Update all blocks batch info
	for _, blockNumber := range batch.FinalizedBlockNumbers {
		sqlStmt = `
		UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batchNumber, batch.Timestamp, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating block: %v", err)
		}
	}

	// Update all transactions batch info
	for _, blockNumber := range batch.FinalizedBlockNumbers {
		sqlStmt = `
		UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
		`
		_, err = db.pool.Exec(ctx, sqlStmt, batchNumber, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("error updating transaction: %v", err)
		}
	}

	log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.FinalizedBlockNumbers), batch.Number)

	return &batchNumber, nil
}

func (db *Client) UpsertL1BatchForScroll(ctx context.Context, batch *ethclient.ScrollBatchEvent) (*uint64, error) {
	// Insert commit batch
	if batch.Status == ethclient.ScrollBatchStatusCommitted {
		sqlStmt := `
				INSERT INTO l1_batches (number, commit_tx_hash, committed_at, status, timestamp)
				VALUES ($1, $2, $3, $4, $5)
				ON CONFLICT (number) DO UPDATE SET
				commit_tx_hash = $2,
				committed_at = $3,
				status = $4,
				timestamp = $5
				`
		committedAt := time.Unix(int64(batch.Timestamp), 0)
		_, err := db.pool.Exec(ctx, sqlStmt, batch.Number, batch.TxHash.Hex(), committedAt, batch.Status, batch.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("[batches] error inserting batch: %v", err)
		}

		// Update all blocks batch info
		for _, blockNumber := range batch.CommittedBlocks {
			sqlStmt = `
					UPDATE blocks SET l1_batch_number = $1, l1_batch_timestamp = $2 WHERE number = $3
					`
			_, err = db.pool.Exec(ctx, sqlStmt, batch.Number, batch.Timestamp, blockNumber)
			if err != nil {
				return nil, fmt.Errorf("error updating block: %v", err)
			}
		}

		// Update all transactions batch info
		for _, blockNumber := range batch.CommittedBlocks {
			sqlStmt = `
					UPDATE transactions SET l1_batch_number = $1 WHERE block_number = $2
					`
			_, err = db.pool.Exec(ctx, sqlStmt, batch.Number, blockNumber)
			if err != nil {
				return nil, fmt.Errorf("error updating transaction: %v", err)
			}
		}

		log.Debug().Msgf("🚧 Updated %d blocks with batch number %d", len(batch.CommittedBlocks), batch.Number)

		return &batch.Number, nil
	}

	// Insert prove batch
	if batch.Status == ethclient.ScrollBatchStatusProven {
		// only update l1_batches table
		provenAt := time.Unix(int64(batch.Timestamp), 0)
		sqlStmt := `
				UPDATE l1_batches SET prove_tx_hash = $1, proven_at = $2, root_hash = $3, status = $4, timestamp = $5 WHERE number = $6
				`
		_, err := db.pool.Exec(ctx, sqlStmt, batch.TxHash.Hex(), provenAt, batch.RootHash.Hex(), batch.Status, batch.Timestamp, batch.Number)
		if err != nil {
			return nil, fmt.Errorf("[batches] error updating batch: %v", err)
		}

		return &batch.Number, nil
	}

	return nil, fmt.Errorf("[batches] unknown batch status: %s", batch.Status)
}

func (db *Client) UpdateTransactionWithReceipt(pgxBatch *pgx.Batch, receipt *ethclient.Receipt, revertReason *string) {
	var fee *big.Int
	if receipt.EffectiveGasPrice == nil {
		fee = big.NewInt(0)
	} else {
		fee = big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
	}
	var l1Fee *big.Int
	if receipt.L1Fee != nil {
		l1Fee = receipt.L1Fee.ToInt()
	}

	sqlStmt := `UPDATE transactions SET transaction_index = $1, status = $2, gas_used = $3, fee = $4, l1fee = $5, revert_reason = $6 WHERE hash = $7`
	pgxBatch.Queue(sqlStmt, receipt.TransactionIndex, receipt.Status, receipt.GasUsed, fee, l1Fee, revertReason, receipt.TxHash.Hex())
}

func (db *Client) InsertLog(ctx context.Context, txLog *types.Log) error {
	sqlStmt := `
	INSERT INTO transaction_logs (transaction_hash, log_index, address, block_number, block_hash, topic1, topic2, topic3, topic4, data, removed) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	ON CONFLICT (transaction_hash, log_index) DO NOTHING
	`
	topics := []interface{}{nil, nil, nil, nil}
	for i, topic := range txLog.Topics {
		if i >= 4 {
			break
		}
		topics[i] = topic.Hex()
	}
	_, err := db.pool.Exec(ctx, sqlStmt,
		txLog.TxHash.Hex(),
		txLog.Index,
		strings.ToLower(txLog.Address.Hex()),
		big.NewInt(int64(txLog.BlockNumber)),
		txLog.BlockHash.Hex(),
		topics[0],
		topics[1],
		topics[2],
		topics[3],
		hexutil.Encode(txLog.Data),
		txLog.Removed,
	)
	if err != nil {
		log.Err(err).Msg("Error inserting log")
		return err
	}
	return nil
}

func (db *Client) InsertLogs(pgxBatch *pgx.Batch, txLogs []*types.Log) {
	for _, l := range txLogs {
		topics := []interface{}{nil, nil, nil, nil}
		for i, topic := range l.Topics {
			if i >= 4 {
				break
			}
			topics[i] = topic.Hex()
		}

		pgxBatch.Queue(`INSERT INTO transaction_logs (transaction_hash, log_index, address, block_number, block_hash, topic1, topic2, topic3, topic4, data, removed) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (transaction_hash, log_index) DO NOTHING`,
			l.TxHash.Hex(),
			l.Index,
			strings.ToLower(l.Address.Hex()),
			big.NewInt(int64(l.BlockNumber)),
			l.BlockHash.Hex(),
			topics[0],
			topics[1],
			topics[2],
			topics[3],
			hexutil.Encode(l.Data),
			l.Removed,
		)
	}
}

func (db *Client) InsertTokenTransfer(ctx context.Context, transfer TokenTransfer, block *ethclient.RpcBlock) error {
	sqlStmt := `
	INSERT INTO token_transfers (transaction_hash, log_index, method_id, token_address, block_number, block_hash, from_address, to_address, value, amount, token_id, amounts, token_ids, token_type, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	ON CONFLICT (transaction_hash, log_index) DO NOTHING
	`
	_, err := db.pool.Exec(ctx, sqlStmt,
		transfer.TransactionHash.Hex(),
		transfer.LogIndex,
		transfer.MethodID,
		strings.ToLower(transfer.TokenAddress.Hex()),
		block.Number.ToInt(),
		block.Hash.Hex(),
		strings.ToLower(transfer.From.Hex()),
		strings.ToLower(transfer.To.Hex()),
		transfer.Value,
		transfer.Amount,
		transfer.TokenID,
		transfer.Amounts,
		transfer.TokenIDs,
		transfer.TokenType,
		block.Timestamp.ToInt(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *Client) InsertTokenTransfers(pgxBatch *pgx.Batch, transfers []TokenTransfer, block *ethclient.RpcBlock) {
	for _, transfer := range transfers {
		pgxBatch.Queue(`INSERT INTO token_transfers (transaction_hash, log_index, method_id, token_address, block_number, block_hash, from_address, to_address, value, amount, token_id, amounts, token_ids, token_type, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14 , $15)
		ON CONFLICT (transaction_hash, log_index) DO NOTHING`,
			transfer.TransactionHash.Hex(),
			transfer.LogIndex,
			transfer.MethodID,
			strings.ToLower(transfer.TokenAddress.Hex()),
			block.Number.ToInt(),
			block.Hash.Hex(),
			strings.ToLower(transfer.From.Hex()),
			strings.ToLower(transfer.To.Hex()),
			transfer.Value,
			transfer.Amount,
			transfer.TokenID,
			transfer.Amounts,
			transfer.TokenIDs,
			transfer.TokenType,
			block.Timestamp.ToInt(),
		)
	}
}

func (db *Client) InsertInternalTransaction(pgxBatch *pgx.Batch, tx *InternalTransaction) {
	stmt := `INSERT INTO internal_transactions (block_number, block_hash, parent_transaction_hash, type, from_address, to_address, value, gas, gas_used, input, output, method, timestamp)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	var output *string
	if tx.Output != nil {
		outputString := tx.Output.String()
		output = &outputString
	}

	// Preparing the parameters for the SQL statement.
	params := []interface{}{
		tx.BlockNumber.ToInt(),         //$1
		tx.BlockHash.Hex(),             //$2
		tx.ParentTransactionHash.Hex(), //$3
		tx.Type,                        //$4
		strings.ToLower(tx.From.Hex()), //$5
		strings.ToLower(tx.To.Hex()),   //$6
		tx.Value.ToInt(),               //$7
		tx.Gas.ToInt(),                 //$8
		tx.GasUsed.ToInt(),             //$9
		tx.Input.String(),              //$10
		output,                         //$11
		tx.Method,                      //$12
		tx.Timestamp.ToInt(),           //$13
	}

	// Adding the command to the batch.
	pgxBatch.Queue(stmt, params...)

}

func (db *Client) HasInternalTransactionByBlockNumber(ctx context.Context, blockNumber string) (bool, error) {
	stmt := `SELECT COUNT(*) FROM internal_transactions WHERE block_number = $1`

	row := db.pool.QueryRow(ctx, stmt, blockNumber)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *Client) HasInternalTransactionCountByBlockNumber(ctx context.Context, blockNumber string) (bool, error) {
	stmt := `SELECT internal_transaction_count FROM blocks WHERE number = $1`
	row := db.pool.QueryRow(ctx, stmt, blockNumber)

	var internalTransactionCount *int
	err := row.Scan(&internalTransactionCount)

	if err != nil {
		return false, err
	}

	if internalTransactionCount == nil {
		return false, nil
	}

	//if *internalTransactionCount == 0 {
	//	return true, nil
	//}

	return true, nil
}

func (db *Client) UpsertAddressBalance(pgxBatch *pgx.Batch, address common.Address, balance *big.Int, blockNumber *big.Int) {
	stmt := `INSERT INTO address_balances (address, balance, updated_block_number) VALUES ($1, $2, $3) ON CONFLICT (address) DO UPDATE SET balance = $2, updated_block_number = $3`
	pgxBatch.Queue(stmt, strings.ToLower(address.Hex()), balance, blockNumber)
}

func (db *Client) UpsertTokenBalance(pgxBatch *pgx.Batch, tokenAddress common.Address, account common.Address, balance *big.Int, blockNumber *big.Int, tokenID *big.Int, tokenType string) {
	if tokenType == "erc20" {
		// Delete first when token_type is erc20 because upsert operation is invalid when token_id is null
		// ! we also need this operation to delete garbage data that has inserted duplicate token_address, address and token_type
		deleteStmt := `DELETE FROM token_balances WHERE token_address = $1 AND address = $2 AND token_type = $3`
		pgxBatch.Queue(deleteStmt, strings.ToLower(tokenAddress.Hex()), strings.ToLower(account.Hex()), tokenType)
		if balance.Cmp(big.NewInt(0)) == 0 {
			return
		}
	}

	upsertStmt := `INSERT INTO token_balances (token_address, address, balance, updated_block_number, token_id, token_type) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (token_address, address, token_id, token_type) DO UPDATE SET balance = $3, updated_block_number = $4, token_id = $5, token_type = $6`
	pgxBatch.Queue(upsertStmt, strings.ToLower(tokenAddress.Hex()), strings.ToLower(account.Hex()), balance, blockNumber, tokenID, tokenType)
}

func (db *Client) UpsertTokenInfo(pgxBatch *pgx.Batch, token Token) {
	stmt := `INSERT INTO tokens (address, name, symbol, decimals, total_supply, token_type) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (address) DO UPDATE SET name = $2, symbol = $3, decimals = $4, total_supply = $5, token_type = $6`
	pgxBatch.Queue(stmt, strings.ToLower(token.Address.Hex()), token.Name, token.Symbol, token.Decimals, token.TotalSupply, token.TokenType)
}

func (db *Client) UpsertContract(pgxBatch *pgx.Batch, contract Contract) {
	stmt := `INSERT INTO contracts (address, creator, creation_tx_hash, creation_bytecode, deployed_bytecode, creation_timestamp)
	VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (address) DO UPDATE SET creator = $2, creation_tx_hash = $3, creation_bytecode = $4, deployed_bytecode = $5, creation_timestamp = $6`
	pgxBatch.Queue(stmt,
		strings.ToLower(contract.Address.Hex()),
		strings.ToLower(contract.Creator.Hex()),
		contract.CreationTxHash.Hex(),
		contract.CreationBytecode.String(),
		contract.DeployedBytecode.String(),
		contract.CreationTimestamp.ToInt(),
	)
}

// GetMissingBlocks retrieves missing block information based on the specified number of the most recent blocks.
// If blocksCount is greater than 0, it only returns the missing blocks within that range.
func (db *Client) GetMissingBlocks(ctx context.Context, blocksCount int) ([]uint64, error) {
	var missingBlocks []uint64

	var stmt string

	if blocksCount > 0 {
		// Modify the query to search for missing blocks within the specified number of the most recent blocks
		// Also, modify the sequence generation part to only include block numbers within this range
		stmt = fmt.Sprintf(`WITH RECENT_BLOCKS AS (
				SELECT number FROM blocks ORDER BY number DESC LIMIT %d
			), MIN_MAX AS (
				SELECT MIN(number) AS min, MAX(number) AS max FROM RECENT_BLOCKS
			)
			SELECT series.number AS missing_block_number
			FROM generate_series((SELECT min FROM MIN_MAX), (SELECT max FROM MIN_MAX)) AS series(number)
			EXCEPT
			SELECT number FROM blocks
			WHERE number IN (SELECT number FROM RECENT_BLOCKS)
			ORDER BY missing_block_number DESC;`, blocksCount)
	} else {
		// If no specific number of blocks is specified, use the base query
		stmt = `SELECT series.number AS missing_block_number
				FROM generate_series((SELECT MIN(number) FROM blocks), (SELECT MAX(number) FROM blocks)) AS series(number)
				EXCEPT
				SELECT number FROM blocks
				ORDER BY missing_block_number DESC;`
	}

	rows, err := db.pool.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err := rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		missingBlocks = append(missingBlocks, number)
	}

	// Handle any errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return missingBlocks, nil
}

// GetMismatchedTransactionBlocks retrieves blocks where the transaction count does not match the expected count.
// If blocksCount is greater than 0, it limits the query to the most recent blocksCount number of blocks.
func (db *Client) GetMismatchedTransactionBlocks(ctx context.Context, blocksCount int) ([]uint64, error) {
	var mismatchedBlocks []uint64

	var stmt string

	if blocksCount > 0 {
		// Modify the query to limit the search within the most recent blocksCount number of blocks
		stmt = fmt.Sprintf(`WITH RECENT_BLOCKS AS (
				SELECT number FROM blocks ORDER BY number DESC LIMIT %d
			)
			SELECT b.number
			FROM blocks b
			LEFT JOIN (
				SELECT block_number, COUNT(*) as transaction_count
				FROM transactions
				GROUP BY block_number
			) tx ON b.number = tx.block_number
			WHERE b.number IN (SELECT number FROM RECENT_BLOCKS) AND b.transaction_count != tx.transaction_count
			ORDER BY b.number;`, blocksCount)
	} else {
		// If blocksCount is not specified or is 0, revert to the original behavior (adjust accordingly)
		stmt = `SELECT b.number
				FROM blocks b
				LEFT JOIN (
					SELECT block_number, COUNT(*) as transaction_count
					FROM transactions
					GROUP BY block_number
				) tx ON b.number = tx.block_number
				WHERE b.transaction_count != tx.transaction_count
				ORDER BY b.number;`
	}

	rows, err := db.pool.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err := rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		mismatchedBlocks = append(mismatchedBlocks, number)
	}

	// Handle any errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mismatchedBlocks, nil
}

func (db *Client) GetMissingBatches(ctx context.Context, startBatchNumber uint64, endBatchNumber uint64) ([]uint64, error) {
	var missingBatches []uint64
	stmt := `SELECT series.number AS missing_batch_number
			FROM generate_series($1::bigint, $2::bigint) AS series(number)
			EXCEPT
			SELECT number FROM l1_batches
			WHERE number BETWEEN $1 AND $2
			ORDER BY missing_batch_number;
			`

	rows, err := db.pool.Query(ctx, stmt, startBatchNumber, endBatchNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err := rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		missingBatches = append(missingBatches, number)
	}

	// Handle any rows error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return missingBatches, nil
}

func (db *Client) GetUnfinalizedL1Batches(ctx context.Context) ([]uint64, error) {
	var batchNumbers []uint64
	stmt := `SELECT number FROM l1_batches WHERE status != 'verified'`

	rows, err := db.pool.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err := rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		batchNumbers = append(batchNumbers, number)
	}

	// Handle any rows error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return batchNumbers, nil
}

func (db *Client) GetUnfinalizedL2Blocks(ctx context.Context) ([]uint64, error) {
	var unfinalizedBlocks []uint64
	stmt := `select distinct block_number from transactions where l1_batch_number IS NULL LIMIT 10000`

	rows, err := db.pool.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err := rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		unfinalizedBlocks = append(unfinalizedBlocks, number)
	}

	// Handle any rows error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unfinalizedBlocks, nil
}

func (db *Client) DeleteBlock(ctx context.Context, blockNumber uint64) error {
	stmt := `DELETE FROM blocks WHERE number = $1`
	_, err := db.pool.Exec(ctx, stmt, blockNumber)
	return err
}

func (db *Client) MaxUnIndexedBlockNumber(ctx context.Context) (blockNumber int64, err error) {
	var value sql.NullInt64
	err = db.pool.QueryRow(ctx, `select max(number) from blocks where trace_checked is not true`).Scan(&value)
	if err != nil {
		return
	}
	if !value.Valid {
		blockNumber = -1
		return
	}
	blockNumber = value.Int64
	return
}

func (db *Client) UpdateBlockTraceChecked(batch *pgx.Batch, blockNumber *big.Int, traceCount int) {
	stmt := `UPDATE blocks SET trace_checked = true, internal_transaction_count = $1 WHERE number = $2`
	batch.Queue(stmt, traceCount, blockNumber.Int64())
}

func (db *Client) BlockTraceHasChecked(ctx context.Context, blockNumber *big.Int) (checked bool, err error) {
	err = db.pool.QueryRow(ctx, `select COALESCE(trace_checked,false) from blocks where number = $1`, blockNumber.Int64()).Scan(&checked)
	return
}

func (db *Client) GetMissedTraceBlockNumbers(ctx context.Context, limit int) (blockNumbers []uint64, err error) {
	var rows pgx.Rows
	if limit == 0 {
		// If limit is 0, select all block numbers
		rows, err = db.pool.Query(ctx, `SELECT number FROM blocks WHERE trace_checked IS FALSE OR trace_checked IS NULL ORDER BY number DESC`)
	} else {
		// Otherwise, select block numbers with the given limit
		rows, err = db.pool.Query(ctx, `SELECT number FROM blocks WHERE trace_checked IS FALSE OR trace_checked IS NULL ORDER BY number DESC LIMIT $1`, limit)
	}

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var number uint64
		err = rows.Scan(&number)
		if err != nil {
			return
		}
		blockNumbers = append(blockNumbers, number)
	}

	return
}

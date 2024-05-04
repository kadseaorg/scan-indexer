package db

import (
	"context"
	"testing"
	"time"
)

func InitTmpPGDB() (*Client, error) {
	pg := "postgres://DcCKwJJDL6YUpH2L@localhost:5432/kadsea_01"

	dbClient, err := NewClient(pg)
	if err != nil {
		return nil, err
	}

	return dbClient, nil
}

func TestSelectNumbersForNullBatches(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	start := time.Now()
	missBatchs, err := dbClient.SelectNumbersForNullBatches(context.Background(), 176395507, 176395537)
	if err != nil {
		t.Errorf("Error Select Numbers For Null Batches: %s", err)
	}

	cost := time.Now().Sub(start).Seconds()

	t.Logf("cost %v miss batchs: [%d]%v", cost, len(missBatchs), missBatchs)
}

// GetLatestBatchNumber
func TestGetLatestBatchNumber(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	start := time.Now()
	batch, err := dbClient.GetLatestBatchNumber(context.Background())
	if err != nil {
		t.Errorf("Error Get Latest Batch Number: %s", err)
	}

	cost := time.Now().Sub(start).Seconds()

	t.Logf("cost %v latest batchs: %v", cost, batch)
}

// GetLatestBlockNumber
func TestGetLatestBlockNumber(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	start := time.Now()
	block, err := dbClient.GetLatestBlockNumber(context.Background())
	if err != nil {
		t.Errorf("Error Get Latest Batch Number: %s", err)
	}

	cost := time.Now().Sub(start).Seconds()

	t.Logf("cost %v latest block number: %v", cost, block)
}

// GetMissingBatches
func TestGetMissingBatches(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	start := time.Now()
	batcher, err := dbClient.GetMissingBatches(context.Background(), 0, 100)
	if err != nil {
		t.Errorf("Error Get Missing Batches: %s", err)
	}

	cost := time.Now().Sub(start).Seconds()

	t.Logf("cost %v Get Missing Batches: %v", cost, batcher)
}

// GetMissingBlocks
func TestGetMissingBlocks(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	start := time.Now()
	batcher, err := dbClient.GetMissingBlocks(context.Background(), 52217872, 82217865)
	if err != nil {
		t.Errorf("Error Get Missing Batches: %s", err)
	}

	cost := time.Now().Sub(start).Seconds()

	t.Logf("cost %v Get Missing Batches: %v", cost, batcher)
}

// CheckL1BlockNumber
func TestCheckL1BlockNumber(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	err = dbClient.CreateFieldsL1BlockNumberForL1Batches()
	if err != nil {
		t.Errorf("Error CheckL1BlockNumber: %s", err)
	}
}

package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
	"testing"
	"time"
)

func TestUpsertAddressBalance(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	bn := big.NewInt(100)
	addr := common.HexToAddress("0xzksync10086")

	wg := sync.WaitGroup{}
	wg.Add(50)

	for i := 1; i < 51; i++ {
		go func(i int) {
			defer wg.Done()
			balacne := big.NewInt(int64(i))
			start := time.Now()
			err = dbClient.UpsertAddressBalance(context.Background(), addr, balacne, bn)
			if err != nil {
				t.Errorf("Error Select Numbers For Null Batches: %s", err)
			}

			cost := time.Now().Sub(start).Seconds()
			t.Logf("cost %v block %d addr %v balance %d", cost, bn, addr, balacne)
		}(i)
	}

	wg.Wait()
}

// GetLatestBatchOfLastL1ProveBlock
func TestGetLatestBatchOfLastL1ProveBlock(t *testing.T) {
	dbClient, err := InitTmpPGDB()
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}

	//l1ProveBlock, err := dbClient.GetLatestBatchOfLastL1ProveBlock(context.Background())
	//if err != nil {
	//	t.Errorf("Error GetLatestBatchOfLastL1ProveBlock : %s", err)
	//}
	//
	//t.Logf("l1ProveBlock:%d", l1ProveBlock)
	//
	//l1CommitBlock, err := dbClient.GetLatestBatchOfLastL1CommitBlock(context.Background())
	//if err != nil {
	//	t.Errorf("Error GetLatestBatchOfLastL1CommitBlock : %s", err)
	//}
	//
	//t.Logf("l1CommitBlock:%d", l1CommitBlock)
	//
	//l1ExecuteBlock, err := dbClient.GetLatestBatchOfLastL1ExecuteBlock(context.Background())
	//if err != nil {
	//	t.Errorf("Error GetLatestBatchOfLastL1ExecuteBlock : %s", err)
	//}
	//
	//t.Logf("l1ExecuteBlock:%d", l1ExecuteBlock)
	//
	//exist, err := dbClient.HasInternalTransactionByBlockNumber(context.Background(), "1817")
	//if err != nil {
	//	t.Errorf("Error HasInternalTransactionByBlockNumber : %s", err)
	//}
	//
	//t.Logf("HasInternalTransactionByBlockNumber:%v", exist)
	//
	//exist, err = dbClient.HasInternalTransactionCountByBlockNumber(context.Background(), "1505")
	//if err != nil {
	//	t.Errorf("Error HasInternalTransactionCountByBlockNumber : %s", err)
	//}
	//t.Logf("HasInternalTransactionCountByBlockNumber:%v", exist)
	//
	//err = dbClient.UpsertBlockInternalTransactionCount(context.Background(), big.NewInt(1505), 0)
	//if err != nil {
	//	t.Errorf("Error UpsertBlockInternalTransactionCount : %s", err)
	//}
	//t.Log("UpsertBlockInternalTransactionCount")
	//
	//exist, err = dbClient.HasInternalTransactionCountByBlockNumber(context.Background(), "1505")
	//if err != nil {
	//	t.Errorf("Error HasInternalTransactionCountByBlockNumber : %s", err)
	//}
	//t.Logf("HasInternalTransactionCountByBlockNumber:%v", exist)

	//AddBlockInternalTransactionCount
	err = dbClient.AddBlockInternalTransactionCount(context.Background(), big.NewInt(1505), 1)
	if err != nil {
		t.Errorf("Error UpsertBlockInternalTransactionCount : %s", err)
	}
	t.Log("UpsertBlockInternalTransactionCount")
}

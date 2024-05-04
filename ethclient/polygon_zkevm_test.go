package ethclient

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPolygonBatchDeliveredEventLogs(t *testing.T) {
	l1Url := "http://10.3.3.94:8545"
	l2Url := "http://10.3.3.42:8545"

	l1Client, err := NewClient(l1Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	//16896722, 16896772

	events, err := l1Client.GetPolygonZkEVMSequenceBatchesEventLogs(context.Background(), PolygonZkEVMSequencerInbox, 16919572, 16919590) //16898936, 16898936) //16896721
	if err != nil {
		t.Errorf("GetPolygonZkEVMSequenceBatchesEventLogs err %v", err)
	}

	if len(events) > 0 {
		t.Logf("Polygon len %d ", len(events))
	} else {
		t.Fatalf("event is empty")
	}

	for _, event := range events {
		batchInfo, err := l2Client.GetBatchByNumber(context.Background(), event.L2FinalBatch)
		if err != nil {
			t.Errorf("GetBatchByNumber err %v", err)
		}

		t.Logf("event: %+v", *event)

		t.Logf("batchInfo.Blocks: %+v", batchInfo.Blocks)

		for _, blockHash := range batchInfo.Blocks {
			blockInfo, err := l2Client.GetBlockByHash(context.Background(), blockHash)
			if err != nil {
				return
			}

			t.Logf("blockInfo.Number: %v blockInfo: %+v ", blockInfo.Number, blockInfo)
		}
	}
}

func TestGetOKX1BatchDeliveredEventLogs(t *testing.T) {
	l1Url := "http://10.3.3.234:8545"
	l2Url := "http://10.3.3.252:8545"

	l1Client, err := NewClient(l1Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	events, err := l1Client.GetOKX1SequenceBatchesEventLogs(context.Background(),
		OKX2SequencerInbox,
		5421721, //5420254,
		5421821) //5420294) //16898936, 16898936) //16896721
	if err != nil {
		t.Errorf("GetPolygonZkEVMSequenceBatchesEventLogs err %v", err)
	}

	if len(events) > 0 {
		t.Logf("OKX1 len %d ", len(events))
	} else {
		t.Fatalf("event is empty")
	}

	L2InitBatch := uint64(306631)

	newBatchEvents := make([]*PolygonZkEVMBatchEvent, 0)
	for _, batch := range events {
		t.Logf("event1 : %+v", *batch)
		for batchNum := L2InitBatch + 1; batchNum <= batch.L2FinalBatch; batchNum++ {
			newBatchEvents = append(newBatchEvents, &PolygonZkEVMBatchEvent{
				L1BlockNumber: batch.L1BlockNumber,
				L1TxHash:      batch.L1TxHash,
				L2InitBatch:   batch.L2InitBatch,
				L2FinalBatch:  batchNum,
				Timestamp:     batch.Timestamp,
				RootHash:      batch.RootHash,
			})
		}
		L2InitBatch = batch.L2FinalBatch
	}

	for _, event := range newBatchEvents {
		batchInfo, err := l2Client.GetBatchByNumber(context.Background(), event.L2FinalBatch)
		if err != nil {
			t.Errorf("GetBatchByNumber err %v", err)
		}
		t.Logf("event2 : %+v", *event)
		t.Logf("len(%d)batchInfo.Blocks: %+v", len(batchInfo.Blocks), batchInfo.Blocks)
		for _, blockHash := range batchInfo.Blocks {
			blockInfo, err := l2Client.GetBlockByHash(context.Background(), blockHash)
			if err != nil {
				return
			}
			t.Logf("blockInfo.Number: %v", blockInfo.Number.ToInt().Int64())
		}
	}
}

func TestGetBsquaredBatchDeliveredEventLogs(t *testing.T) {
	url := "https://haven-b2-nodes.bsquared.network"
	l2Url := "https://haven-rpc.bsquared.network"

	l1Client, err := NewClient(url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	//16896722, 16896772
	events, err := l1Client.GetPolygonZkEVMSequenceBatchesEventLogs(context.Background(), BsquaredSequencerInbox, 3422359, 3422559) //551084
	if err != nil {
		t.Errorf("GetBsquaredSequenceBatchesEventLogs err %v", err)
	}

	if len(events) > 0 {
		t.Logf("Bsquared len %d ", len(events))
	} else {
		t.Fatalf("event is empty")
	}

	for _, event := range events {
		batchInfo, err := l2Client.GetBatchByNumber(context.Background(), event.L2FinalBatch)
		if err != nil {
			t.Errorf("GetBatchByNumber err %v", err)
		}

		t.Logf("event: %+v", *event)

		t.Logf("batchInfo.Blocks: %+v", batchInfo.Blocks)

		for _, blockHash := range batchInfo.Blocks {
			blockInfo, err := l2Client.GetBlockByHash(context.Background(), blockHash)
			if err != nil {
				return
			}

			t.Logf("blockInfo.Number: %v blockInfo: %+v ", blockInfo.Number.ToInt(), blockInfo)
		}
	}

	//l1 3222379 l2 576864
	// 3322587 579071
}

func TestGetBsquaredBatchInfo(t *testing.T) {
	url := "https://haven-b2-nodes.bsquared.network"
	l1Client, err := NewClient(url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	//block, err := l1Client.BlockNumber(context.Background())
	//if err != nil {
	//	t.Errorf("GetBatchByNumber err %v", err)
	//}

	//tx, ispending, err := l1Client.TransactionByHash(context.Background(), common.HexToHash("0x90236be60820e85e82a03028961637d98dda6b04a8b65bc49c5d115d69a2a5b7"))
	//if err != nil {
	//	t.Errorf("TransactionByHash err %v", err)
	//}
	//t.Logf("ispending: %+v tx %+v", ispending, tx)

	//0x1d7852cfd42d1a113063ed17933cafba3272cff242ee964d5d671df13423b5d1  batch：15286	block：4105221
	//0xd90aeda77f53bd44df4aa7cf46be34d451070cdf97b39834367ec514b73f46af	15286  1021280

	receipt, err := l1Client.TransactionReceipt(context.Background(), common.HexToHash("0xd90aeda77f53bd44df4aa7cf46be34d451070cdf97b39834367ec514b73f46af"))
	if err != nil {
		t.Errorf("TransactionByHash err %v", err)
	}

	t.Logf("block %+v eceipt: %+v ", receipt.BlockNumber.Int64(), receipt)
}

//0x90236be60820e85e82a03028961637d98dda6b04a8b65bc49c5d115d69a2a5b7

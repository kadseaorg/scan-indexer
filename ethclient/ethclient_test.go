package ethclient

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTransactionReceipt(t *testing.T) {
	client, err := NewClient("https://scroll-mainnet-public.unifra.io")
	assert.NoError(t, err, "Expected no error while creating client")

	txHash := common.HexToHash("0x0649eba4cb3dbea9617ad4114dbbe1ddf84018ebffda21ab01e4c430358192f2")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)

	assert.NoError(t, err, "Expected no error while getting receipt")

	// check receipt l1fee is not nil
	assert.NotNil(t, receipt.L1Fee, "Expected L1Fee to be not nil")
}

func TestGetOKX1SepoliaBatchDeliveredEventLogs(t *testing.T) {
	l1Url := "http://10.3.3.234:8545"
	l2Url := "http://10.3.3.252:8545"

	l1Client, err := NewClient(l1Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	sttart := uint64(4653166)
	end := uint64(4653628)
	//16896722, 16896772
	events, err := l1Client.GetOKX1SepoliaSequenceBatchesEventLogs(context.Background(), OKX1SepoliaConsensusSmartContract, sttart, end)
	if err != nil {
		t.Errorf("GetOKX1SepoliaSequenceBatchesEventLogs err %v", err)
	}

	if len(events) > 0 {
		t.Logf("okx1-sepolia len %d ", len(events))
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

func TestGetBsquaredL2BlockNumber(t *testing.T) {
	//url := "https://haven-b2-nodes.bsquared.network"
	l2Url := "https://rpc.merlinchain.io" //"https://haven-rpc.bsquared.network" //"https://roles-rpc.bsquared.network"
	//l1url "https://haven-b2-nodes.bsquared.network"
	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	//block, err := l2Client.GetBlockByNumber(context.Background(), big.NewInt(239250))
	//if err != nil {
	//	t.Errorf("Error GetBlockByNumber : %s ", err)
	//	return
	//}
	//t.Logf("block:%+v", *block)

	blockNumber, err := l2Client.BlockNumber(context.Background())
	if err != nil {
		t.Errorf("Error GetBlockByNumber : %s ", err)
		return
	}

	t.Logf("Latest blockNumber:%+v", blockNumber)
}

// GetCallTrace
func TestGetCallTraceByBlock(t *testing.T) {
	l2Url := "http://alb-merlin-unifra-rpc-1442295568.ap-southeast-1.elb.amazonaws.com"
	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	block, err := l2Client.GetBlockByNumber(context.Background(), big.NewInt(239250))
	if err != nil {
		t.Errorf("Error GetBlockByNumber : %s ", err)
		return
	}
	t.Logf("block:%+v", *block)

	for _, tx := range block.Transactions {
		t.Logf("GetCallTrace for tx:%+v", tx)
		trace, err := l2Client.GetCallTrace(context.Background(), tx.Hash)
		if err != nil {
			t.Errorf("Error GetBlockByNumber : %s ", err)
			return
		}
		t.Logf("trace:%+v", *trace)
	}
}

func TestGetCallTraceByHash(t *testing.T) {
	l2Url := "http://alb-merlin-unifra-rpc-1442295568.ap-southeast-1.elb.amazonaws.com"
	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	hash := common.HexToHash("")
	trace, err := l2Client.GetCallTrace(context.Background(), hash)
	if err != nil {
		t.Errorf("Error GetBlockByNumber : %s ", err)
		return
	}
	t.Logf("trace:%+v", *trace)
}

func TestBalanceAt(t *testing.T) {
	l2Url := "https://rpc.merlinchain.io"
	l2Client, err := NewClient(l2Url)
	assert.NoError(t, err, "Expected no error while creating l1Client")

	//0xaC40c276d4368B03c5d0E56e5FdECfD36A8A1D8b
	//0x055Bc6573BcF10715325C5A15E8C48213Eef2E7D
	account := common.HexToAddress("0x62CCfC7B3d821Dfed6F9C93f1B6DAC9df00e4F8b")
	ba, err := l2Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		t.Errorf("Error BalanceAt : %s ", err)
		return
	}
	t.Logf("Balance :%+v", ba.Int64())
}

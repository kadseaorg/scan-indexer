package ethclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetArbOneSequencerBatchDeliveredEventLogs
func TestGetArbOneSequencerBatchDeliveredEventLogs(t *testing.T) {
	client, err := NewClient("http://10.3.3.94:8545")
	assert.NoError(t, err, "Expected no error while creating client")

	for i := 1; i < 2; i++ {
		start := uint64(19169869)
		end := start + 20
		events, err := client.GetArbOneSequencerBatchDeliveredEventLogs(context.Background(), start, end) //15447810
		if err != nil {
			t.Errorf("GetArbOneSequencerBatchDeliveredEventLogs err %v", err)
		}

		for _, event := range events {
			t.Logf("Arb start %d end %d event: %+v", start, end, event)
		}

		start = end
	}
}

func TestGetArbOneBatchContainingBlock(t *testing.T) {
	client, err := NewClient("http://10.3.3.219:8547")
	assert.NoError(t, err, "Expected no error while creating client")

	batchNumber, err := client.GetArbOneBatchContainingBlock(context.Background(), 178499356) //178497817) //178499356
	if err != nil {
		t.Fatalf("Err GetArbOneBatchContainingBlock err %v", err)
	}

	t.Logf("batchNumber:%d", batchNumber)

}

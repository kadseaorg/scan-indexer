package ethclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetArbOneSequencerBatchDeliveredEventLogs
func TestGetBaseOutputProposedEventLogs(t *testing.T) {
	client, err := NewClient("http://10.3.3.94:8545")
	assert.NoError(t, err, "Expected no error while creating client")

	for i := 1; i < 2; i++ {
		start := uint64(19169927)
		end := start + 20
		events, err := client.GetBaseOutputProposedEventLogs(context.Background(), start, end) //15447810
		if err != nil {
			t.Errorf("GetBaseOutputProposedEventLogs err %v", err)
		}

		for _, event := range events {
			t.Logf("Base start %d end %d event: %+v", start, end, event)
		}

		start = end
	}
}

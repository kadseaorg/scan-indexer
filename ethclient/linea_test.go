package ethclient

//import (
//	"context"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//// GetArbOneSequencerBatchDeliveredEventLogs
//func TestGetLineaFinalizedEventLogss(t *testing.T) {
//	client, err := NewClient("http://10.3.3.94:8545")
//	assert.NoError(t, err, "Expected no error while creating client")
//
//	start := uint64(19239933) // 19211678 19239933)
//	end := start + 20
//	events, err := client.GetLineaBatchEventsFromL1(context.Background(), start, end) //15447810
//	if err != nil {
//		t.Errorf("GetArbOneSequencerBatchDeliveredEventLogs err %v", err)
//	}
//
//	for _, event := range events {
//		t.Logf("Linea start %d end %d event: %+v", start, end, event)
//	}
//
//}

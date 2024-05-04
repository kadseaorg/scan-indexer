package ethclient

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestParseChunk(t *testing.T) {
	chunk, _ := hex.DecodeString("0100000000000000010000000064d08870000000000000000000000000000000000000000000000000000000000000000000000000007a120000000000")
	tests := []struct {
		name       string
		chunkData  []byte
		wantBlocks []BlockContext
		wantErr    bool
	}{
		{
			name:      "valid chunk",
			chunkData: chunk,
			wantBlocks: []BlockContext{
				{
					BlockNumber:     1,
					Timestamp:       1691388016,
					BaseFee:         big.NewInt(0),
					GasLimit:        8000000,
					NumTransactions: 0,
					NumL1Messages:   0,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseChunk(tt.chunkData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseChunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(got.Blocks) != len(tt.wantBlocks) {
				t.Errorf("ParseChunk() got %d blocks, want %d", len(got.Blocks), len(tt.wantBlocks))
				return
			}

			for i, wantBlock := range tt.wantBlocks {
				gotBlock := got.Blocks[i]
				if gotBlock.BlockNumber != wantBlock.BlockNumber {
					t.Errorf("ParseChunk() block %d: got BlockNumber %d, want %d", i, gotBlock.BlockNumber, wantBlock.BlockNumber)
				}
				if gotBlock.Timestamp != wantBlock.Timestamp {
					t.Errorf("ParseChunk() block %d: got Timestamp %d, want %d", i, gotBlock.Timestamp, wantBlock.Timestamp)
				}
				if gotBlock.GasLimit != wantBlock.GasLimit {
					t.Errorf("ParseChunk() block %d: got GasLimit %d, want %d", i, gotBlock.GasLimit, wantBlock.GasLimit)
				}
				if gotBlock.NumTransactions != wantBlock.NumTransactions {
					t.Errorf("ParseChunk() block %d: got NumTransactions %d, want %d", i, gotBlock.NumTransactions, wantBlock.NumTransactions)
				}
				if gotBlock.NumL1Messages != wantBlock.NumL1Messages {
					t.Errorf("ParseChunk() block %d: got NumL1Messages %d, want %d", i, gotBlock.NumL1Messages, wantBlock.NumL1Messages)
				}
				if !bytes.Equal(gotBlock.BaseFee.Bytes(), wantBlock.BaseFee.Bytes()) {
					t.Errorf("ParseChunk() block %d: got BaseFee %v, want %v", i, gotBlock.BaseFee, wantBlock.BaseFee)
				}
			}

			if !bytes.Equal(got.L2Transactions, tt.chunkData[1+len(tt.wantBlocks)*BlockContextLength:]) {
				t.Errorf("ParseChunk() got L2Transactions %v, want %v", got.L2Transactions, tt.chunkData[1+len(tt.wantBlocks)*BlockContextLength:])
			}
		})
	}
}

// GetScrollBatches
func TestGetScrollBatches(t *testing.T) {
	client, err := NewClient("http://10.3.3.234:8545")
	assert.NoError(t, err, "Expected no error while creating client")

	for i := 1; i < 2; i++ {
		start := uint64(5364017)
		end := start + 100
		events, err := client.GetScrollBatches(context.Background(), "0x2D567EcE699Eabe5afCd141eDB7A4f2D0D6ce8a0", start, end) //15447810
		if err != nil {
			t.Errorf("GetBaseOutputProposedEventLogs err %v", err)
		}

		for _, event := range events {
			t.Logf("Base start %d end %d event: %+v", start, end, event)
		}

		start = end
	}
}

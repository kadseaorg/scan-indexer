package ethclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeLineaContractAbi(t *testing.T) {
	ethClient, err := NewClient("https://rpc.ankr.com/eth/0706dd4e96b29b1ed4d1d1e5e8bcc3b14823ac2bb3f429b8f6ff245a03fd1729")
	assert.NoError(t, err)

	ethClient.initLineaContract()
	assert.NotNil(t, lineaContract.newLinea, "EthClient create new linea contract failed")
	assert.NotNil(t, lineaContract.oldLinea, "EthClient create old linea contract failed")
}

func TestGetLineaBatchEventsFromL1AfterUpgrade(t *testing.T) {
	ethClient, err := NewClient("https://rpc.ankr.com/eth/0706dd4e96b29b1ed4d1d1e5e8bcc3b14823ac2bb3f429b8f6ff245a03fd1729")
	assert.NoError(t, err, "create client failed")

	// https://etherscan.io/tx/0xa472633897e4e4c9b100dce2cf87e0299fd8240349a0abff6e4f49831c3b5016 last finalize block linea 2325867.
	// The previous tx https://etherscan.io/tx/0x3108ddb5140e8fca1c05c539327ab3865d7bb0f61531f7188aeca9b2f6e19392 last finalize block 2323209
	fromEth := uint64(19217958)
	endEth := uint64(19217975)
	last_l2_block_number := uint64(2238199) //2323210 - 2325867

	lineaBatchEvent, err := ethClient.GetLineaBatchEventsFromL1(context.Background(), &last_l2_block_number, fromEth, endEth)
	assert.NoError(t, err)

	if len(lineaBatchEvent) == 0 {
		return
	}
	t.Logf("lineaBatchEvent:%+v", lineaBatchEvent[0])

	//assert.Equal(t, len(lineaBatchEvent), 1)
	//assert.Equal(t, last_l2_block_number, uint64(2325867))
	//
	//assert.Equal(t, lineaBatchEvent[0].FinalRootHash, common.HexToHash("0D9E4E5185D1CA40FFA3047DF5273E14DB609CB74F91557B96187E152F5FBB6B"))
	//assert.Equal(t, lineaBatchEvent[0].TransactionHash, common.HexToHash("0xa472633897e4e4c9b100dce2cf87e0299fd8240349a0abff6e4f49831c3b5016"))
	//assert.Equal(t, lineaBatchEvent[0].Timestamp, uint64(1708185251))
	//assert.Equal(t, len(lineaBatchEvent[0].FinalizedBlockNumbers), 2325867-2323209)
	//assert.Equal(t, lineaBatchEvent[0].BlockNumber, fromEth)
}

func TestGetLineaBatchEventsFromL1BeforeUpgrade(t *testing.T) {
	ethClient, err := NewClient("http://10.3.3.94:8545")
	assert.NoError(t, err, "create client failed")

	// https://etherscan.io/tx/0x60b2f0ea6fb61fdca95d75d6e02e5370c616f1dea21fa70743528530e5f9fa3b
	// first finalize block 2238182.
	// last finalize block 2238199
	fromEth := uint64(19217957)
	endEth := uint64(19217957)
	last_l2_block_number := uint64(2238180) //2238182-2238199

	lineaBatchEvent, err := ethClient.GetLineaBatchEventsFromL1(context.Background(), &last_l2_block_number, fromEth, endEth)
	assert.NoError(t, err)

	for _, event := range lineaBatchEvent {
		t.Logf("lineaBatchEvent:%+v", event)
	}

	//fromEth = uint64(19217958)
	//endEth = uint64(19217998)
	//last_l2_block_number = uint64(2238200) //2238182-2238199
	//
	//lineaBatchEvent, err = ethClient.GetLineaBatchEventsFromL1(context.Background(), &last_l2_block_number, fromEth, endEth)
	//assert.NoError(t, err)
	//
	//t.Logf("len(%d) lineaBatchEvent:%+v ", len(lineaBatchEvent), lineaBatchEvent[0])

	//assert.Equal(t, len(lineaBatchEvent), 1)
	//// the old event should't change last_l2_block_number
	//assert.Equal(t, last_l2_block_number, uint64(2238182))
	//
	//assert.Equal(t, lineaBatchEvent[0].FinalRootHash, common.HexToHash("30109245D023028993A1AAED16EE0B975D0EE9EAFB1EA7D0E2BA5688FAE5EEA4"))
	//assert.Equal(t, lineaBatchEvent[0].TransactionHash, common.HexToHash("0x60b2f0ea6fb61fdca95d75d6e02e5370c616f1dea21fa70743528530e5f9fa3b"))
	//assert.Equal(t, lineaBatchEvent[0].Timestamp, uint64(1707813551))
	//assert.Equal(t, len(lineaBatchEvent[0].FinalizedBlockNumbers), 2238199-2238182+1)
	//assert.Equal(t, lineaBatchEvent[0].BlockNumber, fromEth)
}

/// The unit test failed because block range is too wide for free eth rpc.
// func TestGetLineaBatchEventsFromL1BetweenUpgrade(t *testing.T) {
// 	ethClient, err := NewClient("https://rpc.ankr.com/eth/0706dd4e96b29b1ed4d1d1e5e8bcc3b14823ac2bb3f429b8f6ff245a03fd1729")
// 	assert.NoError(t, err, "create client failed")

// 	fromEth := uint64(19217957)
// 	// https://etherscan.io/tx/0x8249932922d2634d6b27624afa572c50f02470c14673823b50ce501a7d300e7f
// 	// the first finalize batch tx after upgrade
// 	endEth := uint64(19222438)
// 	last_l2_block_number := uint64(2238182)

// 	lineaBatchEvent, err := ethClient.GetLineaBatchEventsFromL1(context.Background(), &last_l2_block_number, fromEth, endEth)
// 	assert.NoError(t, err)

// 	assert.Equal(t, len(lineaBatchEvent), 2)
// 	// the old event should't change last_l2_block_number
// 	assert.Equal(t, last_l2_block_number, uint64(2250847))

// 	assert.Equal(t, lineaBatchEvent[0].FinalRootHash, common.HexToHash("30109245D023028993A1AAED16EE0B975D0EE9EAFB1EA7D0E2BA5688FAE5EEA4"))
// 	assert.Equal(t, lineaBatchEvent[0].TransactionHash, common.HexToHash("0x60b2f0ea6fb61fdca95d75d6e02e5370c616f1dea21fa70743528530e5f9fa3b"))
// 	assert.Equal(t, lineaBatchEvent[0].Timestamp, uint64(1707813551))
// 	assert.Equal(t, len(lineaBatchEvent[0].FinalizedBlockNumbers), 2238199-2238182+1)
// 	assert.Equal(t, lineaBatchEvent[0].BlockNumber, fromEth)

// 	assert.Equal(t, lineaBatchEvent[0].FinalRootHash, common.HexToHash("1244273FD6C8822CB888E50A5864A14057B0C93536102A850B82C00231B288B6"))
// 	assert.Equal(t, lineaBatchEvent[0].TransactionHash, common.HexToHash("0x8249932922d2634d6b27624afa572c50f02470c14673823b50ce501a7d300e7f"))
// 	assert.Equal(t, lineaBatchEvent[0].Timestamp, uint64(1708813551))
// 	assert.Equal(t, len(lineaBatchEvent[0].FinalizedBlockNumbers), 2250847-2238199+1)
// 	assert.Equal(t, lineaBatchEvent[0].BlockNumber, fromEth)
// }

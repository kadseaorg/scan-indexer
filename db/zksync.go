package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/unifra20/l2scan-indexer/contract"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (db *Client) UpdateZKSyncBridgeDepositHistory(ctx context.Context, bridgeEvents []ethclient.ZKSDepositEvent) error {
	stmt := `INSERT INTO zksync_bridge_deposit_history (l1_tx_hash, l2_tx_hash, l1_tx_timestamp)
	VALUES ($1, $2, $3) ON CONFLICT (l1_tx_hash)DO UPDATE SET l2_tx_hash = $2, l1_tx_timestamp = $3`
	for _, event := range bridgeEvents {
		_, err := db.pool.Exec(ctx, stmt, event.Raw.TxHash.Hex(), hexutil.Encode(event.TxHash[:]), event.BlockTimestamp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Client) InsertZKSyncWithdrawHistory(ctx context.Context, withdrawInput []contract.WithdrawalFinalizerRequestFinalizeWithdrawal, l1TxHash common.Hash, l1TxTimestamp uint64) error {
	stmt := `INSERT INTO zksync_bridge_withdraw_history (l1_batch_number, l1_batch_tx_index, l1_tx_hash, l1_tx_timestamp) VALUES ($1, $2, $3, $4)`
	for _, input := range withdrawInput {
		//! L2BlockNumber means l1_batch_number
		//! L2TxNumberInBlock means l1_batch_tx_index
		_, err := db.pool.Exec(ctx, stmt, input.L2BlockNumber, input.L2TxNumberInBlock, l1TxHash.Hex(), l1TxTimestamp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Client) InsertZKSyncETHWithdrawHistory(ctx context.Context, withdrawInput ethclient.ETHWithdrawalFinalizerData, l1TxHash common.Hash, l1TxTimestamp uint64) error {
	stmt := `INSERT INTO zksync_bridge_withdraw_history (l1_batch_number, l1_batch_tx_index, l1_tx_hash, l1_tx_timestamp) VALUES ($1, $2, $3, $4)`
	//! L2BlockNumber means l1_batch_number
	//! L2TxNumberInBlock means l1_batch_tx_index
	_, err := db.pool.Exec(ctx, stmt, withdrawInput.L2BlockNumber, withdrawInput.L2TxNumberInBlock, l1TxHash.Hex(), l1TxTimestamp)
	if err != nil {
		return err
	}
	return nil
}

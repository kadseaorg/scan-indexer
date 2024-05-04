package indexer

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/alitto/pond"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (i *Indexer) StartL1IndexerForZKSyncEra(worker int, forceStartBatchNumber *uint64) {
	log.Info().Msg("🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢 Starting L1 Indexer 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢")

	// Start from an initial batch, this could either be a batch forced from an argument or the latest indexed
	var startBatchNumber uint64
	if forceStartBatchNumber != nil {
		log.Warn().Msgf("Force starting L1 indexer from batch %d", *forceStartBatchNumber)
		startBatchNumber = *forceStartBatchNumber
	} else {
		var err error
		startBatchNumber, err = i.dbClient.GetLatestBatchNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last indexed L1 batch number: %s", err)
			return
		}
	}

	// Pool setup
	pool := pond.New(worker, worker*2) // worker*2 is the max capacity of tasks in queue

	for {
		// Get latest batch number each time through the loop
		latestBatchNumber, err := i.l2client.GetZKSL1BatchNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting latest zksync era L1 batch number: %s", err)
			time.Sleep(time.Second * 5)
			continue
		}

		// Process new batches, if any
		for n := startBatchNumber + 1; n <= latestBatchNumber; n++ {
			batchNumber := n

			// Asynchronously index each batch
			pool.Submit(func() {
				err := i.indexZKSEraL1Batch(batchNumber)
				if err != nil {
					log.Error().Msgf("Error indexing zksync era L1 batch %d: %s", batchNumber, err)
					os.Exit(-1)
				}
			})
		}

		// Update start block for next iteration
		startBatchNumber = latestBatchNumber

		// Wait and repeat
		time.Sleep(time.Second * 5)
	}

	// This will never happen, but it's here for completeness
	pool.StopAndWait()
}

func (i *Indexer) indexZKSEraL1Batch(batchNumber uint64) error {
	startTime := time.Now()

	batch, err := i.l2client.GetZKSL1BatchDetails(i.ctx, batchNumber)
	if err != nil {
		return fmt.Errorf("error getting zks L1 batch %d: %s", batchNumber, err)
	}

	//exist, err := i.dbClient.IsHashRootExist(i.ctx, batch.RootHash)
	//if err != nil {
	//	return fmt.Errorf("error zks check hash_root, batch %d: %s", batchNumber, err)
	//}
	//
	//if exist {
	//	log.Info().Msgf("🟢 [L1 Indexer] Indexed batch %d, hash_root already exists, skipping", batchNumber)
	//	return nil
	//}

	err = i.dbClient.UpsertL1Batch(i.ctx, batch)
	if err != nil {
		return fmt.Errorf("error upserting zks L1 batch %d: %s", batchNumber, err)
	}

	elapsedTime := time.Since(startTime).Milliseconds()

	log.Info().Msgf("🟢 [L1 Indexer] Indexed batch %d, status: %s, l2_txns_count: %d, time taken: %dms", batchNumber, batch.Status, batch.L2TxCount, elapsedTime)

	return nil
}

func (i *Indexer) StartL1BridgeDepositIndexerForZKSyncEra(worker int, forceStartBlock *uint64) {
	if i.chain.Name != ZkSyncEra.Name && i.chain != ZkSyncEraSepolia {
		return
	}

	log.Info().Msg("🌉🌉🌉🌉🌉🌉🌉🌉🌉🌉 Starting L1 Bridge Indexer[deposit] 🌉🌉🌉🌉🌉🌉🌉🌉🌉🌉")
	var lastSyncedBlockNumber uint64
	var err error

	if forceStartBlock != nil {
		log.Warn().Msgf("🌉 Force starting L1 bridge indexer from block %d", *forceStartBlock)
		lastSyncedBlockNumber = *forceStartBlock
	} else {
		lastSyncedBlockNumber, err = i.dbClient.GetSyncProgress(i.ctx, "l1_bridge_deposit_history_synced_block_number")
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 block number: %s", err)
			return
		}

	}

	const eventStep = 50
	for {
		startBlock := lastSyncedBlockNumber + 1
		for {
			latestBlockNumber, err := i.l1client.BlockNumber(i.ctx)
			if err != nil {
				// exit and let the process restart
				log.Error().Msgf("Error getting latest block number: %s", err)
				os.Exit(-1)
			}

			// reset startBlock if it's greater than latestBlockNumber
			if startBlock > latestBlockNumber {
				startBlock = latestBlockNumber
			}
			endBlock := startBlock + eventStep
			if endBlock > latestBlockNumber {
				endBlock = latestBlockNumber
			}

			eventsLen, err := i.indexL1BridgeDepositEventsForZKSync(startBlock, endBlock)
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
					// has indexed, skip this block range
					log.Info().Msgf("🤷 [L1 Bridge Indexer] Already indexed range: [%d, %d]", startBlock, endBlock)
					startBlock = endBlock + 1
					continue
				}
				log.Error().Msgf("Error indexing L1 bridge for range: [%d, %d]: %s", startBlock, endBlock, err)
				break
			}

			log.Info().Msgf("🌉 [L1 Bridge Indexer[deposit]] Indexed bridge events range: [%d, %d], len: %d", startBlock, endBlock, eventsLen)

			if endBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 Bridge Indexer[deposit]] Have indexed all bridge up to block %d", latestBlockNumber)
				break
			}

			startBlock = endBlock + 1
		}
		lastSyncedBlockNumber, err = i.dbClient.GetSyncProgress(i.ctx, "l1_bridge_deposit_history_synced_block_number")
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 bridge block number in loop: %s", err)
		}
		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
	}
}

func (i *Indexer) indexL1BridgeDepositEventsForZKSync(from uint64, to uint64) (int, error) {
	bridgeEvents, err := i.l1client.GetZKSL1DepositEvents(i.ctx, i.chain.DiamondProxy, from, to)
	if err != nil {
		return 0, fmt.Errorf("error getting L1 bridge events: %w", err)
	}

	if len(bridgeEvents) == 0 {
		return 0, nil
	}

	err = i.dbClient.UpdateZKSyncBridgeDepositHistory(i.ctx, bridgeEvents)
	if err != nil {
		return 0, fmt.Errorf("error upserting L1 bridge events: %w", err)
	}

	err = i.dbClient.UpdateSyncProgress(i.ctx, "l1_bridge_deposit_history_synced_block_number", to)
	if err != nil {
		return 0, fmt.Errorf("error upserting L1 bridge sync progress: %w", err)
	}

	return len(bridgeEvents), nil
}

func (i *Indexer) StartL1BridgeWithdrawIndexerForZKSyncEra(worker int, forceStartBlock *uint64) {
	if i.chain.Name != ZkSyncEra.Name && i.chain != ZkSyncEraSepolia {
		return
	}

	log.Info().Msg("🌉🌉🌉🌉🌉🌉🌉🌉🌉🌉 Starting L1 Bridge Indexer[withdraw] 🌉🌉🌉🌉🌉🌉🌉🌉🌉🌉")
	var lastSyncedBlockNumber uint64
	var err error

	if forceStartBlock != nil {
		log.Warn().Msgf("[L1 Bridge Indexer[withdraw]] Force starting L1 indexer from block %d", *forceStartBlock)
		lastSyncedBlockNumber = *forceStartBlock
	} else {
		lastSyncedBlockNumber, err = i.dbClient.GetSyncProgress(i.ctx, "l1_bridge_withdraw_history_synced_block_number")
		if err != nil {
			log.Error().Msgf("[L1 Bridge Indexer[withdraw]] Error getting last synced L1 block number: %s", err)
			return
		}

	}

	for {
		startBlock := lastSyncedBlockNumber + 1
		for {
			latestBlockNumber, err := i.l1client.BlockNumber(i.ctx)
			if err != nil {
				// exit and let the process restart
				log.Error().Msgf("[L1 Bridge Indexer[withdraw]] Error getting latest block number: %s", err)
				os.Exit(-1)
			}

			// reset startBlock if it's greater than latestBlockNumber
			if startBlock > latestBlockNumber {
				startBlock = latestBlockNumber
			}

			err = i.indexL1BridgeWithdrawBlockForZKSync(big.NewInt(int64(startBlock)))
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
					// has indexed, skip this block range
					log.Info().Msgf("🤷 [L1 Bridge Indexer[withdraw]] Already indexed block: [%d]", startBlock)
					startBlock = startBlock + 1
					continue
				}
				log.Error().Msgf("Error indexing L1 bridge withdraw for block: [%d]: %s", startBlock, err)
				break
			}

			// log.Info().Msgf("🌉 🚀 [L1 Bridge Indexer[withdraw]] Indexed bridge block: [%d]", startBlock)

			if startBlock >= latestBlockNumber {
				log.Info().Msgf("👻 [L1 Bridge Indexer[withdraw]] Have indexed all bridge up to block %d", latestBlockNumber)
				break
			}

			startBlock = startBlock + 1
		}
		lastSyncedBlockNumber, err = i.dbClient.GetSyncProgress(i.ctx, "l1_bridge_withdraw_history_synced_block_number")
		if err != nil {
			log.Error().Msgf("Error getting last synced L1 bridge block number in loop: %s", err)
		}
		time.Sleep(5 * time.Second) // wait for 5 seconds before checking for new blocks
	}
}

func (i *Indexer) indexL1BridgeWithdrawBlockForZKSync(blockNumber *big.Int) error {
	// startTime := time.Now()

	block, err := i.l1client.GetBlockByNumber(i.ctx, blockNumber)
	if err != nil {
		return fmt.Errorf("error getting block %d: %s", blockNumber, err)
	}

	for _, tx := range block.Transactions {
		if err := i.indexL1BridgeWithdrawInput(tx, block.Timestamp.ToInt().Uint64()); err != nil {
			return fmt.Errorf("error indexing transaction %s: %w", tx.Hash.Hex(), err)
		}
	}

	// elapsedTime := time.Since(startTime).Milliseconds()

	err = i.dbClient.UpdateSyncProgress(i.ctx, "l1_bridge_withdraw_history_synced_block_number", blockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("error upserting L1 bridge sync progress: %w", err)
	}

	// log.Info().Msgf("🌉 [L2 Bridge Withdraw Indexer] Indexed block %d: txns count: %d, time taken: %dms", blockNumber, len(block.Transactions), elapsedTime)

	return nil
}

func (i *Indexer) indexL1BridgeWithdrawInput(tx *ethclient.RpcTransaction, blockTimestramp uint64) error {
	if tx.To == nil || (tx.To.Hex() != i.chain.WithdrawalFinalizer && tx.To.Hex() != i.chain.DiamondProxy) {
		// log.Debug().Msgf("🤷 [L2 Bridge Withdraw Indexer] Transaction %s is not a withdraw input, skipping...", tx.Hash.Hex())
		return nil
	}

	if tx.To.Hex() == i.chain.WithdrawalFinalizer {
		withdrawInput, err := ethclient.ParseWithdrawalFinalizationInput(tx.Input.String())
		if err != nil {
			if errors.Is(err, ethclient.ErrWithdrawalFinalizationRequestNotFound) {
				log.Debug().Msgf("🤷 [L2 Bridge Withdraw Indexer] Transaction %s is not a withdraw input, skipping...", tx.Hash.Hex())
				return nil
			}
			return fmt.Errorf("error parsing withdraw input: %w", err)
		}

		err = i.dbClient.InsertZKSyncWithdrawHistory(i.ctx, withdrawInput, tx.Hash, blockTimestramp)
		if err != nil {
			return fmt.Errorf("error inserting withdraw input: %w", err)
		}
	} else if tx.To.Hex() == i.chain.DiamondProxy {
		withdrawInput, err := ethclient.ParseETHWithdrawalFinalizationInput(tx.Input.String())
		if err != nil {
			if errors.Is(err, ethclient.ErrWithdrawalFinalizationRequestNotFound) {
				log.Debug().Msgf("🤷 [L2 Bridge Withdraw Indexer] Transaction %s is not a withdraw input, skipping...", tx.Hash.Hex())
				return nil
			}
			return fmt.Errorf("error parsing withdraw input: %w", err)
		}

		if withdrawInput.L2BlockNumber == nil {
			return nil
		}

		err = i.dbClient.InsertZKSyncETHWithdrawHistory(i.ctx, withdrawInput, tx.Hash, blockTimestramp)
		if err != nil {
			return fmt.Errorf("error inserting withdraw input: %w", err)
		}
	}

	log.Info().Msgf("🌉 [L1 Bridge Indexer[withdraw]] Indexed withdraw tx %s", tx.Hash.Hex())

	return nil
}

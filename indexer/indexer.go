package indexer

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"

	"github.com/alitto/pond"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/contract"
	"github.com/unifra20/l2scan-indexer/db"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

// eventHash represents an event keccak256 hash
type eventHash string

const (
	// transferEventHash represents the keccak256 hash of Transfer(address,address,uint256)
	transferEventHash eventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// transferSingleEventHash represents the keccak256 hash of TransferSingle(address,address,address,uint256,uint256)
	transferSingleEventHash eventHash = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	// transferBatchEventHash represents the keccak256 hash of TransferBatch(address,address,address,uint256[],uint256[])
	transferBatchEventHash eventHash = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
)

const (
	ERC20   string = "erc20"
	ERC721  string = "erc721"
	ERC1155 string = "erc1155"
)

type Indexer struct {
	l1client            *ethclient.Client
	l2client            *ethclient.Client
	debugClient         *ethclient.Client
	dbClient            *db.Client
	erc20Client         *contract.IERC20
	erc721Client        *contract.IERC721
	erc1155Client       *contract.IERC1155
	chain               Chain
	indexedTokens       *lru.Cache[string, string]
	sanitizationPolicy  *bluemonday.Policy
	ctx                 context.Context
	latestL2BlockNumber uint64
}

func NewIndexer(chain Chain, l1client, l2client, debugClient *ethclient.Client, dbClient *db.Client) *Indexer {
	erc20Client, _ := contract.NewIERC20(common.Address{}, l2client)
	erc721Client, _ := contract.NewIERC721(common.Address{}, l2client)
	erc1155Client, _ := contract.NewIERC1155(common.Address{}, l2client)

	// timeoutCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cache, err := lru.New[string, string](1024 * 1024 * 1024 * 4) // 1024 * 1024 * 1024 * 2 = 2GB
	if err != nil {
		log.Fatal().Msgf("Error creating cache: %s", err)
		os.Exit(1)
	}

	var sanitizationPolicy = bluemonday.UGCPolicy()

	latestBlockNumber, err := dbClient.GetLatestBlockNumber(context.Background())
	if err != nil {
		log.Fatal().Msgf("Error get latest blockNumber from db : %s", err)
		os.Exit(1)
	}

	return &Indexer{
		l1client:            l1client,
		l2client:            l2client,
		debugClient:         debugClient,
		dbClient:            dbClient,
		erc20Client:         erc20Client,
		erc721Client:        erc721Client,
		erc1155Client:       erc1155Client,
		chain:               chain,
		indexedTokens:       cache,
		sanitizationPolicy:  sanitizationPolicy,
		ctx:                 context.Background(),
		latestL2BlockNumber: latestBlockNumber,
	}
}

func (i *Indexer) Start(ctx context.Context,
	worker int,
	l1ForceStartBlock *uint64,
	l2ForceStartBlock *uint64,
	checkMisBatchedBlocks bool,
	l1BridgeForceStartBlock *uint64,
	inscriptionForceStartBlock *uint64,
	recoveryStartBlock *uint64,
) {
	go i.StartMultiL1Chain(worker, l1ForceStartBlock, l2ForceStartBlock)
	go i.StartL2Indexer(worker, l2ForceStartBlock)
	go i.L2RecoveryExecutor(worker, checkMisBatchedBlocks)
	go i.RecoverL2MissingTraces(ctx, worker)
	go i.RecoverZKSyncMissingBatches(worker)
	go i.RecoverL1BatchesForArbitrumOne(worker, recoveryStartBlock)
	// go i.StartL1RecoveryForOKX1Sepolia(worker)
	go i.ScheduleZKSL1StatusUpdate(worker)
	go i.ScheduleArbL1StatusUpdate(worker, l2ForceStartBlock)
	go i.CheckAndUpdateLineaUnfinalizedL2Blocks()
	go i.StartL1BridgeDepositIndexerForZKSyncEra(worker, l1BridgeForceStartBlock)
	go i.StartL1BridgeWithdrawIndexerForZKSyncEra(worker, l1BridgeForceStartBlock)
	go i.StartInscriptionIndexer(worker, inscriptionForceStartBlock)
	go i.L2MonitorExecutor(worker)

	<-ctx.Done()
}

func (i *Indexer) StartMultiL1Chain(worker int, l1ForceStartBlock *uint64, l2ForceStartBlock *uint64) {
	switch i.chain.Name {
	case ZkSyncEra.Name, ZkSyncEraSepolia.Name:
		i.StartL1IndexerForZKSyncEra(worker, l1ForceStartBlock)
	case Linea.Name:
		i.StartL1IndexerForLinea(worker, l1ForceStartBlock)

	case Scroll.Name:
		i.StartL1IndexerForScroll(worker, l1ForceStartBlock)

	case ScrollSepolia.Name:
		i.StartL1IndexerForScroll(worker, l1ForceStartBlock)

	case Base.Name:
		i.StartL1IndexerForBase(worker, l1ForceStartBlock)

	case MantaPacific.Name:
		i.StartL1IndexerForMantaPacific(worker, l1ForceStartBlock)

	case ArbitrumOne.Name:
		i.StartL1IndexerForArbitrumOne(worker, l1ForceStartBlock)

	case BsquaredTestnet.Name:
		i.StartL1IndexerForPolygonZkEVM(worker, l1ForceStartBlock)

	case PolygonZkEVM.Name:
		i.StartL1IndexerForPolygonZkEVM(worker, l1ForceStartBlock)

	case OKX1Sepolia.Name:
		i.StartL1IndexerForOKX1Sepolia(worker, l1ForceStartBlock)

	case OroTestnet.Name:
		i.StartL1IndexerForOro(worker, l1ForceStartBlock)
	case Kadsea.Name, KadseaTestnet.Name:
		return

	default:
		log.Fatal().Msgf("Unsupported chain: %s", i.chain)
	}
}

func (i *Indexer) StartL2Indexer(worker int, forceStartBlock *uint64) {
	log.Info().Msg("🔵🔵🔵🔵🔵🔵🔵🔵🔵🔵 Starting L2 Indexer 🔵🔵🔵🔵🔵🔵🔵🔵🔵🔵")

	var lastIndexedBlockNumber uint64
	var err error
	if forceStartBlock != nil {
		log.Warn().Msgf("Force starting L2 indexer from block %d", *forceStartBlock)
		lastIndexedBlockNumber = *forceStartBlock
	} else {
		lastIndexedBlockNumber, err = i.dbClient.GetLatestBlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting last indexed block number: %s", err)
			return
		}
	}

	if i.chain == ArbitrumOne {
		if lastIndexedBlockNumber <= NitroGensisBlockNum {
			lastIndexedBlockNumber = NitroGensisBlockNum + 1
		}
	}
	// Create a worker pool with a capacity of n-worker tasks at a time
	pool := pond.New(worker, worker*2) // worker*2 is the max capacity of tasks in queue

	for {
		latestBlockNumber, err := i.l2client.BlockNumber(i.ctx)
		if err != nil {
			log.Error().Msgf("Error getting latest block number: %s", err)
			os.Exit(-1)
		}

		latestBlockNumber = latestBlockNumber - 10

		log.Info().Msgf("L2 Indexer: lastIndexedBlockNumber %d latestBlockNumber %d", lastIndexedBlockNumber, latestBlockNumber)

		if latestBlockNumber > lastIndexedBlockNumber {
			for n := lastIndexedBlockNumber; n < latestBlockNumber; n++ {
				blockNumber := big.NewInt(int64(n)) // Create a new variable to avoid data race
				pool.Submit(func() {
					i.handleBlock(blockNumber)
				})
			}
			lastIndexedBlockNumber = latestBlockNumber
		}

		time.Sleep(3 * time.Second) // wait for 5 seconds
	}

	// pool.StopAndWait()
}

func (i *Indexer) handleBlock(blockNumber *big.Int) {
	if err := i.indexBlock(blockNumber); err != nil {
		log.Error().Msgf("Error indexing block %d: %s", blockNumber, err)
	}
}

func (i *Indexer) indexBlock(blockNumber *big.Int) error {
	startTime := time.Now()

	block, err := i.l2client.GetBlockByNumber(i.ctx, blockNumber)
	if err != nil {
		return fmt.Errorf("error getting block %d: %s", blockNumber, err)
	}

	pgxBatch := &pgx.Batch{}
	internalTxnCount, traceChecked, err := i.IndexTransactions(pgxBatch, block)
	if err != nil {
		return err
	}

	i.dbClient.UpsertBlock(pgxBatch, block, traceChecked, internalTxnCount)

	// write batch to db
	if err := i.dbClient.ExecBatch(pgxBatch); err != nil {
		// if bulk insert fails, try to rollback the block
		if err := i.dbClient.DeleteBlock(i.ctx, blockNumber.Uint64()); err != nil {
			return fmt.Errorf("error rolling back block %d: %w", blockNumber, err)
		}
		return fmt.Errorf("error executing batch: %w", err)
	}

	//Put it into memory instead of getting it from db；
	//If the updated blockNumber is smaller than the current latestL2BlockNumber in memory,
	//it is considered to be the height of the recovery history and will not be updated.
	//todo If the block is rolled back, do we need to consider
	if blockNumber.Uint64() >= i.latestL2BlockNumber-1 {
		i.latestL2BlockNumber = blockNumber.Uint64()
	}
	elapsedTime := time.Since(startTime).Milliseconds()

	log.Info().Msgf("🔵 [L2 Indexer] Indexed block %d: txns count: %d, internal txns count: %d, trace checked: %t, elapsed time: %d ms", blockNumber, len(block.Transactions), internalTxnCount, traceChecked, elapsedTime)

	go i.AddBatchNumberToL2blocksArbOne(blockNumber)

	return nil
}

func (i *Indexer) IndexTransactions(pgxBatch *pgx.Batch, block *ethclient.RpcBlock) (internalTxnCount int, traceChecked bool, err error) {
	traceChecked = true
	for _, tx := range block.Transactions {
		traceCount, err := i.indexTransaction(pgxBatch, block, tx)
		if err != nil {
			if strings.Contains(err.Error(), ErrGettingCallTrace) {
				log.Warn().Msgf("err getting call trace for transaction %s: %s, skipping", tx.Hash.Hex(), err)
				traceChecked = false
				continue
			}
			return 0, false, fmt.Errorf("error indexing transaction %s: %w", tx.Hash.Hex(), err)
		}
		internalTxnCount += traceCount
	}
	return internalTxnCount, traceChecked, nil
}

func (i *Indexer) indexInternalTransaction(pgxBatch *pgx.Batch, tx *ethclient.RpcTransaction, block *ethclient.RpcBlock, timeoutSeconds int) (int, error) {
	timeoutCtx := i.ctx
	if timeoutSeconds > 0 {
		var cancel context.CancelFunc
		timeoutCtx, cancel = context.WithTimeout(i.ctx, time.Duration(timeoutSeconds)*time.Second)
		defer cancel()
	}

	var callTrace *ethclient.CallTrace
	var err error
	if i.debugClient != nil {
		callTrace, err = i.debugClient.GetCallTrace(timeoutCtx, tx.Hash)
	} else {
		callTrace, err = i.l2client.GetCallTrace(timeoutCtx, tx.Hash)
	}
	if err != nil {
		return 0, fmt.Errorf("[%s] on transaction %s: %w", ErrGettingCallTrace, tx.Hash.Hex(), err)
	}

	if callTrace == nil {
		// zksync will return nil call trace for early transactions
		return 0, nil
	}

	// index call trace(internal transactions)
	// Skip for the zksync for the time being because it will incur a big database.
	// We need to figure out how to optimize it.
	var internalTxCount *int
	if i.chain.Name != ZkSyncEra.Name && i.chain.Name != ZkSyncEraSepolia.Name && i.chain.Name != OKX1Sepolia.Name {
		internalTxCount = new(int)
		i.indexCallTrace(pgxBatch, callTrace, tx.Hash, block, internalTxCount)
	}

	// index tracer contract creation
	i.indexTracerContractCreation(pgxBatch, callTrace, tx, block.Timestamp)

	if internalTxCount != nil {
		return *internalTxCount, nil
	}

	return 0, nil
}

func (i *Indexer) logsToTransfers(ctx context.Context, logs []*types.Log, block *ethclient.RpcBlock, input hexutil.Bytes) []db.TokenTransfer {
	var transfers []db.TokenTransfer
	var methodID string
	if len(input) >= 4 {
		methodID = hexutil.Encode(input[:4])
	}

	for _, txLog := range logs {
		if txLog.Topics == nil || len(txLog.Topics) == 0 {
			continue
		}

		switch {
		case strings.EqualFold(txLog.Topics[0].Hex(), string(transferEventHash)):
			if len(txLog.Topics) < 3 {
				continue
			}

			if len(txLog.Topics) == 3 {
				// ERC20 transfer
				event, err := i.erc20Client.ParseTransfer(*txLog)
				if err != nil {
					log.Error().Msgf("Error parsing ERC20 transfer event: %s", err)
					continue
				}
				transfers = append(transfers, db.TokenTransfer{
					TransactionHash: txLog.TxHash,
					LogIndex:        int(txLog.Index),
					MethodID:        methodID,
					TokenAddress:    txLog.Address,
					From:            event.From,
					To:              event.To,
					Value:           event.Value,
					TokenType:       ERC20,
					BlockNumber:     block.Number.ToInt(),
					BlockHash:       block.Hash,
				})
			}

			if len(txLog.Topics) == 4 {
				// ERC721 transfer
				event, err := i.erc721Client.ParseTransfer(*txLog)
				if err != nil {
					log.Error().Msgf("Error parsing ERC721 transfer event: %s", err)
					continue
				}
				transfers = append(transfers, db.TokenTransfer{
					TransactionHash: txLog.TxHash,
					LogIndex:        int(txLog.Index),
					MethodID:        methodID,
					TokenAddress:    txLog.Address,
					From:            event.From,
					To:              event.To,
					Amount:          big.NewInt(1),
					TokenID:         event.Id,
					TokenType:       ERC721,
					BlockNumber:     block.Number.ToInt(),
					BlockHash:       block.Hash,
				})
			}

		case strings.EqualFold(txLog.Topics[0].Hex(), string(transferSingleEventHash)):
			// ERC1155 transferSingle
			if len(txLog.Topics) < 4 {
				continue
			}

			event, err := i.erc1155Client.ParseTransferSingle(*txLog)
			if err != nil {
				log.Error().Msgf("Error indexing receipt with reverror parsing ERC1155 transferSingle event: %s", err)
				continue
			}
			transfers = append(transfers, db.TokenTransfer{
				TransactionHash: txLog.TxHash,
				LogIndex:        int(txLog.Index),
				MethodID:        methodID,
				TokenAddress:    txLog.Address,
				From:            event.From,
				To:              event.To,
				Amount:          event.Value,
				TokenID:         event.Id,
				TokenType:       ERC1155,
				BlockNumber:     block.Number.ToInt(),
				BlockHash:       block.Hash,
			})

		case strings.EqualFold(txLog.Topics[0].Hex(), string(transferBatchEventHash)):
			// ERC1155 transferBatch
			if len(txLog.Topics) < 4 {
				continue
			}

			event, err := i.erc1155Client.ParseTransferBatch(*txLog)
			if err != nil {
				log.Error().Msgf("Error parsing ERC1155 transferBatch event: %s", err)
				continue
			}
			for j := 0; j < len(event.Ids); j++ {
				transfers = append(transfers, db.TokenTransfer{
					TransactionHash: txLog.TxHash,
					LogIndex:        int(txLog.Index),
					TokenAddress:    txLog.Address,
					From:            event.From,
					To:              event.To,
					Amount:          event.Values[j],
					TokenID:         event.Ids[j],
					TokenType:       ERC1155,
					BlockNumber:     block.Number.ToInt(),
					BlockHash:       block.Hash,
				})
			}
		}
	}

	return transfers
}

// filter and unique token address from token transfers
func filterTokenAddress(tokenTransfers []db.TokenTransfer) map[common.Address]string {
	result := make(map[common.Address]string)
	unique := make(map[common.Address]bool)

	for _, transfer := range tokenTransfers {
		address := transfer.TokenAddress
		tokenType := transfer.TokenType

		if !unique[address] {
			unique[address] = true
			result[address] = tokenType
		}
	}

	return result
}

func (i *Indexer) getBalance(tokenType string, tokenAddress common.Address, account common.Address, tokenID *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	switch tokenType {
	case ERC20:
		erc20, err := contract.NewIERC20(tokenAddress, i.l2client)
		if err != nil {
			return nil, fmt.Errorf("error creating new ERC20 contract: %w", err)
		}
		return erc20.BalanceOf(opts, account)
	case ERC721:
		erc721, err := contract.NewIERC721(tokenAddress, i.l2client)
		if err != nil {
			return nil, fmt.Errorf("error creating new ERC721 contract: %w", err)
		}
		return erc721.BalanceOf(opts, account)
	case ERC1155:
		erc1155, err := contract.NewIERC1155(tokenAddress, i.l2client)
		if err != nil {
			return nil, fmt.Errorf("error creating new ERC1155 contract: %w", err)
		}
		return erc1155.BalanceOf(opts, account, tokenID)
	default:
		return nil, fmt.Errorf("unsupported token type: %s", tokenType)
	}
}

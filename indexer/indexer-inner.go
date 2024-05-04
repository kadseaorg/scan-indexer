package indexer

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"
	"github.com/unifra20/l2scan-indexer/contract"

	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/db"
	"github.com/unifra20/l2scan-indexer/ethclient"
)

func (i *Indexer) indexTransaction(pgxBatch *pgx.Batch, block *ethclient.RpcBlock, tx *ethclient.RpcTransaction) (internalTxnCount int, err error) {
	receipt, err := i.l2client.TransactionReceipt(i.ctx, tx.Hash)
	if err != nil {
		return 0, fmt.Errorf("error getting receipt for transaction: %w", err)
	}

	if !tx.GasPrice.ToInt().IsInt64() {
		tx.GasPrice = (*hexutil.Big)(receipt.EffectiveGasPrice)
	}

	i.dbClient.InsertTransaction(pgxBatch, block, tx)
	i.indexLogs(pgxBatch, receipt.Logs)

	err = i.indexReceiptWithRevertReason(pgxBatch, tx, receipt, block.Number.ToInt())
	if err != nil {
		return 0, fmt.Errorf("error indexing receipt with revert reason for transaction: %w", err)
	}

	tokenTransfers := i.indexTokenTransfers(pgxBatch, receipt.Logs, block, tx.Input)

	err = i.indexAddressBalances(pgxBatch, tx, block.Number.ToInt())
	if err != nil {
		return 0, fmt.Errorf("error indexing address balances for transaction: %w", err)
	}

	err = i.indexTokenBalance(pgxBatch, tokenTransfers)
	if err != nil {
		return 0, fmt.Errorf("error indexing token balances for transaction: %w", err)
	}

	err = i.indexTokenInfo(pgxBatch, tokenTransfers)
	if err != nil {
		return 0, fmt.Errorf("error indexing token info for transaction: %w", err)
	}

	// we still index null to address for contract creation because some chain can't use debug_traceTransaction with callTracer(for index create/create2 instructions),
	// and we can handle duplicate contract creation tx in db correctly
	err = i.indexNullToAddressContractCreation(pgxBatch, tx, receipt, block.Timestamp)
	if err != nil {
		return 0, fmt.Errorf("error indexing contract creation for transaction: %w", err)
	}

	// ! internal call trace indexer must be the last one to be indexed, because we can ignore trace error in the sequence indexer and recover it later
	if i.chain == ArbitrumOne && block.Number.ToInt().Uint64() <= 22207815 {
		err = i.indexArbOneClassicInternalTransaction(tx, block)
	} else {
		internalTxnCount, err = i.indexInternalTransaction(pgxBatch, tx, block, 0)
	}

	if err != nil {
		return 0, fmt.Errorf("error indexing internal transaction for transaction: %w", err)
	}

	// log.Info().Msgf("Indexed transaction %s", tx.Hash.Hex())
	return internalTxnCount, nil
}

func (i *Indexer) indexLogs(pgxBatch *pgx.Batch, logs []*types.Log) {
	i.dbClient.InsertLogs(pgxBatch, logs)
}

func (i *Indexer) indexReceiptWithRevertReason(pgxBatch *pgx.Batch, tx *ethclient.RpcTransaction, receipt *ethclient.Receipt, blockNumber *big.Int) error {
	// get revert reason if transaction failed
	var revertReason *string
	if receipt.Status == types.ReceiptStatusFailed {
		r, err := i.l2client.GetRevertReason(i.ctx, tx, blockNumber)
		if err != nil {
			return fmt.Errorf("error getting revert reason: %w", err)
		}
		revertReason = &r
	}

	if revertReason != nil {
		// strings to valid utf8
		r := strings.ToValidUTF8(*revertReason, "")
		// Remove null byte sequences
		r = strings.ReplaceAll(r, "\x00", "")
		revertReason = &r
	}

	i.dbClient.UpdateTransactionWithReceipt(pgxBatch, receipt, revertReason)

	return nil
}

func (i *Indexer) indexTokenTransfers(pgxBatch *pgx.Batch, logs []*types.Log, block *ethclient.RpcBlock, input hexutil.Bytes) []db.TokenTransfer {
	tokenTransfers := i.logsToTransfers(i.ctx, logs, block, input)
	i.dbClient.InsertTokenTransfers(pgxBatch, tokenTransfers, block)
	return tokenTransfers
}

func (i *Indexer) indexAddressBalances(pgxBatch *pgx.Batch, tx *ethclient.RpcTransaction, blockNumber *big.Int) error {
	if err := i.updateAddressBalance(pgxBatch, tx.From, blockNumber); err != nil {
		return fmt.Errorf("error updating sender balance for address %s at block %d: %w", tx.From.Hex(), blockNumber, err)
	}

	recipient := tx.To
	if recipient != nil {
		if err := i.updateAddressBalance(pgxBatch, *recipient, blockNumber); err != nil {
			return fmt.Errorf("error updating recipient balance for address %s at block %d: %w", recipient.Hex(), blockNumber, err)
		}
	}

	return nil
}

func (i *Indexer) updateAddressBalance(pgxBatch *pgx.Batch, address common.Address, blockNumber *big.Int) error {
	balance, err := i.l2client.BalanceAt(i.ctx, address, nil)
	if err != nil {
		if strings.Contains(err.Error(), "failed to get balance from state") {
			// Kadsea chain will return this error occasionally, we can ignore it, bc our service fetch balance from rpc directly
			log.Warn().Msgf("🔵 [L2 Indexer] get balance on block %d addr %s, failed to get balance from state",
				blockNumber.Int64(), address)
		} else {
			return fmt.Errorf("error getting balance for address %s: %w", address.Hex(), err)
		}
	}

	i.dbClient.UpsertAddressBalance(pgxBatch, address, balance, blockNumber)

	return nil
}

func (i *Indexer) indexCallTrace(pgxBatch *pgx.Batch, callTrace *ethclient.CallTrace, txHash common.Hash, block *ethclient.RpcBlock, internalTraceCount *int) {
	internalTx := db.InternalTransaction{
		BlockNumber:           block.Number,
		BlockHash:             block.Hash,
		ParentTransactionHash: txHash,
		Type:                  callTrace.Type,
		From:                  callTrace.From,
		To:                    callTrace.To,
		Value:                 callTrace.Value,
		Gas:                   callTrace.Gas,
		GasUsed:               callTrace.GasUsed,
		Input:                 callTrace.Input,
		Output:                callTrace.Output,
		Method:                callTrace.Method,
		Timestamp:             block.Timestamp,
	}

	i.dbClient.InsertInternalTransaction(pgxBatch, &internalTx)

	if internalTraceCount == nil {
		internalTraceCount = new(int)
		*internalTraceCount = 1
	} else {
		*internalTraceCount++
	}

	for _, call := range callTrace.Calls {
		i.indexCallTrace(pgxBatch, call, txHash, block, internalTraceCount)
	}
}

func (i *Indexer) indexTokenInfo(pgxBatch *pgx.Batch, tokenTransfers []db.TokenTransfer) error {
	// filter and unique token address from token transfers
	tokenAddresses := filterTokenAddress(tokenTransfers)
	for tokenAddress, tokenType := range tokenAddresses {
		cacheKey := fmt.Sprintf("%s_totalSupply", tokenAddress.Hex())
		cachedTotalSupply, existed := i.indexedTokens.Get(cacheKey)

		erc20, err := contract.NewIERC20(tokenAddress, i.l2client)
		if err != nil {
			return fmt.Errorf("error creating new ERC20 contract: %w", err)
		}

		totalSupply, _ := erc20.TotalSupply(nil)
		// Convert totalSupply to string, use "nil" to represent nil totalSupply
		var totalSupplyStr string
		if totalSupply == nil {
			totalSupplyStr = "nil"
		} else {
			totalSupplyStr = totalSupply.String()
		}

		// If totalSupply is different or token not indexed yet, proceed
		if !existed || (existed && cachedTotalSupply != totalSupplyStr) {
			name, _ := erc20.Name(nil)
			symbol, _ := erc20.Symbol(nil)
			decimals, _ := erc20.Decimals(nil)
			decimalsInt := int(decimals)

			token := db.Token{
				Address:     tokenAddress,
				Name:        i.sanitizationPolicy.Sanitize(strings.ToValidUTF8(name, "")),
				Symbol:      i.sanitizationPolicy.Sanitize(strings.ToValidUTF8(symbol, "")),
				Decimals:    &decimalsInt,
				TotalSupply: totalSupply,
				TokenType:   tokenType,
			}

			i.dbClient.UpsertTokenInfo(pgxBatch, token)

			// Update cache with new totalSupply
			i.indexedTokens.Add(cacheKey, totalSupplyStr)
		}
	}
	return nil
}

type AccountToken struct {
	Account      common.Address
	TokenAddress common.Address
	TokenType    string
	TokenID      *big.Int
}

// Define the pair struct
type AccountTokenPair struct {
	Account      common.Address
	TokenAddress common.Address
}

func (i *Indexer) indexTokenBalance(pgxBatch *pgx.Batch, tokenTransfers []db.TokenTransfer) error {
	accountTokenSet := make(map[AccountToken]bool)

	// Filter out unique account-tokenAddress pairs
	for _, transfer := range tokenTransfers {
		tokenAddress := transfer.TokenAddress
		tokenType := transfer.TokenType
		tokenID := transfer.TokenID

		for _, account := range []common.Address{transfer.From, transfer.To} {
			accToken := AccountToken{Account: account, TokenAddress: tokenAddress, TokenType: tokenType, TokenID: tokenID}
			accountTokenSet[accToken] = true
		}
	}

	blockNumber := new(big.Int)
	if len(tokenTransfers) > 0 {
		blockNumber = tokenTransfers[0].BlockNumber
	}
	opts := &bind.CallOpts{}

	// Create map to keep track of account-tokenaddress pairs that have been processed
	checkedPairMap := make(map[AccountTokenPair]*big.Int)

	// Fetch and update balance for each unique account-tokenAddress pair
	for accToken := range accountTokenSet {
		pair := AccountTokenPair{Account: accToken.Account, TokenAddress: accToken.TokenAddress}

		var balance *big.Int
		var err error

		// If this account-tokenAddress pair has been processed for ERC721, use the stored balance
		if accToken.TokenType == ERC721 && checkedPairMap[pair] != nil {
			balance = checkedPairMap[pair]
		} else {
			// skip zero account
			if accToken.Account.Hex() == "0x0000000000000000000000000000000000000000" ||
				accToken.TokenAddress.Hex() == "0x000000000000000000000000000000000000800A" {
				// zksync - 0x00..800A
				continue
			}

			balance, err = i.getBalance(accToken.TokenType, accToken.TokenAddress, accToken.Account, accToken.TokenID, opts)
			if err != nil {
				return fmt.Errorf("error getting balance of token %s for account %s on block %d: %s", accToken.TokenAddress.Hex(), accToken.Account.Hex(), blockNumber.Int64(), err.Error())
			}

			// If the TokenType is ERC721, store the balance in checkedPairMap for future use
			if accToken.TokenType == ERC721 {
				checkedPairMap[pair] = balance
			}
		}

		i.dbClient.UpsertTokenBalance(pgxBatch, accToken.TokenAddress, accToken.Account, balance, blockNumber, accToken.TokenID, accToken.TokenType)
	}
	return nil
}

func (i *Indexer) indexNullToAddressContractCreation(pgxBatch *pgx.Batch, tx *ethclient.RpcTransaction, receipt *ethclient.Receipt, timestamp *hexutil.Big) error {
	if tx.To != nil || receipt.Status != types.ReceiptStatusSuccessful {
		// not a contract creation or transaction failed
		return nil
	}

	// Fetch deployed bytecode
	deployedBytecode, err := i.l2client.CodeAt(i.ctx, receipt.ContractAddress, receipt.BlockNumber)
	if err != nil {
		return fmt.Errorf("error getting deployed bytecode in transaction %s: %w", tx.Hash.Hex(), err)
	}

	contractAddr := db.Contract{
		Address:           receipt.ContractAddress,
		Creator:           tx.From,
		CreationTxHash:    tx.Hash,
		CreationTimestamp: timestamp,
		CreationBytecode:  tx.Input,
		DeployedBytecode:  deployedBytecode,
	}

	i.dbClient.UpsertContract(pgxBatch, contractAddr)

	log.Info().Msgf("📜 Indexed contract creation %s", contractAddr.Address.Hex())

	return nil
}

func (i *Indexer) indexTracerContractCreation(pgxBatch *pgx.Batch, callTrace *ethclient.CallTrace, tx *ethclient.RpcTransaction, timestamp *hexutil.Big) {
	var calltype = strings.ToUpper(callTrace.Type)
	if calltype == "CREATE" || calltype == "CREATE2" {
		// it's a contract creation
		contractAddr := db.Contract{
			Address:           callTrace.To,
			Creator:           tx.From,
			CreationTxHash:    tx.Hash,
			CreationTimestamp: timestamp,
			CreationBytecode:  callTrace.Input,
			// DeployedBytecode: *callTrace.Output,
		}

		if callTrace.Output != nil {
			contractAddr.DeployedBytecode = *callTrace.Output
		}

		i.dbClient.UpsertContract(pgxBatch, contractAddr)

		log.Info().Msgf("📜 Indexed contract creation in tracer %s on block %d", contractAddr.Address.Hex(), tx.BlockNumber.ToInt())
	}

	// Process calls recursively
	if callTrace.Calls != nil && len(callTrace.Calls) > 0 {
		for _, innerCallTrace := range callTrace.Calls {
			i.indexTracerContractCreation(pgxBatch, innerCallTrace, tx, timestamp)
		}
	}
}

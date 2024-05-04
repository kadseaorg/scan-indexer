package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog/log"
)

type Client struct {
	*ethclient.Client
	rpcClient *rpc.Client
}

func NewClient(url string) (*Client, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Error().Msgf("[ethclient] Error connecting to Ethereum node: %s", err)
		return nil, err
	}

	return &Client{client, rpcClient}, nil
}

type RpcBlock struct {
	BaseFeePerGas    *hexutil.Big   `json:"baseFeePerGas"`
	Difficulty       *hexutil.Big   `json:"difficulty"`
	ExtraData        hexutil.Bytes  `json:"extraData"`
	GasLimit         *hexutil.Big   `json:"gasLimit"`
	GasUsed          *hexutil.Big   `json:"gasUsed"`
	Hash             common.Hash    `json:"hash"`
	L1BatchNumber    *hexutil.Big   `json:"l1BatchNumber"`
	L1BatchTimestamp *hexutil.Big   `json:"l1BatchTimestamp"`
	Miner            common.Address `json:"miner"`
	MixHash          common.Hash    `json:"mixHash"`
	Nonce            *hexutil.Bytes `json:"nonce"`
	Number           *hexutil.Big   `json:"number"`
	ParentHash       common.Hash    `json:"parentHash"`
	ReceiptsRoot     common.Hash    `json:"receiptsRoot"`
	// SealFields       []*etypes.HeaderField `json:"sealFields"`
	Sha3Uncles       common.Hash       `json:"sha3Uncles"`
	Size             *hexutil.Big      `json:"size"`
	StateRoot        common.Hash       `json:"stateRoot"`
	Timestamp        *hexutil.Big      `json:"timestamp"`
	TotalDifficulty  *hexutil.Big      `json:"totalDifficulty"`
	Transactions     []*RpcTransaction `json:"transactions"`
	TransactionsRoot common.Hash       `json:"transactionsRoot"`
	Uncles           []common.Hash     `json:"uncles"`
}

type RpcTransaction struct {
	BlockHash            *common.Hash    `json:"blockHash"`
	BlockNumber          *hexutil.Big    `json:"blockNumber"`
	From                 common.Address  `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Big    `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	Hash                 common.Hash     `json:"hash"`
	Input                hexutil.Bytes   `json:"input"`
	L1BatchNumber        *hexutil.Big    `json:"l1BatchNumber"`
	L1BatchTxIndex       *hexutil.Big    `json:"l1BatchTxIndex"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Nonce                *hexutil.Big    `json:"nonce"`
	TransactionIndex     *hexutil.Big    `json:"transactionIndex"`
	Type                 *hexutil.Big    `json:"type"`
	Value                *hexutil.Big    `json:"value"`
}

func (ec Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*RpcBlock, error) {
	var result RpcBlock
	err := ec.rpcClient.CallContext(ctx, &result, "eth_getBlockByNumber", toBlockNumArg(number), true)
	return &result, err
}

func (ec Client) GetBlockByHash(ctx context.Context, hash string) (*RpcBlock, error) {
	var result RpcBlock
	err := ec.rpcClient.CallContext(ctx, &result, "eth_getBlockByHash", hash, true)
	return &result, err
}

type CallTrace struct {
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	Value   *hexutil.Big   `json:"value"`
	Gas     *hexutil.Big   `json:"gas"`
	GasUsed *hexutil.Big   `json:"gasUsed"`
	Input   hexutil.Bytes  `json:"input"`
	Output  *hexutil.Bytes `json:"output,omitempty"`
	Method  string         `json:"method,omitempty"`
	Calls   []*CallTrace   `json:"calls,omitempty"`
}

func (c *Client) GetCallTrace(ctx context.Context, txHash common.Hash) (*CallTrace, error) {
	var result *CallTrace
	// just get call traces
	err := c.rpcClient.CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
		"tracer": "callTracer",
		// "timeout": "10s",
	})

	return result, err
}

func (c *Client) GetRevertReason(ctx context.Context, tx *RpcTransaction, blockNumber *big.Int) (string, error) {
	callMsg := ethereum.CallMsg{
		From:     tx.From,
		To:       tx.To,
		Gas:      tx.Gas.ToInt().Uint64(),
		GasPrice: tx.GasPrice.ToInt(),
		Value:    tx.Value.ToInt(),
		Data:     tx.Input,
	}

	_, err := c.CallContract(ctx, callMsg, blockNumber)
	if err != nil {
		return err.Error(), nil
	}

	return "", nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}

func (ec *Client) GetBlockTimestampByTxHash(ctx context.Context, txHash common.Hash) (uint64, error) {
	// Get the transaction receipt, which includes the block number
	receipt, err := ec.TransactionReceipt(ctx, txHash)
	if err != nil {
		return 0, err
	}

	return ec.GetBlockTimestampByNumber(ctx, receipt.BlockNumber)
}

func (ec *Client) GetBlockTimestampByNumber(ctx context.Context, blockNumber *big.Int) (uint64, error) {
	// Get the block header details using the block number retrieved
	header, err := ec.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		return 0, err
	}

	// Get the timestamp from the block header and return it
	return header.Time, nil
}

// TransactionReceipt extends types.Receipt with an additional L1Fee field
type Receipt struct {
	types.Receipt              // Embedded to extend the struct
	L1Fee         *hexutil.Big `json:"l1Fee,omitempty"`
}

func (ec *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*Receipt, error) {
	var result Receipt
	err := ec.rpcClient.CallContext(ctx, &result, "eth_getTransactionReceipt", txHash)
	return &result, err
}

type BatchByNumber struct {
	Number              string      `json:"number"`
	Coinbase            string      `json:"coinbase"`
	StateRoot           string      `json:"stateRoot"`
	GlobalExitRoot      string      `json:"globalExitRoot"`
	MainnetExitRoot     string      `json:"mainnetExitRoot"`
	RollupExitRoot      string      `json:"rollupExitRoot"`
	LocalExitRoot       string      `json:"localExitRoot"`
	AccInputHash        string      `json:"accInputHash"`
	Timestamp           string      `json:"timestamp"`
	SendSequencesTxHash string      `json:"sendSequencesTxHash"`
	VerifyBatchTxHash   interface{} `json:"verifyBatchTxHash"`
	Closed              bool        `json:"closed"`
	Blocks              []string    `json:"blocks"`
	Transactions        []string    `json:"transactions"`
	BatchL2Data         string      `json:"batchL2Data"`
}

func (ec *Client) GetBatchByNumber(ctx context.Context, batch uint64) (*BatchByNumber, error) {
	var result BatchByNumber
	hex := hexutil.EncodeUint64(batch)
	err := ec.rpcClient.CallContext(ctx, &result, "zkevm_getBatchByNumber", hex)
	return &result, err
}

func (r *Receipt) UnmarshalJSON(input []byte) error {
	// Create an alias to avoid infinite recursive parsing
	type Alias Receipt
	alias := &struct{ *Alias }{Alias: (*Alias)(r)}

	err := json.Unmarshal(input, &alias)
	if err != nil {
		return err
	}

	// Custom parse L1Fee
	var raw map[string]*json.RawMessage
	err = json.Unmarshal(input, &raw)
	if err != nil {
		return err
	}

	if val, ok := raw["l1Fee"]; ok {
		err = json.Unmarshal(*val, &(r.L1Fee))
	}

	return err
}

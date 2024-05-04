package db

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Token struct {
	// ID                   int64
	Name        string
	Symbol      string
	Address     common.Address
	Decimals    *int
	TotalSupply *big.Int
	TokenType   string
	// InsertedAt time.Time
	// UpdatedAt  time.Time
}

type InternalTransaction struct {
	// ID                    int64
	BlockNumber           *hexutil.Big
	BlockHash             common.Hash
	ParentTransactionHash common.Hash
	Type                  string
	From                  common.Address
	To                    common.Address
	Value                 *hexutil.Big
	Gas                   *hexutil.Big
	GasUsed               *hexutil.Big
	Input                 hexutil.Bytes
	Output                *hexutil.Bytes
	Method                string
	Timestamp             *hexutil.Big
	// InsertedAt            time.Time
	// UpdatedAt             time.Time
}

type TokenTransfer struct {
	// ID              int64
	TransactionHash common.Hash
	LogIndex        int
	MethodID        string
	TokenAddress    common.Address
	BlockNumber     *big.Int
	BlockHash       common.Hash
	From            common.Address
	To              common.Address
	Value           *big.Int
	Amount          *big.Int
	TokenID         *big.Int
	Amounts         []big.Int
	TokenIDs        []big.Int
	TokenType       string
	// InsertedAt      time.Time
	// UpdatedAt       time.Time
}

type Contract struct {
	// ID                   int64
	// Name                 string
	Address           common.Address
	Creator           common.Address
	CreationTxHash    common.Hash
	CreationTimestamp *hexutil.Big
	CreationBytecode  hexutil.Bytes
	DeployedBytecode  hexutil.Bytes
	// ABI                  string
	// ConstructorArguments string
	// SourceCode           string
	// CompilerVersion      string
	// Optimization         bool
	// OptimizationRuns     int
	// EVMVersion           string
	// IsVerified           bool
	// InsertedAt           time.Time
	// UpdatedAt            time.Time
}

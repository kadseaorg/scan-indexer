package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetSender(tx *types.Transaction) (common.Address, error) {
	return types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
}

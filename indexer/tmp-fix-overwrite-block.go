package indexer

import "math/big"

func (i *Indexer) FixOverwriteBlock(blockNumber int64) {
	i.handleBlock(big.NewInt(blockNumber))
}

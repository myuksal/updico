package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func (container *jsonRPCContainer) Block(blockNumber big.Int) *types.Block {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	block, err := _it.client.BlockByNumber(ctx, &blockNumber)
	if err != nil {
		panic(err)
	}

	return block
}

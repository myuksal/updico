package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func Block(blockNumber big.Int) *types.Block {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	block, err := JsonRPC.client.BlockByNumber(ctx, &blockNumber)
	if err != nil {
		panic(err)
	}

	return block
}

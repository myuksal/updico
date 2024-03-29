package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func Header(blockNumber big.Int) *types.Header {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	header, err := JsonRPC.client.HeaderByNumber(ctx, &blockNumber)
	if err != nil {
		panic(err)
	}

	return header
}

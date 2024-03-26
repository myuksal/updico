package jsonrpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (container *jsonRPCContainer) TransactionReceipt(hash common.Hash) *types.Receipt {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	receipt, err := _it.client.TransactionReceipt(ctx, hash)
	if err != nil {
		panic(err)
	}

	return receipt
}

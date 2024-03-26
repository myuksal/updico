package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func (container *jsonRPCContainer) LogsByBlockRange(from big.Int, to big.Int) []types.Log {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	query := ethereum.FilterQuery{
		FromBlock: &from,
		ToBlock:   &to,
	}

	logs, err := _it.client.FilterLogs(ctx, query)
	if err != nil {
		panic(err)
	}

	return logs
}

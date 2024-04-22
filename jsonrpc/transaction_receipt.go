package jsonrpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/metrics"
)

func (rpc *jsonRPCContainer) TransactionReceipt(parentCtx context.Context, hash common.Hash) *types.Receipt {
	rpcConfig := config.GetConfig().JsonRpcConfig

	// prometheus metrics
	start := time.Now()
	defer metrics.JsonRpcHistogram.WithLabelValues(metrics.CollectMetricJsonRpcCallTransactionReceipt).Observe(float64(time.Since(start).Milliseconds()))

	// context
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(rpcConfig.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	// json rpc call
	receipt, err := rpc.client.TransactionReceipt(ctx, hash)
	if err != nil {
		panic(err)
	}

	return receipt
}

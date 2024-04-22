package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/metrics"
)

func (rpc *jsonRPCContainer) Header(parentCtx context.Context, blockNumber big.Int) *types.Header {
	// config
	rpcConfig := config.GetConfig().JsonRpcConfig

	// prometheus metrics
	start := time.Now()
	defer metrics.JsonRpcHistogram.WithLabelValues(metrics.CollectMetricJsonRpcCallHeader).Observe(float64(time.Since(start).Milliseconds()))

	// context
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(rpcConfig.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	// json rpc call
	header, err := rpc.client.HeaderByNumber(ctx, &blockNumber)
	if err != nil {
		panic(err)
	}

	return header
}

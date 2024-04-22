package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/metrics"
)

func (rpc *jsonRPCContainer) Block(parentCtx context.Context, blockNumber big.Int) (*types.Block, error) {
	// config
	rpcConfig := config.GetConfig().JsonRpcConfig

	// prometheus metrics
	start := time.Now()
	defer metrics.JsonRpcHistogram.WithLabelValues(metrics.CollectMetricJsonRpcCallBlock).Observe(float64(time.Since(start).Milliseconds()))

	// context
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(rpcConfig.Timeout.Block)*time.Millisecond)
	defer cancel()

	// json rpc call
	block, err := rpc.client.BlockByNumber(ctx, &blockNumber)
	if err != nil {
		return nil, err
	}

	return block, nil
}

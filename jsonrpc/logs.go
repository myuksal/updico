package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/metrics"
)

func (rpc *jsonRPCContainer) LogsByBlockRange(parentCtx context.Context, from big.Int, to big.Int) []types.Log {
	// config
	rpcConfig := config.GetConfig().JsonRpcConfig

	// prometheus metrics
	start := time.Now()
	defer metrics.JsonRpcHistogram.WithLabelValues(metrics.CollectMetricJsonRpcCallLogs).Observe(float64(time.Since(start).Milliseconds()))

	// context
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(rpcConfig.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	// json rpc call
	query := ethereum.FilterQuery{
		FromBlock: &from,
		ToBlock:   &to,
	}

	logs, err := rpc.client.FilterLogs(ctx, query)
	if err != nil {
		panic(err)
	}

	return logs
}

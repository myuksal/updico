package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/metrics"
)

func (rpc *jsonRPCContainer) Balance(parentCtx context.Context, account common.Address, block big.Int) (*big.Int, error) {
	// config
	rpcConfig := config.GetConfig().JsonRpcConfig

	// prometheus metrics
	start := time.Now()
	defer metrics.JsonRpcHistogram.WithLabelValues(metrics.CollectMetricJsonRpcCallBalance).Observe(float64(time.Since(start).Milliseconds()))

	// context
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(rpcConfig.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	// json rpc call
	balance, err := rpc.client.BalanceAt(ctx, account, &block)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

package jsonrpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func Balance(account common.Address, block big.Int) *big.Int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(_config.Timeout.TransactionReceipt)*time.Millisecond)
	defer cancel()

	balance, err := JsonRPC.client.BalanceAt(ctx, account, &block)
	if err != nil {
		panic(err)
	}

	return balance
}

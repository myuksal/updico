package entities

import (
	"math/big"

	"github.com/myuksal/updico/types"
)

type LogEntity struct {
	address types.Address
	// topics:["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x000000000000000000000000eee0fca5bd69a2da5bba40057b16666d1d74c900","0x000000000000000000000000995d1a0a821cfeeacf21d1d4a23f7aebe2caa90d"],
	data             string
	blockNumber      big.Int
	transactionHash  types.Hash64
	transactionIndex int32
	blockHash        types.Hash64
	logIndex         int32
	removed          bool
}

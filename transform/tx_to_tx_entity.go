package transform

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/myuksal/updico/entities"
)

func TransactionToEntity(receipt types.Receipt) entities.TransactionEntity {
	return entities.TransactionEntity{
		Hash:                 receipt.TxHash,
		Nonce:                0,
		BlockHash:            receipt.BlockHash,
		BlockNumber:          *receipt.BlockNumber,
		Gas:                  *receipt.BlobGasPrice,
		GasPrice:             *receipt.BlobGasPrice,
		MaxPriorityFeePerGas: &big.Int{},
		MaxFeePerGas:         &big.Int{},
		Type:                 0,
		Value:                big.Int{},
		Input:                []byte{},
		V:                    big.Int{},
		R:                    big.Int{},
		S:                    big.Int{},
	}

}

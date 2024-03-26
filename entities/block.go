package entities

import (
	"math/big"

	"github.com/myuksal/updico/types"
)

type BlockEntity struct {
	Hash        types.Hash64  `json:"hash"       			 gencodec:"required"  gorm:"primaryKey"`
	ParentHash  types.Hash64  `json:"parentHash"       gencodec:"required"`
	UncleHash   types.Hash64  `json:"uncleHahs"        gencodec:"required"`
	Coinbase    types.Address `json:"miner"`
	Root        types.Hash32  `json:"stateRoot"        gencodec:"required"`
	TxHash      types.Hash32  `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash types.Hash32  `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom   `json:"logsBloom"        gencodec:"required"`
	Difficulty  big.Int       `json:"difficulty"       gencodec:"required"`
	Number      big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64        `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64        `json:"gasUsed"          gencodec:"required"`
	Time        uint64        `json:"timestamp"        gencodec:"required"`
	Extra       []byte        `json:"extraData"        gencodec:"required"`
	MixDigest   types.Hash32  `json:"mixHash"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash types.Hash32 `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot types.Hash32 `json:"parentBeaconBlockRoot" rlp:"optional"`
}

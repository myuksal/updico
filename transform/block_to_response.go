package transform

import "github.com/ethereum/go-ethereum/core/types"

type blockResponse struct {
	BaseFee    string `json:"baseFee"`
	GasLimit   uint64 `json:"gasLimit"`
	GasUsed    uint64 `json:"gasUsed"`
	Difficulty string `json:"difficulty"`
	ParentHash string `json:"parentHash"`
	Hash       string `json:"hash"`
	Size       uint64 `json:"size"`
}

func BlockToResponse(block *types.Block) blockResponse {
	return blockResponse{
		BaseFee:    block.BaseFee().Text(10),
		GasLimit:   block.GasLimit(),
		GasUsed:    block.GasUsed(),
		Difficulty: block.Difficulty().Text(10),
		ParentHash: block.ParentHash().Hex(),
		Hash:       block.Hash().Hex(),
		Size:       block.Size(),
	}

}

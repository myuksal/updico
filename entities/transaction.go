package entities

import (
	"database/sql"
)

type TransactionEntity struct {
	Hash                 string         `json:"hash" gorm:"type:CHAR(64) CHARACTER SET binary COLLATE binary"`
	Nonce                int8           `json:"nonce"`
	BlockHash            string         `json:"blockhash" gorm:"type:CHAR(64) CHARACTER SET binary COLLATE binary"`
	BlockNumber          string         `json:"blockNumber"`
	From                 sql.NullString `json:"from" gorm:"type:CHAR(20) CHARACTER SET binary COLLATE binary"`
	To                   string         `json:"to" gorm:"type:CHAR(20) CHARACTER SET binary COLLATE binary"`
	Gas                  string         `json:"gas" gorm:"type:DECIMAL(65)"`
	GasPrice             string         `json:"gasPrice" gorm:"type:DECIMAL(65)"`
	MaxPriorityFeePerGas *string        `json:"maxPriorityFeePerGas" gorm:"type:DECIMAL(65)"`
	MaxFeePerGas         *string        `json:"maxFeePerGas" gorm:"type:DECIMAL(65)"`
	Type                 byte           `json:"type"`
	Value                string         `json:"value" gorm:"type:DECIMAL(65)"`
	Input                sql.RawBytes   `json:"input"`
	V                    string         `json:"v" gorm:"type:DECIMAL(65)"`
	R                    string         `json:"r" gorm:"type:DECIMAL(65)"`
	S                    string         `json:"s" gorm:"type:DECIMAL(65)"`
}

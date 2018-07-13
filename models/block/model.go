package block

import (
	"time"
	"explorer-api/models/transaction"
	"explorer-api/models/validator"
)

type Block struct {
	ID           uint                  `gorm:"primary_key"`
	Height       uint                  `json:"height"`
	TxCount      uint                  `json:"num_txs"`
	Size         uint                  `json:"size"`
	BlockTime    uint                  `json:"block_time"`
	BlockReward  uint                  `json:"block_reward" gorm:"type:numeric(50, 0)"`
	Hash         string                `json:"hash" gorm:"type:bytea"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Validators   []validator.Validator `gorm:"many2many:block_validator;"`
	Transactions []transaction.Transaction
}

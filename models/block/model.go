package block

import (
	"time"
	"explorer-api/models/transaction"
)

type Model struct {
	ID          uint   `gorm:"primary_key"`
	Height      uint   `json:"height"`
	TxCount     uint   `json:"num_txs"`
	Size        uint   `json:"size"`
	BlockTime   uint   `json:"block_time"`
	BlockReward uint   `json:"block_reward" gorm:"type:numeric(50, 0)"`
	Hash        string `json:"hash" gorm:"type:bytea"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time

	Transactions []transaction.Model
}

func (Model) TableName() string {
	return "blocks"
}

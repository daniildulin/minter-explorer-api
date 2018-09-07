package block

import (
	"github.com/daniildulin/explorer-api/models/transaction"
	"github.com/daniildulin/explorer-api/models/validator"
	"time"
)

type Block struct {
	ID           uint                  `gorm:"primary_key"`
	Height       uint                  `json:"height"`
	TxCount      uint                  `json:"num_txs"`
	Size         uint                  `json:"size"`
	BlockTime    uint                  `json:"block_time"`
	BlockReward  string                `json:"block_reward" gorm:"type:numeric(50, 0)"`
	Hash         string                `json:"hash" gorm:"type:bytea"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Validators   []validator.Validator `gorm:"many2many:block_validator;"`
	Transactions []transaction.Transaction
}

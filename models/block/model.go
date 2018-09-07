package block

import (
	"github.com/daniildulin/explorer-api/models/transaction"
	"github.com/daniildulin/explorer-api/models/validator"
	"time"
)

type Block struct {
	ID           uint    `gorm:"type:bigint;primary_key"`
	Height       uint    `json:"height"       gorm:"type:bigint;unique_index"`
	Timestamp    float32 `json:"timestamp"    gorm:"type:numeric(20, 10)"`
	TxCount      uint    `json:"tx_count"`
	Size         uint    `json:"size"`
	BlockTime    uint    `json:"block_time"   gorm:"type:numeric(12, 9)"`
	Hash         string  `json:"hash"         gorm:"type:bytea"`
	BlockReward  string  `json:"block_reward" gorm:"type:numeric(50, 0)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Validators   []validator.Validator `gorm:"many2many:block_validator;"`
	Transactions []transaction.Transaction
}

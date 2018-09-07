package validator

import (
	"math/big"
	"time"
)

type Validator struct {
	ID                uint     `gorm:"type:bigint;primary_key"`
	Name              *string  `json:"name"`
	AccumulatedReward *big.Int `json:"accumulated_reward" gorm:"type:numeric(300,0)"`
	AbsentTimes       uint     `json:"absent_times"       gorm:"type:bigint"`
	Address           string   `json:"address"            gorm:"type:varchar(100)"`
	TotalStake        *big.Int `json:"total_stake"        gorm:"type:numeric(300,0)"`
	PubKey            string   `json:"public_key"         gorm:"type:varchar(255)"`
	Commission        uint     `json:"commission"         gorm:"type:bigint"`
	Status            uint     `json:"status"             gorm:"type:smallint"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

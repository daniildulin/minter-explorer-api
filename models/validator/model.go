package validator

import (
	"time"
)

type Validator struct {
	ID        uint   `gorm:"primary_key"`
	Address   string `json:"address"`
	PubKey    string `json:"public_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

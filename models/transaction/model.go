package transaction

import (
	"math/big"
	"time"
)

type Model struct {
	ID                   uint     `gorm:"primary_key"`
	BlockID              uint     `json:"block_id"`
	Type                 uint     `json:"type"`
	Nonce                uint     `json:"nonce"`
	GasPrice             big.Int  `json:"gas_price" gorm:"type:numeric(50,0)"`
	From                 string   `json:"from" gorm:"type:varchar(100)"`
	To                   *string  `json:"to" gorm:"type:varchar(100)"`
	Coin                 *string  `json:"coin" gorm:"type:varchar(255)"`
	Hash                 string   `json:"hash" gorm:"type:varchar(100)"`
	Payload              string   `json:"payload" gorm:"type:varchar(255)"`
	ServiceData          string   `json:"service_data" gorm:"type:varchar(255)"`
	PubKey               *string  `json:"PubKey"`
	Address              *string  `json:"Address"`
	FromCoinSymbol       *string  `json:"FromCoinSymbol"`
	ToCoinSymbol         *string  `json:"ToCoinSymbol"`
	RawCheck             *string  `json:"RawCheck"`
	Proof                *string  `json:"Proof"`
	Name                 *string  `json:"name"`
	Symbol               *string  `json:"symbol"`
	Gas                  big.Int  `json:"gas" gorm:"type:numeric(50,0)"`
	Value                *big.Int `json:"value" gorm:"type:numeric(50,18)"`
	Stake                *big.Int `json:"Stake" gorm:"type:numeric(50,0)"`
	Commission           *uint    `json:"Commission" gorm:"type:numeric(50,0)"`
	InitialAmount        *big.Int `json:"initial_amount" gorm:"type:numeric(50,0)"`
	InitialReserve       *big.Int `json:"initial_reserve" gorm:"type:numeric(50,0)"`
	ConstantReserveRatio *uint    `json:"constant_reserve_ratio" gorm:"type:numeric(50,0)"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

func (Model) TableName() string {
	return "transactions"
}

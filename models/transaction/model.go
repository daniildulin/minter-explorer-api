package transaction

import (
	"math/big"
	"time"
)

type Transaction struct {
	ID                   uint     `gorm:"type:bigint;primary_key"`
	BlockID              uint     `json:"block_id"               gorm:"type:bigint;primary_key"`
	Type                 uint     `json:"type"`
	From                 string   `json:"from"                   gorm:"type:varchar(100)"`
	To                   *string  `json:"to"                     gorm:"type:varchar(100)"`
	Hash                 string   `json:"hash"                   gorm:"type:varchar(100)"`
	PubKey               *string  `json:"pub_key"                gorm:"type:varchar(255)"`
	Value                *big.Int `json:"value"                  gorm:"type:numeric(300,0)"`
	ValueToSell          *big.Int `json:"value_to_sell"          gorm:"type:numeric(300,0)"`
	ValueToBuy           *big.Int `json:"value_to_buy"           gorm:"type:numeric(300,0)"`
	Fee                  big.Int  `json:"fee"                    gorm:"type:numeric(300,0)"`
	Stake                *big.Int `json:"stake"                  gorm:"type:numeric(300,0)"`
	Commission           *big.Int `json:"Commission"             gorm:"type:numeric(300,0)"`
	InitialAmount        *big.Int `json:"initial_amount"         gorm:"type:numeric(300,0)"`
	InitialReserve       *big.Int `json:"initial_reserve"        gorm:"type:numeric(50,0)"`
	ConstantReserveRatio *big.Int `json:"constant_reserve_ratio" gorm:"type:numeric(300,0)"`
	GasWanted            big.Int  `json:"gas_wanted"             gorm:"type:numeric(300,0)"`
	GasUsed              big.Int  `json:"gas_used"               gorm:"type:numeric(300,0)"`
	GasPrice             big.Int  `json:"gas_price"              gorm:"type:numeric(300,0)"`
	GasCoin              *string  `json:"gas_coin"               gorm:"type:varchar(20)"`
	Coin                 *string  `json:"coin"                   gorm:"type:varchar(255)"`
	Nonce                uint     `json:"nonce"`
	Payload              *string  `json:"payload"                gorm:"type:varchar(255)"`
	ServiceData          *string  `json:"service_data"           gorm:"type:varchar(255)"`
	Address              *string  `json:"Address"                gorm:"type:varchar(255)"`
	CoinToSell           *string  `json:"coin_to_sell"           gorm:"type:varchar(25)"`
	CoinToBuy            *string  `json:"coin_to_buy"            gorm:"type:varchar(25)"`
	RawCheck             *string  `json:"raw_check"`
	Proof                *string  `json:"proof"`
	Name                 *string  `json:"name"`
	Log                  *string  `json:"log"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

package minterapi

import (
	"math/big"
	"time"
)

type statusResponse struct {
	Code   uint         `json:"code"`
	Result statusResult `json:"result"`
}

type blockResponse struct {
	Code   uint        `json:"code"`
	Result blockResult `json:"result"`
}

type validatorsResponse struct {
	Code   uint        `json:"code"`
	Result []validator `json:"result"`
}

type statusResult struct {
	LatestBlockHash   []byte `json:"latest_block_hash"`
	LatestAppHash     []byte `json:"latest_app_hash"`
	LatestBlockHeight uint   `json:"latest_block_height"`
}

type blockResult struct {
	Hash         string    `json:"hash"`
	Height       uint      `json:"height"`
	TxCount      uint      `json:"num_txs"`
	TotalTx      uint      `json:"total_txs"`
	Time         time.Time `json:"time"`
	Transactions []transaction
}

type transaction struct {
	Hash        string          `json:"hash"`
	From        string          `json:"from"`
	Nonce       uint            `json:"nonce"`
	GasPrice    big.Int         `json:"gasPrice"`
	Gas         big.Int         `json:"gas"`
	Type        uint            `json:"type"`
	Payload     string          `json:"payload"`
	ServiceData string          `json:"serviceData"`
	Data        transactionData `json:"data"`
	TxResult    *txResultData   `json:"tx_result"`
}

type transactionData struct {
	Coin                 *string  `json:"coin"`
	To                   *string  `json:"to"`
	Value                *big.Int `json:"value"`
	CoinToSell           *string  `json:"coin_to_sell"`
	CoinToBuy            *string  `json:"coin_to_buy"`
	ValueToSell          *big.Int `json:"value_to_sell"`
	ValueToBuy           *big.Int `json:"value_to_buy"`
	Name                 *string  `json:"name"`
	Symbol               *string  `json:"coin_symbol"`
	InitialAmount        *big.Int `json:"initial_amount"`
	InitialReserve       *big.Int `json:"initial_reserve"`
	ConstantReserveRatio *uint    `json:"constant_reserve_ratio"`
	Address              *string  `json:"address"`
	PubKey               *string  `json:"pub_key"`
	Commission           *uint    `json:"commission"`
	Stake                *big.Int `json:"stake"`
	Proof                *string  `json:"proof"`
	RawCheck             *string  `json:"raw_check"`
	ToCoinSymbol         *string  `json:"to_coin_symbol"`
	FromCoinSymbol       *string  `json:"from_coin_symbol"`
}

type txResultData struct {
	GasWanted *uint   `json:"gas_wanted"`
	GasUsed   *uint   `json:"gas_used"`
	Log       *string `json:"log"`
	Code      *uint   `json:"code"`
	//Tags      *[]txResultTag `json:"tags"`
}
type txResultTag struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type validator struct {
	Address           string `json:"candidate_address"`
	PubKey            string `json:"pub_key"`
	TotalStake        uint   `json:"total_stake"`
	Commission        uint   `json:"commission"`
	AccumulatedReward uint   `json:"accumulated_reward"`
}

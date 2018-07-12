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
}

type transactionData struct {
	To                   *string  `json:"to"`
	Coin                 *string  `json:"coin"`
	Name                 *string  `json:"name"`
	Value                *big.Int `json:"value"`
	Stake                *big.Int `json:"Stake"`
	Proof                *string  `json:"Proof"`
	Symbol               *string  `json:"coin_symbol"`
	PubKey               *string  `json:"PubKey"`
	Address              *string  `json:"Address"`
	RawCheck             *string  `json:"RawCheck"`
	Commission           *uint    `json:"Commission"`
	ToCoinSymbol         *string  `json:"to_coin"`
	InitialAmount        *big.Int `json:"initial_amount"`
	InitialReserve       *big.Int `json:"initial_reserve"`
	FromCoinSymbol       *string  `json:"from_coin"`
	ConstantReserveRatio *uint    `json:"constant_reserve_ratio"`
}

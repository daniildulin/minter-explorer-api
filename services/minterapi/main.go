package minterapi

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"

	"encoding/json"
	"github.com/jinzhu/gorm"

	"explorer-api/env"
	"explorer-api/models/block"
	"explorer-api/helpers"
	minterReward "explorer-api/mintersdk/reward"
	validatorModel "explorer-api/models/validator"
	transactionModel "explorer-api/models/transaction"
	validatorRepository "explorer-api/repositories/validator"
)

var httpClient = &http.Client{Timeout: 1 * time.Second}

func Run(config env.Config, db *gorm.DB) {

	currentDBBlock := getLastBlockFromDB(db) + 1

	lastApiBlock := getLastBlockFromMinterAPI(config)

	log.Printf("Connect to %s", config.GetString("minterApi"))

	log.Printf("Start from block %d", currentDBBlock)

	for {
		if currentDBBlock <= lastApiBlock {
			start := time.Now()
			storeDataToDb(config, db, currentDBBlock)
			elapsed := time.Since(start)
			currentDBBlock++

			if config.GetBool(`debug`) {
				log.Printf("Time of processing %s for block %s", elapsed, fmt.Sprint(currentDBBlock))
			}

		} else {
			lastApiBlock = getLastBlockFromMinterAPI(config)
		}
	}
}

//Get JSON response from API
func getJson(url string, target interface{}) error {

	r, err := httpClient.Get(url)

	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// Get last block height from Minter API
func getLastBlockFromMinterAPI(config env.Config) uint {
	statusResponse := statusResponse{}
	getJson(`http://`+config.GetString("minterApi")+`/api/status`, &statusResponse)
	return statusResponse.Result.LatestBlockHeight
}

func getLastBlockFromDB(db *gorm.DB) uint {
	var b block.Block
	db.Last(&b)
	return b.Height
}

// Store data to DB
func storeDataToDb(config env.Config, db *gorm.DB, blockHeight uint) error {
	apiLink := `http://` + config.GetString("minterApi") + `/api/block/` + fmt.Sprint(blockHeight)
	blockResponse := blockResponse{}
	getJson(apiLink, &blockResponse)

	validatorsResponse := validatorsResponse{}
	apiLink = `http://` + config.GetString("minterApi") + `/api/validators/?height=` + fmt.Sprint(blockHeight)
	getJson(apiLink, &validatorsResponse)
	validators := getValidatorModels(db, validatorsResponse.Result)

	storeBlockToDB(db, &blockResponse.Result, validators)

	if config.GetBool(`debug`) {
		log.Printf("Block: %d; Txs: %d; Hash: %s", blockResponse.Result.Height, blockResponse.Result.TxCount, blockResponse.Result.Hash)
	}

	return nil
}

func storeBlockToDB(db *gorm.DB, blockData *blockResult, validators []validatorModel.Validator) {

	blockModel := block.Block{
		Hash:        blockData.Hash,
		Height:      blockData.Height,
		TxCount:     blockData.TxCount,
		CreatedAt:   blockData.Time,
		Size:        0,
		BlockTime:   getBlockTime(db, blockData.Height, blockData.Time),
		BlockReward: getBlockReward(blockData.Height),
	}

	if blockModel.TxCount > 0 {
		blockModel.Transactions = getTransactionModelsFromApiData(blockModel.TxCount, blockData.Time, blockData.Transactions)
	}

	if len(validators) > 0 {
		blockModel.Validators = validators
	}

	db.Create(&blockModel)

}

func getTransactionModelsFromApiData(count uint, blockTime time.Time, transactions []transaction) []transactionModel.Transaction {

	var result = make([]transactionModel.Transaction, count)
	i := 0
	for _, tx := range transactions {
		result[i] = transactionModel.Transaction{
			Hash:                 strings.Title(tx.Hash),
			From:                 strings.Title(tx.From),
			Type:                 tx.Type,
			Nonce:                tx.Nonce,
			GasPrice:             tx.GasPrice,
			Gas:                  tx.Gas,
			Payload:              tx.Payload,
			ServiceData:          tx.ServiceData,
			CreatedAt:            blockTime,
			To:                   tx.Data.To,
			Address:              tx.Data.Address,
			FromCoinSymbol:       tx.Data.FromCoinSymbol,
			ToCoinSymbol:         tx.Data.ToCoinSymbol,
			Name:                 tx.Data.Name,
			Symbol:               tx.Data.Symbol,
			Stake:                tx.Data.Stake,
			Value:                tx.Data.Value,
			Commission:           tx.Data.Commission,
			InitialAmount:        tx.Data.InitialAmount,
			InitialReserve:       tx.Data.InitialReserve,
			ConstantReserveRatio: tx.Data.ConstantReserveRatio,
			RawCheck:             tx.Data.RawCheck,
			Proof:                tx.Data.Proof,
			Coin:                 tx.Data.Coin,
			PubKey:               tx.Data.PubKey,
		}
	}

	return result
}

func getValidatorModels(db *gorm.DB, validatorsData []validator) []validatorModel.Validator {

	var result []validatorModel.Validator

	if len(validatorsData) > 0 {
		for _, v := range validatorsData {
			var vld validatorModel.Validator
			vld = validatorRepository.GetByPubKey(db, v.PubKey)
			if vld.ID == 0 && len(v.PubKey) > 0 {
				result = append(result, validatorModel.Validator{
					Address: v.Address,
					PubKey:  v.PubKey,
				})
			} else if vld.ID != 0 {
				result = append(result, vld)
			}

		}
		return result
	}
	return nil
}

func getBlockTime(db *gorm.DB, currentBlockHeight uint, blockTime time.Time) uint {

	if currentBlockHeight == 1 {
		return 5
	}

	var b block.Block
	db.Where("height = ?", currentBlockHeight-1).First(&b)

	result := blockTime.Second() - b.CreatedAt.Second()
	if result < 0 {
		return 5
	}

	return uint(result)
}

func getBlockReward(blockHeight uint) string {

	blockReward, err := minterReward.Get(blockHeight)
	helpers.CheckErr(err)

	return blockReward.String()
}

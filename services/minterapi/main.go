package minterapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"encoding/json"
	"github.com/jinzhu/gorm"

	"github.com/daniildulin/explorer-api/env"
	"github.com/daniildulin/explorer-api/helpers"
	minterReward "github.com/daniildulin/explorer-api/mintersdk/reward"
	"github.com/daniildulin/explorer-api/models/block"
	modelTransaction "github.com/daniildulin/explorer-api/models/transaction"
	modelValidator "github.com/daniildulin/explorer-api/models/validator"
	validatorRepository "github.com/daniildulin/explorer-api/repositories/validator"
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

func getApiLink(config env.Config) string {
	return `http://` + config.GetString("minterApi")
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
	getJson(getApiLink(config)+`/api/status`, &statusResponse)
	return statusResponse.Result.LatestBlockHeight
}

func getLastBlockFromDB(db *gorm.DB) uint {
	var b block.Block
	db.Last(&b)
	return b.Height
}

// Store data to DB
func storeDataToDb(config env.Config, db *gorm.DB, blockHeight uint) error {
	apiLink := getApiLink(config) + `/api/block/` + fmt.Sprint(blockHeight)
	blockResponse := blockResponse{}
	getJson(apiLink, &blockResponse)
	blockResult := blockResponse.Result

	validatorsResponse := validatorsResponse{}
	apiLink = getApiLink(config) + `/api/validators/?height=` + fmt.Sprint(blockHeight)
	getJson(apiLink, &validatorsResponse)
	validators := getValidatorModels(db, validatorsResponse.Result)

	storeBlockToDB(db, &blockResult, validators)

	if config.GetBool(`debug`) {
		log.Printf("Block: %d; Txs: %d; Hash: %s", blockResult.Height, blockResult.TxCount, blockResponse.Result.Hash)
	}

	return nil
}

func storeBlockToDB(db *gorm.DB, blockData *blockResult, validators []modelValidator.Validator) {

	if blockData.Height <= 0 {
		return
		log.Printf("Block: %d; Txs: %d; Hash: %s", blockData.Height, blockData.TxCount, blockData.Hash)
	}

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

func getTransactionModelsFromApiData(count uint, blockTime time.Time, transactions []transaction) []modelTransaction.Transaction {

	var result = make([]modelTransaction.Transaction, count)
	i := 0
	for _, tx := range transactions {
		result[i] = modelTransaction.Transaction{
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

func getValidatorModels(db *gorm.DB, validatorsData []validator) []modelValidator.Validator {

	var result []modelValidator.Validator

	if len(validatorsData) > 0 {
		for _, v := range validatorsData {
			var vld modelValidator.Validator
			vld = validatorRepository.GetByPubKey(db, v.PubKey)
			if vld.ID == 0 && len(v.PubKey) > 0 {
				result = append(result, modelValidator.Validator{
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

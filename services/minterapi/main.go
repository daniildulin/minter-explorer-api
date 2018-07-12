package minterapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"encoding/json"
	"explorer-api/env"

	"github.com/jinzhu/gorm"
	"strings"
	"explorer-api/models/block"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Run(config env.Config, db *gorm.DB) {

	lastDBBlock := getLastBlockFromDB(db)

	lastApiBlock := getLastBlockFromMinterAPI(config)

	log.Printf("Connect to %s", config.GetString("minterApi"))

	log.Printf("Start from block %d", lastDBBlock)

	for {
		if lastDBBlock <= lastApiBlock {
			start := time.Now()
			storeDataToDb(config, db, lastDBBlock)
			elapsed := time.Since(start)
			lastDBBlock++

			if config.GetBool(`debug`) {
				log.Printf("Time of processing %s for block %s", elapsed, fmt.Sprint(lastDBBlock))
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
	var b block.Model
	db.Last(&b)
	return b.Height
}

// Store data to DB
func storeDataToDb(config env.Config, db *gorm.DB, blockHeight uint) error {
	apiLink := `http://` + config.GetString("minterApi") + `/api/block/` + fmt.Sprint(blockHeight)
	response := blockResponse{}
	getJson(apiLink, &response)

	storeBlockToDB(db, &response.Result)

	if config.GetBool(`debug`) {
		log.Printf("Block: %d; Txs: %d; Hash: %s", response.Result.Height, response.Result.TxCount, response.Result.Hash)
	}

	return nil
}

func storeBlockToDB(db *gorm.DB, blockData *blockResult) {

	blk := block.Model{
		Hash:        strings.Title(blockData.Hash),
		Height:      blockData.Height,
		TxCount:     blockData.TxCount,
		Size:        0,
		BlockTime:   5, //TODO: добавить расчет
		BlockReward: 0,
	}

	blk.CreatedAt = blockData.Time

	db.Create(&blk)
}

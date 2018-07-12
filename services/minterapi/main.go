package minterapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"encoding/json"
	"explorer-api/env"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Run(config env.Config) {

	var lastDBBlock uint = 1

	lastApiBlock := getLastBlockFromMinterAPI(config)

	log.Printf("Connect to %s", config.GetString("minterApi"))

	log.Printf("Start from block %d", lastDBBlock)

	for {
		if lastDBBlock <= lastApiBlock {
			start := time.Now()
			storeDataToDb(config, lastDBBlock)
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

// Store data to DB
func storeDataToDb(config env.Config, blockHeight uint) error {
	apiLink := `http://` + config.GetString("minterApi") + `/api/block/` + fmt.Sprint(blockHeight)
	response := blockResponse{}
	getJson(apiLink, &response)

	if config.GetBool(`debug`) {
		log.Printf("Block: %d; Txs: %d; Hash: %s", response.Result.Height, response.Result.TxCount, response.Result.Hash)
	}

	return nil
}

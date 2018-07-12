package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"explorer-api/env"
	"explorer-api/services/minterapi"
	"explorer-api/helpers"
	"explorer-api/models/transaction"
	"explorer-api/models/block"
)

var Version string   // Version
var GitCommit string // Git commit
var BuildDate string // Build date
var AppName string   // Application name
var config env.Config

var version = flag.Bool("version", false, "Prints current version")

// Initialize app.
func init() {
	config = env.NewViperConfig()
	AppName = config.GetString("name")

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("%s v%s Commit %s builded %s\n", AppName, Version, GitCommit, BuildDate)
		os.Exit(0)
	}

	db, err := gorm.Open("postgres", config.GetString(`database.url`))
	helpers.CheckErr(err)
	defer db.Close()

	migrate(db)

	minterapi.Run(config, db)
}

func migrate(db *gorm.DB) {
	// Use GORM automigrate for models
	fmt.Println("Automigrate database schema.")
	db.AutoMigrate(&block.Model{}, &transaction.Model{})
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/daniildulin/explorer-api/env"
	"github.com/daniildulin/explorer-api/helpers"
	"github.com/daniildulin/explorer-api/models/block"
	"github.com/daniildulin/explorer-api/models/transaction"
	"github.com/daniildulin/explorer-api/models/validator"
	"github.com/daniildulin/explorer-api/services/minterapi"
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
	db.AutoMigrate(&block.Block{}, &transaction.Transaction{}, &validator.Validator{})
	db.Exec("CREATE TABLE IF NOT EXISTS block_validator( block_id INT NOT NULL, validator_id INT NOT NULL)")
}

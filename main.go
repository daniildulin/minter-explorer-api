package main

import (
	"flag"
	"fmt"
	"os"

	"explorer-api/env"
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
}

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("%s v%s Commit %s builded %s\n", AppName, Version, GitCommit, BuildDate)
		os.Exit(0)
	}
}

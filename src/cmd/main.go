package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"timetracker/cmd/time_tracker"
)

var (
	configPath string
	sqlDB      string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config.toml", "path to config file")
	flag.StringVar(&sqlDB, "sql-db", "postgres", "what sql-db the application uses")
}

func main() {
	flag.Parse()

	timeTracker := time_tracker.TimeTracker{}
	_, err := toml.DecodeFile(configPath, &timeTracker)

	if err != nil {
		log.Fatal(err)
	}

	err = timeTracker.Run()

	if err != nil {
		return
	}
}

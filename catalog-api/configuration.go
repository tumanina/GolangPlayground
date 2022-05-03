package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	ConnectionString string
	DatabaseName     string
}

func GetDatabaseConfiguration() DatabaseConfiguration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration.Database
}

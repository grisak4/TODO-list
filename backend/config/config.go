package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
}

var AppConfig Config

func Load() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config JSON: %v", err)
	}
}

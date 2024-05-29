package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ChromeDriverPath string `json:"/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver"`
	Port             int    `json:"8080"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}

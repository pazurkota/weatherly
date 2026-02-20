package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var apiKey string

func setApiKey() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	key := os.Getenv("API_KEY")
	apiKey = key
}

func GetApiKey() string {
	if apiKey == "" {
		setApiKey()
	}

	return apiKey
}

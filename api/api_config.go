package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getApiKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	key := os.Getenv("API_KEY")
	return key
}

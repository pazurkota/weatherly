package api

import (
	"fmt"
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
	fmt.Println(key)
	return key
}

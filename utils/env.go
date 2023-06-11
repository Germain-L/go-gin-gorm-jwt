package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Only load .env file if ENV=dev.
	// This is to prevent loading .env file in production and other environments.

	isProd := os.Getenv("ENV") == "prod"

	if !isProd {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}
}

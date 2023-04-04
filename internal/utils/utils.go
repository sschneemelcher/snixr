package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(fileNames ...string) {
    var err error

    if len(fileNames) > 0 {
        err = godotenv.Load(fileNames...)
    } else {
        err = godotenv.Load(".env")
    }

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}


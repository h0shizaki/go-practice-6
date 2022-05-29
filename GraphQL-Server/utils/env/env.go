package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MustGet(key string) string {
	val := os.Getenv(key)

	if val == "" {
		log.Panic("Env key is missing " + key)
	}

	return val
}

func CheckENV() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
}

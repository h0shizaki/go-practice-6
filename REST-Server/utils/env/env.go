package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func MustGet(key string) string {
	val := os.Getenv(key)

	if val == "" {
		fmt.Println("Env key is missing " + key)
	}

	return val
}

func CheckENV() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

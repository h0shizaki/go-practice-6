package env

import (
	"os"
	"server/log"

	"github.com/joho/godotenv"
)

func MustGet(key string) string {
	val := os.Getenv(key)

	if val == "" {
		// fmt.Println()
		log.Log.LogWarning("Env key is missing " + key)
	}

	return val
}

func CheckENV() {
	err := godotenv.Load()

	if err != nil {
		// fmt.Println("Error loading .env file")
		log.Log.LogDanger("Error loading .env file")
	}
}

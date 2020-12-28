package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config func to get env value from key ---
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)

}

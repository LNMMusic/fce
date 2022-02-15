package config

import (
	"os"
	"log"
	dotenv "github.com/joho/godotenv"
)

const filename = ".env"
func EnvLoad() {
	if err := dotenv.Load(filename); err != nil {
		log.Fatalf("Failed to load the .env file: %v", err)
	}
}

func EnvGet(key string) string {
	return os.Getenv(key)
}
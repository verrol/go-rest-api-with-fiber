package config

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("unable to load env file, falling back to default values", "file", ".env")
	}

	return os.Getenv(key)
}

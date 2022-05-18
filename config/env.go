package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvConfig(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	return os.Getenv(name)
}

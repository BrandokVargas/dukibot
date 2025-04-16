package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Token  string
	Prefix string
}

func Load() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error: can't load .env")
	}

	token := os.Getenv("TOKEN_DISCORD")
	prefix := os.Getenv("PREFIX")

	if token == "" {
		log.Fatalf("TOKEN_DISCORD is empty")
	}

	if prefix == "" {
		log.Fatalf("PREFIX is empty")
	}

	return &Config{Token: token, Prefix: prefix}
}

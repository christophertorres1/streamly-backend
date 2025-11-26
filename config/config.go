package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port     string
	MongoURI string
}

var C AppConfig

func Load() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("MONGO_URI must be set")
	}

	C = AppConfig{
		Port:     port,
		MongoURI: mongoURI,
	}
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Just a class (struct) that creates a new object with variables AppPort & DatabaseURL

type Config struct {
	AppPort     string
	DatabaseURL string
}

// Function Load that returns a type Config by looking up and calling function getEnv that takes
// AppPort and DatabaseURL from .env or assigns 3030 if it is empty in .env or does not exist there
func Load() Config {

	_ = godotenv.Load()

	return Config{
		AppPort:     getEnv("APP_PORT", "3030"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

// The function getEnv mentioned earlier that returns either data from .env or our set port (3030)
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type yacaldavConfig struct {
	YAUSER string
	CALPWD string
	YACAL  string
}

type Config struct {
	YaAuth yacaldavConfig
}

// New returns a new Config struct
func New() *Config {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	return &Config{
		YaAuth: yacaldavConfig{
			YAUSER: getEnv("YAUSER", ""),
			CALPWD: getEnv("CALPWD", ""),
			YACAL:  getEnv("YACAL", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

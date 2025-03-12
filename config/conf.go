package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type yacaldavConfig struct {
	YAUSER string
	CALPWD string
	YACAL  string
}

type net struct {
	Timeout time.Duration
}
type Config struct {
	YaAuth yacaldavConfig
	Net    net
}

// New returns a new Config struct
func LoadConifg() *Config {
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
		Net: net{
			Timeout: time.Millisecond * 3000,
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

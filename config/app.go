package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	IsDev         bool
	ApiVersion    string
	RedisHost     string
	RedisUsername string
	RedisPassword string
}

// envs
const (
	DEV            = "DEV"
	API_VERSION    = "API_VERSION"
	REDIS_HOST     = "REDIS_HOST"
	REDIS_USERNAME = "REDIS_USERNAME"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

func New() (*AppConfig, error) {

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	isDev, err := strconv.ParseBool(os.Getenv("DEV"))
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		IsDev:         isDev,
		ApiVersion:    os.Getenv(API_VERSION),
		RedisHost:     os.Getenv(REDIS_HOST),
		RedisUsername: os.Getenv(REDIS_USERNAME),
		RedisPassword: os.Getenv(REDIS_PASSWORD),
	}, nil
}

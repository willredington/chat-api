package redis

import (
	"com/willredington/chat-api/config"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func ConnectRedis(appConfig *config.AppConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisHost,
		Username: appConfig.RedisUsername,
		Password: appConfig.RedisPassword,
	})
}

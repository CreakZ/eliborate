package database

import (
	"fmt"
	"yurii-lib/pkg/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func InitRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%d", viper.GetString(config.RedisHost), viper.GetInt(config.RedisPort))

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString(config.RedisPassword),
		DB:       0,
	})
}

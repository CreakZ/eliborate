package database

import (
	"context"
	"fmt"
	"time"
	"yurii-lib/pkg/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitCache() *redis.Client {
	addr := fmt.Sprintf("%s:%d", viper.GetString(config.RedisHost), viper.GetInt(config.RedisPort))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString(config.RedisPassword),
		DB:       0,
	})

	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, err := client.Ping(c).Result(); err != nil {
		panic(fmt.Sprintf("failed to init redis client: %s", err.Error()))
	}

	return client
}

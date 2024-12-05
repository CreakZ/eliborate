package storage

import (
	"context"
	"eliborate/pkg/config"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCacheManager() *RedisCache {
	addr := fmt.Sprintf("%s:%d", viper.GetString(config.RedisHost), viper.GetInt(config.RedisPort))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString(config.RedisPassword),
		DB:       0,
	})

	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if result, err := client.Ping(c).Result(); err != nil {
		panic(fmt.Sprintf("Result: %s\nFailed to init redis client: %s", result, err.Error()))
	}

	return &RedisCache{
		client: client,
	}
}

func (rc RedisCache) SetInt(key string, val int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	status := rc.client.Set(ctx, key, val, 0)
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

func (rc RedisCache) GetInt(key string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	status := rc.client.Get(ctx, key)
	if err := status.Err(); err != nil {
		return 0, err
	}

	val, err := status.Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (rc RedisCache) Incr(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	status := rc.client.Incr(ctx, key)
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

func (rc RedisCache) Decr(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	status := rc.client.Decr(ctx, key)
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wichadak/eDNA/config"
)

func NewRedisConnection(config *config.RedisConfig) *redis.Client {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		fmt.Println("[Redis] Connection failed")
	} else {
		fmt.Println("[Redis] Connection to redis is established")
	}
	return rdb
}

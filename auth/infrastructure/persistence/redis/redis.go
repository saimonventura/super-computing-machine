package redis

import (
	"context"
	"log"

	"super-computing-machine/auth/infrastructure/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitializeRedis() *redis.Client {
	redisAddr := config.GetEnvWithDefault("REDIS_ADDR", "localhost:6379")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	return rdb
}

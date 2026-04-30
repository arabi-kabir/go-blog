package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RDB *redis.Client

func ConnectRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: redisHost + ":6379",
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic("Redis connection failed")
	}
}

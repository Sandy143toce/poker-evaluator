package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if redisHost == "" {
		redisHost = "localhost"
	}
	if redisPort == "" {
		redisPort = "6379"
	}

	redisURL := fmt.Sprintf("redis://%s:%s", redisHost, redisPort)

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		fmt.Printf("Warning: Failed to parse Redis URL: %v\n", err)
		return nil
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Warning: Failed to connect to Redis: %v\n", err)
		return nil
	}

	RedisClient = client
	fmt.Println("Successfully connected to Redis")
	return client
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

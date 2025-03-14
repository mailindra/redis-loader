package config

import (
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	redisAddr = "localhost:6379"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     "",
		DB:           0,
		PoolSize:     10,              // Number of pool connections
		MinIdleConns: 5,               // Minimum number of idle connections
		MaxRetries:   3,               // Maximum number of retries
		ReadTimeout:  time.Second * 2, // Read timeout
		WriteTimeout: time.Second * 2, // Write timeout
		PoolTimeout:  time.Second * 3, // Pool timeout
	})

	return client, nil
}

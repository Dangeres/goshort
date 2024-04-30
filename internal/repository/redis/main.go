// Package redis for redis usages
package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

// New redis constructor
func New() *redis.Client {
	url := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	return client
}

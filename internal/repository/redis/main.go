// Package redis for redis usages
package redis

import (
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

	return redis.NewClient(opts)
}

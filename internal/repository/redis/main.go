// Package redis for redis usages
package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis struct needs for custom uses redis client
type Redis struct {
	client *redis.Client
}

// Set uses for set key value to redis with ttl
func (r Redis) Set(ctx context.Context, key string, bdata []byte, ttl time.Duration) (string, error) {
	return r.client.Set(ctx, key, bdata, ttl).Result()
}

// Get uses for get value by key
func (r Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Close uses for close client
func (r Redis) Close() error {
	return r.client.Close()
}

// Ping uses for ping to redis server
func (r Redis) Ping() error {
	_, err := r.client.Ping(context.Background()).Result()

	return err
}

// New redis constructor
func New() *Redis {
	url := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	client := &Redis{
		client: redis.NewClient(opts),
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}

	return client
}

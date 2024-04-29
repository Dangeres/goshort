// Package handlers for all handlers app
package handlers

import "github.com/redis/go-redis/v9"

// URL struct for url handlers
type URL struct {
	redis *redis.Client
}

// New constructor
func New(redis *redis.Client) *URL {
	return &URL{
		redis: redis,
	}
}

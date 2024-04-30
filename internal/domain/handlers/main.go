// Package handlers for all handlers app
package handlers

import (
	"context"
	"time"
)

//go:generate minimock -i github.com/Dangeres/goshort/internal/domain/handlers.Rediser -o ./mocks/rediser_mock.go -n RediserMock -p mocks

// Rediser custom redis struct interface
type Rediser interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, []byte, time.Duration) (string, error)
	Ping() error
	Close() error
}

// URL struct for url handlers
type URL struct {
	redis Rediser
}

// New constructor
func New(redis Rediser) *URL {
	return &URL{
		redis: redis,
	}
}

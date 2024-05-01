// Package actions uses for action in app
package actions

import (
	"context"
	"time"
)

//go:generate minimock -i github.com/Dangeres/goshort/internal/domain/actions.Rediser -o ./mocks/rediser_mock.go -n RediserMock -p mocks

// Rediser custom redis struct interface
type Rediser interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, []byte, time.Duration) (string, error)
	Expire(context.Context, string) (int64, error)
	Ping() error
	Close() error
}

// Actions base struct with redis
type Actions struct {
	redis Rediser
}

// New constructor
func New(redis Rediser) *Actions {
	return &Actions{
		redis: redis,
	}
}

// Package handlers for all handlers app
package handlers

import (
	"context"
	"time"

	"github.com/Dangeres/goshort/internal/domain/actions"
	"github.com/Dangeres/goshort/internal/structures"
)

//go:generate minimock -i github.com/Dangeres/goshort/internal/domain/handlers.Rediser -o ./mocks/rediser_mock.go -n RediserMock -p mocks

// Rediser custom redis struct interface
type Rediser interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, []byte, time.Duration) (string, error)
	Expire(context.Context, string) (int64, error)
	Ping() error
	Close() error
}

//go:generate minimock -i github.com/Dangeres/goshort/internal/domain/handlers.Actionser -o ./mocks/actionser_mock.go -n ActionserMock -p mocks

// Actionser uses for action with data
type Actionser interface {
	Get(context.Context, structures.GetIn) (structures.InRedisData, error)
	Post(context.Context, structures.ShortIn) (structures.ShortOut, error)
}

// HandlerURL struct for url handlers
type HandlerURL struct {
	// redis Rediser
	actions Actionser
}

// New constructor
func New(redis Rediser) *HandlerURL {
	return &HandlerURL{
		actions: actions.New(redis),
	}
}

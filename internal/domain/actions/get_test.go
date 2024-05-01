package actions

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Dangeres/goshort/internal/domain/handlers/mocks"
	"github.com/Dangeres/goshort/internal/structures"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetSuccess(t *testing.T) {
	var (
		mc      = minimock.NewController(t)
		redis   = mocks.NewRediserMock(mc)
		actions = New(redis)
		ctx     = context.Background()
		getin   = structures.GetIn{URL: "sJDCKQmz2x", Redirect: false}
		url     = "https://ya.ru/"
		ttl     = 1440
		count   = 0
		result  = fmt.Sprintf("{\"ttl\":%d,\"count\":%d,\"url\":\"%s\"}", ttl, count, url)
		resultS = structures.InRedisData{TTL: int64(ttl), Count: uint(count), URL: url}
	)

	redis.GetMock.Expect(ctx, getin.URL).Return(result, nil)

	res, err := actions.Get(ctx, getin)

	assert.NoError(t, err)
	assert.Equal(t, res, resultS)
}

func TestGetError(t *testing.T) {
	var (
		mc      = minimock.NewController(t)
		redis   = mocks.NewRediserMock(mc)
		actions = New(redis)
		ctx     = context.Background()
		getin   = structures.GetIn{URL: "sJDCKQmz2x", Redirect: false}
		result  = ""
		resultS = structures.InRedisData{}
		rerr    = errors.New("RedisError")
	)

	redis.GetMock.Expect(ctx, getin.URL).Return(result, fmt.Errorf("%w: not found", rerr))

	res, err := actions.Get(ctx, getin)

	assert.ErrorIs(t, err, rerr)
	assert.Equal(t, res, resultS)
}

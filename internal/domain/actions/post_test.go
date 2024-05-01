package actions

import (
	"context"
	"testing"
	"time"

	"github.com/Dangeres/goshort/internal/constants"
	"github.com/Dangeres/goshort/internal/domain/handlers/mocks"
	"github.com/Dangeres/goshort/internal/structures"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestPostSuccessBasic(t *testing.T) {
	var (
		mc      = minimock.NewController(t)
		redis   = mocks.NewRediserMock(mc)
		actions = New(redis)
		ctx     = context.Background()
		url     = "https://ya.ru/"
		short   = "sJDCKQmz2x"
		ttl     = 1440
		count   = 0
		now     = time.Now()
		shortin = structures.ShortIn{TTL: int64(ttl), Count: uint(count), URL: url}
	)

	redis.SetMock.Set(
		func(_ context.Context, _ string, _ []byte, _ time.Duration) (string, error) {
			return short, nil
		},
	)

	redis.ExpireMock.Set(
		func(_ context.Context, _ string) (int64, error) {
			return now.Unix(), nil
		},
	)

	res, err := actions.Post(ctx, shortin)

	assert.NoError(t, err)

	assert.Equal(t, len(res.SURL), constants.LenShort)
	assert.Equal(t, res.TTL, now.Unix())
}

func TestPostSuccessOverTTL(t *testing.T) {
	var (
		mc      = minimock.NewController(t)
		redis   = mocks.NewRediserMock(mc)
		actions = New(redis)
		ctx     = context.Background()
		url     = "https://ya.ru/"
		short   = "sJDCKQmz2x"
		ttl     = constants.MaxTTL * 10
		count   = 0
		now     = time.Now()
		shortin = structures.ShortIn{TTL: int64(ttl), Count: uint(count), URL: url}
	)

	redis.SetMock.Set(
		func(_ context.Context, _ string, _ []byte, _ time.Duration) (s2 string, err error) {
			return short, nil
		},
	)

	redis.ExpireMock.Set(
		func(_ context.Context, _ string) (int64, error) {
			return now.Unix() + constants.MaxTTL, nil
		},
	)

	res, err := actions.Post(ctx, shortin)

	assert.NoError(t, err)

	assert.Equal(t, len(res.SURL), constants.LenShort)
	assert.Equal(t, res.TTL-now.Unix(), int64(constants.MaxTTL))
}

// Package shorter uses for short url
package shorter

import (
	"testing"

	"github.com/Dangeres/goshort/internal/cerrors"
	"github.com/Dangeres/goshort/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestShortLinkSuccess(t *testing.T) {
	t.Parallel()

	res := ShortLink()

	assert.Equal(t, len(res), constants.LenShort)
}

func TestGenerateRandomBytesError(t *testing.T) {
	t.Parallel()

	var (
		n = -1
	)

	_, err := GenerateRandomString(n)

	assert.ErrorIs(t, err, cerrors.ErrorNegativeN)
}

func TestGenerateRandomBytesZero(t *testing.T) {
	t.Parallel()

	var (
		n = 0
	)

	res, err := GenerateRandomString(n)

	assert.NoError(t, err)

	assert.Equal(t, res, "")
}

func TestGenerateRandomBytesSuccess(t *testing.T) {
	t.Parallel()

	var (
		n = 10
	)

	res, err := GenerateRandomString(n)

	assert.NoError(t, err)

	assert.Equal(t, len(res), n)
}

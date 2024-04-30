// Package shorter uses for short url
package shorter

import (
	"testing"

	"github.com/Dangeres/goshort/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestShortLink(t *testing.T) {
	t.Parallel()

	res := ShortLink()

	assert.Equal(t, len(res), constants.LenShort)
}

// Package shorter uses for short url
package shorter

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/Dangeres/goshort/internal/cerrors"
	"github.com/Dangeres/goshort/internal/constants"
)

var (
	letters  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	maxValue = big.NewInt(int64(len(letters)))
)

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("%w: must be non negative", cerrors.ErrorNegativeN)
	}

	ret := make([]rune, n)

	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, maxValue)

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// ShortLink uses for generate short link with len = LenShort
func ShortLink() string {
	res, _ := GenerateRandomString(constants.LenShort)

	return res
}

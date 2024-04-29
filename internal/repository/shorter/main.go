// Package shorter uses for short url
package shorter

import (
	"crypto/rand"
	"math/big"

	"github.com/Dangeres/goshort/internal/constants"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(n int) (string, error) {
	ret := make([]rune, n)

	maxValue := big.NewInt(int64(len(letters)))

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, maxValue)

		if err != nil {
			return "", err
		}

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// ShortLink uses for generate short link with len = LenShort
func ShortLink() string {
	var (
		res string
		err error
	)

	res, err = GenerateRandomString(constants.LenShort)

	for err != nil {
		res, err = GenerateRandomString(constants.LenShort)
	}

	return res
}

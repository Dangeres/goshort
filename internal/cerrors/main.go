// Package cerrors uses for custom errors in app
package cerrors

import "errors"

var (
	// ErrorNegativeN uses when N must be >= 0
	ErrorNegativeN = errors.New("NegativeN")
)

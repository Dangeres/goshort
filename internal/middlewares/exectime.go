// Package middlewares uses for middlewares
package middlewares

import (
	"log"
	"net/http"
	"time"
)

// CalcTime for calculating time execution
func CalcTime(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		f(w, r)

		log.Printf(
			"[%s] %s - %d ms\n",
			r.Method,
			r.RequestURI,
			time.Since(timeStart)/time.Millisecond,
		)
	}
}

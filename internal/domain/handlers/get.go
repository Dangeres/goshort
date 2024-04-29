// Package handlers for all handlers app
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dangeres/goshort/internal/constants"
	"github.com/Dangeres/goshort/internal/structures"
)

// HUnShort uses for unshort url
func (url URL) HUnShort(w http.ResponseWriter, r *http.Request) {
	// dat, _ := io.ReadAll(r.Body)

	// log.Println(string(dat))

	ctx := r.Context()

	surl := r.PathValue(constants.PathURL)

	rget := url.redis.Get(ctx, surl)

	rdata, err := rget.Result()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	boutdata := []byte(rdata)

	pr := structures.InRedisData{}

	err = json.Unmarshal(boutdata, &pr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if pr.Redirect {
		http.Redirect(w, r, pr.URL, http.StatusSeeOther)
		return
	}

	_, err = w.Write(boutdata)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
}

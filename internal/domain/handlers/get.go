// Package handlers for all handlers app
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dangeres/goshort/internal/structures"
	"github.com/gorilla/schema"
)

// HUnShort uses for unshort url
func (url URL) HUnShort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	getin := structures.GetIn{}

	decoder := schema.NewDecoder()

	err := decoder.Decode(&getin, r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	rdata, err := url.redis.Get(ctx, getin.URL)

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

	if getin.Redirect {
		http.Redirect(w, r, pr.URL, http.StatusSeeOther)
		return
	}

	_, err = w.Write(boutdata)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

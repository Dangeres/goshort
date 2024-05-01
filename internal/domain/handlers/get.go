// Package handlers for all handlers app
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dangeres/goshort/internal/structures"
	"github.com/gorilla/schema"
)

// UnShort uses for unshort url
func (hu HandlerURL) UnShort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	getin := structures.GetIn{}

	decoder := schema.NewDecoder()

	err := decoder.Decode(&getin, r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	pr, err := hu.actions.Get(ctx, getin)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	boutdata, err := json.Marshal(&pr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if getin.Redirect {
		http.Redirect(w, r, pr.URL, http.StatusPermanentRedirect)
		return
	}

	_, err = w.Write(boutdata)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

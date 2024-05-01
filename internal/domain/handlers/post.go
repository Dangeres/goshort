// Package handlers for all handlers app
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Dangeres/goshort/internal/structures"
)

// HShort uses for short link
func (hu HandlerURL) Short(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bdata, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := structures.ShortIn{}

	err = json.Unmarshal(bdata, &data)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pr, err := hu.actions.Post(ctx, data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	obdata, err := json.Marshal(&pr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = w.Write(obdata)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
}

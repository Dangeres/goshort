// Package handlers for all handlers app
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Dangeres/goshort/internal/repository/shorter"
	"github.com/Dangeres/goshort/internal/structures"
)

// HShort uses for short link
func (url URL) HShort(w http.ResponseWriter, r *http.Request) {
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

	if data.TTL == 0 {
		data.TTL = 24 * 60
	}

	hashedURL := shorter.ShortLink()

	inredisdata, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	res := url.redis.Set(ctx, hashedURL, inredisdata, time.Minute*time.Duration(data.TTL))

	if res.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	result := structures.ShortOut{
		TTL:  time.Now().Unix() + data.TTL*60,
		SURL: hashedURL,
	}

	obdata, err := json.Marshal(&result)

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

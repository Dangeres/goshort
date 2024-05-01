package actions

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Dangeres/goshort/internal/constants"
	"github.com/Dangeres/goshort/internal/repository/shorter"
	"github.com/Dangeres/goshort/internal/structures"
)

func (a Actions) Post(ctx context.Context, data structures.ShortIn) (structures.ShortOut, error) {
	if data.TTL == 0 {
		data.TTL = 24 * 60
	}

	if data.TTL > constants.MaxTTL {
		data.TTL = constants.MaxTTL
	}

	shortURL := shorter.ShortLink()

	inredisdata, err := json.Marshal(data)

	if err != nil {
		return structures.ShortOut{}, err
	}

	_, err = a.redis.Set(ctx, shortURL, inredisdata, time.Minute*time.Duration(data.TTL))

	if err != nil {
		return structures.ShortOut{}, err
	}

	dd, err := a.redis.Expire(ctx, shortURL)

	if err != nil {
		return structures.ShortOut{}, err
	}

	result := structures.ShortOut{
		TTL:  int64(dd),
		SURL: shortURL,
	}

	return result, nil
}

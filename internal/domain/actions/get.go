package actions

import (
	"context"
	"encoding/json"

	"github.com/Dangeres/goshort/internal/structures"
)

func (a Actions) Get(ctx context.Context, getin structures.GetIn) (structures.InRedisData, error) {
	rdata, err := a.redis.Get(ctx, getin.URL)

	if err != nil {
		return structures.InRedisData{}, err
	}

	boutdata := []byte(rdata)

	pr := structures.InRedisData{}

	err = json.Unmarshal(boutdata, &pr)

	return pr, err
}

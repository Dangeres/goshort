// Package structures uses for all structures
package structures

// ShortIn uses like input schema for short request
type ShortIn struct {
	TTL   int64  `json:"ttl"`
	Count uint   `json:"count"`
	URL   string `json:"url"`
}

// ShortOut uses like output schema for short request
type ShortOut struct {
	TTL  int64  `json:"ttl"`
	SURL string `json:"surl"`
}

// GetIn uses like input schema for get link request
type GetIn struct {
	URL      string `json:"url" schema:"url"`
	Redirect bool   `json:"redirect" schema:"redirect"`
}

// GetOut uses like output schema for get link request
type GetOut struct {
	URL string `json:"url"`
}

// InRedisData uses like schema for save in redis storage
type InRedisData struct {
	ShortIn
}

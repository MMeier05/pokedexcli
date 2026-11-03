package pokeapi

import (
	"time"
	"github.com/MMeier05/pokedexcli/internal/pokecache"
)
type Client struct {
	cache *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
	}
}

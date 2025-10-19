package pokeapi

import (
    "net/http"
    "time"

    "github.com/Ichise5/pokedexcli/internal/pokecache"
)

// Client wraps an HTTP client for PokeAPI requests
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with the specified timeout
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
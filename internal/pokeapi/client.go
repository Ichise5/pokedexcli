package pokeapi

import (
    "net/http"
    "time"
)

// Client wraps an HTTP client for PokeAPI requests
type Client struct {
    httpClient http.Client
}

// NewClient creates a new PokeAPI client with the specified timeout
func NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}
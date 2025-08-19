package pokeapi

// RespShallowLocations represents the response from the location-area endpoint
type RespShallowLocations struct {
    Count    int     `json:"count"`
    Next     *string `json:"next"`     // Note: pointer to handle null values
    Previous *string `json:"previous"` // Note: pointer to handle null values
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}
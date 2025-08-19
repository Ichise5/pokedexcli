package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"` // Fixed syntax: no space after colon
	Results  []LocationName `json:"results"`  // The API's JSON key is "results"
}

// LocationName struct fields must also be capitalized.
type LocationName struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


func GetMap(config *conf) error{
	var url string
	//PokeAPIUrl := "https://pokeapi.co/api/v2/location-area/"

	if config.forward{
		if config.next == ""{
			url = "https://pokeapi.co/api/v2/location-area/"
		}else{
			url = config.next
		}
	}else{
		if config.previous == ""{
			fmt.Println("you're on the first page")
			return nil
		}else{
			url = config.previous
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)	
	}

	if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)	
		}


	var locs Location

	if err := json.Unmarshal(body, &locs); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for _, locat := range locs.Results {
		fmt.Println(locat.Name)
	}
	
	config.next = locs.Next
	config.previous = locs.Previous

	return nil
}


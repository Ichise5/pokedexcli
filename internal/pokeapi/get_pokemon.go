package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (RespPokemon, error) {
	url := baseURL + "/pokemon"
	if name != nil {
		url = url + "/" + *name
	} else {
		return RespPokemon{}, fmt.Errorf("location was not provided")
	}

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	Pokemon := RespPokemon{}
	err = json.Unmarshal(dat, &Pokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	return  Pokemon, nil
}
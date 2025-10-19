package main

import (
    "time"

    "github.com/Ichise5/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
    cfg := &config{
        pokeapiClient: pokeClient,
        maxExp: 500,
    }

    startRepl(cfg)
}
package main

import "fmt"

func commandExplore(cfg *config, param *string) error {

	locationInfo, err := cfg.pokeapiClient.ExploreLocation(param)
    if err != nil {
        return err
    }

	fmt.Printf("Exploring %s...\n", locationInfo.Name)
	fmt.Println("Found Pokemon:")

	for _, pok := range locationInfo.PokemonEncounters {
        fmt.Printf(" - %s\n", pok.Pokemon.Name)
    }
    return nil
}
package main

import(
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param *string) error{
	fmt.Printf("Throwing a Pokeball at %s...\n", *param)

	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(param)
    if err != nil {
        return err
    }

	for{
		
	if pokemonInfo.BaseExperience > cfg.maxExp{
			cfg.maxExp += 100
 		} else{
			break
		}
	}
	
	pokemonInfo.Seen = true
	if rand.Intn(cfg.maxExp) > pokemonInfo.BaseExperience{
		pokemonInfo.Caught = true
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		cfg.pokemons[pokemonInfo.Name] = pokemonInfo
	}else{
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}

	return nil
}
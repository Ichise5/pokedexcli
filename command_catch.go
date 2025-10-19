package main

import(
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param *string) error{
	fmt.Printf("Throwing a Pokeball at %s...", *param)

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

	if rand.Intn(cfg.maxExp) > pokemonInfo.BaseExperience{
		fmt.Printf("%s was caught!", pokemonInfo.Name)
		cfg.pokemon[pokemonInfo.Name] = pokemonInfo
	}else{
		fmt.Printf("%s escaped!", pokemonInfo.Name)
	}

	return nil
}
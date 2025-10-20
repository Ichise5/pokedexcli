package main

import (
	"fmt"

)

func commandPokedex(cfg *config, param *string)error{

	if len(cfg.pokemons) == 0{

		fmt.Println("You have not seen any pokemon yet.")
		return nil
	}

	for key, value := range cfg.pokemons {
		status := ""

		if value.Seen{
			status = "(seen)"
		}
		if value.Caught{
			status = "(caught)"
		}
    	fmt.Printf(" - %s %s\n", key, status)
	}
	return nil
}
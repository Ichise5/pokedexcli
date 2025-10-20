package main

import "fmt"

func commandInspect(cfg *config, param *string) error{

	val, ok := cfg.pokemons[*param]

	if !ok {
		fmt.Println("Pokemon has not been found - it was not seen nor caught.")
		return nil
	}



	fmt.Printf("Name: %s\n", val.Name)
	if val.Caught{
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Println("Stats:")
		for i := 0; i < len(val.Stats); i++ {
			fmt.Printf("  -%s: %d\n", val.Stats[i].Stat.Name, val.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(val.Types); i++ {
			fmt.Printf("  - %s\n", val.Types[i].Type.Name)
		}
	}
	return nil
}
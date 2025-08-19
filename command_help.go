package main

import "fmt"

func commandHelp(cfg *config) error{
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands(){
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
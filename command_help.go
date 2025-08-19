package main

import "fmt"

func commandHelp(config *conf) error{
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands(){
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
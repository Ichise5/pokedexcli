package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"log"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	var config conf
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		
		line := scanner.Text()
		if line == ""{
			fmt.Println("Your have not entered a command")
			continue
		}
		words := cleanInput(line)

		foundCommand := false
		for _,command := range commands(){
			if command.name == words[0]{

				switch words[0] {
					case "map":
						config.forward = true
					case "mapb":
						config.forward = false
				}

				err := command.callback(&config)
				if err != nil{
					fmt.Println(err)
				}
				foundCommand = true
			}
		}
		if !foundCommand{
			fmt.Println("Unknown command")
		}

		
	}

}

func cleanInput(text string) []string {
	var res []string

	seq := strings.SplitSeq(text, " ")

	for word := range seq {
		if word != "" {
			res = append(res, strings.ToLower(word))
		}

	}

	return res
}

type cliCommand struct {
	name        string
	description string
	callback    func(*conf) error
}

type conf struct{
	next 		string
	previous	string
	forward		bool
}

func commands() map[string]cliCommand{
	return map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
	"help":{
		name:		 "help",
		description: "Displays a help message",
		callback: 	 commandHelp,
	},
	"map":{
		name:		 "map",
		description: "Displays next 20 locations",
		callback: 	 GetMap,
	},
	"mapb":{
		name:		 "mapb",
		description: "Display previous 20 locations",
		callback: 	 GetMap,
	},
}

}




package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/Ichise5/pokedexcli/internal/pokeapi"
)

// Updated config struct
type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
    maxExp           int
    pokemons         map[string] pokeapi.RespPokemon
}

func startRepl(cfg *config) {  // Now takes config parameter
    scanner := bufio.NewScanner(os.Stdin)
    
    cfg.pokemons = make(map[string] pokeapi.RespPokemon)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan() 
        if err := scanner.Err(); err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }
        
        line := scanner.Text()
        if line == "" {
            continue
        }
        words := cleanInput(line)
        if len(words) == 1{
            words = append(words, "")
        }

        foundCommand := false
        for _, command := range getCommands() {  // Renamed from commands()
            if command.name == words[0] {
                err := command.callback(cfg, &words[1])  // Pass config and parameter
                if err != nil {
                    fmt.Println(err)
                }
                foundCommand = true
                break
            }
        }
        if !foundCommand {
            fmt.Println("Unknown command")
        }
    }
}

// Renamed and simplified
func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message.",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Get the next page of locations.",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Get the previous page of locations.",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Get list of pokemons in a location. This command requires a parameter.",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Try to catch a pokemon.",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "Retuns pokemon information in case it has been seen or caught. This command requires a parameter.",
            callback:    commandInspect,
        },
    }
}

// Updated cliCommand type
type cliCommand struct {
    name        string
    description string
    callback    func(*config, *string) error  // Now takes config parameter
}

// Keep your cleanInput function as is - it's fine
func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}
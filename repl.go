package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/bootdotdev/pokedexcli/internal/pokeapi"
)

// Updated config struct
type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
}

func startRepl(cfg *config) {  // Now takes config parameter
    scanner := bufio.NewScanner(os.Stdin)
    
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

        foundCommand := false
        for _, command := range getCommands() {  // Renamed from commands()
            if command.name == words[0] {
                err := command.callback(cfg)  // Pass config
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
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Get the next page of locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Get the previous page of locations",
            callback:    commandMapb,
        },
    }
}

// Updated cliCommand type
type cliCommand struct {
    name        string
    description string
    callback    func(*config) error  // Now takes config parameter
}

// Keep your cleanInput function as is - it's fine
func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}
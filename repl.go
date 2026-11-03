package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/MMeier05/pokedexcli/internal/pokeapi"
)
type config struct{
	commands map[string]cliCommand
	areaNext string
	areaPrev any 
	pokeapiClient pokeapi.Client
	pokedex map[string]pokeapi.PokemonData
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
} 

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	lowered := strings.ToLower(trimmed)
	result := strings.Fields(lowered)
	return result
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning: %v", err)
			return
		}
		rawInput := scanner.Text()
		cleanInput := cleanInput(rawInput)
		cmd, ok := conf.commands[cleanInput[0]]
		if !ok {
			fmt.Print("Unknown command\n")
		} else {
			// Commands with arg handeled here
			if len(cleanInput) == 2 {
				err := cmd.callback(conf, cleanInput[1])
				if err != nil {
					fmt.Printf("Error: %v", err)
				}
			} else {
				// Commands without arg handeled here
				err := cmd.callback(conf, "")
				if err != nil {
					fmt.Printf("Error: %v", err)
				}
			}
		}
	}
}

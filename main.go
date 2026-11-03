package main

import (
	"time"
	"github.com/MMeier05/pokedexcli/internal/pokeapi"
)
func main() {
	conf := config{
		commands: map[string]cliCommand{
			"exit": {
				name:	"exit",
				description: "Exit the Pokedex",
				callback: commandExit,
			},
			"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
			"map": {
				name: "map",
				description: "Display next 20 location areas",
				callback: commandMap,
			},
			"mapb": {
				name: "mapb",
				description: "Display previous 20 location areas",
				callback: commandMapb,
			},
			"explore": {
				name: "explore",
				description: "Display the Pokemon found in the specified area",
				callback: commandExplore,
			},
			"catch": {
				name: "catch",
				description: "Catch a Pokemon",
				callback: commandCatch,
			},
			"inspect": {
				name: "inspect",
				description: "Display details about the Pokemon",
				callback: commandInspect,
			},
		},
		areaNext: "https://pokeapi.co/api/v2/location-area/",
		areaPrev: nil,
		pokeapiClient: pokeapi.NewClient(6 * time.Second),
		pokedex: map[string]pokeapi.PokemonData{},
	}
	startRepl(&conf)
}


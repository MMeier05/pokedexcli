package main

import (
	"fmt"
)


func commandExplore(conf *config, arg string) error {
	fullURL := "https://pokeapi.co/api/v2/location-area/" + arg
	pokemon, err := conf.pokeapiClient.GetEncounters(fullURL)
	if err != nil {
		return err
	}
	for _, p := range pokemon.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil

}

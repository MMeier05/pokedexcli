package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, pokemonName string) error {
	fullURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	data, err := conf.pokeapiClient.GetPokemonData(fullURL)
	if err != nil {
		return err 
	}
	if data.Name != pokemonName {
		return nil
	}
	randInt := rand.Intn(data.BaseExperience / 17)
	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name) 
	if randInt == 0 {
		fmt.Printf("%s was caught!\n", data.Name)
		conf.pokedex[data.Name] = data
		return nil
	}
	fmt.Printf("escaped!\n")


	return nil
}

package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

type Encounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
} `json:"pokemon_encounters"`
}

func commandExplore(conf *config, arg string) error {
	fullURL := "https://pokeapi.co/api/v2/location-area/" + arg
	// First check cache
	var data []byte
	data, ok := conf.cache.Get(fullURL)
	// Add area to cache on miss
	if !ok {
		res ,err := http.Get(fullURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.cache.Add(fullURL, data)
	}
	pokemon := Encounters{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}
	for _, p := range pokemon.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil

}

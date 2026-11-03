package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

type AreaData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} 
}

type Encounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonData struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		Stat struct {
			StatName string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetLocations(url string) (AreaData, error) {
	data, ok := c.cache.Get(url)
	if ok {
		areaDat := AreaData{}
		if err := json.Unmarshal(data, &areaDat); err != nil {
			return AreaData{}, err 
		}
		return areaDat, nil 
	}
	res, err := http.Get(url)
	if err != nil {
		return AreaData{} , err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return AreaData{}, err
	}
	areaDat := AreaData{}
	if err := json.Unmarshal(data, &areaDat); err != nil {
		return AreaData{}, err 
	}
	return areaDat, nil 
}

func (c *Client) GetEncounters(url string) (Encounters, error) {
	data, ok := c.cache.Get(url)
	if ok {
		encounterDat := Encounters{}
		if err := json.Unmarshal(data, &encounterDat); err != nil {
			return Encounters{}, err 
		}
		return encounterDat, nil 
	}
	res, err := http.Get(url)
	if err != nil {
		return Encounters{} , err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Encounters{}, err
	}
	encounterDat := Encounters{}
	if err := json.Unmarshal(data, &encounterDat); err != nil {
		return Encounters{}, err 
	}
	return encounterDat, nil 
}

func (c *Client) GetPokemonData(url string) (PokemonData, error) {
	data, ok := c.cache.Get(url)
	if ok {
		pokemonData := PokemonData{}
		if err := json.Unmarshal(data, &pokemonData); err != nil {
			return PokemonData{}, err 
		}
		return pokemonData, nil 
	}
	res, err := http.Get(url)
	if err != nil {
		return PokemonData{} , err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}
	pokemonData := PokemonData{}
	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return PokemonData{}, err 
	}
	return pokemonData, nil 
}

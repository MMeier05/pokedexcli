package main

import (
	"fmt"
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


func commandMap(conf *config, _ string) error {
	areaDat, err := conf.pokeapiClient.GetLocations(conf.areaNext)
	if err != nil {
		return err
	}
	conf.areaNext = areaDat.Next
	conf.areaPrev = areaDat.Previous
	for _, name := range areaDat.Results {
		fmt.Printf("%s\n", name.Name)
	}
	return nil
}

func commandMapb(conf *config, _ string) error {
	prev, ok := conf.areaPrev.(string)
	if !ok {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	areaDat, err := conf.pokeapiClient.GetLocations(prev)
	if err != nil {
		return err 
	}
	conf.areaNext = areaDat.Next
	conf.areaPrev = areaDat.Previous
	for _, name := range areaDat.Results {
		fmt.Printf("%s\n", name.Name)
	}
	return nil
}

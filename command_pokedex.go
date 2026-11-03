package main

import (
	"fmt"
)

func commandPokedex(conf *config, _ string) error {
	fmt.Print("Your Pokedex:\n")
	for _, p := range conf.pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}

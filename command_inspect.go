package main

import (
	"fmt"
)

func commandInspect(conf *config, pokemonName string) error{
	pokemon, ok := conf.pokedex[pokemonName]
	if !ok {
		fmt.Print("you have not caught that pokemon\n")
		return nil
	}
	details := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\n", 
		pokemon.Name, 
		pokemon.Height,
		pokemon.Weight,
	)
	stats := "Stats:\n"
	for _, s := range pokemon.Stats {
		stats += fmt.Sprintf("  -%s: %d\n", s.Stat.StatName, s.BaseStat)
	}
	types := "Types:\n"
	for _, t := range pokemon.Types {
		types += fmt.Sprintf("  - %s\n", t.Type.Name)
	}
	fmt.Print(details, stats, types)
	return nil
}

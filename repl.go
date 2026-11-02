package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/MMeier05/pokedexcli/internal"
)
type config struct{
	commands map[string]cliCommand
	areaNext string
	areaPrev any 
	cache *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
} 

type AreaData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} 
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
			fmt.Print("Unknown command")
		} else {
			err := cmd.callback(conf)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	}
}

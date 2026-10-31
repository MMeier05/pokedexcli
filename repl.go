package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
	"net/http"
	"encoding/json"
)
type config struct{
	commands map[string]cliCommand
	areaNext string
	areaPrev any 
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

//-----COMMANDS-----
func commandExit(conf *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	usageString := "Welcome to the Pokedex!\nUsage:\n\n"
	replCommands := conf.commands
	for cmd := range replCommands {
		commandRepr := fmt.Sprintf("%s: %s\n", replCommands[cmd].name, replCommands[cmd].description)
		usageString += commandRepr
	}
	fmt.Print(usageString)
	return nil
}
func AreaGet(url string) (AreaData, error) {
	res, err := http.Get(url)
	if err != nil {
		return AreaData{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaData{}, err
	}
	areaDat := AreaData{}
	if err = json.Unmarshal(data, &areaDat); err != nil {
		return AreaData{}, err 
	}
	return areaDat, nil

}

func commandMap(conf *config) error {
	areaDat, err := AreaGet(conf.areaNext)
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

func commandMapb(conf *config) error {
	prev, ok := conf.areaPrev.(string)
	if !ok {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	areaDat, err := AreaGet(prev)
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

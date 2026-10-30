package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)
type config struct{
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

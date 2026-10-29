package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
} 

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	lowered := strings.ToLower(trimmed)
	result := strings.Fields(lowered)
	return result
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func startRepl() {
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
		cmd, ok := getCommands()[cleanInput[0]]
		if !ok {
			fmt.Print("Unknown command")
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	}
}

//-----COMMANDS-----
func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	usageString := "Welcome to the Pokedex!\nUsage:\n\n"
	replCommands := getCommands()
	for cmd := range replCommands {
		commandRepr := fmt.Sprintf("%s: %s\n", replCommands[cmd].name, replCommands[cmd].description)
		usageString += commandRepr
	}
	fmt.Print(usageString)
	return nil

}

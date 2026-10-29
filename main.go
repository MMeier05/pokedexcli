package main

import (
	"bufio"
	"fmt"
	"os"
)
type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func main() {

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
		replCommands := getCommands()
		cmd, ok := replCommands[cleanInput[0]]
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


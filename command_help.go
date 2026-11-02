package main

import (
	"fmt"
)

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

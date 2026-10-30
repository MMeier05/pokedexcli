package main

func main() {
	conf := config{
			commands: map[string]cliCommand{
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
			},
		}

	startRepl(&conf)
}


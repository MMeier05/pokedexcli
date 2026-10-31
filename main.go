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
				"map": {
					name: "map",
					description: "Display next 20 location areas",
					callback:commandMap,
				},
				"mapb": {
					name: "mapb",
					description: "Display previous 20 location areas",
					callback:commandMapb,
				},
			},
		areaNext: "https://pokeapi.co/api/v2/location-area/",
		areaPrev: nil,
		}

	startRepl(&conf)
}


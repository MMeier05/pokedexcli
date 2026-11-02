package main

import (
	"os"
	"fmt"
)

func commandExit(conf *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

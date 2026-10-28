package main

import (
		"fmt"
		"bufio"
		"os"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ;i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning: %v", err) 
			return
		}
		inp := scanner.Text()
		result := cleanInput(inp)
		inp = ""
		fmt.Printf("Your command was: %s\n",result[0])
	}
}


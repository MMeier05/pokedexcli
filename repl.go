package main

import (
		"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	result := strings.Split(trimmed, " ")
	return result
}

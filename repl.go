package main

import "strings"

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	fields := strings.Fields(lowerText)
	return fields
}

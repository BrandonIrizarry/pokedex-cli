package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) (words []string) {
	lowerText := strings.ToLower(text)
	lowerTextTrimmed := strings.TrimSpace(lowerText)
	words = strings.Split(lowerTextTrimmed, " ")

	return words
}

func main() {
	fmt.Println("Hello, World!")
}

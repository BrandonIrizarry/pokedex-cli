package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) (words []string) {
	lowerText := strings.ToLower(text)
	lowerTextTrimmed := strings.TrimSpace(lowerText)
	words = strings.Split(lowerTextTrimmed, " ")

	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "

	for fmt.Print(prompt); scanner.Scan(); fmt.Printf("\n%s", prompt) {
		line := scanner.Text()
		textTokens := cleanInput(line)

		fmt.Printf("Your command was: %s", textTokens[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}
}

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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	return nil
}

var commandRegistry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "

	for fmt.Print(prompt); scanner.Scan(); fmt.Printf("\n%s", prompt) {
		line := scanner.Text()
		textTokens := cleanInput(line)
		command := textTokens[0]

		what := commandRegistry[command]

		if what.name == "" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}
}

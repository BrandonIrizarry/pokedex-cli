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

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}
}

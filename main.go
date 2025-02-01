package main

import (
	"bufio"
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokeapi"
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
	callback    func(*pokeapi.OverworldPage, *pokeapi.RegionInfoPage, ...string) error
}

var commandRegistry = make(map[string]cliCommand)

// For now, this is just a dummy command.
func commandExit(_ *pokeapi.OverworldPage, _ *pokeapi.RegionInfoPage, _ ...string) error {
	return nil
}

func commandHelp(_ *pokeapi.OverworldPage, _ *pokeapi.RegionInfoPage, _ ...string) error {
	fmt.Printf("Usage:\n\n")
	for commandName, clicmd := range commandRegistry {
		fmt.Printf("%s: %s\n", commandName, clicmd.description)
	}

	return nil
}

// List the placenames found in the current page.
func commandMapForward(page *pokeapi.OverworldPage, _ *pokeapi.RegionInfoPage, _ ...string) error {
	var loader func(*pokeapi.OverworldPage) error

	// If 'map' is called for the first time, we bootstrap into the
	// forward/backward pagination by listing the first page of
	// results.
	if page.Next == nil {
		loader = pokeapi.LoadFirstURL
	} else {
		loader = pokeapi.LoadNextURL
	}

	err := loader(page)

	if err != nil {
		return err
	}

	placeNames := pokeapi.GetPlaceNames(page)

	for _, placeName := range placeNames {
		fmt.Println(placeName)
	}

	return nil
}

func commandMapBackward(page *pokeapi.OverworldPage, _ *pokeapi.RegionInfoPage, _ ...string) error {
	if page.Previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	err := pokeapi.LoadPreviousURL(page)

	if err != nil {
		return err
	}

	placeNames := pokeapi.GetPlaceNames(page)

	for _, placeName := range placeNames {
		fmt.Println(placeName)
	}

	return nil
}

func commandExplore(page *pokeapi.OverworldPage, regionInfo *pokeapi.RegionInfoPage, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("Wrong number of arguments to 'explore': %v\n", len(args))
		return nil
	}

	err := pokeapi.LoadRegionInfo(page, regionInfo, args[0])

	if err != nil {
		return err
	}

	fmt.Println(regionInfo)

	return nil
}

func main() {
	// Define the registry here, in main.
	commandRegistry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commandRegistry["help"] = cliCommand{
		name:        "help",
		description: "Display names of commands with their descriptions",
		callback:    commandHelp,
	}

	commandRegistry["map"] = cliCommand{
		name:        "map",
		description: "paginate forwards",
		callback:    commandMapForward,
	}

	commandRegistry["mapb"] = cliCommand{
		name:        "mapb",
		description: "paginate backwards",
		callback:    commandMapBackward,
	}

	commandRegistry["explore"] = cliCommand{
		name:        "explore",
		description: "list the pokemon found in the place-name given as the first argument",
		callback:    commandExplore,
	}

	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "

	// This is where the currently loaded page will reside.
	var page pokeapi.OverworldPage

	// This is where we will keep information about the current
	// region.
	var regionInfo pokeapi.RegionInfoPage

	for fmt.Print(prompt); scanner.Scan(); fmt.Printf("\n%s", prompt) {
		line := scanner.Text()
		textTokens := cleanInput(line)
		command := textTokens[0]

		// This slice holds the arguments passed to the given command
		// (e.g., the name of region to explore when invoking
		// 'explore'.)
		args := textTokens[1:]

		what := commandRegistry[command]

		// Check for a missing hashmap value. Unfortunately, we can't
		// compare 'what' directly with 'cliCommand{}', since
		// 'cliCommand' includes a func field which breaks
		// comparability.
		if what.name == "" {
			fmt.Printf("Invalid command '%s'\n", command)
			continue
		}

		if err := what.callback(&page, &regionInfo, args...); err != nil {
			fmt.Fprintf(os.Stderr, "Error in command '%s': %v\n", what.name, err)
		}

		// In case 'commandExit' should contain extra logic in the
		// future, we check for the presence of the exit command
		// _after_ running the current command's callback.
		if what.name == "exit" {
			fmt.Println("Closing the Pokedex... Goodbye!")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}
}

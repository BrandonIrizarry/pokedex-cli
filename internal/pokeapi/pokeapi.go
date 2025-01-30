package pokeapi

import (
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokecache"
)

var firstLoaded bool = false

// We include query parameters here, since calls to 'mapb' from the
// second page will add these anyway.
const pokeapiURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

// Load the previous page in the sequence.
func LoadPreviousURL(page *OverworldPage) error {
	previousURL := page.Previous

	if previousURL == nil {
		return fmt.Errorf("No previous URL")
	}

	return loadFromURL(*previousURL, page)
}

// Load the next page in the sequence.
func LoadNextURL(page *OverworldPage) error {
	// If 'map' is called for the first time, we bootstrap into the
	// forward/backward pagination by listing the first page of
	// results.
	if !firstLoaded {
		firstLoaded = true
		go pokecache.InitCacheCleanup(5000, nil)
		return loadFromURL(pokeapiURL, page)
	}

	nextURL := page.Next

	if nextURL == nil {
		return fmt.Errorf("No next URL")
	}

	return loadFromURL(*nextURL, page)
}

// Given the current page, return a slice of the listed Pokemon
// universe place names.
func GetPlaceNames(page *OverworldPage) (placeNames []string) {
	placeNames = make([]string, 0, len(page.Results))

	for _, result := range page.Results {
		placeNames = append(placeNames, result.Name)
	}

	return placeNames
}

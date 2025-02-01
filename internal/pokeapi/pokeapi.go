package pokeapi

import (
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokecache"
)

// We include query parameters here, since calls to 'mapb' from the
// second page will add these anyway.
const pokeapiURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

func LoadFirstURL(page *OverworldPage) error {
	if page == nil {
		return fmt.Errorf("Fatal: 'page' parameter is nil")
	}

	if page.Next != nil {
		return fmt.Errorf("Fatal: second call to 'LoadFirstURL")
	}

	go pokecache.InitCacheCleanup(5000, nil)
	return loadFromURL(pokeapiURL, page)
}

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

	return
}

// Search for 'regionName' among the Results structs, and load the
// JSON data pointed to by the corresponding URL into 'regionInfo'
func LoadRegionInfo(page *OverworldPage, regionInfo *RegionInfoPage, regionName string) error {
	for _, result := range page.Results {
		if result.Name == regionName {
			return loadFromURL(result.URL, regionInfo)
		}
	}

	return fmt.Errorf("Name '%s' not found", regionName)
}

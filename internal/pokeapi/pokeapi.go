package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Page struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var firstLoaded bool = false

// Load the previous page in the sequence.
func LoadPreviousURL(page *Page) error {
	previousURL := page.Previous

	if previousURL == nil {
		return fmt.Errorf("No previous URL")
	}

	return loadFromURL(*previousURL, page)
}

// Load the next page in the sequence.
func LoadNextURL(page *Page) error {
	// If 'map' is called for the first time, we bootstrap into the
	// forward/backward pagination by listing the first page of
	// results.
	if !firstLoaded {
		firstLoaded = true
		return loadFromURL("https://pokeapi.co/api/v2/location-area/", page)
	}

	nextURL := page.Next

	if nextURL == nil {
		return fmt.Errorf("No next URL")
	}

	return loadFromURL(*nextURL, page)
}

// Given the current page, return a slice of the listed Pokemon
// universe place names.
func GetPlaceNames(page *Page) (placeNames []string) {
	placeNames = make([]string, 0, len(page.Results))

	for _, result := range page.Results {
		placeNames = append(placeNames, result.Name)
	}

	return placeNames
}

// Unmarshal contents of URL into the given page.
func loadFromURL(url string, page *Page) error {
	if page == nil {
		return fmt.Errorf("Fatal: page pointer is nil")
	}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("Fatal: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}

	decoder := json.NewDecoder(response.Body)

	if err = decoder.Decode(page); err != nil {
		return fmt.Errorf("Decoding failed: %v", err)
	}

	return nil
}

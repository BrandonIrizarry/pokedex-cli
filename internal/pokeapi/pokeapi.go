package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokecache"
	"io"
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

const pokeapiURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

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

	jsonBytes, found := pokecache.GetEntry(url)

	if found {
		fmt.Println("In the cache.")
		return unmarshal(page, jsonBytes)
	}

	response, err := makeGETRequest(url)

	if err != nil {
		return err
	}

	jsonBytes, err = io.ReadAll(response.Body)

	if err != nil {
		return fmt.Errorf("Reading bytes from response failed: %v", err)
	}

	// Don't forget to add the url to the cache!
	pokecache.AddEntry(url, jsonBytes)

	return unmarshal(page, jsonBytes)
}

// Make a GET request to the given URL.
func makeGETRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Fatal: failure creating connection: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	return response, nil
}

// Unmarshal the given byte slice, representing JSON data, into the
// given page.
func unmarshal(page *Page, jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, page); err != nil {
		return err
	}

	return nil
}

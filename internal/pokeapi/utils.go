package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

// Unmarshal contents of URL into the given destination (a struct for
// unmarshalling.)
func loadFromURL(url string, dest interface{}) error {
	if dest == nil {
		return fmt.Errorf("Fatal: page pointer is nil")
	}

	jsonBytes, found := pokecache.GetEntry(url)

	if found {
		fmt.Println("In the cache.")
		return unmarshal(dest, jsonBytes)
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

	return unmarshal(dest, jsonBytes)
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
func unmarshal(dest interface{}, jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, dest); err != nil {
		return err
	}

	return nil
}

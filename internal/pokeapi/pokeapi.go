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

var pokedexURL string = "https://pokeapi.co/api/v2/location-area/"

func GetFirstPage() (Page, error) {
	request, err := http.NewRequest("GET", pokedexURL, nil)

	if err != nil {
		return Page{}, fmt.Errorf("Fatal: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return Page{}, fmt.Errorf("Request failed: %v", err)
	}

	var newPage Page

	decoder := json.NewDecoder(response.Body)

	if err = decoder.Decode(&newPage); err != nil {
		return Page{}, fmt.Errorf("Decoding failed: %v", err)
	}

	return newPage, nil
}

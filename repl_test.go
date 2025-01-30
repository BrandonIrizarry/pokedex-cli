package main

import (
	"fmt"
	"github.com/BrandonIrizarry/pokedexcli/internal/pokecache"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Anonymous struct reflecting the signature of the function under
	// test.
	testCases := []struct {
		input         string
		expectedWords []string
	}{
		{
			input:         "  Hello World  ",
			expectedWords: []string{"hello", "world"},
		},
		{
			input:         "Charmander Bulbasaur PIKACHU",
			expectedWords: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, tcase := range testCases {
		actualWords := cleanInput(tcase.input)

		if len(actualWords) != len(tcase.expectedWords) {
			t.Errorf("Unequal result lengths")
			return
		}

		for i := 0; i < len(actualWords); i++ {
			actualWord := actualWords[i]
			expectedWord := tcase.expectedWords[i]

			if actualWord != expectedWord {
				t.Errorf("Unequal words at position %d", i)
				return
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	// Note that this is in milliseconds.
	const lifetime = 5000

	testCases := []struct {
		url       string
		someBytes []byte
	}{
		{
			url:       "https://example.com",
			someBytes: []byte("fake JSON data"),
		},

		{
			url:       "https://example.com/path",
			someBytes: []byte("more fake JSON data"),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			pokecache.AddEntry(testCase.url, testCase.someBytes)

			someBytes, found := pokecache.GetEntry(testCase.url)

			if !found {
				t.Errorf("URL not added to cache: '%s'", testCase.url)
				return
			}

			if string(someBytes) != string(testCase.someBytes) {
				t.Errorf("Data added to cache (%s) doesn't match retrieved data (%s)",
					string(testCase.someBytes), string(someBytes))
			}
		})
	}
}

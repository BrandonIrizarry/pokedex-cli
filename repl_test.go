package main

import (
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

package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Struct type-definition reflecting the signature of the function
	// under test.
	type testCase struct {
		input         string
		expectedWords []string
	}

	testCases := []testCase{
		testCase{
			input:         "  Hello World  ",
			expectedWords: []string{"hello", "world"},
		},

		testCase{
			input:         "Charmander Bulbasaur PIKACHU",
			expectedWords: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, tcase := range testCases {
		actualWords := cleanInput(tcase.input)

		if len(actualWords) != len(tcase.expectedWords) {
			fmt.Errorf("Unequal result lengths")
		}

		for i := 0; i < len(actualWords); i++ {
			actualWord := actualWords[i]
			expectedWord := tcase.expectedWords[i]

			if actualWord != expectedWord {
				fmt.Errorf("Unequal words at position %d", i)
			}
		}
	}
}

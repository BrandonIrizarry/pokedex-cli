package pokecache

import (
	"testing"
)

func TestAddGet(t *testing.T) {
	testCases := []struct {
		description string
		url         string
		someBytes   []byte
	}{
		{
			description: "Add arbitrary string to cache",
			url:         "grape juice",
			someBytes:   []byte("fake JSON data"),
		},

		{
			description: "Add another arbitrary string to cache",
			url:         "orange juice",
			someBytes:   []byte("more fake JSON data"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			AddEntry(testCase.url, testCase.someBytes)

			someBytes, found := GetEntry(testCase.url)

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

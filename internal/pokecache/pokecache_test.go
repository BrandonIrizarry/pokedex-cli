package pokecache

import (
	"testing"
)

func TestAddGet(t *testing.T) {
	// Note that this is in milliseconds.
	const lifetime = 5000

	testCases := []struct {
		description string
		url         string
		someBytes   []byte
	}{
		{
			description: "Test an arbitrary string (no HTTP request)",
			url:         "grape juice",
			someBytes:   []byte("fake JSON data"),
		},

		{
			description: "Test another arbitrary string (no HTTP request)",
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

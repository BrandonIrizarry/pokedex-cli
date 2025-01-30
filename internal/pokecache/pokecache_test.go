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

func TestCacheCleanup(t *testing.T) {
	const lifetimeMillis = 5
	tick := make(chan struct{})

	go InitCacheCleanup(lifetimeMillis, tick)

	key := "cheeseburger"
	AddEntry(key, []byte("bread, meat, and cheese"))

	_, found := GetEntry(key)

	if !found {
		t.Errorf("Key '%s' shouldn't've been deleted yet", key)
		return
	}

	// Wait for a tick to occur, so that anything that should be
	// purged, will in fact be.
	<-tick

	_, found = GetEntry(key)

	if found {
		t.Errorf("Key '%s' should've been deleted already", key)
		return
	}
}

package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	timeOfCreation time.Time
	bytes          []byte
}

// A mutex used to prevent concurrent access to the cache.
var mutex sync.Mutex

// The cache itself, mapping URLs to cache entries (basically, JSON
// byte data.)
var cache = make(map[string]cacheEntry)

// Register some JSON bytes in the cache, using the 'cacheEntry'
// format.
func AddEntry(url string, bytes []byte) {
	entry := cacheEntry{
		timeOfCreation: time.Now(),
		bytes:          bytes,
	}

	mutex.Lock()
	cache[url] = entry
	mutex.Unlock()

	fmt.Printf("Added URL %s to cache\n", url)
}

// Fetch a cache entry (as its collection of bytes.)
func GetEntry(url string) ([]byte, bool) {
	mutex.Lock()
	what, ok := cache[url]
	mutex.Unlock()

	if !ok {
		return nil, false
	}

	return what.bytes, true
}

func InitCacheCleanup(_lifetime int, tick chan struct{}) {
	lifetime := time.Duration(_lifetime) * time.Millisecond

	ticker := time.NewTicker(lifetime)
	defer ticker.Stop()

	for currentTime := range ticker.C {
		// Signal to the outside world that a check for cleaning up is
		// to be performed.
		//
		// We mostly need this for our unit tests, so that the test
		// runner doesn't finish without giving this loop a chance to
		// potentially clean the cache.
		//
		// Outside of unit testing, 'tick' should be nil, and so we
		// should check for this.
		if tick != nil {
			tick <- struct{}{}
		}

		for url, entry := range cache {
			if currentTime.Sub(entry.timeOfCreation) >= lifetime {
				mutex.Lock()
				delete(cache, url)
				mutex.Unlock()

				fmt.Printf("Deleted URL %s from cache\n", url)
			}
		}
	}
}

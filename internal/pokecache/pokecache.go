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

func InitCacheCleanup(_lifetime int) {
	lifetime := time.Duration(_lifetime) * time.Millisecond

	ticker := time.NewTicker(lifetime)
	defer ticker.Stop()

	for range ticker.C {
		for url, entry := range cache {
			if time.Since(entry.timeOfCreation) > lifetime {
				fmt.Printf("URL %s is old\n", url)
			}
		}
	}
}

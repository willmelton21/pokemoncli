package pokecache 

import (
	"fmt"
	"time"
    "sync"
	)

//caching goes here

type cacheEntry struct {
   createdAt time.Time 
   val []byte
}

type Cache struct {
    m map[string]cacheEntry
    mu sync.Mutex
    lifetime time.Duration
    func (c *Cache) reapLoop()
}

func reapLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
		
	for range ticker.C { //clean cache entries
		m.mu.Lock()		
		for key, value := range cache.m {
			if(value.Duration > m.lifetime) {
				delete(cache.m,key)
			}

		}
	}
	m.mu.Unlock()

}

func NewCache(time time.Duration) *Cache {
	m := make(map[string]cacheEntry)

	cache := &Cache{m: m, lifetime: time}

	go cache.reapLoop()

	return cache 

}

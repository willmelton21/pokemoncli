package pokecache 

import (
	"fmt"
	"time"
    "byte"
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
}

func NewCache(time time.Duration) *Cache {
	m := make(map[string]cacheEntry

	cache := &Cache{m: m, lifetime: time}

	go cache.reapLoop()

	return cache 

}

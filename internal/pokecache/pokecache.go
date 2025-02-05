package pokecache 

import (
	//"fmt"
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
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.lifetime)
	defer ticker.Stop()
		
	for range ticker.C { //clean cache entries
		c.mu.Lock()		
		for key, value := range c.m {
			timeSinceCreation := time.Since(value.createdAt)	
			if(timeSinceCreation > c.lifetime) {
				delete(c.m,key)
			}

		}
	}
	c.mu.Unlock()

}

func (c *Cache) Add(key string , val []byte)  {
	//create cacheEntry
	newEntry := cacheEntry{val: val, createdAt: time.Now()}
	c.m[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte,bool) {


	ret, ok := c.m[key]

	if !ok {
		return nil, false
	
	} else {
		return ret.val, true 
	}
}

func NewCache(time time.Duration) *Cache {
	m := make(map[string]cacheEntry)

	cache := &Cache{m: m, lifetime: time}

	go cache.reapLoop()

	return cache 

}

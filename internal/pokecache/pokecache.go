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

}

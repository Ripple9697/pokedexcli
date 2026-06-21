package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mapResp  map[string]cacheEntry
	mut      sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mapResp:  make(map[string]cacheEntry),
		mut:      sync.RWMutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock() 
	c.mapResp[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.RLock() 
	entry, ok := c.mapResp[key]
	c.mut.RUnlock()
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mut.Lock()
		for key, entry := range c.mapResp {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.mapResp, key)
			}
		}
		c.mut.Unlock()
	}
}

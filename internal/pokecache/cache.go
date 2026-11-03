package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache : make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mu.Lock()
		for i, elem := range c.cache {
			if time.Since(elem.createdAt) > interval {
				delete(c.cache, i)
			}
		}
		c.mu.Unlock()
	}
}

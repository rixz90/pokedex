package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]CacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	defer c.mutex.Unlock()

	c.mutex.Lock()
	c.entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	defer c.mutex.Unlock()

	c.mutex.Lock()
	value, ok := c.entries[key]

	if !ok {
		return nil, false
	}
	return value.Val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.entries {
			if time.Since(v.CreatedAt) > c.interval {
				delete(c.entries, k)
			}
		}
		c.mutex.Unlock()
	}
}

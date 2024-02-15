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
	mu      sync.Mutex
	ttl     time.Duration
	entries map[string]cacheEntry
}

func NewCache(ttl time.Duration) *Cache {
	newCache := &Cache{
		ttl:     ttl,
		mu:      sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}
	go newCache.readLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.entries[key]

	if !found {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for range ticker.C {
		c.expireTtl()
	}
}

func (c *Cache) expireTtl() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key := range c.entries {
		if c.entries[key].createdAt.Add(c.ttl).Before(time.Now()) {
			delete(c.entries, key)
		}
	}
}

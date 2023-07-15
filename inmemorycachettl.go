package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	value    interface{}
	duration time.Duration
	time     time.Time
}

type Cache struct {
	storage map[string]CacheItem
	mu      *sync.RWMutex
}

func NewCacheItem(value interface{}, ttl time.Duration) CacheItem {
	return CacheItem{value, ttl, time.Now()}
}

func NewCache() *Cache {
	return &Cache{storage: make(map[string]CacheItem), mu: new(sync.RWMutex)}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.storage[key] = NewCacheItem(value, ttl)
	time.AfterFunc(ttl, func() { c.Delete(key) })
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.storage[key]
	if !ok {
		return nil, fmt.Errorf("Cache item \"%s\" not found", key)
	}

	return item.value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}

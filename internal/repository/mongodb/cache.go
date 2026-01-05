// Package mongodb provides caching layer for MongoDB operations.
package mongodb

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// CacheEntry represents a cached item.
type CacheEntry struct {
	Value     interface{}
	ExpiresAt time.Time
}

// Cache provides an in-memory cache for MongoDB queries.
type Cache struct {
	mu      sync.RWMutex
	entries map[string]*CacheEntry
	ttl     time.Duration
	maxSize int
}

// CacheConfig configures the cache.
type CacheConfig struct {
	TTL     time.Duration
	MaxSize int
}

// NewCache creates a new cache.
func NewCache(config CacheConfig) *Cache {
	if config.TTL <= 0 {
		config.TTL = 5 * time.Minute
	}
	if config.MaxSize <= 0 {
		config.MaxSize = 10000
	}

	c := &Cache{
		entries: make(map[string]*CacheEntry),
		ttl:     config.TTL,
		maxSize: config.MaxSize,
	}

	// Start cleanup goroutine
	go c.cleanup()

	return c
}

// Get retrieves an item from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(entry.ExpiresAt) {
		return nil, false
	}

	return entry.Value, true
}

// Set stores an item in the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, c.ttl)
}

// SetWithTTL stores an item with a custom TTL.
func (c *Cache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Evict if at max size
	if len(c.entries) >= c.maxSize {
		c.evictOldest()
	}

	c.entries[key] = &CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// Delete removes an item from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, key)
}

// DeletePrefix removes all items with a key prefix.
func (c *Cache) DeletePrefix(prefix string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key := range c.entries {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			delete(c.entries, key)
		}
	}
}

// Clear removes all items from the cache.
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]*CacheEntry)
}

// Size returns the number of items in the cache.
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}

func (c *Cache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, entry := range c.entries {
		if oldestKey == "" || entry.ExpiresAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.ExpiresAt
		}
	}

	if oldestKey != "" {
		delete(c.entries, oldestKey)
	}
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.After(entry.ExpiresAt) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

// CacheKey generates a cache key for a document query.
func CacheKey(database, collection string, id bson.ObjectID) string {
	return database + ":" + collection + ":" + id.Hex()
}

// CacheKeyQuery generates a cache key for a query.
func CacheKeyQuery(database, collection string, filter interface{}) string {
	// Simple hash based on BSON encoding
	data, _ := bson.Marshal(filter)
	return database + ":" + collection + ":q:" + string(data)
}

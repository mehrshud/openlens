package utils

import (
	"log"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheConfig represents the configuration for the cache utility
type CacheConfig struct {
	// DefaultExpiration is the default expiration time for cache entries
	DefaultExpiration time.Duration
	// CleanupInterval is the interval at which cache entries are cleaned up
	CleanupInterval time.Duration
}

// CacheUtility provides a thread-safe cache implementation
type CacheUtility struct {
	// cache is the underlying cache store
	cache *cache.Cache
	// mu is a mutex for thread safety
	mu sync.RWMutex
	// config is the cache configuration
	config CacheConfig
}

// NewCacheUtility returns a new instance of the CacheUtility
func NewCacheUtility(config CacheConfig) *CacheUtility {
	return &CacheUtility{
		cache:  cache.New(config.DefaultExpiration, config.CleanupInterval),
		config: config,
	}
}

// Get retrieves a value from the cache
func (c *CacheUtility) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache.Get(key)
}

// Set sets a value in the cache
func (c *CacheUtility) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Set(key, value, c.config.DefaultExpiration)
}

// Delete removes a value from the cache
func (c *CacheUtility) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Delete(key)
}

// CacheUser caches a User object
func (c *CacheUtility) CacheUser(user User) error {
	c.Set("user_"+string(user.ID), user)
	log.Printf("Cached user with ID %d", user.ID)
	return nil
}

// GetCachedUser retrieves a cached User object
func (c *CacheUtility) GetCachedUser(id int) (User, bool) {
	value, found := c.Get("user_" + string(id))
	if found {
		return value.(User), true
	}
	return User{}, false
}

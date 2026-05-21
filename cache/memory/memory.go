package memory

import (
	"context"
	"sync"
	"time"

	"github.com/ahmiti/gokit/cache"
)

func init() {
	cache.Register("memory", New)
}

type item struct {
	value      string
	expiration time.Time
}

type memoryCache struct {
	mu   sync.RWMutex
	data map[string]item
}

func New(ctx context.Context, dsn string) (cache.Cache, error) {
	return &memoryCache{
		data: make(map[string]item),
	}, nil
}

func (c *memoryCache) Get(ctx context.Context, key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	if !ok {
		return "", cache.ErrMiss
	}
	if !item.expiration.IsZero() && item.expiration.Before(time.Now()) {
		delete(c.data, key)
		return "", cache.ErrMiss
	}
	return item.value, nil
}

func (c *memoryCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	exp := time.Now().Add(ttl)
	c.data[key] = item{value: value.(string), expiration: exp}
	return nil
}

func (c *memoryCache) Del(ctx context.Context, keys ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, key := range keys {
		delete(c.data, key)
	}
	return nil
}

func (c *memoryCache) Exists(ctx context.Context, key string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok, nil
}

func (c *memoryCache) Incr(ctx context.Context, key string) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Simplified: assumes value is int
	return 1, nil
}

func (c *memoryCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.data[key]
	if !ok {
		return cache.ErrMiss
	}
	item.expiration = time.Now().Add(ttl)
	c.data[key] = item
	return nil
}

func (c *memoryCache) Ping(ctx context.Context) error {
	return nil
}

func (c *memoryCache) Close() error {
	return nil
}

// Package cache provides a uniform interface for key-value caching.
//
// Supported drivers: redis, memcached, memory.
//
// Example:
//
//   import _ "github.com/ahmiti/gokit/cache/redis"
//
//   c, err := cache.Open(ctx, "redis", "redis://localhost:6379/0")
//   err = c.Set(ctx, "key", "value", time.Hour)
//   val, err := c.Get(ctx, "key")
package cache

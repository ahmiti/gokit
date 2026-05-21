package redis

import (
	"context"
	"time"

	"github.com/ahmiti/gokit/cache"
	"github.com/redis/go-redis/v9"
)

func init() {
	cache.Register("redis", New)
}

type redisCache struct {
	client *redis.Client
}

func New(ctx context.Context, dsn string) (cache.Cache, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &redisCache{client: client}, nil
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", cache.ErrMiss
	}
	return val, err
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *redisCache) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *redisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	return n > 0, err
}

func (r *redisCache) Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

func (r *redisCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return r.client.Expire(ctx, key, ttl).Err()
}

func (r *redisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisCache) Close() error {
	return r.client.Close()
}

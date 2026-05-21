package storage

import (
	"context"
	"io"
	"time"
)

type ObjectInfo struct {
	Key          string
	Size         int64
	ContentType  string
	ETag         string
	LastModified time.Time
}

type Storage interface {
	Put(ctx context.Context, key string, reader io.Reader, contentType string) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Stat(ctx context.Context, key string) (*ObjectInfo, error)
	Delete(ctx context.Context, key string) error
	SignedURL(ctx context.Context, key string, ttl time.Duration) (string, error)
	List(ctx context.Context, prefix string) ([]*ObjectInfo, error)
	Close() error
}

package local

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/ahmiti/gokit/storage"
)

func init() {
	storage.Register("local", New)
}

type localStorage struct {
	basePath string
}

func New(ctx context.Context, dsn string) (storage.Storage, error) {
	// dsn = /tmp/storage
	return &localStorage{basePath: dsn}, nil
}

func (l *localStorage) Put(ctx context.Context, key string, reader io.Reader, contentType string) error {
	fullPath := filepath.Join(l.basePath, key)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, reader)
	return err
}

func (l *localStorage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	fullPath := filepath.Join(l.basePath, key)
	return os.Open(fullPath)
}

func (l *localStorage) Stat(ctx context.Context, key string) (*storage.ObjectInfo, error) {
	fullPath := filepath.Join(l.basePath, key)
	info, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}
	return &storage.ObjectInfo{
		Key:          key,
		Size:         info.Size(),
		LastModified: info.ModTime(),
	}, nil
}

func (l *localStorage) Delete(ctx context.Context, key string) error {
	return os.Remove(filepath.Join(l.basePath, key))
}

func (l *localStorage) SignedURL(ctx context.Context, key string, ttl time.Duration) (string, error) {
	return "file://" + filepath.Join(l.basePath, key), nil
}

func (l *localStorage) List(ctx context.Context, prefix string) ([]*storage.ObjectInfo, error) {
	return nil, nil
}

func (l *localStorage) Close() error {
	return nil
}

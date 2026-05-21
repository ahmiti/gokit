package cache

import (
	"context"
	"fmt"
	"sync"
)

type Constructor func(ctx context.Context, dsn string) (Cache, error)

var (
	drivers = make(map[string]Constructor)
	mu      sync.RWMutex
)

func Register(name string, ctor Constructor) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := drivers[name]; exists {
		panic(fmt.Sprintf("cache: driver %s already registered", name))
	}
	drivers[name] = ctor
}

func Open(ctx context.Context, name, dsn string) (Cache, error) {
	mu.RLock()
	ctor, ok := drivers[name]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("cache: unknown driver %q", name)
	}
	return ctor(ctx, dsn)
}

func Drivers() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(drivers))
	for name := range drivers {
		names = append(names, name)
	}
	return names
}

package db

import (
	"context"
	"fmt"
	"sync"
)

type Constructor func(ctx context.Context, dsn string) (DB, error)

var (
	drivers = make(map[string]Constructor)
	mu      sync.RWMutex
)

// Register adds a driver constructor to the registry.
// Called by driver packages in their init() function.
func Register(name string, ctor Constructor) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := drivers[name]; exists {
		panic(fmt.Sprintf("db: driver %s already registered", name))
	}
	drivers[name] = ctor
}

// Open opens a database connection using the named driver.
func Open(ctx context.Context, name, dsn string) (DB, error) {
	mu.RLock()
	ctor, ok := drivers[name]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("db: unknown driver %q (available: %v)", name, Drivers())
	}
	return ctor(ctx, dsn)
}

// Drivers returns a list of registered driver names.
func Drivers() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(drivers))
	for name := range drivers {
		names = append(names, name)
	}
	return names
}

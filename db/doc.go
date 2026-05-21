// Package db provides a uniform interface for relational databases.
//
// The driver pattern:
//   - Interface lives in db.go
//   - Registry in registry.go
//   - Drivers in subpackages (postgres, memory, etc.)
//
// Example:
//
//   import _ "github.com/ahmiti/gokit/db/postgres"
//
//   conn, err := db.Open(ctx, "postgres", "postgres://...")
//   defer conn.Close()
//   rows, err := conn.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
package db

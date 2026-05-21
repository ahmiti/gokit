package db

import (
	"context"
)

// DB is the primary interface for database operations.
type DB interface {
	// Ping checks the database connectivity.
	Ping(ctx context.Context) error

	// Exec executes a query without returning rows.
	Exec(ctx context.Context, query string, args ...interface{}) (Result, error)

	// Query executes a query that returns rows.
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)

	// QueryRow executes a query that expects at most one row.
	QueryRow(ctx context.Context, query string, args ...interface{}) Row

	// Begin starts a new transaction.
	Begin(ctx context.Context) (Tx, error)

	// Close closes the database connection pool.
	Close() error
}

// Tx represents a database transaction.
type Tx interface {
	Exec(ctx context.Context, query string, args ...interface{}) (Result, error)
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// Result is the result of an Exec query.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows is an iterator over query results.
type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
	Err() error
}

// Row is a single row result.
type Row interface {
	Scan(dest ...interface{}) error
}

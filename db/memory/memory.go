package memory

import (
	"context"
	"sync"

	"github.com/ahmiti/gokit/db"
)

func init() {
	db.Register("memory", New)
}

type memoryDB struct {
	mu    sync.RWMutex
	data  map[string]map[string]interface{}
	query func(string, ...interface{}) (db.Rows, error)
}

func New(ctx context.Context, dsn string) (db.DB, error) {
	return &memoryDB{
		data: make(map[string]map[string]interface{}),
	}, nil
}

func (m *memoryDB) Ping(ctx context.Context) error {
	return nil
}

func (m *memoryDB) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	return &memoryResult{rowsAffected: 1}, nil
}

func (m *memoryDB) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	return &memoryRows{}, nil
}

func (m *memoryDB) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return &memoryRow{}
}

func (m *memoryDB) Begin(ctx context.Context) (db.Tx, error) {
	return &memoryTx{}, nil
}

func (m *memoryDB) Close() error {
	return nil
}

type memoryResult struct{ rowsAffected int64 }

func (r *memoryResult) LastInsertId() (int64, error) { return 0, nil }
func (r *memoryResult) RowsAffected() (int64, error) { return r.rowsAffected, nil }

type memoryRows struct{ closed bool }

func (r *memoryRows) Next() bool                     { return false }
func (r *memoryRows) Scan(dest ...interface{}) error { return nil }
func (r *memoryRows) Close() error                   { r.closed = true; return nil }
func (r *memoryRows) Err() error                     { return nil }

type memoryRow struct{}

func (r *memoryRow) Scan(dest ...interface{}) error { return db.ErrNoRows }

type memoryTx struct{}

func (t *memoryTx) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	return &memoryResult{}, nil
}
func (t *memoryTx) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	return &memoryRows{}, nil
}
func (t *memoryTx) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return &memoryRow{}
}
func (t *memoryTx) Commit(ctx context.Context) error   { return nil }
func (t *memoryTx) Rollback(ctx context.Context) error { return nil }

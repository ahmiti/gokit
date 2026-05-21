package postgres

import (
	"context"
	"log/slog"

	"github.com/ahmiti/gokit/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	db.Register("postgres", New)
}

type postgresDB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (db.DB, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &postgresDB{pool: pool}, nil
}

func (p *postgresDB) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *postgresDB) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	tag, err := p.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresResult{tag: tag}, nil
}

func (p *postgresDB) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresRows{rows: rows}, nil
}

func (p *postgresDB) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return &postgresRow{row: p.pool.QueryRow(ctx, query, args...)}
}

func (p *postgresDB) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &postgresTx{tx: tx}, nil
}

func (p *postgresDB) Close() error {
	p.pool.Close()
	return nil
}

type postgresResult struct{ tag pgx.CommandTag }

func (r *postgresResult) LastInsertId() (int64, error) { return 0, nil }
func (r *postgresResult) RowsAffected() (int64, error) { return r.tag.RowsAffected(), nil }

type postgresRows struct {
	rows pgx.Rows
}

func (r *postgresRows) Next() bool                { return r.rows.Next() }
func (r *postgresRows) Scan(dest ...interface{}) error { return r.rows.Scan(dest...) }
func (r *postgresRows) Close() error              { r.rows.Close(); return nil }
func (r *postgresRows) Err() error                { return r.rows.Err() }

type postgresRow struct{ row pgx.Row }

func (r *postgresRow) Scan(dest ...interface{}) error {
	err := r.row.Scan(dest...)
	if err == pgx.ErrNoRows {
		return db.ErrNoRows
	}
	return err
}

type postgresTx struct {
	tx pgx.Tx
}

func (t *postgresTx) Exec(ctx context.Context, query string, args ...interface{}) (db.Result, error) {
	tag, err := t.tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresResult{tag: tag}, nil
}
func (t *postgresTx) Query(ctx context.Context, query string, args ...interface{}) (db.Rows, error) {
	rows, err := t.tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresRows{rows: rows}, nil
}
func (t *postgresTx) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	return &postgresRow{row: t.tx.QueryRow(ctx, query, args...)}
}
func (t *postgresTx) Commit(ctx context.Context) error   { return t.tx.Commit(ctx) }
func (t *postgresTx) Rollback(ctx context.Context) error { return t.tx.Rollback(ctx) }

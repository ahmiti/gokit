package db

import "errors"

var (
	ErrNoRows     = errors.New("db: no rows in result set")
	ErrTxClosed   = errors.New("db: transaction is closed")
	ErrConnClosed = errors.New("db: connection is closed")
)

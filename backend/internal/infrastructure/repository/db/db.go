package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Transaction(ctx context.Context, txOps *sql.TxOptions, fn func(ctx context.Context) error) error
}
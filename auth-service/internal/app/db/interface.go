package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type PgxPoolIface interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Close()
}

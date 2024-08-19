package db

import (
	"context"
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool PgxPoolIface
}

func New(cf *config.DBConfig) (*DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cf.Conn.User,
		cf.Conn.Password,
		cf.Conn.Host,
		cf.Conn.Port,
		cf.Conn.DB)
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

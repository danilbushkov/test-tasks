package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	_ "github.com/lib/pq"
)

type DB struct {
	sqlDB *sql.DB
	db    *reform.DB
}

func New(cf *config.DBConfig, log *logrus.Logger) (*DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cf.Conn.User,
		cf.Conn.Password,
		cf.Conn.Host,
		cf.Conn.Port,
		cf.Conn.DB)
	sqlDB, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, errors.New("database unreachable")
	}

	db := reform.NewDB(
		sqlDB,
		postgresql.Dialect,
		reform.NewPrintfLogger(log.Printf))
	return &DB{sqlDB: sqlDB, db: db}, nil
}

func (db *DB) Orm() *reform.DB {
	return db.db
}

func (db *DB) Close() {
	db.sqlDB.Close()
}

package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DSN     string
	MinOpen int = 5
	MaxOpen int = 10
	DB      *sql.DB
)

var DefaultConnectTimeout = 10 * time.Second

// Build build sqlite3
func Build() error {
	pool, err := buildSQLite3(DSN)
	if err != nil {
		return err
	}
	DB = pool
	return nil
}

// Close close sqlite3
func Close() error {
	if DB == nil {
		return nil
	}

	return DB.Close()
}

func buildSQLite3(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("DSN no set")
	}
	pool, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(MaxOpen)
	pool.SetMaxIdleConns(MinOpen)

	ctx, cancel := context.WithTimeout(context.Background(), DefaultConnectTimeout)
	defer cancel()

	if err = pool.PingContext(ctx); err != nil {
		return nil, err
	}
	return pool, nil
}

// Exec exec sql
type Exec interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

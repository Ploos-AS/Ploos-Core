package storage

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func OpenPostgres(dsn string) (*sql.DB, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("database DSN missing")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.New("database configuration invalid")
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}

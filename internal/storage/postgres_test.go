package storage

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestPostgresIntegration(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not configured")
	}
	db, err := OpenPostgres(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := CheckSchema(ctx, db); err != nil {
		t.Fatal(err)
	}
}

func TestMissingDSN(t *testing.T) {
	if _, err := OpenPostgres(""); err == nil {
		t.Fatal("empty DSN accepted")
	}
}

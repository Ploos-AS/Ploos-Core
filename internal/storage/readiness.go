package storage

import (
	"context"
	"database/sql"
	"errors"
)

// CheckSchema verifies connectivity and the expected migration version.
func CheckSchema(ctx context.Context, db *sql.DB) error {
	if db == nil {
		return errors.New("database not configured")
	}
	var version int
	if err := db.QueryRowContext(ctx, "SELECT COALESCE(max(version), 0) FROM core_schema_migrations").Scan(&version); err != nil {
		return errors.New("database schema unavailable")
	}
	if version != 4 {
		return errors.New("unexpected schema version")
	}
	return nil
}

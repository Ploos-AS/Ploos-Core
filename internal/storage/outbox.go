package storage

import (
	"context"
	"database/sql"
)

// LockNextOutbox selects one pending event inside a caller-owned transaction.
// The lock is held until commit or rollback. Never call external services while holding it.
func LockNextOutbox(ctx context.Context, tx *sql.Tx) (string, error) {
	var id string
	err := tx.QueryRowContext(ctx, "SELECT id FROM provisioning_outbox WHERE delivered_at IS NULL AND dead_at IS NULL AND lease_id IS NULL AND attempts < 10 AND available_at <= now() ORDER BY available_at,id LIMIT 1 FOR UPDATE SKIP LOCKED").Scan(&id)
	return id, err
}

// DeferLockedOutbox updates a previously locked event within the same transaction.
func DeferLockedOutbox(ctx context.Context, tx *sql.Tx, id string) error {
	result, err := tx.ExecContext(ctx, "UPDATE provisioning_outbox SET attempts=attempts+1,available_at=now()+interval '1 minute' WHERE id=$1 AND delivered_at IS NULL AND dead_at IS NULL AND lease_id IS NULL AND attempts < 10", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return sql.ErrNoRows
	}
	return nil
}

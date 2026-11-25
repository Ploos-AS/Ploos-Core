package storage

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

// CompleteOutboxLease atomically records the result and marks the event delivered.
func CompleteOutboxLease(ctx context.Context, db *sql.DB, leaseID, externalResourceID string) error {
	if strings.TrimSpace(externalResourceID) == "" {
		return errors.New("missing external resource ID")
	}
	result, err := db.ExecContext(ctx, `UPDATE provisioning_outbox
 SET delivered_at=now(),external_resource_id=$2,lease_id=NULL,lease_until=NULL
 WHERE lease_id=$1 AND lease_until>now() AND delivered_at IS NULL AND dead_at IS NULL`, leaseID, externalResourceID)
	return requireLeaseUpdate(result, err)
}

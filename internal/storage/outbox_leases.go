package storage

import (
	"context"
	"database/sql"
)

// ClaimOutboxLease reserves a due event for two minutes. It does not authorize execution.
func ClaimOutboxLease(ctx context.Context, db *sql.DB) (eventID, leaseID string, err error) {
	err = db.QueryRowContext(ctx, `WITH candidate AS (
 SELECT id FROM provisioning_outbox
 WHERE delivered_at IS NULL AND dead_at IS NULL AND available_at <= now()
 AND (lease_until IS NULL OR lease_until <= now()) AND attempts < 10
 ORDER BY available_at,id LIMIT 1 FOR UPDATE SKIP LOCKED
 )
 UPDATE provisioning_outbox o SET lease_id=gen_random_uuid(),lease_until=now()+interval '2 minutes',attempts=attempts+1
 FROM candidate c WHERE o.id=c.id RETURNING o.id,o.lease_id`).Scan(&eventID, &leaseID)
	return
}

// RetryOutboxLease rejects expired or superseded lease tokens.
func RetryOutboxLease(ctx context.Context, db *sql.DB, leaseID string) error {
	result, err := db.ExecContext(ctx, `UPDATE provisioning_outbox
 SET lease_id=NULL,lease_until=NULL,available_at=now()+interval '1 minute',
 dead_at=CASE WHEN attempts>=10 THEN now() ELSE NULL END
 WHERE lease_id=$1 AND lease_until>now() AND delivered_at IS NULL AND dead_at IS NULL`, leaseID)
	return requireLeaseUpdate(result, err)
}

func requireLeaseUpdate(result sql.Result, err error) error {
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

// ReapExhaustedOutbox marks exhausted jobs after their final lease expires.
func ReapExhaustedOutbox(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `UPDATE provisioning_outbox SET dead_at=now(),lease_id=NULL,lease_until=NULL
 WHERE delivered_at IS NULL AND dead_at IS NULL AND attempts>=10
 AND (lease_until IS NULL OR lease_until<=now())`)
	return err
}

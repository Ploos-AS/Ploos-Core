package storage

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

// ResolveIdentity maps a verified issuer and subject to an existing user.
// It must never be called with unverified token claims.
func ResolveIdentity(ctx context.Context, db *sql.DB, issuer, subject string) (string, error) {
	if db == nil || strings.TrimSpace(issuer) == "" || strings.TrimSpace(subject) == "" {
		return "", errors.New("invalid identity")
	}
	var id string
	if err := db.QueryRowContext(ctx, "SELECT id FROM users WHERE issuer=$1 AND subject=$2", issuer, subject).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

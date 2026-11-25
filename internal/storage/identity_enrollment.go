package storage

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

// EnrollIdentity is an operator-only primitive for an independently verified
// issuer and subject. It does not grant organization membership.
func EnrollIdentity(ctx context.Context, db *sql.DB, issuer, subject string) (string, error) {
	if db == nil || strings.TrimSpace(issuer) == "" || strings.TrimSpace(subject) == "" {
		return "", errors.New("invalid identity")
	}
	var id string
	err := db.QueryRowContext(ctx, "INSERT INTO users(issuer,subject) VALUES ($1,$2) ON CONFLICT (issuer,subject) DO UPDATE SET subject=EXCLUDED.subject RETURNING id", issuer, subject).Scan(&id)
	return id, err
}

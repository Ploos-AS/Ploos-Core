package storage

import (
	"context"
	"database/sql"
)

// ResolveDomainsOrganization returns a trusted mapping only for a current member.
func ResolveDomainsOrganization(ctx context.Context, db *sql.DB, organizationID, userID string) (string, error) {
	var upstreamID string
	err := db.QueryRowContext(ctx, "SELECT l.domains_organization_id FROM domains_organization_links l JOIN memberships m ON m.organization_id=l.organization_id WHERE l.organization_id=$1 AND m.user_id=$2", organizationID, userID).Scan(&upstreamID)
	return upstreamID, err
}

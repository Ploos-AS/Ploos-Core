package storage

import (
	"context"
	"database/sql"
)

type OrganizationSummary struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Kind        string `json:"kind"`
	Role        string `json:"role"`
}

// ListOrganizations returns only organizations with current user membership.
func ListOrganizations(ctx context.Context, db *sql.DB, userID string) ([]OrganizationSummary, error) {
	rows, err := db.QueryContext(ctx, "SELECT o.id,o.display_name,o.kind,m.role FROM organizations o JOIN memberships m ON m.organization_id=o.id WHERE m.user_id=$1 ORDER BY o.display_name,o.id", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]OrganizationSummary, 0)
	for rows.Next() {
		var item OrganizationSummary
		if err := rows.Scan(&item.ID, &item.DisplayName, &item.Kind, &item.Role); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

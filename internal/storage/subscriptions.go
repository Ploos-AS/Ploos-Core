package storage

import (
	"context"
	"database/sql"
)

type SubscriptionSummary struct {
	ID     string `json:"id"`
	PlanID string `json:"plan_id"`
	Status string `json:"status"`
}

// ListSubscriptions scopes the query to current organization membership.
func ListSubscriptions(ctx context.Context, db *sql.DB, organizationID, userID string) ([]SubscriptionSummary, error) {
	rows, err := db.QueryContext(ctx, "SELECT s.id,s.plan_id,s.status FROM subscriptions s JOIN memberships m ON m.organization_id=s.organization_id WHERE s.organization_id=$1 AND m.user_id=$2 ORDER BY s.created_at,s.id", organizationID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]SubscriptionSummary, 0)
	for rows.Next() {
		var s SubscriptionSummary
		if err := rows.Scan(&s.ID, &s.PlanID, &s.Status); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, rows.Err()
}

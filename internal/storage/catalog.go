package storage

import (
	"context"
	"database/sql"
)

type CatalogPlan struct {
	ID        string `json:"id"`
	Product   string `json:"product"`
	Provider  string `json:"provider"`
	Code      string `json:"code"`
	SeatLimit *int   `json:"seat_limit"`
}

func ListCatalog(ctx context.Context, db *sql.DB) ([]CatalogPlan, error) {
	rows, err := db.QueryContext(ctx, "SELECT p.id,p.product_code,p.provider_code,p.code,p.seat_limit FROM plans p JOIN providers v ON v.code=p.provider_code WHERE p.enabled AND v.enabled ORDER BY p.product_code,p.provider_code,p.code")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	plans := make([]CatalogPlan, 0)
	for rows.Next() {
		var p CatalogPlan
		if err := rows.Scan(&p.ID, &p.Product, &p.Provider, &p.Code, &p.SeatLimit); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, rows.Err()
}

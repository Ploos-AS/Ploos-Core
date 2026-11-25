package server

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Ploos-AS/Ploos-Core/internal/auth"
	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

// NewPortalHandler exposes read-only portal APIs behind verified identity.
func NewPortalHandler(db *sql.DB, verifier auth.Verifier) http.Handler {
	mux := http.NewServeMux()
	if db == nil || verifier == nil {
		return NewHandler()
	}
	mux.Handle("/", NewHandlerWithReadiness(func(ctx context.Context) error {
		return storage.CheckSchema(ctx, db)
	}))
	resolver := auth.DatabaseResolver{DB: db}
	protect := func(h http.Handler) http.Handler {
		return auth.RequireIdentity(verifier, resolver, h)
	}
	mux.Handle("/api/v1/catalog", protect(CatalogHandler(func(ctx context.Context) ([]storage.CatalogPlan, error) {
		return storage.ListCatalog(ctx, db)
	})))
	mux.Handle("/api/v1/organizations/{organizationID}/subscriptions", protect(SubscriptionsHandler(func(ctx context.Context, org, user string) ([]storage.SubscriptionSummary, error) {
		return storage.ListSubscriptions(ctx, db, org, user)
	})))
	mux.Handle("/api/v1/organizations", protect(OrganizationsHandler(func(ctx context.Context, user string) ([]storage.OrganizationSummary, error) {
		return storage.ListOrganizations(ctx, db, user)
	})))
	return mux
}

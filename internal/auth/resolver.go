package auth

import (
	"context"
	"database/sql"

	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

type DatabaseResolver struct {
	DB *sql.DB
}

func (r DatabaseResolver) Resolve(ctx context.Context, identity Identity) (string, error) {
	return storage.ResolveIdentity(ctx, r.DB, identity.Issuer, identity.Subject)
}

var _ Resolver = DatabaseResolver{}

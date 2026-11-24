package storage

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

func TestOrganizationMembershipIntegration(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not configured")
	}
	db, err := OpenPostgres(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var owner, second, orgID string
	for _, id := range []*string{&owner, &second} {
		if err := db.QueryRowContext(ctx, "INSERT INTO users(issuer,subject) VALUES ('integration',gen_random_uuid()::text) RETURNING id").Scan(id); err != nil {
			t.Fatal(err)
		}
	}
	if err := db.QueryRowContext(ctx, "SELECT gen_random_uuid()").Scan(&orgID); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_, _ = db.Exec("DELETE FROM memberships WHERE organization_id=$1", orgID)
		_, _ = db.Exec("DELETE FROM organizations WHERE id=$1", orgID)
		_, _ = db.Exec("DELETE FROM users WHERE id IN ($1,$2)", owner, second)
	}()
	org := domain.Organization{ID: orgID, DisplayName: "Integration association", Kind: domain.KindAssociation}
	if err := CreateOrganization(ctx, db, org, owner); err != nil {
		t.Fatal(err)
	}
	if ChangeMembership(ctx, db, orgID, owner, owner, nil) == nil {
		t.Fatal("last owner removed")
	}
	role := domain.RoleOwner
	if err := ChangeMembership(ctx, db, orgID, owner, second, &role); err != nil {
		t.Fatal(err)
	}
	if err := ChangeMembership(ctx, db, orgID, owner, owner, nil); err != nil {
		t.Fatal(err)
	}
	if ChangeMembership(ctx, db, orgID, owner, owner, &role) == nil {
		t.Fatal("former owner regained access")
	}
	if ChangeMembership(ctx, db, orgID, second, second, nil) == nil {
		t.Fatal("last remaining owner removed")
	}
}

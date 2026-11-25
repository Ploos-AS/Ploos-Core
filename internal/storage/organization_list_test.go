package storage

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

func TestOrganizationDiscoveryIsolation(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not configured")
	}
	db, err := OpenPostgres(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var users, orgs [2]string
	defer func() {
		for _, user := range users {
			if user != "" {
				_, _ = db.Exec("DELETE FROM users WHERE id=$1", user)
			}
		}
	}()
	for i := range users {
		if err := db.QueryRowContext(ctx, "INSERT INTO users(issuer,subject) VALUES ('isolation-test',gen_random_uuid()::text) RETURNING id").Scan(&users[i]); err != nil {
			t.Fatal(err)
		}
		user := users[i]
		if err := db.QueryRowContext(ctx, "SELECT gen_random_uuid()").Scan(&orgs[i]); err != nil {
			t.Fatal(err)
		}
		org := domain.Organization{ID: orgs[i], DisplayName: "Isolation test", Kind: domain.KindAssociation}
		if err := CreateOrganization(ctx, db, org, user); err != nil {
			t.Fatal(err)
		}
		defer func(id string) {
			_, _ = db.Exec("DELETE FROM audit_events WHERE organization_id=$1", id)
			_, _ = db.Exec("DELETE FROM memberships WHERE organization_id=$1", id)
			_, _ = db.Exec("DELETE FROM organizations WHERE id=$1", id)
		}(org.ID)
		var count int
		if err := db.QueryRowContext(ctx, "SELECT count(*) FROM audit_events WHERE organization_id=$1 AND actor_user_id=$2 AND action='organization.created'", org.ID, user).Scan(&count); err != nil || count != 1 {
			t.Fatalf("creation audit: count=%d err=%v", count, err)
		}
	}
	for i, user := range users {
		items, err := ListOrganizations(ctx, db, user)
		if err != nil {
			t.Fatal(err)
		}
		if len(items) != 1 || items[0].ID != orgs[i] || items[0].Role != "owner" {
			t.Fatalf("unexpected organization visibility: %+v", items)
		}
	}
	role := domain.RoleViewer
	if err := ChangeMembership(ctx, db, orgs[0], users[0], users[1], &role); err != nil {
		t.Fatal(err)
	}
	items, err := ListOrganizations(ctx, db, users[1])
	if err != nil || len(items) != 2 {
		t.Fatalf("shared membership: count=%d err=%v", len(items), err)
	}
	if err := ChangeMembership(ctx, db, orgs[0], users[0], users[1], nil); err != nil {
		t.Fatal(err)
	}
	items, err = ListOrganizations(ctx, db, users[1])
	if err != nil || len(items) != 1 || items[0].ID != orgs[1] {
		t.Fatalf("revoked membership still visible: %+v err=%v", items, err)
	}
}

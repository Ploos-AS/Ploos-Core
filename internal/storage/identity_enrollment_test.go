package storage

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestIdentityEnrollmentIntegration(t *testing.T) {
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
	var subject string
	if err := db.QueryRowContext(ctx, "SELECT gen_random_uuid()::text").Scan(&subject); err != nil {
		t.Fatal(err)
	}
	issuer := "https://identity-test.invalid"
	defer db.Exec("DELETE FROM users WHERE subject=$1 AND issuer IN ($2,$3)", subject, issuer, issuer+"/other")
	first, err := EnrollIdentity(ctx, db, issuer, subject)
	if err != nil {
		t.Fatal(err)
	}
	again, err := EnrollIdentity(ctx, db, issuer, subject)
	if err != nil || again != first {
		t.Fatalf("enrollment not idempotent: %v", err)
	}
	other, err := EnrollIdentity(ctx, db, issuer+"/other", subject)
	if err != nil || other == first {
		t.Fatalf("issuer isolation failed: %v", err)
	}
	resolved, err := ResolveIdentity(ctx, db, issuer, subject)
	if err != nil || resolved != first {
		t.Fatalf("resolution failed: %v", err)
	}
	items, err := ListOrganizations(ctx, db, first)
	if err != nil || len(items) != 0 {
		t.Fatalf("enrollment unexpectedly grants membership: %v", err)
	}
}

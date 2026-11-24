package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ploos-AS/Ploos-Core/internal/auth"
	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

type subscriptionVerifier struct{}

func (subscriptionVerifier) Verify(context.Context, string) (auth.Identity, error) {
	return auth.Identity{Issuer: "issuer", Subject: "subject"}, nil
}

type subscriptionResolver struct{}

func (subscriptionResolver) Resolve(context.Context, auth.Identity) (string, error) {
	return "user-a", nil
}

func TestSubscriptionsHandler(t *testing.T) {
	h := SubscriptionsHandler(func(_ context.Context, org, user string) ([]storage.SubscriptionSummary, error) {
		if org != "org-a" || user != "user-a" {
			t.Fatal("incorrect scope")
		}
		return nil, nil
	})
	r := httptest.NewRequest("GET", "/", nil)
	r.SetPathValue("organizationID", "org-a")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusUnauthorized {
		t.Fatal("unauthenticated request accepted")
	}
	r.Header.Set("Authorization", "Bearer test")
	w = httptest.NewRecorder()
	auth.RequireIdentity(subscriptionVerifier{}, subscriptionResolver{}, h).ServeHTTP(w, r)
	if w.Code != http.StatusOK || w.Body.String() != "{\"subscriptions\":[]}\n" {
		t.Fatal("unexpected subscription response")
	}
}

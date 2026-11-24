package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type verifierStub struct{}

func (verifierStub) Verify(_ context.Context, token string) (Identity, error) {
	if token != "valid" {
		return Identity{}, errors.New("invalid")
	}
	return Identity{Issuer: "issuer", Subject: "subject"}, nil
}

type resolverStub struct{}

func (resolverStub) Resolve(_ context.Context, identity Identity) (string, error) {
	if identity.Issuer != "issuer" || identity.Subject != "subject" {
		return "", errors.New("unknown")
	}
	return "user-a", nil
}

func TestRequireIdentity(t *testing.T) {
	h := RequireIdentity(verifierStub{}, resolverStub{}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id, ok := UserID(r.Context()); !ok || id != "user-a" {
			t.Error("missing verified identity")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	for _, tc := range []struct {
		header string
		status int
	}{
		{"", 401},
		{"Bearer invalid", 401},
		{"Bearer valid", 204},
	} {
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", tc.header)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		if w.Code != tc.status {
			t.Errorf("got %d, want %d", w.Code, tc.status)
		}
	}
}

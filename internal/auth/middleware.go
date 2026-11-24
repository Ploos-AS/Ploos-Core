package auth

import (
	"context"
	"net/http"
	"strings"
)

type Identity struct {
	Issuer  string
	Subject string
}

// Verifier must validate signature, issuer, audience, expiry and token type.
type Verifier interface {
	Verify(context.Context, string) (Identity, error)
}

type Resolver interface {
	Resolve(context.Context, Identity) (string, error)
}

type userKey struct{}

func UserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userKey{}).(string)
	return id, ok && id != ""
}

func RequireIdentity(verifier Verifier, resolver Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		if verifier == nil || resolver == nil || next == nil {
			http.Error(w, "authentication unavailable", http.StatusServiceUnavailable)
			return
		}
		parts := strings.Fields(r.Header.Get("Authorization"))
		if len(r.Header.Values("Authorization")) != 1 || len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			w.Header().Set("WWW-Authenticate", "Bearer")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		identity, err := verifier.Verify(r.Context(), parts[1])
		if err != nil || strings.TrimSpace(identity.Issuer) == "" || strings.TrimSpace(identity.Subject) == "" {
			w.Header().Set("WWW-Authenticate", "Bearer")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		id, err := resolver.Resolve(r.Context(), identity)
		if err != nil || strings.TrimSpace(id) == "" {
			http.Error(w, "identity unavailable", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey{}, id)))
	})
}

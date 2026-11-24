package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIntrospection(t *testing.T) {
	for _, active := range []bool{true, false} {
		s := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, password, ok := r.BasicAuth()
			if !ok || user != "core" || password != "secret" || r.Method != "POST" || r.FormValue("token") != "opaque" {
				t.Error("invalid introspection request")
			}
			_ = json.NewEncoder(w).Encode(map[string]any{"active": active, "iss": "https://issuer.example", "sub": "user", "aud": []string{"core-api"}, "token_type": "Bearer", "exp": time.Now().Add(time.Minute).Unix()})
		}))
		v, err := NewIntrospector(s.URL, "https://issuer.example", "core-api", "core", "secret")
		if err != nil {
			s.Close()
			t.Fatal(err)
		}
		v.client.Transport = s.Client().Transport
		id, err := v.Verify(context.Background(), "opaque")
		s.Close()
		if active && (err != nil || id.Subject != "user") {
			t.Fatal("valid token rejected")
		}
		if !active && err == nil {
			t.Fatal("inactive token accepted")
		}
	}
}

package domains

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPortfolio(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/v1/organizations/org-a/portfolio" || r.Header.Get("Authorization") != "Bearer delegated" {
			t.Error("unexpected request")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		_, _ = w.Write([]byte(`{"domains":[]}`))
	}))
	defer server.Close()
	c, err := NewClient(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	c.http.Transport = server.Client().Transport
	result, err := c.Portfolio(context.Background(), "org-a", "delegated")
	if err != nil || result == nil || len(result) != 0 {
		t.Fatalf("unexpected portfolio: %v %v", result, err)
	}
	if _, err := c.Portfolio(context.Background(), "../other", "delegated"); err == nil {
		t.Fatal("invalid organization accepted")
	}
	if _, err := c.Portfolio(context.Background(), "org-a", ""); err == nil {
		t.Fatal("missing token accepted")
	}
}

func TestHTTPSRequired(t *testing.T) {
	if _, err := NewClient("http://example.org"); err == nil {
		t.Fatal("insecure URL accepted")
	}
}

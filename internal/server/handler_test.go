package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOperationalEndpoints(t *testing.T) {
	h := NewHandler()
	for _, tc := range []struct {
		method string
		path   string
		status int
	}{
		{"GET", "/healthz", http.StatusOK},
		{"GET", "/readyz", http.StatusServiceUnavailable},
		{"POST", "/healthz", http.StatusMethodNotAllowed},
		{"GET", "/api/v1/organizations", http.StatusNotFound},
	} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		if w.Code != tc.status {
			t.Errorf("%s %s: got %d, want %d", tc.method, tc.path, w.Code, tc.status)
		}
	}
}

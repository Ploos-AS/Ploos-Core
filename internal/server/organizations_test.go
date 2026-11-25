package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrganizationsRequiresIdentity(t *testing.T) {
	w := httptest.NewRecorder()
	OrganizationsHandler(nil).ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/v1/organizations", nil))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", w.Code)
	}
}

func TestOrganizationsRejectsWrite(t *testing.T) {
	w := httptest.NewRecorder()
	OrganizationsHandler(nil).ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/v1/organizations", nil))
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d", w.Code)
	}
}

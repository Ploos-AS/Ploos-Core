package server

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

func TestCatalogHandler(t *testing.T) {
	h := CatalogHandler(func(context.Context) ([]storage.CatalogPlan, error) {
		return nil, nil
	})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/catalog", nil))
	if w.Code != http.StatusOK || w.Body.String() != "{\"plans\":[]}\n" {
		t.Fatal("unexpected catalog response")
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/catalog", nil))
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatal("write accepted")
	}
	h = CatalogHandler(func(context.Context) ([]storage.CatalogPlan, error) {
		return nil, errors.New("database failure")
	})
	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/catalog", nil))
	if w.Code != http.StatusServiceUnavailable {
		t.Fatal("database failure not handled")
	}
}

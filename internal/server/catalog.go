package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

// CatalogHandler is read-only. Mount behind authentication when enabling the API.
func CatalogHandler(list func(context.Context) ([]storage.CatalogPlan, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if list == nil {
			http.Error(w, "catalog unavailable", http.StatusServiceUnavailable)
			return
		}
		plans, err := list(r.Context())
		if err != nil {
			http.Error(w, "catalog unavailable", http.StatusServiceUnavailable)
			return
		}
		if plans == nil {
			plans = []storage.CatalogPlan{}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(struct {
			Plans []storage.CatalogPlan `json:"plans"`
		}{Plans: plans})
	})
}

package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ploos-AS/Ploos-Core/internal/auth"
	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

func SubscriptionsHandler(list func(context.Context, string, string) ([]storage.SubscriptionSummary, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userID, ok := auth.UserID(r.Context())
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		orgID := r.PathValue("organizationID")
		if orgID == "" {
			http.Error(w, "missing organization", http.StatusBadRequest)
			return
		}
		if list == nil {
			http.Error(w, "subscriptions unavailable", http.StatusServiceUnavailable)
			return
		}
		items, err := list(r.Context(), orgID, userID)
		if err != nil {
			http.Error(w, "subscriptions unavailable", http.StatusServiceUnavailable)
			return
		}
		if items == nil {
			items = []storage.SubscriptionSummary{}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(struct {
			Subscriptions []storage.SubscriptionSummary `json:"subscriptions"`
		}{Subscriptions: items})
	})
}

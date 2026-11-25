package server

import (
	"context"
	"crypto/rand"
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/Ploos-AS/Ploos-Core/internal/auth"
	"github.com/Ploos-AS/Ploos-Core/internal/domain"
	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

func CreateOrganizationHandler(create func(context.Context, domain.Organization, string) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		user, ok := auth.UserID(r.Context())
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var input struct {
			DisplayName string                  `json:"display_name"`
			Kind        domain.OrganizationKind `json:"kind"`
		}
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096))
		decoder.DisallowUnknownFields()
		if decoder.Decode(&input) != nil || decoder.Decode(new(any)) != io.EOF {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		var id [16]byte
		if _, err := rand.Read(id[:]); err != nil {
			http.Error(w, "organization creation unavailable", http.StatusServiceUnavailable)
			return
		}
		id[6] = (id[6] & 0x0f) | 0x40
		id[8] = (id[8] & 0x3f) | 0x80
		org := domain.Organization{ID: fmt.Sprintf("%x-%x-%x-%x-%x", id[:4], id[4:6], id[6:8], id[8:10], id[10:]), DisplayName: strings.TrimSpace(input.DisplayName), Kind: input.Kind}
		if org.Validate() != nil || utf8.RuneCountInString(org.DisplayName) > 200 || strings.ContainsRune(org.DisplayName, 0) {
			http.Error(w, "invalid organization", http.StatusBadRequest)
			return
		}
		if create == nil || create(r.Context(), org, user) != nil {
			http.Error(w, "organization creation unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(storage.OrganizationSummary{ID: org.ID, DisplayName: org.DisplayName, Kind: string(org.Kind), Role: "owner"})
	})
}

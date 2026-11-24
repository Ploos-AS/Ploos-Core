package server

import (
	"context"
	"net/http"
	"time"
)

// NewHandler exposes only operational endpoints until authentication and
// persistent storage are configured.
func NewHandler() http.Handler {
	return NewHandlerWithReadiness(nil)
}

func NewHandlerWithReadiness(check func(context.Context) error) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = w.Write([]byte("{\"status\":\"ok\"}\n"))
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if check != nil && check(ctx) == nil {
			_, _ = w.Write([]byte("{\"status\":\"ok\"}\n"))
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("{\"status\":\"not_ready\"}\n"))
	})
	return mux
}

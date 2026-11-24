package server

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadinessCheck(t *testing.T) {
	for _, tc := range []struct {
		err  error
		want int
	}{
		{nil, http.StatusOK},
		{errors.New("database unavailable"), http.StatusServiceUnavailable},
	} {
		h := NewHandlerWithReadiness(func(context.Context) error { return tc.err })
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/readyz", nil))
		if w.Code != tc.want {
			t.Errorf("got %d, want %d", w.Code, tc.want)
		}
	}
}

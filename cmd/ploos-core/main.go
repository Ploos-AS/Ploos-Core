package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ploos-AS/Ploos-Core/internal/server"
	"github.com/Ploos-AS/Ploos-Core/internal/storage"
)

func main() {
	addr := os.Getenv("PLOOS_CORE_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	handler := server.NewHandler()
	if dsn := os.Getenv("PLOOS_CORE_DATABASE_URL"); dsn != "" {
		db, err := storage.OpenPostgres(dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		handler = server.NewHandlerWithReadiness(func(ctx context.Context) error {
			return storage.CheckSchema(ctx, db)
		})
	}
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	done := make(chan error, 1)
	go func() { done <- s.ListenAndServe() }()
	log.Printf("Ploos Core listening on %s", addr)
	select {
	case err := <-done:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	case <-ctx.Done():
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(shutdown); err != nil {
			log.Fatal(err)
		}
	}
}

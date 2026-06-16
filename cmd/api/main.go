package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tolbra/gowallet/internal/config"
)

func main() {

	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	// temporary check
	log.Println("Database_URL is set")
	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to create databasse pool", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	log.Println("Successfully connected to PostgreSQL")

	defer dbPool.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "ok",
		}
		json.NewEncoder(w).Encode(response)

	})

	log.Println("server started on :" + cfg.AppPort)

	err = http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		log.Fatal(err)
	}

}

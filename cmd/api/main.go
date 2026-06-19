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

	// Config part
	cfg := config.Load()       // Variable cfg that loads config
	if cfg.DatabaseURL == "" { // Check if Database URL is empty
		log.Fatal("DATABASE_URL is not set") // if so log.Fatal
	}

	log.Println("Database_URL is set") // if ok, print database url is ok
	// Creating database
	ctx := context.Background() // creates base context: context is used by database/network operations to carry
	// cancellation, timeout, or request-scoped values. Here it means "run normally with no timeout or cancellation"
	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL) // creates postgresql connection pool using database URL. dbPool is
	// reused for database queries instead of opening a new DB connection every time.
	if err != nil { // if err is detected
		log.Fatal("failed to create databasse pool", err) // log.Fatal & err
	}

	if err := dbPool.Ping(ctx); err != nil { // Tests database connection.
		log.Fatal("Failed to connect to PostgreSQL:", err) // if err -> log.Fatal & err
	}
	defer dbPool.Close()                                // close dbPool after the program stops
	log.Println("Successfully connected to PostgreSQL") // if no err -> print ok

	r := chi.NewRouter() // create chi router

	r.Use(middleware.Logger) // Adds Chi's request logger middleware. Every HTTP request will be logged in the terminal
	// such as method, path, status code, and duration

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) { // endpoint
		w.Header().Set("Content-Type", "application/json") // Tells the client/browser that the response body is JSON

		response := map[string]string{ // Creates a small JSON-like response object.
			"status": "ok", // Key "status" has value "ok"
		}
		json.NewEncoder(w).Encode(response) // Converts the response map into JSON and writesi t to the HTTP response.
		// THe client recieves: {"status": "ok"}

	})

	log.Println("server started on :" + cfg.AppPort) // server started message w port

	err = http.ListenAndServe(":"+cfg.AppPort, r) // check for errors during server life
	if err != nil {
		log.Fatal(err) // if err stop, and output it
	}

}

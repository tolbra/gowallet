package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "ok",
		}
		json.NewEncoder(w).Encode(response)

	})

	log.Println("server started on :3030")

	err := http.ListenAndServe(":3030", r)
	if err != nil {
		log.Fatal(err)
	}

}

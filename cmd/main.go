package main

import (
	"log"
	"net/http"

	"streamly-backend/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.Load()
	config.ConnectMongo()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	addr := ":" + config.C.Port
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

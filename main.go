package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/simple-store/store/handler"
	"github.com/simple-store/store/repository"
)

func main() {
	port := "8080"
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Timeout(60 * time.Second))
	storeHandler := handler.NewStoreHandler(r, repository.NewStore())
	r.Route("/", func(r chi.Router) {
		r.Post("/store", storeHandler.Save)
	})
	log.Printf("Starting up on http://localhost:%s", port)
	err := http.ListenAndServe(fmt.Sprintf("%s%s", "127.0.0.1:", port), r)
	if err != nil {
		log.Panic(err)
	}
}

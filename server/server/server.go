package server

import (
	"fmt"
	"net/http"

	"github.com/ara-o/gren/routes/register"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type server struct {
	address string
}

func (s *server) Start() error {
	fmt.Println("Starting server...")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	loadRoutes(r)

	fmt.Println("Server listening on", s.address)
	return http.ListenAndServe(s.address, r)
}

func New(address string) *server {
	return &server{
		address: address,
	}
}

func loadRoutes(r *chi.Mux) {
	r.Post("/api/register", register.Register)
}

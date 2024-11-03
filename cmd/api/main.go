package main

import (
	"github.com/Waelson/go-grafana-loki/internal/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	controller := handlers.NewController()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", controller.Routes())

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

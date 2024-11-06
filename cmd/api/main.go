package main

import (
	"github.com/Waelson/go-grafana-loki/internal/handlers"
	"github.com/Waelson/go-grafana-loki/internal/logger"
	"github.com/Waelson/go-grafana-loki/internal/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	log := log.NewZapLogger()
	controller := handlers.NewController(log)

	r := chi.NewRouter()
	r.Use(middleware.NewContextMiddleware())
	r.Mount("/", controller.Routes())

	log.Info("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

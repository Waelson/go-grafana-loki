package main

import (
	"fmt"
	"github.com/Waelson/go-grafana-loki/internal/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	controller := handlers.NewController()

	go func() {
		for {
			fmt.Println(`{"message":"log message","level":"info","thread":"main"}`)
			time.Sleep(1 * time.Second)
		}
	}()

	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Mount("/", controller.Routes())

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

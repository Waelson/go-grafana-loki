package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Home route accessed"
	log.Println(msg) // Log para stdout
	w.Write([]byte(msg))
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {
	msg := "Health route accessed"
	log.Println(msg) // Log para stdout
	w.Write([]byte(msg))
}

func (c *Controller) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.Home)
	r.Get("/health", c.Health)
	return r
}

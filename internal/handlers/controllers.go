package handlers

import (
	"github.com/Waelson/go-grafana-loki/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Controller struct {
	log log.Logger
}

func NewController(log log.Logger) *Controller {

	return &Controller{log}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Home route accessed"
	c.log.InfoCtx(r.Context(), msg)
	w.Write([]byte(msg))
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {
	msg := "Health route accessed"
	c.log.InfoCtx(r.Context(), msg)
	w.Write([]byte(msg))
}

func (c *Controller) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.Home)
	r.Get("/health", c.Health)
	return r
}

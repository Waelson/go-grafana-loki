package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type LogMessage struct {
	App     string `json:"app"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func logJSON(app, level, message string) {
	logMsg := LogMessage{App: app, Level: level, Message: message}
	jsonData, err := json.Marshal(logMsg)
	if err != nil {
		log.Printf("Error marshalling JSON log: %v", err)
		return
	}
	log.Println(string(jsonData))
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Home route accessed"
	logJSON("golang-app", "info", msg)
	w.Write([]byte(msg))
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {
	msg := "Health route accessed"
	logJSON("golang-app", "info", msg)
	w.Write([]byte(msg))
}

func (c *Controller) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.Home)
	r.Get("/health", c.Health)
	return r
}

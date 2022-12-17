package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct {
	endpoint string
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{endpoint: "/healthz"}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &model.HealthzResponse{Message: "OK"}
	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		log.Println("Error:", err)
	}
}

package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"./models"
)

// Define a type for the HTTP API handler
type HTTPHandler struct {
	config *models.Config
}

// NewHTTPHandler returns a new instance of the HTTP API handler
func NewHTTPHandler(config *models.Config) *HTTPHandler {
	return &HTTPHandler{config: config}
}

// GetDashboard handles GET requests for dashboards
func (h *HTTPHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request for dashboard")

	// Try to get the dashboard from the database
	dashboard, err := models.GetDashboard(h.config.DBURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the dashboard to JSON
	jsonDashboard, err := json.Marshal(dashboard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDashboard)
}

// ServeHTTP serves the HTTP API
func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/dashboard" {
			h.GetDashboard(w, r)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

package routes

import (
	"net/http"
)

// GetDocsHandler docs route handler
func GetDocsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

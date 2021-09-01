package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "album API is running as expected",
		}

		json.NewEncoder(w).Encode(response)
	}
}

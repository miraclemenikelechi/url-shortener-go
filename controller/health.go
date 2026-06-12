package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &HealthCheckResponse{
		Message: "healthy",
		Status:  "Ok",
	}
	json.NewEncoder(w).Encode(response)
}

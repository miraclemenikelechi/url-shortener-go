package controller

import (
	"net/http"

	"github.com/miraclemenikelechi/url-shortener-go/utils"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	utils.RespondWithJSON(w, &HealthCheckResponse{
		Message: "healthy", Status: "Ok",
	})
}

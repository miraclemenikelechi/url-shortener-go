package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miraclemenikelechi/url-shortner-go/utils"
)

type LongTextRequest struct {
	RawURL string `json:"raw_url"`
}

type LongTextResponse struct {
	ShortURL string `json:"short_url"`
}

func HandleLongTextFromClient(w http.ResponseWriter, r *http.Request) {
	var request LongTextRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("received: %s\n", request.RawURL)
	utils.RespondWithJSON(w, &LongTextResponse{ShortURL: utils.GenerateRandomString(6)})
}

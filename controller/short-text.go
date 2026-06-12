package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ShortTextRequest struct {
	RawURL string `json:"raw_url"`
}

type ShortTextResponse struct {
	ShortURL string `json:"short_url"`
}

func HandleShortTextFromClient(w http.ResponseWriter, r *http.Request) {
	var request ShortTextRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("received: %s\n", request.RawURL)

	response := &ShortTextResponse{ShortURL: ""}
}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miraclemenikelechi/url-shortner-go/memory"
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

	if !utils.IsValidURL(request.RawURL) {
		http.Error(w, "invalid URL", http.StatusUnprocessableEntity)
		return
	}

	log.Printf("received: %s\n", request.RawURL)
	generatedCode := utils.GenerateRandomString(6)
	memory.URLS_DB[generatedCode] = request.RawURL

	utils.RespondWithJSON(w, &LongTextResponse{ShortURL: generatedCode})
}

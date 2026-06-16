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

func (c *Controller) HandleLongTextFromClient(w http.ResponseWriter, r *http.Request) {
	var request LongTextRequest
	ctx := r.Context()

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

	if err := c.DB.CreateUrl(ctx, memory.CreateUrlParams{
		ShortenedCode: generatedCode, OriginalUrl: request.RawURL,
	}); err != nil {
		http.Error(w, "failed to create URL", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, &LongTextResponse{ShortURL: generatedCode})
}

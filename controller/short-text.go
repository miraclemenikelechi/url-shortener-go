package controllers

import (
	"net/http"
)

type ShortTextRequest struct {
	ShortURL string `json:"short_url"`
}

type ShortTextResponse struct {
	RawURL string `json:"raw_url"`
}

func HandleShortTextFromClient(w http.ResponseWriter, r *http.Request) {
	// var request ShortTextRequest

	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	http.Error(w, "invalid request", http.StatusBadRequest)
	// 	return
	// }

	// log.Printf("received: %s\n", request.ShortURL)

	// response := &ShortTextResponse{RawURL: ""}
	//
	// code := r.PathValue("code")
}

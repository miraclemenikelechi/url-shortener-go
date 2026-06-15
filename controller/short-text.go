package controllers

import (
	"log"
	"net/http"

	"github.com/miraclemenikelechi/url-shortner-go/memory"
)

func HandleShortTextFromClient(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	log.Printf("received: %s\n", code)

	if rawURL, ok := memory.URLS_DB[code]; ok {
		http.Redirect(w, r, rawURL, http.StatusFound)
	} else {
		http.Error(w, "not found", http.StatusNotFound)
	}
}

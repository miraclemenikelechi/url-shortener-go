package controllers

import (
	"log"
	"net/http"
)

func (c *Controller) HandleShortTextFromClient(w http.ResponseWriter, r *http.Request) {
	code, ctx := r.PathValue("code"), r.Context()
	log.Printf("received: %s\n", code)

	rawURL, err := c.DB.GetUrl(ctx, code)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, rawURL, http.StatusFound)
}

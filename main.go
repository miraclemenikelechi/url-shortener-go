package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/miraclemenikelechi/url-shortner-go/controller"
)

func main() {
	mux := http.NewServeMux()

	// app health check
	mux.HandleFunc("GET /health", controllers.HealthCheckHandler)

	// send raw text to the url shortner engine
	mux.HandleFunc("POST /send-raw-text", controllers.HandleLongTextFromClient)

	// forward transformed url to original destination
	mux.HandleFunc("GET /{code}", controllers.HandleShortTextFromClient)

	fmt.Println("server running on http://localhost:8649")
	log.Fatal(http.ListenAndServe(":8649", mux))
}

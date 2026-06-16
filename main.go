package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/miraclemenikelechi/url-shortner-go/controller"
	"github.com/miraclemenikelechi/url-shortner-go/memory"
)

func main() {
	db, port := memory.ConnectDB(), os.Getenv("PORT")
	memory.MigrateDB(db)

	c, mux := &controllers.Controller{DB: memory.New(db)}, http.NewServeMux()

	// app health check
	mux.HandleFunc("GET /health", controllers.HealthCheckHandler)

	// send raw text to the url shortner engine
	mux.HandleFunc("POST /send-raw-text", c.HandleLongTextFromClient)

	// forward transformed url to original destination
	mux.HandleFunc("GET /{code}", c.HandleShortTextFromClient)

	fmt.Println("server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

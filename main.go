package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/ve6om/go-crypto-api/handlers"
)

// main function initializes the server and sets up routes
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port: %s\n", port)

	r := chi.NewRouter()

	// Routes
	r.Get("/price", handlers.GetPriceHandler)
	r.Get("/coins", handlers.GetCoinsHandler)

	// Start server
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

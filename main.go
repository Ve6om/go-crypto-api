package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/ve6om/go-crypto-api/handlers"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/price", handlers.GetPriceHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

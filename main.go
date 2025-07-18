package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/ve6om/go-crypto-api/handlers"
)

func main() {
	_ = godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No PORT set — using default port 8080")
	}

	frontendURL := os.Getenv("FRONTEND_URL")

	if frontendURL != "" {
		fmt.Printf("Allowed CORS origin: %s\n", frontendURL)
	} else {
		fmt.Println("No FRONTEND_URL set — CORS headers disabled")
	}

	fmt.Printf("Starting server on port: %s\n", port)

	r := chi.NewRouter()

	r.Get("/price", handlers.GetPriceHandler)
	r.Get("/coins", handlers.GetCoinsHandler)

	handler := withConditionalCORS(r, frontendURL)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

func withConditionalCORS(next http.Handler, allowedOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigin != "" && origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

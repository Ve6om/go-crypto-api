package handlers

import (
	"encoding/json"
	"net/http"

	coins_client "github.com/ve6om/go-crypto-api/clients"
)

func GetCoinsHandler(w http.ResponseWriter, r *http.Request) {
	coins, err := coins_client.FetchCoins()
	if err != nil {
		http.Error(w, "Failed to fetch coins", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coins)
}
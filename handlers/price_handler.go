package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ve6om/go-crypto-api/services"
)

func GetPriceHandler(w http.ResponseWriter, r *http.Request) {
	coin := r.URL.Query().Get("coin")
	if coin == "" {
		http.Error(w, "coin query param required", http.StatusInternalServerError)
		return
	}

	currency := r.URL.Query().Get("currency")
	if currency == "" {
		currency = "usd"
	}

	timezone := r.URL.Query().Get("timezone")

	priceResult, err := services.GetPrice(coin, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if timezone != "" {
		loc, err := time.LoadLocation(timezone)
		if err == nil {
			priceResult.Time = priceResult.Time.In(loc)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(priceResult)
}

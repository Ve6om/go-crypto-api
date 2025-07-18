package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ve6om/go-crypto-api/models"
)

func GetPrice(coin string, currency string) (models.PriceResult, error) {
	if currency == "" {
		currency = "usd"
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", coin, currency)

	resp, err := http.Get(url)
	if err != nil {
		return models.PriceResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PriceResult{}, fmt.Errorf("bad stutus: %s", resp.Status)
	}

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.PriceResult{}, fmt.Errorf("price not found for %s", coin)
	}

	price, ok := result[coin][currency]
	if !ok {
		return models.PriceResult{}, fmt.Errorf("price not found for %s in %s", currency)
	}

	return models.PriceResult{
		Coin:     coin,
		Currency: currency,
		Price:    price,
		Time:     time.Now(),
	}, nil
}

package coins_client

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/ve6om/go-crypto-api/models"
)

var (
	cachedCoins []models.Coin
	lastUpdated time.Time
	cacheMutex  sync.Mutex
	cacheTTL    = 6 * time.Hour
)

func FetchCoins() ([]models.Coin, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if time.Since(lastUpdated) < cacheTTL && cachedCoins != nil {
		return cachedCoins, nil
	}

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch coins from CoinGecko")
	}

	var coins []models.Coin
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		return nil, err
	}

	cachedCoins = coins
	lastUpdated = time.Now()

	return coins, nil
}
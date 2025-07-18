package models

import "time"

type PriceResult struct {
	Coin string `json:"coin"`
	Currency string `json:"currency"`
	Price float64 `json:"price"`
	Time time.Time `json:"timestamp"`
}
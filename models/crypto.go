package models

type Crypto struct {
	Currency string `json:"currency"`
	Price    float32 `json:"price"`
}
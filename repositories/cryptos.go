package repositories

import (
	"gin/models"
	"os"
	"strconv"
)

type Cryptos interface {
	Find(coin string) models.Crypto
}

type InMemoryCryptos struct{}

var prices = map[string]models.Crypto{}

// Init BTC/USD and ETH/USD pairs
func (c InMemoryCryptos) Init() {
	btcPrice, err := strconv.ParseFloat(os.Getenv("BTC"), 32)
	if err != nil {
		btcPrice = 1000
	}

	ethPrice, err := strconv.ParseFloat(os.Getenv("ETH"), 32)
	if err != nil {
		ethPrice = 500
	}

	prices["BTC"] = models.Crypto{
		Currency: "BTC",
		Price:    float32(btcPrice),
	}
	prices["ETH"] = models.Crypto{
		Currency: "ETH",
		Price:    float32(ethPrice),
	}
}

func (c InMemoryCryptos) Find(coin string) models.Crypto {
	return prices[coin]
}

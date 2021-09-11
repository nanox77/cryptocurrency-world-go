package models

import (
	"errors"
	"fmt"
)

const USD = "USD"
const BTC = "BTC"
const ETH = "ETH"
const CENT = 100

type Coin struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Wallet struct {
	Coins map[string]Coin `json:"coins"`
}

func (w Wallet) AddOrUpdate(coin Coin) {
	elem, exist := w.Coins[coin.Name]
	if exist {
		elem.Amount += coin.Amount
		w.Coins[coin.Name] = elem
	} else {
		w.Coins[coin.Name] = coin
	}
}

func (w Wallet) ExtractFromWallet(coin Coin) error {
	elem, exist := w.Coins[coin.Name]
	if exist {
		isNegative := elem.Amount-coin.Amount < 0
		if isNegative {
			return errors.New("insufficient balance")
		}
		elem.Amount -= coin.Amount
		w.Coins[coin.Name] = elem
	} else {
		w.Coins[coin.Name] = coin
	}
	return nil
}

type User struct {
	UserId string `json:"userId"`
	Wallet Wallet `json:"wallet"`
}

// AddToWallet amount in cents
func (u User) AddToWallet(name string, amount int) {
	u.Wallet.AddOrUpdate(Coin{
		Name:   name,
		Amount: amount,
	})
}

// ExtractFromWallet amount in cents
func (u User) ExtractFromWallet(name string, amount int) error {
	return u.Wallet.ExtractFromWallet(Coin{
		Name:   name,
		Amount: amount,
	})
}

func (u User) GetCoin(coinName string) (*Coin, error) {
	coin, exist := u.Wallet.Coins[coinName]
	if !exist {
		return nil, errors.New(fmt.Sprintf("%s not supported", coinName))
	}
	return &coin, nil
}

func NewUser(userId string) User {
	return User{
		UserId: userId,
		Wallet: newWallet(),
	}
}

func newWallet() Wallet {
	return Wallet{Coins: map[string]Coin{USD: {
		Name:   USD,
		Amount: 0,
	}, BTC: {
		Name:   BTC,
		Amount: 0,
	}, ETH: {
		Name:   ETH,
		Amount: 0,
	}}}
}

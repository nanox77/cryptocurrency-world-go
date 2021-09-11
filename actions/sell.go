package actions

import (
	"gin/models"
	"gin/repositories"
)

type Sell struct {
	Users   repositories.Users
	Cryptos repositories.Cryptos
}

func (s Sell) Sell(userId string, coinName string, amount int) error {
	user, err := s.Users.FindById(userId)
	if err != nil {
		return err
	}

	coin := s.Cryptos.Find(coinName)
	cost := int(coin.Price) * models.CENT * amount

	err = user.ExtractFromWallet(coinName, amount)
	if err != nil {
		return err
	}

	user.AddToWallet("USD", cost)
	s.Users.Store(*user)

	return nil
}

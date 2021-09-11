package actions

import (
	"gin/models"
	"gin/repositories"
)

type Buy struct {
	Users   repositories.Users
	Cryptos repositories.Cryptos
}

func (b Buy) Buy(userId string, coinName string, amount int) error {
	user, err := b.Users.FindById(userId)
	if err != nil {
		return err
	}

	coin := b.Cryptos.Find(coinName)
	cost := int(coin.Price) * models.CENT * amount

	err = user.ExtractFromWallet("USD", cost)
	if err != nil {
		return err
	}

	user.AddToWallet(coinName, amount)
	b.Users.Store(*user)

	return nil
}

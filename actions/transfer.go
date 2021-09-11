package actions

import (
	"gin/repositories"
)

type Transfer struct {
	Users   repositories.Users
	Cryptos repositories.Cryptos
}

func (t Transfer) Transfer(userId string, toUser string, coinName string, amount int) error {
	userFrom, err := t.Users.FindById(userId)
	if err != nil {
		return err
	}
	userTo, err := t.Users.FindById(toUser)
	if err != nil {
		return err
	}

	err = userFrom.ExtractFromWallet(coinName, amount)
	if err != nil {
		return err
	}
	userTo.AddToWallet(coinName, amount)

	t.Users.Store(*userFrom)
	t.Users.Store(*userTo)

	return nil
}

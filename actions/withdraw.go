package actions

import (
	"gin/repositories"
)

type Withdraw struct {
	Users repositories.Users
}

func (w Withdraw) Withdraw(userId string, amount int) error {
	user, err := w.Users.FindById(userId)
	if err != nil {
		return err
	}

	err = user.ExtractFromWallet("USD", amount)
	if err != nil {
		return err
	}

	w.Users.Store(*user)

	return nil
}

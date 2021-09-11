package actions

import (
	"gin/repositories"
)

type Deposit struct {
	Users repositories.Users
}

func (d Deposit) Deposit(userId string, amount int) error {
	user, err := d.Users.FindById(userId)

	if err != nil {
		return err
	}

	user.AddToWallet("USD", amount)
	d.Users.Store(*user)

	return nil
}

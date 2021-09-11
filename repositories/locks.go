package repositories

import (
	"errors"
	"fmt"
)

type Locks interface {
	Lock(userId string) error
	Unlock(userId string)
}

type InMemoryLocks struct{}

var usersLocks = map[string]string{}

func (InMemoryLocks) Lock(userId string) error {
	_, exist := usersLocks[userId]

	if exist {
		return errors.New(fmt.Sprintf("user %s is in another operation", userId))
	}

	usersLocks[userId] = userId
	return nil
}

func (InMemoryLocks) Unlock(userId string) {
	delete(usersLocks, userId)
}

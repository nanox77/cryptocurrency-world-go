package repositories

import (
	"errors"
	"fmt"
	"gin/models"
	"sort"
)

// Users Repository
type Users interface {
	FindById(id string) (*models.User, error)
	FindAll() []models.User
	Store(user models.User)
}

type InMemoryUsers struct{}

var users = map[string]models.User{}

func (u InMemoryUsers) Init() {
	users["1"] = models.NewUser("1")
	users["2"] = models.NewUser("2")
	users["3"] = models.NewUser("3")
}

func (InMemoryUsers) FindById(userId string) (*models.User, error) {
	element, exist := users[userId]

	if exist {
		return &element, nil
	}

	return nil, errors.New(fmt.Sprintf("user %s not found", userId))
}

func (InMemoryUsers) Store(user models.User) {
	users[user.UserId] = user
}

func (InMemoryUsers) FindAll() []models.User {
	var values []models.User
	for _, v := range users {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].UserId < values[j].UserId
	})

	return values
}

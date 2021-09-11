package main

import (
	"gin/actions"
	"gin/handlers"
	"gin/repositories"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	r := gin.Default()

	users := repositories.InMemoryUsers{}
	cryptos := repositories.InMemoryCryptos{}
	locks := repositories.InMemoryLocks{}
	users.Init()
	cryptos.Init()

	deposit := actions.Deposit{
		Users: users,
	}
	withdraw := actions.Withdraw{
		Users: users,
	}
	buy := actions.Buy{
		Users:   users,
		Cryptos: cryptos,
	}
	sell := actions.Sell{
		Users:   users,
		Cryptos: cryptos,
	}
	transfer := actions.Transfer{
		Users:   users,
		Cryptos: cryptos,
	}

	r.GET("/", handlers.Greetings())
	r.GET("/users", handlers.GetAll(users))
	r.POST("/users/:userId/deposit", handlers.Lock(locks), handlers.Deposit(deposit), handlers.Unlock(locks))
	r.POST("/users/:userId/withdraw", handlers.Lock(locks), handlers.Withdraw(withdraw), handlers.Unlock(locks))
	r.POST("/users/:userId/buy", handlers.Lock(locks), handlers.Buy(buy), handlers.Unlock(locks))
	r.POST("/users/:userId/sell", handlers.Lock(locks), handlers.Sell(sell), handlers.Unlock(locks))
	r.POST("/users/:userId/transfer", handlers.Lock(locks), handlers.Transfer(transfer), handlers.Unlock(locks))
	r.Run()
}

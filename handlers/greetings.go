package handlers

import (
	"github.com/gin-gonic/gin"
)

func Greetings() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Welcome to CryptoCurrency world!")
	}
}

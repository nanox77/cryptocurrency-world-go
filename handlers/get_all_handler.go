package handlers

import (
	"gin/repositories"
	"github.com/gin-gonic/gin"
)

func GetAll(users repositories.Users) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, users.FindAll())
	}
}

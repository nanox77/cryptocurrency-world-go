package handlers

import (
	"gin/repositories"
	"github.com/gin-gonic/gin"
)

func Unlock(locks repositories.Locks) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		locks.Unlock(userId)
		c.Next()
	}
}

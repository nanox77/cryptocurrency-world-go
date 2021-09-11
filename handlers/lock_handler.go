package handlers

import (
	"gin/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Lock(locks repositories.Locks) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		err := locks.Lock(userId)

		if err != nil {
			buildError(c, err)
			c.AbortWithStatusJSON(http.StatusProcessing, err)
		}

		c.Next()
	}
}

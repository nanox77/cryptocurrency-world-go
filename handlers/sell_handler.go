package handlers

import (
	"gin/actions"
	"gin/input"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Sell(sellAction actions.Sell) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		sellRequest := input.BuyRequest{}
		err := c.ShouldBindBodyWith(&sellRequest, binding.JSON)
		if err != nil {
			buildError(c, err)
			return
		}

		err = validator.New().Struct(sellRequest)
		if err != nil {
			buildFromValidationError(c, err)
			return
		}

		err = sellAction.Sell(userId, sellRequest.Coin, sellRequest.Amount)
		if err != nil {
			buildError(c, err)
			return
		}
		c.Status(200)
		return
	}
}

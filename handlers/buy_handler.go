package handlers

import (
	"gin/actions"
	"gin/input"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Buy(buyAction actions.Buy) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		buyRequest := input.BuyRequest{}
		err := c.ShouldBindBodyWith(&buyRequest, binding.JSON)
		if err != nil {
			buildError(c, err)
			return
		}

		err = validator.New().Struct(buyRequest)
		if err != nil {
			buildFromValidationError(c, err)
			return
		}

		err = buyAction.Buy(userId, buyRequest.Coin, buyRequest.Amount)
		if err != nil {
			buildError(c, err)
			return
		}
		c.Status(200)
		return
	}
}

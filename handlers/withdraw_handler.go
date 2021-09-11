package handlers

import (
	"gin/actions"
	"gin/input"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Withdraw(withdrawAction actions.Withdraw) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		withdrawRequest := input.WithdrawRequest{}
		err := c.ShouldBindBodyWith(&withdrawRequest, binding.JSON)
		if err != nil {
			buildError(c, err)
			return
		}

		err = validator.New().Struct(withdrawRequest)
		if err != nil {
			buildFromValidationError(c, err)
			return
		}

		err = withdrawAction.Withdraw(userId, withdrawRequest.Amount)
		if err != nil {
			buildError(c, err)
			return
		}
		c.Status(200)
		return
	}
}

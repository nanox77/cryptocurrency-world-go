package handlers

import (
	"gin/actions"
	"gin/input"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Deposit(depositAction actions.Deposit) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		depositRequest := input.DepositRequest{}
		err := c.ShouldBindBodyWith(&depositRequest, binding.JSON)
		if err != nil {
			buildError(c, err)
			return
		}

		err = validator.New().Struct(depositRequest)
		if err != nil {
			buildFromValidationError(c, err)
			return
		}

		err = depositAction.Deposit(userId, depositRequest.Amount)
		if err != nil {
			buildError(c, err)
			return
		}
		c.Status(200)
		return
	}
}

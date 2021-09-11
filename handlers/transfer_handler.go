package handlers

import (
	"gin/actions"
	"gin/input"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Transfer(transferAction actions.Transfer) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		transferRequest := input.TransferRequest{}
		err := c.ShouldBindBodyWith(&transferRequest, binding.JSON)
		if err != nil {
			buildError(c, err)
			return
		}

		err = validator.New().Struct(transferRequest)
		if err != nil {
			buildFromValidationError(c, err)
			return
		}

		err = transferAction.Transfer(userId, transferRequest.ToUser, transferRequest.Coin, transferRequest.Amount)
		if err != nil {
			buildError(c, err)
			return
		}
		c.Status(200)
		return
	}
}

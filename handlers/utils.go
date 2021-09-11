package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func buildError(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error_message": err.Error(),
	})
}

func buildFromValidationError(c *gin.Context, err error) {
	errs := err.(validator.ValidationErrors)
	c.JSON(400, gin.H{
		"error_message": fmt.Sprintf("%s is required", errs[0].StructField()),
	})
}

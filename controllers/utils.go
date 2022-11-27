package controllers

import (
	"context"
	"gin-api-template/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func SendJSON(c *gin.Context, response *responses.Response) {
	c.JSON(response.Status, response)
}

func GetAndValidateBody[T any](c *gin.Context) (*T, error) {
	var model T

	if err := c.BindJSON(&model); err != nil {
		SendJSON(c, responses.BadRequest(err))
		return &model, err
	}

	if err := validate.Struct(&model); err != nil {
		SendJSON(c, responses.BadRequest(err))
		return &model, err
	}

	return &model, nil
}

package controllers

import (
	"context"
	"errors"
	"gin-api-template/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func SendJSON(c *gin.Context, response *responses.Response) {
	c.JSON(response.Status, response)
}

func GetAndValidateQueryId(c *gin.Context, idKey string) (primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(c.Param(idKey))
	if err != nil {
		SendJSON(c, responses.BadRequest(err))
		return primitive.ObjectID{}, errors.New("could not parse ID")
	}

	return objId, nil
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

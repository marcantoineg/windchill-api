package controllers

import (
	"gin-api-template/configs"
	"gin-api-template/models"
	"gin-api-template/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetAllUser(c *gin.Context) {
	ctx, cancel := GetContext()
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		SendJSON(c, responses.ServerError(err))
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			SendJSON(c, responses.ServerError(err))
		}

		users = append(users, singleUser)
	}

	SendJSON(c, responses.OK(users))
}

func GetUser(c *gin.Context) {
	ctx, cancel := GetContext()
	var user models.User
	defer cancel()

	objId, err := GetAndValidateQueryId(c, "userId")
	if err != nil {
		return
	}

	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		SendJSON(c, responses.ServerError(err))
		return
	}

	SendJSON(c, responses.OK(user))
}

func CreateUser(c *gin.Context) {
	ctx, cancel := GetContext()
	defer cancel()

	user, err := GetAndValidateBody[models.User](c)
	if err != nil {
		return
	}

	newUser := models.User{
		Id:   primitive.NewObjectID(),
		Name: user.Name,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		SendJSON(c, responses.ServerError(err))
		return
	}

	SendJSON(c, responses.OK(result))
}

func UpdateUser(c *gin.Context) {
	ctx, cancel := GetContext()
	defer cancel()

	objId, err := GetAndValidateQueryId(c, "userId")
	if err != nil {
		return
	}

	user, err := GetAndValidateBody[models.User](c)
	if err != nil {
		return
	}

	update := bson.M{"$set": bson.M{"name": user.Name}}
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		SendJSON(c, responses.ServerError(err))
		return
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
		if err != nil {
			SendJSON(c, responses.ServerError(err))
			return
		}
	} else if result.MatchedCount > 1 {
		SendJSON(c, responses.ServerErrorWithMessage("multiple document found."))
		return
	} else {
		SendJSON(c, responses.NotFound())
		return
	}

	SendJSON(c, responses.OK(updatedUser))
}

func DeleteUser(c *gin.Context) {
	ctx, cancel := GetContext()
	defer cancel()

	objId, err := GetAndValidateQueryId(c, "userId")
	if err != nil {
		return
	}

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		SendJSON(c, responses.ServerError(err))
		return
	}

	if result.DeletedCount < 1 {
		SendJSON(c, responses.NotFound())
		return
	}

	SendJSON(c, responses.OK("user have been deleted."))
}

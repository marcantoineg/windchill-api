package routes

import (
	"gin-api-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetAllUser)
	r.GET("/users/:userId", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:userId", controllers.UpdateUser)
	r.DELETE("/users/:userId", controllers.DeleteUser)
}

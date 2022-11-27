package api

import (
	"gin-api-template/configs"
	"gin-api-template/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	configs.ConnectDB()

	routes.RegisterUserRoutes(r)

	routes.RegisterPingPong(r)

	return r
}

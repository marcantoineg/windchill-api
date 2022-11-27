package api

import (
	"gin-api-template/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	routes.RegisterPingPong(r)
	routes.RegisterWindchillRoutes(r)

	return r
}

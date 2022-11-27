package routes

import (
	"gin-api-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPingPong(r *gin.Engine) {
	r.GET("/ping", controllers.PingPong)
}

package routes

import (
	"gin-api-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterWindchillRoutes(r *gin.Engine) {
	r.POST("/windchill", controllers.WindChill)
}

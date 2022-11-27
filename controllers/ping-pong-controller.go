package controllers

import (
	"gin-api-template/responses"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	SendJSON(c, responses.OK("pong"))
}

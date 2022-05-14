package router

import (
	"Lovers/handler"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.POST("/user/login", handler.LoginHandler)
}

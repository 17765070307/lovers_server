package main

import (
	"Lovers/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Register(app)
	app.Run(":10001")
}

package router

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/adapter/controller"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", controller.HelloWorld)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myapp/config"
	"myapp/internal/framework/db"
	"myapp/internal/framework/middleware"
	"myapp/internal/framework/router"
)

func main() {
	// Initialize database
	db.SetupDB()

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction())
	app.Use(middleware.Cors())
	router.SetupRoutes(app)
	app.Run(fmt.Sprintf(":%d", config.Port))
}

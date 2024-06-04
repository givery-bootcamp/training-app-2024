package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"myapp/internal/config"
	"myapp/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutes(config *config.Config, app *gin.Engine) error {
	api, err := InjectAPI(config)
	if err != nil {
		return fmt.Errorf("failed to inject api: %w", err)
	}

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", api.HelloWorld)

	return nil
}

func Run(ctx context.Context, conf *config.Config) error {
	app := newGin()
	if err := setupRoutes(conf, app); err != nil {
		return fmt.Errorf("failed to setup routes: %w", err)
	}

	address := fmt.Sprintf("%s:%d", conf.Server.HostName, conf.Server.Port)
	port := fmt.Sprintf(":%d", conf.Server.Port)
	log.Printf("Starting server on %s...\n", address)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// addressだとなぜか動かなかった(ERR_CONNECTION_REFUSED)
	return app.Run(port)
}

func newGin() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	return r
}

package middleware

import (
	"myapp/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{config.GetConfig().Server.CorsAllowOrigin}
	conf.AllowCredentials = true
	return cors.New(conf)
}

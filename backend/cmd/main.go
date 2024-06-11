package main

import (
	"context"
	"log"

	"myapp/internal/config"
	"myapp/internal/infrastructure/repository/gorm"
	"myapp/internal/presentation/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	gorm.NewGormDB(ctx, conf.DB)

	if err := server.Run(ctx, conf); err != nil {
		log.Panic(err)
	}
}

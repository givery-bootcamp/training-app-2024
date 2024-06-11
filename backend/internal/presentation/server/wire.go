//go:build wireinject
// +build wireinject

package server

import (
	"myapp/internal/config"
	"myapp/internal/infrastructure/repository/gorm"
	"myapp/internal/presentation/hello"
	"myapp/internal/usecase"
	domainhello "myapp/internal/domain/hello"

	"github.com/google/wire"
)


func InjectAPI(config *config.Config) (*API, error) {
	wire.Build(

		wire.Bind(new(domainhello.HelloWorldRepository), new(*gorm.HelloWorldRepository)),
		gorm.NewHelloWorldRepository,
		usecase.NewGetHelloUsecase,
		hello.NewHandler,
		NewAPI,
	)
	return nil, nil
}

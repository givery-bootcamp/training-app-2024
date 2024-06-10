package repository

import (
	"myapp/internal/domain/entity"
)

type HelloWorldRepository interface {
	Get(lang string) (*entity.HelloWorld, error)
}

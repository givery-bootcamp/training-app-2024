package usecase

import (
	"myapp/internal/domain/entity"
	repositories "myapp/internal/domain/repository"
)

type HelloWorldUsecase struct {
	repository repositories.HelloWorldRepository
}

func NewHelloWorldUsecase(r repositories.HelloWorldRepository) *HelloWorldUsecase {
	return &HelloWorldUsecase{
		repository: r,
	}
}

func (u *HelloWorldUsecase) Execute(lang string) (*entity.HelloWorld, error) {
	return u.repository.Get(lang)
}

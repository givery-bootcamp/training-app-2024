package usecase

import (
	"context"

	"myapp/internal/domain/hello"
)

type GetHelloUsecase struct {
	hr hello.HelloWorldRepository
}

func NewGetHelloUsecase(r hello.HelloWorldRepository) *GetHelloUsecase {
	return &GetHelloUsecase{
		hr: r,
	}
}

type HelloWorldDTO struct {
	Lang    string
	Message string
}

func (u *GetHelloUsecase) Exec(ctx context.Context, lang string) (*HelloWorldDTO, error) {
	helloWorld, err := u.hr.Get(ctx, hello.Lang(lang))
	if err != nil {
		return nil, err
	}

	HelloWorldDTO := &HelloWorldDTO{
		Lang:    helloWorld.Lang(),
		Message: helloWorld.Message(),
	}

	return HelloWorldDTO, nil
}

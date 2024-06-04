package gorm

import (
	"context"
	"errors"
	"fmt"

	"myapp/internal/domain/hello"
	"myapp/internal/infrastructure/repository/gorm/schema"

	"gorm.io/gorm"
)

type HelloWorldRepository struct{}

func NewHelloWorldRepository() *HelloWorldRepository {
	return &HelloWorldRepository{}
}

func (r *HelloWorldRepository) Get(ctx context.Context, lang hello.Lang) (*hello.HelloWorld, error) {
	query := GetQuery(ctx)
	var helloObj schema.HelloWorld
	if err := query.Where("lang = ?", lang).First(&helloObj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("not found hello world message for lang: %s", lang)
		}
		return nil, fmt.Errorf("failed to get hello world message for lang: %w", err)
	}

	result, err := hello.NewHelloWorld(lang, helloObj.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to create hello world message for lang: %w", err)
	}

	return result, nil
}

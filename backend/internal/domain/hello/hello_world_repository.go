package hello

import "context"

type HelloWorldRepository interface {
	Get(ctx context.Context,lang Lang) (*HelloWorld, error)
}
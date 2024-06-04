package server

import "myapp/internal/presentation/hello"

type API struct {
	*hello.Handler
}

func NewAPI(h *hello.Handler) *API {
	return &API{h}
}

package hello

import (
	err "myapp/internal/domain/error"
)

type HelloWorld struct {
	lang    Lang
	message string
}

func (h *HelloWorld) Lang() string {
	return h.lang.String()
}

func (h *HelloWorld) Message() string {
	return h.message
}

func NewHelloWorld(lang Lang, message string) (*HelloWorld, error) {
	return newHelloWorld(lang, message)
}

func newHelloWorld(lang Lang, message string) (*HelloWorld, error) {
	if !lang.IsValid() {
		return nil, err.InvalidLangParameter
	}

	return &HelloWorld{
		lang:    lang,
		message: message,
	}, nil
}

package hello

type HelloWorldParams struct {
	Lang string `json:"lang" validate:"required,oneof=ja en,lte=2"`
}
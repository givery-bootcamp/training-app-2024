package error

type Error struct {
	description string
}

func (e *Error) Error() string {
	return e.description
}

func NewError(s string) *Error {
	return &Error{
		description: s,
	}
}

var (
	NotFoundErr = NewError("not found")
	InvalidParameterErr = NewError("invalid parameter")
	InternalServerError = NewError("internal server error")
	InvalidLangParameter = NewError("invalid lang parameter")
)



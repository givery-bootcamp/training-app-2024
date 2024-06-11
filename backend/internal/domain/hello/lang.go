package hello

type Lang string

const (
	ja Lang = "ja"
	en Lang = "en"
)

func (l Lang) IsValid() bool {
	switch l {
	case ja:
		return true
	case en:
		return true
	}
	return false
}

func (l Lang) String() string {
	switch l {
	case ja:	return "ja"
	case en:	return "en"
	}
	return ""
}
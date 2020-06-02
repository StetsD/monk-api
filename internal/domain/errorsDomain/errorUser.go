package errorsDomain

type ErrorUser string

func (eu ErrorUser) Error() string {
	return string(eu)
}

var (
	ErrorEmailExists = ErrorUser("the user with that email already exists")
)

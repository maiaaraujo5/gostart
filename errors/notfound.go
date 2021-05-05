package errors

type notfound struct {
	Err
}

func NotFound(message string) error {
	return &notfound{
		Err{message: message},
	}
}

func IsNotFound(err error) bool {
	_, ok := err.(*notfound)
	return ok
}

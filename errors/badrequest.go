package errors

type badrequest struct {
	Err
}

func BadRequest(message string) error {
	return &badrequest{
		Err{message: message},
	}
}

func IsBadRequest(err error) bool {
	_, ok := err.(*badrequest)
	return ok
}

package rest

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewHttpError(statusCode int, message string) *httpError {
	return &httpError{
		StatusCode: statusCode,
		Message:    message,
	}
}

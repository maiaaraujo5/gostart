package httperror

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func New(statusCode int, message string) *httpError {
	return &httpError{
		StatusCode: statusCode,
		Message:    message,
	}
}

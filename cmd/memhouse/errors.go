package main

type HttpError struct {
	Status int
	err    string
}

func NewHttpError(status int, err string) *HttpError {
	return &HttpError{status, err}
}

// HttpError implements Error method
func (e *HttpError) Error() string {
	return e.err
}

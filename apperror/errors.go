package apperror

import (
	"errors"
	"fmt"
)

// Error types
var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal server error")
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

// ValidationError membuat error baru untuk validasi
func ValidationError(field string) error {
	return &CustomError{
		Code:    400,
		Message: fmt.Sprintf("Field '%s' is required", field),
	}
}

func ValidationErrorWithMessage(message string) error {
	return fmt.Errorf("%s", message)
}

// DetermineErrorType mengecek jenis error dan balikin status + message
func DetermineErrorType(err error) (int, string) {
	switch e := err.(type) {
	case *CustomError:
		return e.Code, e.Message
	default:
		switch {
		case errors.Is(err, ErrBadRequest):
			return 400, "Bad Request"
		case errors.Is(err, ErrNotFound):
			return 404, "Not Found"
		default:
			return 500, "Internal Server Error"
		}
	}
}

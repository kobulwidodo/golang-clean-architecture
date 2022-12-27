package errors

import (
	"fmt"
	"go-clean/src/lib/codes"
)

type CustomError struct {
	Code  codes.Message
	Cause string
}

func (ce *CustomError) Error() string {
	return fmt.Sprint(ce.Cause)
}

func NewWithCode(codes codes.Message, err string) error {
	newErr := &CustomError{
		Code:  codes,
		Cause: err,
	}

	return newErr
}

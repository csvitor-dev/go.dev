package validations

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	FieldName string
	Errors    []error
}

func (e *ValidationError) error(message string) {
	e.Errors = append(e.Errors, errors.New(message))
}

func (e *ValidationError) errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	e.error(message)
}

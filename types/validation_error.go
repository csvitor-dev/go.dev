package types

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	FieldName string
	Errors    []error
}

func (e *ValidationError) Error(message string) {
	e.Errors = append(e.Errors, errors.New(message))
}

func (e *ValidationError) Errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	e.Error(message)
}

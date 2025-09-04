package errors

import (
	"errors"
)

var (
	ErrTooManyFields    = errors.New("db(users): too many fields provided")
	ErrNoFieldsToUpdate = errors.New("db(users): no fields to update")
	ErrModelNotFound    = errors.New("db: model has not found with information provided")
)

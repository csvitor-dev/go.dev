package errors

import (
	"errors"
)

var (
	ErrTooManyFields    = errors.New("db(users): too many fields provided")
	ErrModelNotFound    = errors.New("db: model has not found with information provided")
	ErrNoFieldsToUpdate = errors.New("db(users): no fields to update")
)

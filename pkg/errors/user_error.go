package errors

import "errors"

var (
	ErrTooManyFields = errors.New("user: too many fields provided")
)

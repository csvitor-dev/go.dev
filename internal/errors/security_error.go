package errors

import "errors"

var (
	ErrPasswordNoProvided = errors.New("security: the password field hasn't been provided")
	ErrHashPassword       = errors.New("security: failed to hash the password")
)

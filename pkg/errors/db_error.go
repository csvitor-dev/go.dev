package errors

import "errors"

var (
	ErrTooManyFields = errors.New("db(users): too many fields provided")
	ErrUserNotFound  = errors.New("db(users): user not found")
	ErrNoFieldsToUpdate = errors.New("db(users): no fields to update")
)

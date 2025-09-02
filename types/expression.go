package types

type Expression interface {
	GetValidationErrorParams() (string, []error)
	Result() *ValidationError
}

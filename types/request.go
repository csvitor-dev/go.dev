package types

type Request interface {
	Validate() RequestValidationGuard
}

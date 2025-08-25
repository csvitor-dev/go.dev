package types

type RequestValidationGuard struct {
	Payload map[string][]error
}

func (r *RequestValidationGuard) HasErrors() bool {
	for _, errors := range r.Payload {
		if len(errors) > 0 {
			return true
		}
	}
	return false
}

func Throw(fields ...*ValidationError) RequestValidationGuard {
	result := map[string][]error{}

	for _, field := range fields {
		if len(field.Errors) > 0 {
			result[field.FieldName] = field.Errors
		}
	}
	return RequestValidationGuard{
		Payload: result,
	}
}

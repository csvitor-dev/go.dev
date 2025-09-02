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

func Throw(expressions ...Expression) RequestValidationGuard {
	result := map[string][]error{}

	for _, exp := range expressions {
		fieldName, errors := exp.GetValidationErrorParams()
		if len(errors) > 0 {
			result[fieldName] = errors
		}
	}
	return RequestValidationGuard{
		Payload: result,
	}
}

package requests

import "github.com/csvitor-dev/social-media/utils/validations"

type RequestOutput struct {
	Payload map[string][]error
}

func (r *RequestOutput) HasErrors() bool {
	for _, errors := range r.Payload {
		if len(errors) > 0 {
			return true
		}
	}
	return false
}

func GenerateOutput(fields ...*validations.ValidationError) RequestOutput {
	result := map[string][]error{}

	for _, field := range fields {
		if len(field.Errors) > 0 {
			result[field.FieldName] = field.Errors
		}
	}
	return RequestOutput{
		Payload: result,
	}
}

package validations

import (
	"github.com/csvitor-dev/go.dev/types"
)

type StringExpression struct {
	validation *types.ValidationError
	payload    string
	isOptional bool
}

func NewString(input, fieldName string) *StringExpression {
	return &StringExpression{
		payload: input,
		validation: &types.ValidationError{
			FieldName: fieldName,
		},
	}
}

func (exp *StringExpression) GetValidationErrorParams() (string, []error) {
	return exp.validation.FieldName, exp.validation.Errors
}

func (exp *StringExpression) Result() *types.ValidationError {
	return exp.validation
}

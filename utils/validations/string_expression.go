package validations

import (
	"net/mail"

	"github.com/dgrijalva/jwt-go"
)

type StringExpression struct {
	*ValidationError
	payload    string
	isOptional bool
}

func NewString(input, fieldName string) *StringExpression {
	return &StringExpression{
		payload: input,
		ValidationError: &ValidationError{
			FieldName: fieldName,
		},
	}
}

func (exp *StringExpression) IsOptional() *StringExpression {
	exp.isOptional = true
	return exp
}

func (exp *StringExpression) IsNotEmpty() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if exp.payload == "" {
		exp.error("cannot be empty")
	}
	return exp
}

func (exp *StringExpression) MinLength(min int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min {
		exp.errorf("must be at least %d characters long", min)
	}
	return exp
}

func (exp *StringExpression) MaxLength(max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) > max {
		exp.errorf("must be at most %d characters long", max)
	}
	return exp
}

func (exp *StringExpression) Between(min, max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min || len(exp.payload) > max {
		exp.errorf("must be between %d and %d characters long", min, max)
	}
	return exp
}

func (exp *StringExpression) Email() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if _, err := mail.ParseAddress(exp.payload); err != nil {
		exp.error("must be a valid email address")
	}
	return exp
}

func (exp *StringExpression) JWT() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if _, err := jwt.Parse(
		exp.payload,
		func(t *jwt.Token) (any, error) {
			return nil, nil
		}); err != nil {
		exp.error("must be a valid JWT format")
	}
	return exp
}

func (exp *StringExpression) Refine(fn func(input string) (string, error)) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}
	value, err := fn(exp.payload)

	if err != nil {
		exp.error(err.Error())
	}
	exp.payload = value
	return exp
}

func (exp *StringExpression) Result() *ValidationError {
	return exp.ValidationError
}

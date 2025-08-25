package validations

import (
	"encoding/base64"
	"net/mail"
	"strings"

	"github.com/csvitor-dev/social-media/types"
)

type StringExpression struct {
	*types.ValidationError
	payload    string
	isOptional bool
}

func NewString(input, fieldName string) *StringExpression {
	return &StringExpression{
		payload: input,
		ValidationError: &types.ValidationError{
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
		exp.Error("cannot be empty")
	}
	return exp
}

func (exp *StringExpression) MinLength(min int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min {
		exp.Errorf("must be at least %d characters long", min)
	}
	return exp
}

func (exp *StringExpression) MaxLength(max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) > max {
		exp.Errorf("must be at most %d characters long", max)
	}
	return exp
}

func (exp *StringExpression) Between(min, max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min || len(exp.payload) > max {
		exp.Errorf("must be between %d and %d characters long", min, max)
	}
	return exp
}

func (exp *StringExpression) Email() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if _, err := mail.ParseAddress(exp.payload); err != nil {
		exp.Error("must be a valid email address")
	}
	return exp
}

func (exp *StringExpression) JWT() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}
	segments := strings.Split(exp.payload, ".")

	err := validateJWTSegments(segments)

	if len(segments) != 3 || err != nil {
		exp.Error("must be a valid JWT format")
	}
	return exp
}

func validateJWTSegments(segments []string) error {
	for _, segment := range segments {
		_, err := base64.RawURLEncoding.DecodeString(segment)

		if err != nil {
			return err
		}
	}
	return nil
}

func (exp *StringExpression) Refine(fn func(input string) (string, error)) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}
	value, err := fn(exp.payload)

	if err != nil {
		exp.Error(err.Error())
	}
	exp.payload = value
	return exp
}

func (exp *StringExpression) Result() *types.ValidationError {
	return exp.ValidationError
}

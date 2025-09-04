package validations

import (
	"encoding/base64"
	"errors"
	"net/mail"
	"regexp"
	"strings"

	"github.com/csvitor-dev/go.dev/types"
	"github.com/csvitor-dev/go.dev/utils/slices"
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

func (exp *StringExpression) IsOptional() *StringExpression {
	exp.isOptional = true
	return exp
}

func (exp *StringExpression) IsNotEmpty() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if exp.payload == "" {
		exp.validation.Error("cannot be empty")
	}
	return exp
}

func (exp *StringExpression) MinLength(min int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min {
		exp.validation.Errorf("must be at least %d characters long", min)
	}
	return exp
}

func (exp *StringExpression) MaxLength(max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) > max {
		exp.validation.Errorf("must be at most %d characters long", max)
	}
	return exp
}

func (exp *StringExpression) Between(min, max int) *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if len(exp.payload) < min || len(exp.payload) > max {
		exp.validation.Errorf("must be between %d and %d characters long", min, max)
	}
	return exp
}

func (exp *StringExpression) Email() *StringExpression {
	if exp.isOptional && exp.payload == "" {
		return exp
	}

	if _, err := mail.ParseAddress(exp.payload); err != nil {
		exp.validation.Error("must be a valid email address")
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
		exp.validation.Error("must be a valid JWT format")
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
		exp.validation.Error(err.Error())
	}
	exp.payload = value
	return exp
}

func (exp *StringExpression) TrimRefine() *StringExpression {
	removeExtraWhiteSpacesRegex, _ := regexp.Compile(`\s+`)
	exp.payload = strings.
		TrimSpace(removeExtraWhiteSpacesRegex.ReplaceAllString(exp.payload, " "))

	return exp
}

func AllOptionalExpressionsAreValid(expressions ...*StringExpression) *StringExpression {
	hasErrors := slices.Some(expressions,
		func(exp *StringExpression, _ int) bool {
			return len(exp.validation.Errors) > 0
		})

	if hasErrors {
		return nil
	}
	allAreInvalid := slices.Every(expressions,
		func(exp *StringExpression, _ int) bool {
			return exp.payload == ""
		})

	if allAreInvalid {
		return &StringExpression{
			validation: &types.ValidationError{
				FieldName: "bad_request",
				Errors:    []error{errors.New("at least one field must be provided")},
			},
		}
	}
	return nil
}

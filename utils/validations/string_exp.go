package validations

import (
	"fmt"
	"net/mail"
)

type StringExp struct {
	payload          string
	fieldName        string
	validationErrors []error
	isOptional       bool
}

func NewString(input, fieldName string) *StringExp {
	return &StringExp{
		payload:   input,
		fieldName: fieldName,
	}
}

func (s *StringExp) error(message string) error {
	return fmt.Errorf("%s: %s", s.fieldName, message)
}

func (s *StringExp) errorf(format string, args ...any) error {
	return fmt.Errorf("%s: %s", s.fieldName, fmt.Sprintf(format, args...))
}

func (s *StringExp) IsOptional() *StringExp {
	s.isOptional = true
	return s
}

func (s *StringExp) IsNotEmpty() *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}

	if s.payload == "" {
		s.validationErrors = append(s.validationErrors, s.error("cannot be empty"))
	}
	return s
}

func (s *StringExp) MinLength(min int) *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}

	if len(s.payload) < min {
		s.validationErrors = append(s.validationErrors, s.errorf("must be at least %d characters long", min))
	}
	return s
}

func (s *StringExp) MaxLength(max int) *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}

	if len(s.payload) > max {
		s.validationErrors = append(s.validationErrors, s.errorf("must be at most %d characters long", max))
	}
	return s
}

func (s *StringExp) Between(min, max int) *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}

	if len(s.payload) < min || len(s.payload) > max {
		s.validationErrors = append(s.validationErrors, s.errorf("must be between %d and %d characters long", min, max))
	}
	return s
}

func (s *StringExp) Email() *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}

	if _, err := mail.ParseAddress(s.payload); err != nil {
		s.validationErrors = append(s.validationErrors, s.error("must be a valid email address"))
	}
	return s
}

func (s *StringExp) Refine(fn func(input string) (string, error)) *StringExp {
	if s.isOptional && s.payload == "" {
		return s
	}
	value, err := fn(s.payload)

	if err != nil {
		s.validationErrors = append(s.validationErrors, s.error(err.Error()))
	}
	s.payload = value
	return s
}

func (s *StringExp) Result() []error {
	return s.validationErrors
}

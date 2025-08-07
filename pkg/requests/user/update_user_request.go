package user

import (
	"slices"
	"strings"

	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/pkg/utils/validations"
)

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (r *UpdateUserRequest) Validate() []error {
	nameErrors := validations.NewString(r.Name, "name").IsOptional().Between(3, 50).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	nickErrors := validations.NewString(r.Nickname, "nickname").IsOptional().Between(3, 50).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	emailErrors := validations.NewString(r.Email, "email").IsOptional().Between(12, 50).Email().
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	return slices.Concat(nameErrors, nickErrors, emailErrors)
}

func (r *UpdateUserRequest) Map() models.User {
	return models.User{
		Name:     r.Name,
		Nickname: r.Nickname,
		Email:    r.Email,
	}
}

package user

import (
	"slices"
	"strings"

	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/pkg/utils/validations"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUserRequest) Validate() []error {
	nameErrors := validations.NewString(r.Name, "name").IsNotEmpty().Between(3, 50).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	nickErrors := validations.NewString(r.Nickname, "nickname").IsNotEmpty().Between(3, 50).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	emailErrors := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	passwordErrors := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25).
		Refine(func(input string) (string, error) {
			return strings.TrimSpace(input), nil
		}).Result()

	return slices.Concat(nameErrors, nickErrors, emailErrors, passwordErrors)
}

func (r *RegisterUserRequest) Map() models.User {
	return models.User{
		Name:     r.Name,
		Nickname: r.Nickname,
		Email:    r.Email,
		Password: r.Password,
	}
}

package user

import (
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUserRequest) Validate() map[string][]error {
	nameErrors := validations.NewString(r.Name, "name").IsNotEmpty().Between(3, 50).Result()

	nickErrors := validations.NewString(r.Nickname, "nickname").IsNotEmpty().Between(3, 50).Result()

	emailErrors := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email().Result()

	passwordErrors := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25).Result()

	return map[string][]error{
		nameErrors.FieldName:     nameErrors.Errors,
		nickErrors.FieldName:     nickErrors.Errors,
		emailErrors.FieldName:    emailErrors.Errors,
		passwordErrors.FieldName: passwordErrors.Errors,
	}
}

func (r *RegisterUserRequest) Map() (models.User, error) {
	return models.NewUser(r.Name, r.Nickname, r.Email, r.Password)
}

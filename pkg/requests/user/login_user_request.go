package user

import (
	"github.com/csvitor-dev/social-media/utils/validations"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() map[string][]error {
	emailErrors := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email().Result()

	passwordErrors := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25).Result()

	return map[string][]error{
		emailErrors.FieldName:    emailErrors.Errors,
		passwordErrors.FieldName: passwordErrors.Errors,
	}
}

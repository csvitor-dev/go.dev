package user

import (
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() types.RequestValidationGuard {
	email := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email()
	password := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25)

	return types.Throw(email, password)
}

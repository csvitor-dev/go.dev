package user

import (
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUserRequest) Validate() types.RequestValidationGuard {
	name := validations.NewString(r.Name, "name").IsNotEmpty().Between(3, 50).TrimRefine()
	nick := validations.NewString(r.Nickname, "nickname").IsNotEmpty().Between(3, 50).TrimRefine()
	email := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email()
	password := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25)

	return types.Throw(name, nick, email, password)
}

func (r *RegisterUserRequest) Map() (models.User, error) {
	return models.NewUser(r.Name, r.Nickname, r.Email, r.Password)
}

package user

import (
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (r *UpdateUserRequest) Validate() types.RequestValidationGuard {
	nameErrors := validations.NewString(r.Name, "name").IsOptional().Between(3, 50).Result()

	nickErrors := validations.NewString(r.Nickname, "nickname").IsOptional().Between(3, 50).Result()

	emailErrors := validations.NewString(r.Email, "email").IsOptional().Between(12, 50).Email().Result()

	return types.Throw(nameErrors, nickErrors, emailErrors)
}

func (r *UpdateUserRequest) Map() (models.User, error) {
	return models.NewUser(r.Name, r.Nickname, r.Email)
}

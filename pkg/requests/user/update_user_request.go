package user

import (
	"github.com/csvitor-dev/go.dev/internal/models"
	"github.com/csvitor-dev/go.dev/types"
	"github.com/csvitor-dev/go.dev/utils/validations"
)

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (r *UpdateUserRequest) Validate() types.RequestValidationGuard {
	name := validations.NewString(r.Name, "name").IsOptional().Between(3, 50).TrimRefine()
	nickname := validations.NewString(r.Nickname, "nickname").IsOptional().Between(3, 50).TrimRefine()
	email := validations.NewString(r.Email, "email").IsOptional().Between(12, 50).Email()

	if optional := validations.
		AllOptionalExpressionsAreValid(
			name,
			nickname,
			email,
		); optional != nil {
		return types.Throw(optional)
	}
	return types.Throw(name, nickname, email)
}

func (r *UpdateUserRequest) Map() (models.User, error) {
	return models.NewUser(r.Name, r.Nickname, r.Email)
}

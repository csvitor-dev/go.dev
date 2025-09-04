package user

import (
	"github.com/csvitor-dev/go.dev/types"
	"github.com/csvitor-dev/go.dev/utils/validations"
)

type ResetUserPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (r *ResetUserPasswordRequest) Validate() types.RequestValidationGuard {
	token := validations.NewString(r.Token, "token").IsNotEmpty().JWT()
	password := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25)

	return types.Throw(token, password)
}

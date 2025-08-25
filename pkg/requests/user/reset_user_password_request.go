package user

import (
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type ResetUserPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (r *ResetUserPasswordRequest) Validate() types.RequestValidationGuard {
	tokenErrors := validations.NewString(r.Token, "token").IsNotEmpty().JWT().Result()
	passwordErrors := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25).Result()

	return types.Throw(tokenErrors, passwordErrors)
}

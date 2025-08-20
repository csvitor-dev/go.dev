package user

import (
	"github.com/csvitor-dev/social-media/pkg/requests"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type ResetUserPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (r *ResetUserPasswordRequest) Validate() requests.RequestOutput {
	tokenErrors := validations.NewString(r.Token, "token").IsNotEmpty().JWT().Result()
	passwordErrors := validations.NewString(r.Password, "password").IsNotEmpty().Between(8, 25).Result()

	return requests.GenerateOutput(tokenErrors, passwordErrors)
}

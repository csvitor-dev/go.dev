package user

import (
	"github.com/csvitor-dev/social-media/pkg/requests"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type RecoverUserPasswordRequest struct {
	Email string `json:"email"`
}

func (r *RecoverUserPasswordRequest) Validate() requests.RequestOutput {
	emailErrors := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email().Result()

	return requests.GenerateOutput(emailErrors)
}

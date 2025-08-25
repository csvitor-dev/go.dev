package user

import (
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type RecoverUserPasswordRequest struct {
	Email string `json:"email"`
}

func (r *RecoverUserPasswordRequest) Validate() types.RequestValidationGuard {
	emailErrors := validations.NewString(r.Email, "email").IsNotEmpty().Between(12, 50).Email().Result()

	return types.Throw(emailErrors)
}

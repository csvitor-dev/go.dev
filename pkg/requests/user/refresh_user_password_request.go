package user

import (
	"github.com/csvitor-dev/social-media/types"
	"github.com/csvitor-dev/social-media/utils/validations"
)

type RefreshUserPasswordRequest struct {
	CurrentPassword string `json:"current"`
	NewPassword     string `json:"new"`
}

func (r *RefreshUserPasswordRequest) Validate() types.RequestValidationGuard {
	current := validations.NewString(r.CurrentPassword, "current").IsNotEmpty().Between(8, 25)
	new := validations.NewString(r.NewPassword, "new").IsNotEmpty().Between(8, 25)

	return types.Throw(current, new)
}

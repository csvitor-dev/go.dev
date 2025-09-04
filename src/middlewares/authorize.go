package middlewares

import (
	"errors"
	"net/http"

	res "github.com/csvitor-dev/go.dev/pkg/responses"
	"github.com/csvitor-dev/go.dev/src/services/auth"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := utils.ExtractToken(r)

		if err := auth.ValidateToken(token); err != nil {
			res.SingleError(w, http.StatusUnauthorized, errors.New("auth: you're not authenticated"))
			return
		}
		next(w, r)
	}
}

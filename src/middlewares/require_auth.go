package middlewares

import (
	"net/http"

	"fmt"

	"github.com/csvitor-dev/go.dev/src/services/auth"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")

		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/auth/login?err=You+are+not+authenticated", http.StatusSeeOther)
			return
		}

		if err := auth.ValidateToken(cookie.Value); err != nil {
			path := fmt.Sprintf("/auth/login?err=%s", err.Error())
			http.Redirect(w, r, path, http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

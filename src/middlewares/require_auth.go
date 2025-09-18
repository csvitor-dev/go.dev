package middlewares

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/src/services/clients/api"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")

		if err != nil || cookie.Value == "" {
			utils.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		client := api.NewApiClient(env.WebEnv.API_URL).WithToken(cookie.Value)

		if _, err := client.Do(
			api.RequestOptions{
				Method:      http.MethodGet,
				Path:        "/auth/verify-token",
				Body:        nil,
				RequireAuth: true,
			}).Done(); err != nil {
			utils.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

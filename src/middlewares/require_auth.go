package middlewares

import (
	"fmt"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/services/clients/api"
	utils "github.com/csvitor-dev/social-media/utils/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")

		if err != nil || cookie.Value == "" {
			utils.Redirect(w, r, "/auth/login?err=You+are+not+authenticated", http.StatusSeeOther)
			return
		}
		client := api.NewApiClient(config.WebEnv.API_URL).WithToken(cookie.Value)

		if _, err := api.ExecuteRequest[any](
			client,
			api.RequestOptions{
				Method:      http.MethodGet,
				Path:        "/auth/verify-token",
				Body:        nil,
				RequireAuth: true,
			}); err != nil {
			path := fmt.Sprintf("/auth/login?err=%s", err.Error())
			utils.Redirect(w, r, path, http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

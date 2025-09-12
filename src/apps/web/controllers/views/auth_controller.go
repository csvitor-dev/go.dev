package views

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/views"
)

func GetLoginView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.login",
		})
}

func GetRegisterView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.register",
			Layout:     "default",
			Data: map[string]any{
				"ScriptUI": "register-user-form",
			},
		})
}

func GetForgotPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.forgot-password",
		})
}

func GetResetPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.reset-password",
		})
}

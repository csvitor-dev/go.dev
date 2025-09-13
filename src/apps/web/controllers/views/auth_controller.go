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
			Layout:     "default",
			Data: map[string]any{
				"ScriptUI": "login-user-form-page",
			},
		})
}

func GetRegisterView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.register",
			Layout:     "default",
			Data: map[string]any{
				"ScriptUI": "register-user-form-page",
			},
		})
}

func GetForgotPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.forgot-password",
			Layout:     "default",
			Data: map[string]any{
				"ScriptUI": "forgot-password-form-page",
			},
		})
}

func GetResetPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "auth.reset-password",
			Layout:     "default",
			Data: map[string]any{
				"ScriptUI": "reset-password-form-page",
			},
		})
}

package views

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/views"
)

func GetLoginView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.login", nil)
}

func GetRegisterView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.register", nil)
}

func GetForgotPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.forgot-password", nil)
}

func GetResetPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.reset-password", nil)
}

package controllers

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/apps/web/views"
)

func GetLoginView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.login", nil)
}

func GetForgotPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.forgot-password", nil)
}

func GetResetPasswordView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "auth.reset-password", nil)
}

func FetchApiForTokenValidation(w http.ResponseWriter, r *http.Request) {

}

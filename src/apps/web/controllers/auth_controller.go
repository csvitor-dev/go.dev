package controllers

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/apps/web/views"
)

func GetLoginView(w http.ResponseWriter, r *http.Request) {
	views.Render(w, http.StatusOK, "login", nil)
}

func FetchApiForTokenValidation(w http.ResponseWriter, r *http.Request) {

}

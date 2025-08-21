package views

import (
	"net/http"

	res "github.com/csvitor-dev/social-media/pkg/responses"
)

func Render(w http.ResponseWriter, status int, view string, data map[string]any) {
	fileName := view + ".html"

	res.View(w, status)
	templates.ExecuteTemplate(w, fileName, data)
}

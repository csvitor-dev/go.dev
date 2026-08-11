package views

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/views"
)

func GetHomeView(w http.ResponseWriter, r *http.Request) {
	views.Render(w,
		views.ViewOptions{
			StatusCode: http.StatusOK,
			View:       "guard.home",
		})
}

package responses

import (
	"net/http"

	utils "github.com/csvitor-dev/social-media/utils/http"
)

// View: prepare HTML output
func View(w http.ResponseWriter, status int) {
	utils.WriteStatus(w, status)
	w.Header().Set("Content-Type", "text/html")
}

package responses

import (
	"fmt"
	"net/http"

	utils "github.com/csvitor-dev/social-media/utils/http"
)

// View: prepare HTML output
func View(w http.ResponseWriter, status int) {
	utils.WriteStatus(w, status)
	w.Header().Set("Content-Type", "text/html")
}

// ErrorView: prepare HTML error output
func ErrorView(w http.ResponseWriter, message string, status int) {
	View(w, status)

	http.Error(w, fmt.Sprintf("%s\n%s", http.StatusText(status), message), status)
}

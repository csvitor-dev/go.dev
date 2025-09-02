package views

import (
	"errors"
	"net/http"
	"strings"

	res "github.com/csvitor-dev/social-media/pkg/responses"
)

func Render(w http.ResponseWriter, status int, view string, data map[string]any) {
	folder, fileName, err := getViewPattern(view)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template, err := LoadTemplateFrom(folder, fileName)

	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res.View(w, status)

	if err := template.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func getViewPattern(view string) (string, string, error) {
	parts := strings.Split(view, ".")

	if len(parts) != 2 {
		return "", "", errors.New("views: invalid view format, expected 'folder.filename'")
	}
	return parts[0], parts[1], nil
}

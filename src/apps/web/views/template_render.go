package views

import (
	"net/http"
	"strings"

	res "github.com/csvitor-dev/go.dev/pkg/responses"
)

func Render(w http.ResponseWriter, status int, view string, data map[string]any) {
	path, fileName := getViewPattern(view)

	template, err := LoadTemplateFrom(path, fileName)

	if err != nil {
		res.ErrorView(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res.View(w, status)

	if err := template.Execute(w, data); err != nil {
		res.ErrorView(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func getViewPattern(view string) (string, string) {
	parts := strings.Split(view, ".")
	lastIndex := len(parts) - 1
	path := strings.Join(parts[:lastIndex], "/")

	return path, parts[lastIndex]
}

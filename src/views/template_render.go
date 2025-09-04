package views

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	res "github.com/csvitor-dev/go.dev/pkg/responses"
)

func getViewPattern(view string) (string, string) {
	parts := strings.Split(view, ".")
	lastIndex := len(parts) - 1
	path := strings.Join(parts[:lastIndex], "/")

	return path, parts[lastIndex]
}

func writeView(w io.Writer, view string, data map[string]any) error {
	path, fileName := getViewPattern(view)
	template, err := LoadTemplateFrom(path, fileName)

	if err != nil {
		return err
	}

	if err := template.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func Render(w http.ResponseWriter, status int, view string, data map[string]any) {
	res.View(w, status)

	if err := writeView(w, view, data); err != nil {
		res.ErrorView(w, "view: "+err.Error(), http.StatusInternalServerError)
	}
}

func Get(view string, data map[string]any) (string, error) {
	var buf bytes.Buffer

	if err := writeView(&buf, view, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

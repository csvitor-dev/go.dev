package views

import (
	"bytes"
	"io"
	"net/http"

	res "github.com/csvitor-dev/go.dev/pkg/responses"
)

func writeView(w io.Writer, options ViewOptions) error {
	template, err := LoadTemplateFrom(options)

	if err != nil {
		return err
	}

	if err := template.Execute(w, options.Data); err != nil {
		return err
	}
	return nil
}

func Render(w http.ResponseWriter, options ViewOptions) {
	res.View(w, options.StatusCode)

	if err := writeView(w, options); err != nil {
		res.ErrorView(w, "view: "+err.Error(), http.StatusInternalServerError)
	}
}

func Get(options ViewOptions) (string, error) {
	var buf bytes.Buffer

	if err := writeView(&buf, options); err != nil {
		return "", err
	}
	return buf.String(), nil
}

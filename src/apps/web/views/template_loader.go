package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

func LoadTemplateFrom(folder, fileName string) (*template.Template, error) {
	basePath, err := filepath.Abs("./src/apps/web/views")

	if err != nil {
		return nil, err
	}
	pattern := fmt.Sprintf("%s/%s/%s.html", basePath, folder, fileName)

	return template.ParseFiles(pattern)
}

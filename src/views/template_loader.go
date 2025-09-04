package views

import (
	"fmt"
	"html/template"
)

func LoadTemplateFrom(folder, fileName string) (*template.Template, error) {
	pattern := fmt.Sprintf("./src/views/%s/%s.html", folder, fileName)
	return template.ParseFiles(pattern)
}

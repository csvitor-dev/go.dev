package views

import (
	"fmt"
	"html/template"
)

func LoadTemplateFrom(options ViewOptions) (*template.Template, error) {
	folder, fileName := options.GetViewPattern()
	pagePath := fmt.Sprintf("./src/views/pages/%s/%s.html", folder, fileName)

	if !options.IsThereLayout() {
		return template.ParseFiles(pagePath)
	}
	layoutPath := options.GetLayoutPattern()
	return template.ParseFiles(layoutPath, pagePath)
}

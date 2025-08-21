package views

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

var templates *template.Template

func LoadTemplates() {
	basePath, err := filepath.Abs("./src/apps/web/views")

	if err != nil {
		log.Fatalf("No such path: %v\n", basePath)
	}
	pattern := fmt.Sprintf("%s/**/*.html", basePath)
	templatesTemp := template.Must(template.ParseGlob(pattern))

	templates = templatesTemp

	log.Println("Templates loaded successfully!")
}

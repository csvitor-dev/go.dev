package views

import (
	"fmt"
	"strings"
)

type ViewOptions struct {
	StatusCode int
	View       string
	Layout     string
	Data       map[string]any
}

func (options ViewOptions) IsThereLayout() bool {
	return options.Layout != ""
}

func (options ViewOptions) GetViewPattern() (string, string) {
	parts := strings.Split(options.View, ".")
	lastIndex := len(parts) - 1
	path := strings.Join(parts[:lastIndex], "/")

	return path, parts[lastIndex]
}

func (options ViewOptions) GetLayoutPattern() string {
	return fmt.Sprintf("./src/views/layouts/%s.html", options.Layout)
}

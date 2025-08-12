package cli

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func StatusText(format string, status int) string {
	message := fmt.Sprintf(format, status, http.StatusText(status))

	switch {
	case 200 <= status && status <= 299:
		return color.GreenString(message)
	case status >= 400:
		return color.RedString(message)
	default:
		return message
	}
}

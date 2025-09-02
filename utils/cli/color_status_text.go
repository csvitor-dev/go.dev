package cli

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func StatusText(format string, status int) string {
	message := fmt.Sprintf(format, status, http.StatusText(status))

	switch {
	case 100 <= status && status <= 199:
		return color.YellowString(message)
	case 200 <= status && status <= 299:
		return color.GreenString(message)
	case 300 <= status && status <= 399:
		return color.BlueString(message)
	case 400 <= status && status <= 499:
		return color.MagentaString(message)
	case status >= 500:
		return color.RedString(message)
	default:
		return message
	}
}

package http

import (
	"net/http"
	"strings"
)

func ExtractToken(request *http.Request) string {
	raw := request.Header.Get("Authorization")
	parts := strings.Split(raw, " ")

	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

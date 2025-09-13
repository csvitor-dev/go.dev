package api

import (
	"bytes"
)

type RequestOptions struct {
	Body        *bytes.Buffer
	RequireAuth bool
	Method      string
	Path        string
}

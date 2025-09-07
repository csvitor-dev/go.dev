package api

type RequestOptions struct {
	Body        any
	RequireAuth bool
	Method      string
	Path        string
}

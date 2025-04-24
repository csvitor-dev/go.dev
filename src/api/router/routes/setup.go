package routes

import "net/http"

// Route: represents an allowed API route
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

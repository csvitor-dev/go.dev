package types

import (
	"net/http"
)

// Route: represents an allowed API route
type Route struct {
	Uri     string
	Method  string
	Handler http.HandlerFunc
	*MiddlewarePipeline
}

func (route *Route) GetHandler() http.HandlerFunc {
	return route.MiddlewarePipeline.Apply(route.Handler)
}

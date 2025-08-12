package routes

import (
	"net/http"
	"slices"

	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

// Route: represents an allowed API route
type Route struct {
	Uri     string
	Method  string
	Handler http.HandlerFunc
	*middlewares.MiddlewareTags
}

func (route *Route) GetHandler() http.HandlerFunc {
	if !route.HasTag("log") {
		route.AddTag("log", true)
	}
	return middlewares.Apply(route.Handler, route.AllTags()...)
}

// All: returns all avaliable routes
func All() []Route {
	return slices.Concat(userRoutes, authRoutes)
}

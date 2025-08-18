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
	*middlewares.MiddlewarePipeline
}

func (route *Route) GetHandler() http.HandlerFunc {
	route.AddLogger()
	return middlewares.Apply(route.Handler, route.MiddlewarePipeline.All())
}

// All: returns all avaliable routes
func All() []Route {
	return slices.Concat(userRoutes, authRoutes)
}

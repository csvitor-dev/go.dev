package router

import (
	"github.com/csvitor-dev/social-media/src/api/router/routes"
	"github.com/gorilla/mux"
)

// Configure: sets up the API routes
func Configure(r *mux.Router) {
	for _, route := range routes.UserRoutes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
}

// Generate: returns a new HTTP Handler router
func Generate() *mux.Router {
	r := mux.NewRouter()
	Configure(r)

	return r
}

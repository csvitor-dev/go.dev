package router

import (
	"github.com/csvitor-dev/social-media/types"
	"github.com/gorilla/mux"
)

// configure: sets up the API routes
func configure(r *mux.Router, routes []types.Route) {
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.GetHandler()).Methods(route.Method)
	}
}

// Generate: returns a new HTTP Handler router
func Generate(routes []types.Route) *mux.Router {
	r := mux.NewRouter()
	configure(r, routes)

	return r
}

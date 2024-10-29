package router

import (
	"github.com/csvitor-dev/social-media/api/src/router/routes"
	"github.com/gorilla/mux"
)

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

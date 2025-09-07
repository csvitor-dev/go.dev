package router

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/src/views"
	"github.com/csvitor-dev/go.dev/types"
	"github.com/gorilla/mux"
)

// configure: sets up the API routes
func configure(r *mux.Router, routes []types.Route) {
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.EnqueueHandler()).Methods(route.Method)
	}
}

// Generate: returns a new HTTP Handler router
func Generate(routes []types.Route, extraConfig ...func(router *mux.Router)) *mux.Router {
	r := mux.NewRouter()
	configure(r, routes)

	for _, config := range extraConfig {
		config(r)
	}
	return r
}

func MapDefaultRoutes(r *mux.Router) {
	notFound := func(w http.ResponseWriter, r *http.Request) {
		views.Render(w,
			views.ViewOptions{
				StatusCode: http.StatusNotFound,
				View:       "errors.404",
			},
		)
	}
	r.NotFoundHandler = http.HandlerFunc(middlewares.Logger(notFound))
}

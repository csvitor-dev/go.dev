package routes

import (
	"log"
	"net/http"
	"slices"

	"github.com/csvitor-dev/social-media/pkg/types"
)

// Route: represents an allowed API route
type Route struct {
	Uri         string
	Method      string
	Handler     func(http.ResponseWriter, *http.Request) types.StatusCode
	RequireAuth bool
}

func (route *Route) Call(w http.ResponseWriter, r *http.Request) {
	log.Printf("<< %s %s\n", r.Method, r.URL.Path)

	status := route.Handler(w, r)

	log.Printf(">> %s %s %v\n", r.Method, r.URL.Path, status)
}

func GetAll() []Route {
	return slices.Concat(userRoutes, authRoutes)
}

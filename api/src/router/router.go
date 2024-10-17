package router

import (
	"github.com/gorilla/mux"
)

// Generate: returns a new HTTP Handler router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return r
}

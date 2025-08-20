package middlewares

import (
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func Apply(target http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		target = middleware(target)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		target(w, r)
	}
}

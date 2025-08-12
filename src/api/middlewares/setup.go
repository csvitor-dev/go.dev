package middlewares

import (
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var allMiddlewares = map[string]Middleware{
	"log":    Logger,
	"auth-z": Authorize,
}

func Apply(target http.HandlerFunc, tags ...string) http.HandlerFunc {
	for _, tag := range tags {
		middleware := allMiddlewares[tag]
		target = middleware(target)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		target(w, r)
	}
}

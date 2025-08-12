package middlewares

import (
	"log"
	"net/http"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth...")
		next(w, r)
	}
}

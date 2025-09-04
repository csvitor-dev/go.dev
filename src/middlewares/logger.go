package middlewares

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/go.dev/utils/cli"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("<< %s %s\n", r.Method, r.RequestURI)

		next(w, r)
		status := utils.Recorder.Status
		message := cli.StatusText("%d %s", status)

		log.Printf(">> %s %s %v\n", r.Method, r.RequestURI, message)
	}
}

package middlewares

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/utils/cli"
	utils "github.com/csvitor-dev/social-media/utils/http"
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

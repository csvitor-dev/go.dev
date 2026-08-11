package http

import (
	"net/http"
)

type StatusCodeRecorder struct {
	Status int
}

var Recorder StatusCodeRecorder

func WriteStatus(w http.ResponseWriter, status int) {
	Recorder.Status = status
	w.WriteHeader(status)
}

func Redirect(w http.ResponseWriter, r *http.Request, url string, status int) {
	Recorder.Status = status
	http.Redirect(w, r, url, status)
}

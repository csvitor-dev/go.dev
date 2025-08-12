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

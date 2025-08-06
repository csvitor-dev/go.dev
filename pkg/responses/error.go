package responses

import (
	"net/http"

	"github.com/csvitor-dev/social-media/pkg/types"
)

func Error(w http.ResponseWriter, status types.StatusCode, errs []error) types.StatusCode {
	var hook []string

	for _, v := range errs {
		hook = append(hook, v.Error())
	}
	return Json(w, status, struct {
		Errors []string `json:"errors"`
	}{
		Errors: hook,
	})
}

func SingleError(w http.ResponseWriter, status types.StatusCode, err error) types.StatusCode {
	return Json(w, status, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

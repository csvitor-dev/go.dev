package responses

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/utils/slices"
)

func ValidationErrors(w http.ResponseWriter, status int, errs map[string][]error) {
	hook := map[string][]string{}

	for key, errors := range errs {
		hook[key] = slices.Map(errors, func(err error, i int) string {
			return err.Error()
		})
	}
	Json(w, status, struct {
		Errors map[string][]string `json:"errors"`
	}{
		Errors: hook,
	})
}

func SingleError(w http.ResponseWriter, status int, err error) {
	Json(w, status, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

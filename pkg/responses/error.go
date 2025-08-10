package responses

import (
	"net/http"

	"github.com/csvitor-dev/social-media/pkg/types"
	"github.com/csvitor-dev/social-media/utils/slices"
)

func ValidationErrors(w http.ResponseWriter, status types.StatusCode, errs map[string][]error) types.StatusCode {
	hook := map[string][]string{}

	for key, errors := range errs {
		hook[key] = slices.Map(errors, func(err error, i int) string {
			return err.Error()
		})
	}
	return Json(w, status, struct {
		Errors map[string][]string `json:"errors"`
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

/*
	"errors": {
		"1": [],
		"2": [],
	}
	map[string][]string
*/

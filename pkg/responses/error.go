package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func ClientError(w http.ResponseWriter, r *http.Response) {
	defer r.Body.Close()
	var apiError struct {
		Errors map[string][]string `json:"errors,omitempty"`
		Error  string              `json:"error,omitempty"`
	}

	json.NewDecoder(r.Body).Decode(&apiError)

	if apiError.Errors != nil {
		var fields []string

		for field := range apiError.Errors {
			fields = append(fields, field)
		}
		apiError.Error = fmt.Sprintf(
			"client: validation errors in fields (%s)",
			strings.Join(fields, ", "),
		)
	}

	Json(w, r.StatusCode,
		map[string]string{
			"error": apiError.Error,
		},
	)
}

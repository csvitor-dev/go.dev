package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/csvitor-dev/go.dev/pkg/responses"
	"github.com/csvitor-dev/go.dev/types"
)

// MapRequestErrorWriter: defines a function that writes an error to an HTTP response
type MapRequestErrorWriter func(http.ResponseWriter)

// MapToRequest: maps the given data to the request and validates it, returning any error encountered
func MapToRequest(request types.Request, data io.ReadCloser) MapRequestErrorWriter {
	body, err := io.ReadAll(data)

	if err != nil {
		return func(w http.ResponseWriter) {
			responses.SingleError(w, http.StatusInternalServerError, err)
		}
	}

	if err := json.Unmarshal(body, &request); err != nil {
		return func(w http.ResponseWriter) {
			responses.SingleError(w, http.StatusBadRequest, err)
		}
	}

	if errs := request.Validate(); errs.HasErrors() {
		return func(w http.ResponseWriter) {
			responses.ValidationErrors(w, http.StatusUnprocessableEntity, errs.Payload)
		}
	}
	return nil
}

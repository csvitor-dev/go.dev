package requests

import (
	"encoding/json"
	"net/http"

	"github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/types"
)

func MapToRequest(w http.ResponseWriter, request types.Request, data []byte) {
	if err := json.Unmarshal(data, &request); err != nil {
		responses.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if errs := request.Validate(); errs.HasErrors() {
		responses.ValidationErrors(w, http.StatusUnprocessableEntity, errs.Payload)
		return
	}
}

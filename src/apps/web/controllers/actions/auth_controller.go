package actions

import (
	"bytes"
	"io"
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	res "github.com/csvitor-dev/go.dev/pkg/responses"
	"github.com/csvitor-dev/go.dev/src/services/clients/api"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

func RegisterUserAction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	client := api.NewApiClient(env.WebEnv.API_URL)

	response, err := client.Do(
		api.RequestOptions{
			Path:   "/auth/register",
			Method: "POST",
			Body:   bytes.NewBuffer(body),
		},
	).Done()

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if utils.IsErrorResponse(response) {
		res.ClientError(w, response)
		return
	}
	res.Json(w, response.StatusCode, nil)
}

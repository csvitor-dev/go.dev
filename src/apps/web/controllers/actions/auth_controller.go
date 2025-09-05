package actions

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	res "github.com/csvitor-dev/go.dev/pkg/responses"
	"github.com/csvitor-dev/go.dev/src/services/clients/api"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

func RegisterUserAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	client := api.NewApiClient(env.WebEnv.API_URL)

	response, err := client.Do(
		api.RequestOptions{
			Path:   "/auth/register",
			Method: "POST",
			Body: map[string]string{
				"name":     r.FormValue("name"),
				"nickname": r.FormValue("nickname"),
				"email":    r.FormValue("email"),
				"password": r.FormValue("password"),
			},
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

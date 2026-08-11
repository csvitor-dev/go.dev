package actions

import (
	"bytes"
	"io"
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	res "github.com/csvitor-dev/go.dev/pkg/responses"
	"github.com/csvitor-dev/go.dev/pkg/responses/user"
	"github.com/csvitor-dev/go.dev/src/services/clients/api"
	"github.com/csvitor-dev/go.dev/src/services/cookies"
)

func RegisterUserAction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	client := api.NewApiClient(env.WebEnv.API_URL, w)

	response, status := api.ExecuteRequest[user.TokenResponse](
		client,
		api.RequestOptions{
			Path:   "/auth/register",
			Method: "POST",
			Body:   bytes.NewBuffer(body),
		},
	)

	res.Json(w, status, response)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	client := api.NewApiClient(env.WebEnv.API_URL, w)

	response, status := api.ExecuteRequest[user.TokenResponse](
		client,
		api.RequestOptions{
			Path:   "/auth/login",
			Method: "POST",
			Body:   bytes.NewBuffer(body),
		},
	)

	if err := cookies.Save(
		w,
		response.UserId,
		response.Token,
	); err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	res.Json(w, status, nil)
}

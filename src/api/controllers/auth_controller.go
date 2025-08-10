package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	"github.com/csvitor-dev/social-media/pkg/requests/user"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/pkg/security"
	"github.com/csvitor-dev/social-media/pkg/types"
	"github.com/csvitor-dev/social-media/src/api/services/auth"
)

func Login(w http.ResponseWriter, r *http.Request) types.StatusCode {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return res.SingleError(w, http.StatusUnprocessableEntity, err)
	}
	var request user.LoginUserRequest

	if err = json.Unmarshal(body, &request); err != nil {
		return res.SingleError(w, http.StatusBadRequest, err)
	}

	if errs := request.Validate(); len(errs) > 0 {
		return res.ValidationErrors(w, http.StatusBadRequest, errs)
	}
	db, err := db.Connect()

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	user, err := repo.FindByEmail(request.Email)

	if err != nil {
		return res.SingleError(w, http.StatusNotFound, err)
	}

	if err = security.VerifyPassword(user.Password, request.Password); err != nil {
		return res.SingleError(w, http.StatusUnauthorized, err)
	}
	token, err := auth.CreateToken(user.Id)

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	return res.Json(w, http.StatusOK, map[string]any{
		"user_id": user.Id,
		"token":   token,
	})
}

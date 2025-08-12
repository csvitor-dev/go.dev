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
	"github.com/csvitor-dev/social-media/src/api/services/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request user.LoginUserRequest

	if err = json.Unmarshal(body, &request); err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if errs := request.Validate(); errs.HasErrors() {
		res.ValidationErrors(w, http.StatusBadRequest, errs.Payload)
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	user, err := repo.FindByEmail(request.Email)

	if err != nil {
		res.SingleError(w, http.StatusNotFound, err)
		return
	}

	if err = security.VerifyPassword(user.Password, request.Password); err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	token, err := auth.CreateToken(user.Id)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusOK, map[string]any{
		"user_id": user.Id,
		"token":   token,
	})
}

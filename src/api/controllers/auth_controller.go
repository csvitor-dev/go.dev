package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	pkg "github.com/csvitor-dev/social-media/pkg/errors"
	"github.com/csvitor-dev/social-media/pkg/requests/user"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/pkg/security"
	"github.com/csvitor-dev/social-media/src/api/services/auth"
)

// Register: creates a user and delegates its persistence
func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request user.RegisterUserRequest

	if err = json.Unmarshal(body, &request); err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if errs := request.Validate(); errs.HasErrors() {
		res.ValidationErrors(w, http.StatusBadRequest, errs.Payload)
		return
	}
	user, err := request.Map()

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	result, err := repo.Create(user)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusCreated, map[string]any{
		"id": result,
	})
}

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

func RefreshPassword(w http.ResponseWriter, r *http.Request) {
	authUserId, err := auth.GetUserId()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request user.RefreshUserPasswordRequest

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
	currentPassword, err := repo.FindPasswordFromUser(authUserId)

	if err != nil {
		var status int

		if errors.Is(err, pkg.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		res.SingleError(w, status, err)
		return
	}

	if err = security.VerifyPassword(
		currentPassword,
		request.CurrentPassword,
	); err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	hashedPassword, err := security.Cryptify(request.NewPassword)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	
	if err = repo.RefreshPasswordFromUser(authUserId, hashedPassword); err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

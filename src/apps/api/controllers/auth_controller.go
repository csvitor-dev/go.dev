package controllers

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	pkg "github.com/csvitor-dev/social-media/pkg/errors"
	"github.com/csvitor-dev/social-media/pkg/requests"
	"github.com/csvitor-dev/social-media/pkg/requests/user"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/pkg/security"
	"github.com/csvitor-dev/social-media/src/services/auth"
	"github.com/csvitor-dev/social-media/src/services/email"
	"github.com/csvitor-dev/social-media/types"
	utils "github.com/csvitor-dev/social-media/utils/http"
)

// Register: creates a user and delegates its persistence
func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request user.RegisterUserRequest
	requests.MapToRequest(w, &request, body)
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
	requests.MapToRequest(w, &request, body)

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
	token, err := auth.CreateToken(user, time.Hour)

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
	authUserId, err := auth.GetUserIdFromToken()

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
	requests.MapToRequest(w, &request, body)

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

		if errors.Is(err, pkg.ErrModelNotFound) {
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
	auth.InvalidateToken()
	res.Json(w, http.StatusNoContent, nil)
}

func RecoverPassword(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	var request user.RecoverUserPasswordRequest
	requests.MapToRequest(w, &request, body)

	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	targetUser, err := repo.FindByEmail(request.Email)

	if err != nil {
		if errors.Is(err, pkg.ErrModelNotFound) {
			res.Json(w, http.StatusNoContent, nil)
			return
		}
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	token, err := auth.CreateToken(targetUser, time.Minute*15)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	toSendEmail := types.Email{
		To:      request.Email,
		Subject: "Recuperação de senha",
	}
	err = email.SendEmailForPasswordReset(toSendEmail, token)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	token := utils.ExtractToken(r)
	err := auth.ValidateToken(token)

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	var request user.ResetUserPasswordRequest
	requests.MapToRequest(w, &request, body)

	err = auth.ValidateToken(request.Token)

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	userId, err := auth.GetUserIdFromToken()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	hashedPassword, err := security.Cryptify(request.Password)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.RefreshPasswordFromUser(userId, hashedPassword); err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	auth.InvalidateToken()

	res.Json(w, http.StatusNoContent, nil)
}

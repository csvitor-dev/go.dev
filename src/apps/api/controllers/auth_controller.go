package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/csvitor-dev/go.dev/internal/db"
	repos "github.com/csvitor-dev/go.dev/internal/db/repositories"
	internal_errors "github.com/csvitor-dev/go.dev/internal/errors"
	"github.com/csvitor-dev/go.dev/internal/security"
	"github.com/csvitor-dev/go.dev/pkg/requests"
	"github.com/csvitor-dev/go.dev/pkg/requests/user"
	res "github.com/csvitor-dev/go.dev/pkg/responses"
	user_res "github.com/csvitor-dev/go.dev/pkg/responses/user"
	"github.com/csvitor-dev/go.dev/src/services/auth"
	"github.com/csvitor-dev/go.dev/src/services/email"
	"github.com/csvitor-dev/go.dev/types"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

// Register: creates a user and delegates its persistence
func Register(w http.ResponseWriter, r *http.Request) {
	var request user.RegisterUserRequest

	if writer := requests.
		MapToRequest(&request, r.Body); writer != nil {
		writer(w)
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
	var request user.LoginUserRequest

	if writer := requests.
		MapToRequest(&request, r.Body); writer != nil {
		writer(w)
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
	token, err := auth.CreateToken(user, time.Hour)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusOK,
		user_res.TokenResponse{
			UserId: user.Id,
			Token:  token,
		})
}

func RefreshPassword(w http.ResponseWriter, r *http.Request) {
	authUserId, err := auth.GetUserIdFromToken()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	var request user.RefreshUserPasswordRequest

	if writer := requests.
		MapToRequest(&request, r.Body); writer != nil {
		writer(w)
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

		if errors.Is(err, internal_errors.ErrModelNotFound) {
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
	var request user.RecoverUserPasswordRequest

	if writer := requests.
		MapToRequest(&request, r.Body); writer != nil {
		writer(w)
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	targetUser, err := repo.FindByEmail(request.Email)

	if err != nil {
		if errors.Is(err, internal_errors.ErrModelNotFound) {
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
	var request user.ResetUserPasswordRequest

	if writer := requests.
		MapToRequest(&request, r.Body); writer != nil {
		writer(w)
		return
	}
	err := auth.ValidateToken(request.Token)

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

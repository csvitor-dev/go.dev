package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	pkg "github.com/csvitor-dev/social-media/pkg/errors"
	"github.com/csvitor-dev/social-media/src/api/services/auth"
	"github.com/gorilla/mux"

	"github.com/csvitor-dev/social-media/pkg/requests/user"
	res "github.com/csvitor-dev/social-media/pkg/responses"
)

// GetAllUsers: retrieves all persisted users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	result, err := repo.FindAll()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusOK, result)
}

// GetUserById: retrieves a persisted user via a given id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

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
	user, err := repo.FindById(userId)

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
	res.Json(w, http.StatusOK, user)
}

// UpdateUserById: updates a user based on the provided id
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	authId, err := auth.GetUserId()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}

	if userId != authId {
		res.SingleError(w, http.StatusForbidden, errors.New("auth: target 'user_id' mismatch"))
		return
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request user.UpdateUserRequest

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
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)

	if err := repo.Update(userId, user); err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

// DeleteUserById: deletes a user based on the provided id
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}
	authId, err := auth.GetUserId()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}

	if userId != authId {
		res.SingleError(w, http.StatusForbidden, errors.New("auth: target 'user_id' mismatch"))
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)

	if err := repo.Delete(userId); err != nil {
		res.SingleError(w, http.StatusNotFound, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

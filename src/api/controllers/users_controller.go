package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/gorilla/mux"

	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/pkg/types"
)

// GetAllUsers: retrieves all persisted users
func GetAllUsers(w http.ResponseWriter, r *http.Request) types.StatusCode {
	db, err := db.Connect()

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	result, err := repo.GetUsers()

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	return res.Json(w, http.StatusOK, result)
}

// GetUserById: retrieves a persisted user via a given id
func GetUserById(w http.ResponseWriter, r *http.Request) types.StatusCode {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		return res.SingleError(w, http.StatusBadRequest, err)
	}
	db, err := db.Connect()

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repos.NewUsersRepository(db)
	user, err := repo.GetById(userId)

	if err != nil {
		return res.SingleError(w, http.StatusNotFound, err)
	}
	return res.Json(w, http.StatusOK, user)
}

// CreateUser: creates a user and delegates its persistence
func CreateUser(w http.ResponseWriter, r *http.Request) types.StatusCode {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return res.SingleError(w, http.StatusUnprocessableEntity, err)
	}
	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		return res.SingleError(w, http.StatusBadRequest, err)
	}

	if errs := user.Prepare(true); errs != nil {
		return res.Error(w, http.StatusBadRequest, errs)
	}
	db, err := db.Connect()

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)
	result, err := repo.CreateUser(user)

	if err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}

	return res.Json(w, http.StatusCreated, struct {
		Id uint64 `json:"id"`
	}{
		Id: result,
	})
}

// UpdateUserById: updates a user based on the provided id
func UpdateUserById(w http.ResponseWriter, r *http.Request) types.StatusCode {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		return res.SingleError(w, http.StatusBadRequest, err)
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return res.SingleError(w, http.StatusUnprocessableEntity, err)
	}
	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		return res.SingleError(w, http.StatusBadRequest, err)
	}

	if errs := user.Prepare(false); errs != nil {
		return res.Error(w, http.StatusBadRequest, errs)
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repo := repos.NewUsersRepository(db)

	if err := repo.UpdateUserById(userId, user); err != nil {
		return res.SingleError(w, http.StatusInternalServerError, err)
	}
	return res.Json(w, http.StatusNoContent, nil)
}

// DeleteUserById: deletes a user based on the provided id
func DeleteUserById(w http.ResponseWriter, r *http.Request) types.StatusCode {
	return 0
}

package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
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
		res.SingleError(w, http.StatusNotFound, err)
		return
	}
	res.Json(w, http.StatusOK, user)
}

// CreateUser: creates a user and delegates its persistence
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

// UpdateUserById: updates a user based on the provided id
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
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

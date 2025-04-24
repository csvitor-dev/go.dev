package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	"github.com/csvitor-dev/social-media/internal/db/repos"
	"github.com/csvitor-dev/social-media/internal/models"
	"github.com/gorilla/mux"

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
	repo := repos.NewUserRepo(db)
	result, err := repo.GetUsers()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.JSON(w, http.StatusOK, result)
}

// GetUserByID: retrieves a persisted user via a given ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	repo := repos.NewUserRepo(db)
	user, err := repo.GetById(userId)

	if err != nil {
		res.SingleError(w, http.StatusNotFound, err)
		return
	}
	res.JSON(w, http.StatusOK, user)
}

// CreateUser: creates a user and delegates its persistence
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.Error(w, http.StatusUnprocessableEntity, []error{err})
		return
	}
	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		res.Error(w, http.StatusBadRequest, []error{err})
		return
	}

	if errs := user.Prepare(); errs != nil {
		res.Error(w, http.StatusBadRequest, errs)
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewUserRepo(db)
	result, err := repo.CreateUser(user)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}

	res.JSON(w, http.StatusCreated, struct {
		ID uint64 `json:"id"`
	}{
		ID: result,
	})
}

// UpdateUser: updates a user based on the provided ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser: deletes a user based on the provided ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

}

package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/csvitor-dev/social-media/api/src/db"
	"github.com/csvitor-dev/social-media/api/src/db/repos"
	"github.com/csvitor-dev/social-media/api/src/models"

	res "github.com/csvitor-dev/social-media/api/src/responses"
)

// GetAllUsers: retrieves all persisted users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

// GetUserByID: retrieves a persisted user via a given ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {

}

// CreateUser: creates a user and delegates its persistence
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.Error(w, http.StatusUnprocessableEntity, []error{err})
		return;
	}
	var user models.User
	
	if err = json.Unmarshal(body, &user); err != nil {
		res.Error(w, http.StatusBadRequest, []error{err})
		return;
	}
	db, err := db.Connect()
	
	if err != nil {
		res.Error(w, http.StatusInternalServerError, []error{err})
		return;
	}
	defer db.Close()
	repo := repos.NewUserRepo(db)
	result, err := repo.CreateUser(user)

	if err != nil {
		res.Error(w, http.StatusInternalServerError, []error{err})
		return;
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
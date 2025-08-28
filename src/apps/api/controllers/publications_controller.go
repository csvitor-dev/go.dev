package controllers

import (
	"io"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	"github.com/csvitor-dev/social-media/pkg/requests"
	"github.com/csvitor-dev/social-media/pkg/requests/publication"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/src/services/auth"
)

func Publish(w http.ResponseWriter, r *http.Request) {
	authUserId, err := auth.GetUserIdFromToken()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	var request publication.CreatePubRequest
	requests.MapToRequest(w, &request, body)
	pub, err := request.Map(authUserId)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	repo := repos.NewPublicationsRepository(db)
	pubId, err := repo.Create(pub)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusCreated, map[string]any{
		"author_id": authUserId,
		"pub_id":    pubId,
	})
}

func GetAllPubs(w http.ResponseWriter, r *http.Request) {

}

func GetPubById(w http.ResponseWriter, r *http.Request) {

}

func UpdatePubById(w http.ResponseWriter, r *http.Request) {

}

func DeletePubById(w http.ResponseWriter, r *http.Request) {

}

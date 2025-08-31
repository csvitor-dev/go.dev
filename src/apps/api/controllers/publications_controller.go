package controllers

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	pkg "github.com/csvitor-dev/social-media/pkg/errors"
	"github.com/csvitor-dev/social-media/pkg/requests"
	"github.com/csvitor-dev/social-media/pkg/requests/publication"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/src/services/auth"
	"github.com/gorilla/mux"
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
	authUserId, err := auth.GetUserIdFromToken()

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
	repo := repos.NewPublicationsRepository(db)
	pubs, err := repo.SearchPubsByUserId(authUserId)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusOK, pubs)
}

func GetPubById(w http.ResponseWriter, r *http.Request) {
	pubId, err := strconv.ParseUint(mux.Vars(r)["pubId"], 10, 64)

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
	repo := repos.NewPublicationsRepository(db)
	pub, err := repo.FindById(pubId)

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
	res.Json(w, http.StatusOK, pub)
}

func UpdatePubById(w http.ResponseWriter, r *http.Request) {
	authUserId, err := auth.GetUserIdFromToken()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	pubId, err := strconv.ParseUint(mux.Vars(r)["pubId"], 10, 64)

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
	repo := repos.NewPublicationsRepository(db)

	if err := repo.IsAuthorOfPub(authUserId, pubId); err != nil {
		var status int

		if errors.Is(err, pkg.ErrModelNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusForbidden
		}
		res.SingleError(w, status, err)
		return
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		res.SingleError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var request publication.UpdatePubRequest
	requests.MapToRequest(w, &request, body)
	pub, err := request.Map(authUserId)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if err := repo.Update(pubId, pub); err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

func DeletePubById(w http.ResponseWriter, r *http.Request) {
	authUserId, err := auth.GetUserIdFromToken()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	pubId, err := strconv.ParseUint(mux.Vars(r)["pubId"], 10, 64)

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
	repo := repos.NewPublicationsRepository(db)

	if err := repo.IsAuthorOfPub(authUserId, pubId); err != nil {
		var status int

		if errors.Is(err, pkg.ErrModelNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusForbidden
		}
		res.SingleError(w, status, err)
		return
	}

	if err := repo.Delete(pubId); err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

func GetAllPubsForUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["userId"], 10, 64)

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
	repo := repos.NewPublicationsRepository(db)
	pubs, err := repo.FilterPubsByUserId(userId)

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	res.Json(w, http.StatusOK, pubs)
}

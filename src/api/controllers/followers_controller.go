package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/csvitor-dev/social-media/internal/db"
	repos "github.com/csvitor-dev/social-media/internal/db/repositories"
	pkg "github.com/csvitor-dev/social-media/pkg/errors"
	res "github.com/csvitor-dev/social-media/pkg/responses"
	"github.com/csvitor-dev/social-media/src/api/services/auth"
	"github.com/gorilla/mux"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	userToFollowId, err := strconv.ParseUint(mux.Vars(r)["userId"], 10, 64)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userToFollowId {
		res.SingleError(w, http.StatusBadRequest, errors.New("controllers: following yourself isn't avaliable"))
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewFollowersRepository(db)

	if err = repo.Follow(userToFollowId, followerId); err != nil {
		var status int

		if errors.Is(err, pkg.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		res.SingleError(w, status, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId()

	if err != nil {
		res.SingleError(w, http.StatusUnauthorized, err)
		return
	}
	userToUnfollowId, err := strconv.ParseUint(mux.Vars(r)["userId"], 10, 64)

	if err != nil {
		res.SingleError(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userToUnfollowId {
		res.SingleError(w, http.StatusBadRequest, errors.New("controllers: unfollowing yourself isn't avaliable"))
		return
	}
	db, err := db.Connect()

	if err != nil {
		res.SingleError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repos.NewFollowersRepository(db)

	if err = repo.Unfollow(userToUnfollowId, followerId); err != nil {
		var status int

		if errors.Is(err, pkg.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		res.SingleError(w, status, err)
		return
	}
	res.Json(w, http.StatusNoContent, nil)
}

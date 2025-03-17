package handler

import (
	"log"
	"net/http"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

// send follow / unfollow
func (Handler *Handler) HandleFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var follow models.Follower
	var err error

	err = utils.ParseBody(r, &follow)

	follow.Follower = user.ID
	follow.CreatedAt = int(time.Now().Unix())
	follow.ModifiedAt = follow.CreatedAt

	// follow.UserID, err = strconv.Atoi(r.FormValue("uesr_id"))

	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = Handler.Service.Follow(&follow)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
}

// accept/reject follow request
func (Handler *Handler) HandleFollowRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var follow models.Follower
	var err error

	err = utils.ParseBody(r, &follow)
	follow.UserID = user.ID
	follow.ModifiedAt = int(time.Now().Unix())
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = Handler.Service.FollowDecision(&follow)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
}

func (Handler *Handler) HandleGetFollowRequests(w http.ResponseWriter, r *http.Request) {
	user, timeBefore, err := Handler.AfterGet(w, r)
	if err != nil {
		return
	}

	requesters, err := Handler.Service.GetFollowRequests(&user, timeBefore.Before)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	utils.WriteJson(w, http.StatusOK, requesters)
}

func (Handler *Handler) HandleGetFollowers(w http.ResponseWriter, r *http.Request) {
	user, timeBefore, err := Handler.AfterGet(w, r)
	if err != nil {
		return
	}

	followers, err := Handler.Service.GetFollowers(&user, timeBefore.Before)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	utils.WriteJson(w, http.StatusOK, followers)
}

func (Handler *Handler) HandleGetFollowings(w http.ResponseWriter, r *http.Request) {
	user, timeBefore, err := Handler.AfterGet(w, r)
	if err != nil {
		return
	}

	followings, err := Handler.Service.GetFollowings(&user, timeBefore.Before)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	utils.WriteJson(w, http.StatusOK, followings)
}

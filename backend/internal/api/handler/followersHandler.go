package handler

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Println(user.ID , follow.UserID)
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

func (Handler *Handler) GetFollowers(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) GetFollowings(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) GetRequestsFollowers(w http.ResponseWriter, r *http.Request) {
}

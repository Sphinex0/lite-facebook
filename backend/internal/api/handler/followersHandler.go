package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddFollow(w http.ResponseWriter, r *http.Request) {
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

	follow.Follower = user.ID
	follow.UserID, err = strconv.Atoi(r.FormValue("target"))
	
	if err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = Handler.Service.CreateFollow(&follow)

	if err != nil {
		fmt.Println(err)
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

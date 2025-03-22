package handler

import (
	"log"
	"net/http"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
)

// send follow / unfollow
func (Handler *Handler) HandleFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var follow models.Follower
	var err error

	err = utils.ParseBody(r, &follow)

	follow.Follower = user.ID
	follow.CreatedAt = int(time.Now().UnixMilli())
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
	utils.WriteJson(w, http.StatusOK, follow)
}

// accept/reject follow request
func (Handler *Handler) HandleFollowRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var follow models.Follower
	var err error

	err = utils.ParseBody(r, &follow)
	follow.UserID = user.ID
	follow.ModifiedAt = int(time.Now().UnixMilli())
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
	user, timeBefore, err1 := Handler.AfterGet(w, r)
	if err1.Err != nil {
		return
	}

	requesters, err := Handler.Service.GetFollowRequests(&user, timeBefore.Before, user.ID)
	if err != nil {
		log.Println(err)
		if err.Error() == "profile is private, follow to see"{
			utils.WriteJson(w, http.StatusBadRequest, "information locked")
			return
		}
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	utils.WriteJson(w, http.StatusOK, requesters)
}

func (Handler *Handler) HandleGetFollowers(w http.ResponseWriter, r *http.Request) {
	user, data, err1 := Handler.AfterGet(w, r)
	if err1.Err != nil {
		return
	}
	var targetUser models.UserInfo
	targetUser.ID = user.ID
	if(data.UserID != 0){
		targetUser.ID =data.UserID
	}
	followers, err := Handler.Service.GetFollowers(&targetUser, data.Before, user.ID)
	if err != nil {
		log.Println(err)
		if err.Error() == "profile is private, follow to see"{
			utils.WriteJson(w, http.StatusBadRequest, "information locked")
			return
		}
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	utils.WriteJson(w, http.StatusOK, followers)
}

func (Handler *Handler) HandleGetFollowings(w http.ResponseWriter, r *http.Request) {
	user, data, err1 := Handler.AfterGet(w, r)
	if err1.Err != nil {
		return
	}
	var targetUser models.UserInfo
	targetUser.ID = user.ID
	if(data.UserID != 0){
		targetUser.ID =data.UserID
	}
	followings, err := Handler.Service.GetFollowings(&targetUser, data.Before, user.ID)
	if err != nil {
		log.Println(err)
		if err.Error() == "profile is private, follow to see"{
			utils.WriteJson(w, http.StatusBadRequest, "information locked")
			return
		}
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	utils.WriteJson(w, http.StatusOK, followings)
}

func (Handler *Handler) HandleGetGroupInvitable(w http.ResponseWriter, r *http.Request) {
	user, data, err1 := Handler.AfterGet(w, r)
	if err1.Err != nil {
		return
	}

	users, err := Handler.Service.GetGroupInvitables(data.BeforeID, user.ID, data.GroupID)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	log.Println("users are :", data.GroupID, data.BeforeID)
	utils.WriteJson(w, http.StatusOK, users)
}
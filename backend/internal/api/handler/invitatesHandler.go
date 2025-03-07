package handler

import (
	"fmt"
	"log"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddInviteByReceiver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Invite models.Invite
	err := utils.ParseBody(r, &Invite)
	fmt.Println(err)
	if err != nil || Invite.Receiver == 0 || Invite.GroupID == 0 {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	Invite.Sender = user.ID
	if err := Handler.Service.CreateInvite(Invite); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) HandleInviteRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var invite models.Invite
	var err error

	err = utils.ParseBody(r, &invite)
	// fmt.Println(user.ID , follow.UserID)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = Handler.Service.FollowDecision(&invite)
	if err != nil {
		log.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
}

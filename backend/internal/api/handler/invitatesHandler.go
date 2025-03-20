package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (Handler *Handler) AddInvite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Invite models.Invite
	err := utils.ParseBody(r, &Invite)
	Invite.Sender = user.ID

	if err != nil || Invite.Receiver == 0 || Invite.GroupID == 0 || Invite.Sender == Invite.Receiver {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := Handler.Service.CreateInvite(Invite); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) HandleInviteRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var Invite models.Invite
	var err error

	Invite.Sender = user.ID
	err = utils.ParseBody(r, &Invite)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = Handler.Service.InviderDecision(&Invite)
	if err != nil {
		log.Println("ttttttt", err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
}

func (Handler *Handler) GetInvites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	Invites, err := Handler.Service.AllInvites(user.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Invites)
}

func (Handler *Handler) GetMembers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var Invite models.Invite
	err := utils.ParseBody(r, &Invite)
	fmt.Println(Invite)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	Invites, err := Handler.Service.AllMembers(Invite.GroupID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	fmt.Println(Invites)
	valid, err := Handler.Service.Members(Invites)
	utils.WriteJson(w, http.StatusOK, valid)
}

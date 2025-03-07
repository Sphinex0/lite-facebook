package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddInviteByReceiver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	IdGroup, err := strconv.Atoi(r.PathValue("IdGroup"))
	if err != nil {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	IdReciver, err := strconv.Atoi(r.PathValue("IdReciver"))
	if err != nil {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if err := Handler.Service.CreateInviteByReceiver(IdGroup, IdReciver); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) AddInviteBySender(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	IdSender :=user.ID
	fmt.Println(IdSender)
	IdGroup, err := strconv.Atoi(r.PathValue("IdGroup"))
	if err != nil {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	IdReciver, err := strconv.Atoi(r.PathValue("IdReciver"))
	if err != nil {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if err := Handler.Service.CreateInviteBySernder(IdGroup, IdSender,IdReciver); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) GetGroupsInvites(w http.ResponseWriter, r *http.Request) {
}

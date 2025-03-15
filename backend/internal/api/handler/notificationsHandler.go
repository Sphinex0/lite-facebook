package handler

import (
	"net/http"

	utils "social-network/pkg"
)

func (H *Handler) HandleGetNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var usrId string
	err := utils.ParseBody(r, &usrId)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	notifications,err := H.Service.GetUserNotifications(usrId); if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	utils.WriteJson(w, http.StatusOK, notifications)
}
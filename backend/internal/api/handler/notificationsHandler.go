package handler

import (
	"net/http"
	"strconv"

	utils "social-network/pkg"
)

func (H *Handler) HandleGetNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := utils.GetUserFromContext(r.Context()); if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "User not Found")
		return
	}
	id := strconv.Itoa(user.ID)
	notifications,err := H.Service.GetUserNotifications(id); if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	utils.WriteJson(w, http.StatusOK, notifications)
}
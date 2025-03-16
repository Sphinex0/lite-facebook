package handler

import (
	"net/http"
	"strconv"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (H *Handler) HandleGetNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "User not Found")
		return
	}

	id := strconv.Itoa(user.ID)
	notifications, count, err := H.Service.GetUserNotifications(id)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	response := struct {
		Notifications []models.Notification `json:"notifications"`
		Unseen        int                   `json:"unseen"`
	}{
		Notifications: notifications,
		Unseen:        count,
	}

	utils.WriteJson(w, http.StatusOK, response)
}

func (H *Handler) MarkNotificationAsSeen(w http.ResponseWriter, r *http.Request) {
	// get the notification id from body
	
	// check it if exists
	// mark as seen
}

package handler

import (
	"fmt"
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

	user, _, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "User not Found")
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	id := strconv.Itoa(user.ID)
	notifications, count, err := H.Service.GetUserNotifications(id, page)
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

func (H *Handler) HandleDeleteNotification(w http.ResponseWriter, r *http.Request) {
	// get the notification id from body
	var ntfID struct {
		ID int `json:"id"`
	}
	err := utils.ParseBody(r, &ntfID)
	if err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	user, _, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "unothorized")
		return
	}

	err = H.Service.Deletentfc(ntfID.ID, user.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusUnauthorized, "unothorized")
		return
	}

	utils.WriteJson(w, http.StatusOK, "Marked succesfuly")
}

func (H *Handler) MarkNotificationAsSeen(w http.ResponseWriter, r *http.Request) {
	// get the notification id from body
	var ntfID int
	err := utils.ParseBody(r, ntfID)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
		return
	}

	user, _, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "unothorized")
		return
	}

	err = H.Service.MarkAsseen(ntfID, user.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusUnauthorized, "unothorized")
		return
	}

	utils.WriteJson(w, http.StatusOK, "Marked succesfuly")
}

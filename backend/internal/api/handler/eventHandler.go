package handler

import (
	"fmt"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Event models.Event
	err := utils.ParseBody(r, &Event)
	fmt.Println(user.ID)
	Event.UserID = user.ID
	if err != nil || Event.UserID == 0 || Event.GroupID == 0 {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := Handler.Service.CreateEvent(Event); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
}

// skip

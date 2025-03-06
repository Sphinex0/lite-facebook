package handler

import (
	"net/http"
	"strings"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
	"social-network/pkg/middlewares"
)

func (Handler *Handler) AddGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	User, ok := r.Context().Value(middlewares.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Group models.Group
	Group.Creator = User.ID
	Group.Title = strings.TrimSpace(r.FormValue("Title"))
	Group.Description = strings.TrimSpace(r.FormValue("Description"))
	Group.CreatedAt=int(time.Now().Unix())
	
	
}

func (Handler *Handler) GetGroups(w http.ResponseWriter, r *http.Request) {
}

func (Handler *Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
}

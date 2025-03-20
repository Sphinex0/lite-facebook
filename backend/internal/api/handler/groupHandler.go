package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (Handler *Handler) AddGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var Group models.Group
	Group.Creator = user.ID
	Group.Title = strings.TrimSpace(r.FormValue("Title"))
	Group.Description = strings.TrimSpace(r.FormValue("Description"))
	Group.CreatedAt = int(time.Now().UnixMilli())
	fmt.Println(Group.Title)
	fmt.Println(Group.Description)
	if Group.Title == "" || len(Group.Title) > 50 || Group.Description == "" || len(Group.Description) > 250 {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusInternalServerError))
		return
	}
	if err := Handler.Service.GreatedGroup(&Group); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (Handler *Handler) GetGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	grp, err := Handler.Service.AllGroups()
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grp)
}

func (Handler *Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var Groups models.Group
	err := utils.ParseBody(r, &Groups)
	group, err := Handler.Service.GetGroupsById(&Groups)
	types, err := Handler.Service.TypeInvate(user.ID, group.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusNotAcceptable, "Not Acceptable")
		return
	}
	var group_info models.GroupInfo
	group_info.Group = *group
	group_info.Action = types
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group_info)
}

func (Handler *Handler) GetMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	group, err := Handler.Service.GetMemberById(user.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

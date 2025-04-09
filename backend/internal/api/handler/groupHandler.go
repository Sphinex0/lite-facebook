package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	// Extract profile picture (optional)
	file, handler, err := r.FormFile("image")
	Group.Image = "default-group.png"
	if err == nil {
		defer file.Close()
		Group.Image, err = utils.StoreThePic("public/pics", file, handler)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "internalserver error")
			return
		}
	}

	Group.CreatedAt = int(time.Now().UnixMilli())
	if Group.Title == "" || len(Group.Title) > 50 || Group.Description == "" || len(Group.Description) > 250 {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Println(Group)
	if err := Handler.Service.GreatedGroup(&Group); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	utils.WriteJson(w, http.StatusOK, Group)
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
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Status Bad Request")
		return
	}
	group, err := Handler.Service.GetGroupsById(&Groups)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Status Bad Request")
		return
	}
	types, err := Handler.Service.TypeInvate(user.ID, group.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Status Bad Request")
		return
	}
	var group_info models.GroupInfo
	group_info.Group = *group
	group_info.Action = types

	//
	creator_info, err := Handler.Service.GetUserByID(group.Creator)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server")
		return
	}
	group_info.CreatorName = fmt.Sprintf("%v %v", creator_info.First_Name, creator_info.Last_Name)
	//
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

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

	var Group models.Group
	Group.Creator = 1
	Group.Title = strings.TrimSpace(r.FormValue("Title"))
	Group.Description = strings.TrimSpace(r.FormValue("Description"))
	Group.CreatedAt = int(time.Now().Unix())
	fmt.Println(Group.Title)
	fmt.Println(Group.Description)
	if err := Handler.Service.GreatedGroup(&Group); err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (Handler *Handler) GetGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var Groups []models.Group
	grp, err := Handler.Service.AllGroups(&Groups)
	if err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grp)
}

func (Handler *Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	idStr := r.PathValue("id")
	var Groups models.Group
	group, err:=Handler.Service.GetGroupsById(&Groups, idStr)
	if err != nil {
		fmt.Println(err)
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

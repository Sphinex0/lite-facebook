package handler

import (
	"encoding/json"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (H *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var user models.User
	if err := utils.ParseBody(r, &user); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err := H.Service.LoginUser(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	userinfo := models.UserInfo{
		Nickname:   user.Nickname,
		First_Name: user.First_Name,
		Last_Name:  user.Last_Name,
		Image:      user.Image,
	}

	utils.WriteJson(w, http.StatusOK, userinfo)
}

func (H *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var user models.User
	if erro := json.NewDecoder(r.Body).Decode(&user); erro != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}
	// Proccess Data and Insert it
	err := H.Service.RegisterUser(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
	}

	utils.WriteJson(w, http.StatusOK, "You'v loged in succesfuly")
}

func (H *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
	}

	utils.DeleteSessionCookie(w, user.Uuid)
	utils.WriteJson(w, http.StatusOK, "You Logged Out Successfuly!")
}

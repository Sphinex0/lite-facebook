package handler

import (
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

	Uuid, err := H.Service.LoginUser(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	userinfo, err := H.Service.Database.GetuserInfo(user.ID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "internal server error")
	}
	utils.SetSessionCookie(w, Uuid)
	utils.WriteJson(w, http.StatusOK, userinfo)
}

func (H *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	user := H.Service.Extractuser(r)

	// Parse the multipart form (10MB max file size)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	// Extract profile picture (optional)
	file, handler, err := r.FormFile("avatar")
	if err == nil {
		defer file.Close()
		user.Image, err = utils.StoreThePic("../front-end/public/pics", file, handler)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "internalserver error")
		}
	}
	// Proccess Data and Insert it
	Uuid, err := H.Service.RegisterUser(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	// some data that to make it easy in the front-end
	userinfo := models.UserInfo{
		First_Name: user.First_Name,
		Last_Name:  user.Last_Name,
		Image:      user.Image,
	}

	utils.SetSessionCookie(w, Uuid)
	utils.WriteJson(w, http.StatusOK, userinfo)
}

func (H *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	_,uuid, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "User not Found")
		return
	}
	err := H.Service.DeleteSessionCookie(w, uuid)
	if err != nil {
		utils.WriteJson(w, http.StatusOK, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, "You Logged Out Successfuly!")
}

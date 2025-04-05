package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		fmt.Println("err",err)
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
	var err1 error
	user.DateBirth , err1 = strconv.Atoi( r.FormValue("dob"))
	if err1 != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	// Parse the multipart form (10MB max file size)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	// Extract profile picture (optional)
	file, handler, err := r.FormFile("avatar")
	if err == nil {
		// user.Image, err = utils.StoreThePic("../front-end/public/pics", file, handler)
		defer file.Close()
		user.Image, err = utils.StoreThePic("public/pics", file, handler)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "internalserver error")
			return
		}
	} else {
		user.Image = "default-profile.png"
	}

	fmt.Println("user", user)
	// Proccess Data and Insert it
	_, err, _ = H.Service.RegisterUser(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	// some data that to make it easy in the front-end
	// userinfo := models.UserInfo{
	// 	ID:         id,
	// 	First_Name: user.First_Name,
	// 	Last_Name:  user.Last_Name,
	// 	Image:      user.Image,
	// }
	// utils.SetSessionCookie(w, Uuid)
	utils.WriteJson(w, http.StatusOK, "successfully")
}

func (H *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	_, uuid, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "User not found")
		return
	}

	err := H.Service.DeleteSessionCookie(w, uuid)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "Error logging out: "+err.Error())
		return
	}

	utils.WriteJson(w, http.StatusOK, "You logged out successfully!")
}
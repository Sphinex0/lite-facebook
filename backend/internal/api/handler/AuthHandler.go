package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
		fmt.Println("yey", err)
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	err := H.Service.LoginUser(&user)
	if err != nil {
		fmt.Println("err",err.Error())
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
	// Parse the multipart form (10MB max file size)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "file too big")
		return
	}

	user := H.Service.Extractuser(r)

	// Extract profile picture (optional)
	var filePath string
	file, handler, err := r.FormFile("avatar")
	fmt.Println(err)
	if err == nil { // No error means a file was uploaded
		defer file.Close()

		// Ensure Profile directory exists
		uploadDir := "../backend/internal/repository/profile"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			fmt.Println("yes it does make a file")
			os.Mkdir(uploadDir, os.ModePerm)
		}

		// Save with a unique filename (e.g., user.UUID + filename)
		filePath = filepath.Join(uploadDir, handler.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			utils.WriteJson(w, http.StatusInternalServerError, "Could not save file")
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "Failed to save file")
			return
		}

		// Assign file path to user struct
		user.Image = filePath
		fmt.Println("image path", filePath)
	}

	// Proccess Data and Insert it
	err = H.Service.RegisterUser(&user)
	if err != nil {
		fmt.Println("grvrgvr", err.Error())
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, "You'v loged in succesfuly")
}

func (H *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
	}
	err = H.Service.DeleteSessionCookie(w, user.Uuid)
	if err != nil {
		utils.WriteJson(w, http.StatusOK, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, "You Logged Out Successfuly!")
}

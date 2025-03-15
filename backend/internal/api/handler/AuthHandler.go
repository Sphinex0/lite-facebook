package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"

	"github.com/mattn/go-sqlite3"
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

	userinfo, err := H.Service.Database.GetuserInfo(user.ID); if err != nil {
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
	var filePath string
	file, handler, err := r.FormFile("avatar")
	if err == nil { // No error means a file was uploaded

		// Ensure Profile directory exists
		uploadDir := "../backend/internal/repository/profile"
		defer file.Close()
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		// Save with a unique filename (e.g., user.UUID + filename)
		// uuid.New() + "."
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
	}

	// Proccess Data and Insert it
	Uuid, err := H.Service.RegisterUser(&user)
	if err != nil {
		fmt.Println("yes", err.Error())
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	user.Uuid = Uuid
	utils.SetSessionCookie(w, Uuid)
	fmt.Println("user", user)
	utils.WriteJson(w, http.StatusOK, user)
}

func (H *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var Uuid string
	err := json.NewDecoder(r.Body).Decode(&Uuid)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "bad request")
	}
	err = H.Service.DeleteSessionCookie(w, Uuid)
	if err != nil {
		utils.WriteJson(w, http.StatusOK, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, "You Logged Out Successfuly!")
}

func (H *Handler) CheckUserValidity(w http.ResponseWriter, r *http.Request) {
	var Authorized bool
	// parse user uid
	userUID, err := r.Cookie("session_token")
	if err != nil {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Get Info Data
	Authorized, err = H.Service.GetInfoData(userUID.Value)
	if err != nil {
		if err == sqlite3.ErrLocked {
			utils.WriteJson(w, http.StatusLocked, struct {
				Message string `json:"message"`
			}{Message: "Database Locked"})
			return
		}

		utils.WriteJson(w, http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{Message: "Internal Server Error"})
		return
	}

	if !H.Service.Database.CheckExpiredCookie(userUID.Value, time.Now()) {
		Authorized = false
		err = H.Service.DeleteSessionCookie(w, userUID.Value)
		if err != nil {
			utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}

	utils.WriteJson(w, http.StatusOK, Authorized)
}

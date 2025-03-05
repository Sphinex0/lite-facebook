package handler

import (
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (H *Handler) Login(w http.ResponseWriter, r *http.Request) {
	utils.SetSessionCookie(w, "550e8400-e29b-41d4-a716-446655440000")
	utils.WriteJson(w, 200, "nice")
	return
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
		utils.WriteJson(w, http.StatusBadRequest, "Error While logging To An  Account.")
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
	w.Write([]byte("Hellolooo"))
	// if r.Method != http.MethodPost {
	// 	utils.WriteJson(w, http.StatusMethodNotAllowed, "Method not allowed")
	// 	return
	// }

	// var user models.User
	// if erro := json.NewDecoder(r.Body).Decode(&user); erro != nil {
	// 	utils.WriteJson(w, http.StatusBadRequest, "Bad request")
	// 	return
	// }
	// // Proccess Data and Insert it
	// err := H.Service.RegisterUser(&user)
	// if err != nil {
	// 	if err == sqlite3.ErrLocked {
	// 		utils.WriteJson(w, http.StatusLocked, "Database Is Busy!")
	// 		return
	// 	}
	// 	// Username
	// 	if err.Error() == models.Errors.InvalidUsername {
	// 		utils.WriteJson(w, http.StatusBadRequest, err.Error())

	// 		return
	// 	}
	// 	// gender
	// 	if err.Error() == models.Errors.InvalidCredentials {
	// 		utils.WriteJson(w, http.StatusBadRequest, "bad request gender!")
	// 		return
	// 	}

	// 	// Age
	// 	if err.Error() == models.UserErrors.InvalideAge {
	// 		utils.WriteJson(w, http.StatusBadRequest, err.Error())

	// 		return
	// 	}

	// 	// Password
	// 	if err.Error() == models.Errors.InvalidPassword {
	// 		utils.WriteJson(w, http.StatusBadRequest, err.Error())

	// 		return
	// 	}
	// 	// Email
	// 	if err.Error() == models.Errors.InvalidEmail {
	// 		utils.WriteJson(w, http.StatusBadRequest, err.Error())

	// 		return
	// 	}
	// 	if err.Error() == models.Errors.LongEmail {
	// 		utils.WriteJson(w, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// 	// General
	// 	if err.Error() == models.Errors.UserAlreadyExist {
	// 		utils.WriteJson(w, http.StatusConflict, models.Errors.UserAlreadyExist)
	// 		return
	// 	}

	// 	utils.WriteJson(w, http.StatusInternalServerError, "Error While Registering The User.")
	// 	return
	// }
	// utils.WriteJson(w, http.StatusOK, "You'v loged succesfuly")
}

func (Handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

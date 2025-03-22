package service

import (
	"social-network/internal/models"
)

func (service *Service) GetUserByID(id int) (user models.UserInfo, err error) {
	user, err = service.Database.GetUserByID(id)
	return
}


func (service *Service) GetAllUser(before int, currentUser int) (users []models.UserInfo, err error) {
	users, err = service.Database.GetAllUsers(before, currentUser)
	return
}
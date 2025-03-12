package service

import (
	"social-network/internal/models"
)

func (service *Service) GetUserByID(id int) (user models.UserInfo, err error) {
	user, err = service.Database.GetUserByID(id)
	return
}

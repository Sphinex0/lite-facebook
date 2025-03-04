package service

import "social-network/internal/models"

func (S *Service) LoginUser(User *models.User) error {
	//

	//
	S.Database.StoreSession(*User)
	return nil
}
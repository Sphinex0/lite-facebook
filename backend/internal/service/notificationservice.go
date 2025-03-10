package service

import (
	"errors"

	"social-network/internal/models"
	repository 	"social-network/internal/repository"

)

func (S *Service) GetUserNotifications(usrId string) ([]models.Notification, (error)) {
	if !repository.CheckIfUserExistsById(usrId, S.Database.Db){
		return []models.Notification{}, errors.New("invaibale user")
	}

	return S.Database.GetUserNotifications(usrId)
}

func (s *Service) AddNotification(notification models.Notification) error {
	switch notification.Type {
	case "follow-request", "invitation-request":
		if !repository.CheckIfUserExistsById(notification.InvokerID,s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID,s.Database.Db) {
			return errors.New("invalid users")
		}

	case "event-created":
		if !s.Database.CheckIfEventExists(notification.EventID) {
			return errors.New("invalid event")
		}
	}
	
	err := s.Database.InsertNotification(notification); if err != nil {
		return err
	}
	
	return nil
}

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
	case "follow-request":
		if !repository.CheckIfUserExistsById(notification.InvokerID,s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID,s.Database.Db) {
			return errors.New("invalid users")
		}
	case "invitation-request":
		if !repository.CheckIfUserExistsById(notification.InvokerID,s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID,s.Database.Db) || !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db){
			return errors.New("invalid users or a group")
		} 

	case "event-created":
		if !s.Database.CheckIfEventExists(notification.EventID) || !repository.CheckIfUserExistsById(notification.InvokerID, s.Database.Db) || !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db){
			return errors.New("invalid event")
		}
	
	case "join":
		if !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID, s.Database.Db) {
			return errors.New("invalide user or group")
		}

	default:
		return errors.New("invalide type")
	}

	err := s.Database.InsertNotification(notification); if err != nil {
		return err
	}
	
	return nil
}

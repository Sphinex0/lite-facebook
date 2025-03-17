package service

import (
	"errors"

	"social-network/internal/models"
	repository "social-network/internal/repository"
)

func (S *Service) GetUserNotifications(usrId string) ([]models.Notification, int, error) {
	if !repository.CheckIfUserExistsById(usrId, S.Database.Db) {
		return []models.Notification{}, 0, errors.New("invaibale user")
	}

	notifications, err := S.Database.GetUserNotifications(usrId)
	if err != nil {
		return []models.Notification{}, 0, errors.New("error while counting unseen notifications")
	}

	count, err := S.Database.CountUnSeenNotifications(usrId)
	if err != nil {
		return []models.Notification{}, 0, errors.New("error while counting unseen notifications")
	}
	return notifications, count, nil
}

func (s *Service) AddNotification(notification models.Notification) error {
	switch notification.Type {
	case "follow-request":
		if !repository.CheckIfUserExistsById(notification.InvokerID, s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID, s.Database.Db) {
			return errors.New("invalid users")
		}
	case "invitation-request":
		if !repository.CheckIfUserExistsById(notification.InvokerID, s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID, s.Database.Db) || !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db) {
			return errors.New("invalid users or a group")
		}

	case "event-created":
		if !s.Database.CheckIfEventExists(notification.EventID) || !repository.CheckIfUserExistsById(notification.InvokerID, s.Database.Db) || !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db) {
			return errors.New("invalid event")
		}

	case "join":
		if !repository.CheckGroupIfExistsById(notification.GroupID, s.Database.Db) || !repository.CheckIfUserExistsById(notification.UserID, s.Database.Db) {
			return errors.New("invalide user or group")
		}

	default:
		return errors.New("invalide type")
	}

	err := s.Database.InsertNotification(notification)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) MarkAsseen(ntfId, userID int) error {
	if !s.Database.CheckNotifValidation(ntfId) {
		return errors.New("not valide notification")
	}

	err := s.Database.MarkAsseen(ntfId, userID); if err != nil {
		return errors.New("error while ineraction with db")
	}

	return nil
}
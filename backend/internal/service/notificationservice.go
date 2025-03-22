package service

import (
	"errors"
	"fmt"
	"strconv"

	"social-network/internal/models"
	repository "social-network/internal/repository"
)

func (S *Service) GetUserNotifications(usrId string, page int) ([]models.Notification, int, error) {
	if !repository.CheckIfUserExistsById(usrId, S.Database.Db) {
		return []models.Notification{}, 0, errors.New("invaibale user")
	}

	if page <= 0 {
		return []models.Notification{}, 0, errors.New("invalide page")
	}

	// Transfer "page" to "from" (page 1 mean page one that has 100 comment from 1 mean comment 1)
	from := (10 * page) - 10

	allntf, err := S.Database.Countallusernotif(usrId)
	if err != nil {
		return []models.Notification{}, 0, err
	}

	if from > allntf {
		return nil, 0, errors.New("invalide page")
	}

	notifications, err := S.Database.GetUserNotifications(usrId, from)
	if err != nil {
		return []models.Notification{}, 0, errors.New("error while counting unseen notifications")
	}

	count, err := S.Database.CountUnSeenNotifications(usrId)
	if err != nil {
		return []models.Notification{}, 0, errors.New("error while counting unseen notifications")
	}

	usrID, err := strconv.Atoi(usrId); if err != nil {
		return nil, 0, errors.New("converting userid to int")
	}

	for _,notification := range notifications {
		err = S.Database.MarkAsseen(notification.ID, usrID); if err != nil {
			return []models.Notification{}, 0, errors.New("error while turning notifications as seen")
		}
	}

	return notifications, count, nil
}

func (s *Service) AddNotification(notification models.Notification) error {
	switch notification.Type {
	case "follow-request", "follow":
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
	fmt.Print("here")
	err := s.Database.InsertNotification(notification)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Service) MarkAsseen(ntfId, userID int) error {
	if !s.Database.CheckNotifValidation(ntfId) {
		return errors.New("not valide notification")
	}

	err := s.Database.MarkAsseen(ntfId, userID)
	if err != nil {
		return errors.New("error while ineraction with db")
	}

	return nil
}

func (s *Service) Deletentfc(ntfId, userID int) error {
	if !s.Database.CheckifUsrMatchNtfc(ntfId, userID) {
		return errors.New("invalide notification")
	}

	err := s.Database.DeleteNotification(ntfId, userID); if err != nil {
		return errors.New("database error")
	}
	return nil
}


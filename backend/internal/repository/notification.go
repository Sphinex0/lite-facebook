package repository

import (
	"social-network/internal/models"
)



func (database *Database) GetUserNotifications(userID string) ([]models.Notification, error) {
	/*
		1- invoker valide group not valide
		2- invoker valide group valide event not valide
	*/
	rows, err := database.Db.Query(``, userID)
	if err != nil {
		return []models.Notification{}, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.Type, &notification.InvokerID, &notification.GroupID, &notification.EventID)
		if err != nil {
			return []models.Notification{}, err
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (database *Database) CheckIfEventExists(EventId int) bool {
	var exists bool

	err := database.Db.QueryRow("SELECT EXISTS(SELECT title FROM events WHERE id = ?)", EventId).Scan(&exists)

	return err == nil
}

func (database *Database) InsertNotification(notification models.Notification) error {
	_,err := database.Db.Exec("INSERT INTO notifications user_id, type, invoker_id, group_id, event_id", 
	notification.UserID, notification.Type, notification.InvokerID, notification.GroupID, notification.EventID); if err != nil {
		return err
	}
	return nil
}
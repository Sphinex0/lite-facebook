package repository

import (
	"fmt"

	"social-network/internal/models"
)

func (database *Database) GetUserNotifications(userID string, start int) ([]models.Notification, error) {
	rows, err := database.Db.Query(`
	SELECT notification_id, notification_type, invoker_name, group_name, group_id, event_name 
	FROM user_notifications
	WHERE notified_user_id = ?
	LIMIT ? OFFSET ?
`, userID, 10, start);if err != nil{
	return []models.Notification{}, err
} 

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(
			&notification.ID,
			&notification.Type,
			&notification.InvokerName,
			&notification.GroupTitle,
			&notification.EventID,
			&notification.EventName,
		)
		
		if err != nil {
			return []models.Notification{}, err
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("err rows",err)
		return nil, err
	}
	rows.Close()

	return notifications, nil
}

func (database *Database) CheckIfEventExists(EventId int) bool {
	var exists bool

	err := database.Db.QueryRow("SELECT EXISTS(SELECT title FROM events WHERE id = ?)", EventId).Scan(&exists)

	return err == nil
}

func (database *Database) InsertNotification(notification models.Notification) error {
	_, err := database.Db.Exec("INSERT INTO notifications user_id, type, invoker_id, group_id, event_id",
		notification.UserID, notification.Type, notification.InvokerID, notification.GroupID, notification.EventID)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) CountUnSeenNotifications(userID string) (int, error) {
	var num int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id = ? AND seen = 0", userID).Scan(&num)
	return num, err
}

func (database *Database) Countallusernotif(userID string) (int, error) {
	var num int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id = ?", userID).Scan(&num)
	return num, err
}

func (database *Database) CheckNotifValidation(ntfId int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM notifications WHERE id = ?)`
	err := database.Db.QueryRow(query, ntfId).Scan(&exists)
	if err != nil {
		fmt.Println("err checknotification", err)
		return false
	}
	return exists
}

func (database *Database) MarkAsseen(ntfId, userID int) error {
	_, err := database.Db.Exec(`UPDATE notifications SET seen = 1 WHERE id = ? AND user_id = ?`, ntfId, userID)
	return err
}
